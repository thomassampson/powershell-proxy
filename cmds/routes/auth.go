package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	cmds "powershell-proxy/cmds"

	"github.com/gofiber/fiber/v2"
)

func CallOktaTokenEndpoint(body []byte) (res *http.Response, err error) {

	var reqBody cmds.TokenRequestBody
	json.Unmarshal(body, &reqBody)

	payload := fmt.Sprintf("grant_type=urn:ietf:params:oauth:grant-type:device_code&client_id=%s&device_code=%s",
		cmds.OktaClientId,
		reqBody.DeviceCode)

	return http.Post(fmt.Sprintf("%s/v1/token", cmds.OktaIssuer),
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(payload))
}

func CallOktaAuthorizeEndpoint() (res *http.Response, err error) {

	payload := fmt.Sprintf("scope=openid profile offline_access&client_id=%s", cmds.OktaClientId)

	return http.Post(fmt.Sprintf("%s/v1/device/authorize", cmds.OktaIssuer),
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(payload))
}

func AuthorizeRoute(c *fiber.Ctx) error {
	res, err := CallOktaAuthorizeEndpoint()
	if err != nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: err.Error(), Level: "error"})
	}

	authBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: err.Error(), Level: "error"})
	}

	return c.Send(authBody)
}

func TokenRoute(c *fiber.Ctx) error {

	res, err := CallOktaTokenEndpoint(c.Body())
	if err != nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: err.Error(), Level: "error"})
	}

	tokenBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: err.Error(), Level: "error"})
	}

	if res.StatusCode != 200 {
		var tokenRes cmds.TokenErrorResponseBody
		json.Unmarshal(tokenBody, &tokenRes)
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: tokenRes.Error, Level: "error"})
	} else {
		var tokenRes cmds.TokenResponseBody
		json.Unmarshal(tokenBody, &tokenRes)

		c.Status(201)
		return c.JSON(cmds.CommandResponseBody{Message: cmds.TokenResponseBody{
			TokenType:    tokenRes.TokenType,
			ExpiresIn:    tokenRes.ExpiresIn,
			AccessToken:  tokenRes.AccessToken,
			Scope:        tokenRes.Scope,
			RefreshToken: tokenRes.RefreshToken},
			Level: "info"})
	}

}
