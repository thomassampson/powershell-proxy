package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

type ResponsePayload struct {
	Message interface{} `json:"message"`
	Level   string      `json:"level"`
}

type CommandRequestBody struct {
	Commands []string `json:"commands"`
}

var (
	ListenAddress        string = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort           string = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName              string = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType       string = os.Getenv("PWSHPRXY_TYPE")
	OktaClientId         string = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaIssuer           string = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaAudience         string = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")
	Shell                string
	JSONDepth            int = 4
	Auth                 map[string]interface{}
	ListenAddressDefault string = "0.0.0.0"
	ListenPortDefault    string = "8000"
	AppNameDefault       string = "Powershell Proxy API"
)

func validateJwt(jwt string) (token *jwtverifier.Jwt, err error) {

	toValidate := map[string]string{}
	toValidate["aud"] = OktaAudience
	toValidate["cid"] = OktaClientId

	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           OktaIssuer,
		ClaimsToValidate: toValidate,
	}

	verifier := jwtVerifierSetup.New()

	return verifier.VerifyAccessToken(jwt)
}

func secureRoute(c *fiber.Ctx) (err error) {
	jwt := strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)
	token, err := validateJwt(jwt)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return c.SendStatus(401)
	}
	Auth = token.Claims
	log.Printf("INFO: Token found for user: %s", Auth["sub"])
	return c.Next()

}

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		log.Printf("INFO: IP Address: %s is Valid", ip)
		return true
	}
}

func validateConfig() (err error) {
	if ListenAddress == "" {
		log.Printf("INFO: Env Variable 'PWSHPRXY_LISTEN_ADDR' not set, defaulting to %s", ListenAddressDefault)
		ListenAddress = ListenAddressDefault
	}

	if !checkIPAddress(ListenAddress) {
		return fmt.Errorf("FATAL: Env Variable 'PWSHPRXY_LISTEN_ADDR' is set, but '%s' is not a valid ipv4 address", ListenAddress)
	}

	if ListenPort == "" {
		log.Printf("INFO: Env Variable 'PWSHPRXY_LISTEN_PORT' not set, defaulting to %s", ListenPortDefault)
		ListenPort = ListenPortDefault
	} else {
		_, err := strconv.Atoi(ListenPort)
		if err != nil {
			return fmt.Errorf("FATAL: Env Variable 'PWSHPRXY_LISTEN_PORT' is set, but '%s' is not a number", ListenPort)
		}
	}

	if OktaClientId == "" {
		return errors.New("FATAL: Env Variable 'PWSHPRXY_OKTA_CLIENT_ID' not set")
	}

	if OktaAudience == "" {
		return errors.New("FATAL: Env Variable 'PWSHPRXY_OKTA_AUDIENCE' not set")
	}

	if OktaIssuer == "" {
		return errors.New("FATAL: Env Variable 'PWSHPRXY_OKTA_ISSUER' not set")
	}

	if AppName == "" {
		log.Printf("INFO: Env Variable 'PWSHPRXY_APP_NAME' not set, defaulting to %s", AppNameDefault)
		AppName = AppNameDefault
	}

	if PowerShellType == "" {
		return errors.New("FATAL: Env Variable 'PWSHPRXY_TYPE' not set, cannot start webserver")
	}

	if strings.ToLower(PowerShellType) == "core" {
		Shell = "pwsh"
	} else if strings.ToLower(PowerShellType) == "windows" {
		Shell = "powershell"
	} else {
		return errors.New("FATAL: Env Variable 'PWSHPRXY_TYPE' must be set to either 'core' or 'powershell'")
	}

	log.Printf("INFO: Using Powershell Type: %s", Shell)
	log.Printf("INFO: Using AppName: %s", AppName)
	log.Printf("INFO: Using ListenPort: %s", ListenPort)
	log.Printf("INFO: Using ListenAddress: %s", ListenAddress)
	log.Printf("INFO: Using OktaClientId: %s", OktaClientId)
	log.Printf("INFO: Using OktaAudience: %s", OktaAudience)
	log.Printf("INFO: Using OktaIssuer: %s", OktaIssuer)

	return nil
}

func convertDepthString(sDepth string) (depth int, err error) {
	if sDepth == "" {
		log.Print(strings.ToUpper(fmt.Sprintf("INFO: Depth Not Set, using default: %d", JSONDepth)))
		return JSONDepth, nil
	}
	qry, err := strconv.Atoi(sDepth)
	if err != nil {
		log.Print(strings.ToUpper(fmt.Sprintf("INFO: [user: %s] Depth Must Be A Number", Auth["sub"])))
		return -1, errors.New(strings.ToUpper("Depth Must Be A Number"))
	}

	if qry <= 0 || qry > 6 {
		log.Print(strings.ToUpper(fmt.Sprintf("INFO: [user: %s] Depth must be a number between 1 and 6, defaulting to %d", Auth["sub"], JSONDepth)))
		return JSONDepth, nil
	}

	log.Printf("INFO: [user: %s] Depth Set by User: %d", Auth["sub"], qry)
	return qry, nil
}

func executeCommand(body CommandRequestBody, depth int) (output bytes.Buffer, err error) {
	var args []string

	args = append(args, "-c", "$WarningPreference = 'SilentlyContinue';")

	args = append(args, body.Commands...)

	args = append(args, fmt.Sprintf("| ConvertTo-Json -depth %d -Compress", depth))

	log.Printf("INFO: [user: %s] Running Commands: %s %v", Auth["sub"], Shell, args)

	cmd := exec.Command(Shell, args...)

	cmd.Stdout = &output

	err = cmd.Run()
	if err != nil {
		log.Print(err.Error())
		return bytes.Buffer{}, err
	}

	return output, nil
}

func main() {
	log.Printf("🔵 Starting %s", AppNameDefault)

	err := validateConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// GET /api
	app.Get("/api", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", AppName)
		return c.SendString(msg)
	})

	// POST /api/command
	app.Post("/api/command", secureRoute, func(c *fiber.Ctx) error {

		depth, err := convertDepthString(c.Query("depth"))
		if err != nil {
			c.Status(400)
			return c.JSON(ResponsePayload{Message: err.Error(), Level: "error"})
		}

		var body CommandRequestBody
		json.Unmarshal(c.Body(), &body)

		if body.Commands == nil {
			c.Status(400)
			return c.JSON(ResponsePayload{Message: strings.ToUpper("You must specify the 'commands' property in the request body."), Level: "error"})
		}

		output, err := executeCommand(body, depth)
		if err != nil {
			c.Status(400)
			return c.JSON(ResponsePayload{Message: err.Error(), Level: "error"})
		}

		c.Status(200)
		c.Set("content-type", "application/json")
		return c.Send(output.Bytes())
	})

	log.Printf("🟢 Started %s", AppName)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", ListenAddress, ListenPort)))

}
