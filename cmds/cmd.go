package cmds

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

type ExecContext = func(name string, arg ...string) *exec.Cmd

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
	ExecCommand          ExecContext
)

func ValidateJwt(jwt string) (token *jwtverifier.Jwt, err error) {

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

func CheckIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		log.Printf("INFO: IP Address: %s is Valid", ip)
		return true
	}
}

func ValidateConfig() (err error) {
	if ListenAddress == "" {
		log.Printf("INFO: Env Variable 'PWSHPRXY_LISTEN_ADDR' not set, defaulting to %s", ListenAddressDefault)
		ListenAddress = ListenAddressDefault
	}

	if !CheckIPAddress(ListenAddress) {
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

func ConvertDepthString(sDepth string) (depth int, err error) {
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
		log.Print(strings.ToUpper(fmt.Sprintf("INFO: [user: %s] Depth must be a number between 1 and 6, defaulting to %d", Auth["sub"],
			JSONDepth)))
		return JSONDepth, nil
	}

	log.Printf("INFO: [user: %s] Depth Set by User: %d", Auth["sub"], qry)
	return qry, nil
}

func ExecuteCommand(body CommandRequestBody, depth int) (output bytes.Buffer, err error) {
	var args []string

	args = append(args, "-c", "$WarningPreference = 'SilentlyContinue';")

	args = append(args, body.Commands...)

	args = append(args, fmt.Sprintf("| ConvertTo-Json -depth %d -Compress", depth))

	log.Printf("INFO: [user: %s] Running Commands: %s %v", Auth["sub"], Shell, args)

	cmd := ExecCommand(Shell, args...)
	cmdError := bytes.Buffer{}
	cmd.Stdout = &output
	cmd.Stderr = &cmdError
	err = cmd.Run()
	if err != nil {
		log.Print(err.Error())
		return bytes.Buffer{}, err
	}

	if cmdError.Len() != 0 {
		return bytes.Buffer{}, errors.New(cmdError.String())
	}

	return output, nil
}
