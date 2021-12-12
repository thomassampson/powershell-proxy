package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"os/exec"
	cmds "powershell-proxy/cmds"
	routes "powershell-proxy/cmds/routes"
	middleware "powershell-proxy/cmds/routes/middleware"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	TEST_DEVICE_CODE             string = "cd302bd4-86a5-4d9c-b6e4-0446f4793a9a"
	TEST_AUTHORIZE_RESPONSE_BODY string = fmt.Sprintf(`{\r\n  \"device_code\": \"%s\",\r\n  
	\"user_code\": \"VNKBHRSV\",\r\n  \"verification_uri\": \"https:\/\/tenant.okta.com\/activate\",\r\n  
	\"verification_uri_complete\": \"https:\/\/tenant.okta.com\/activate?user_code=VNKBHRSV\",\r\n  
	\"expires_in\": 600,\r\n  \"interval\": 5\r\n}`, TEST_DEVICE_CODE)

	TEST_OKTA_TOKEN_RESPONSE_BODY cmds.TokenResponseBody = cmds.TokenResponseBody{
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		Scope:        "offline_access profile",
		AccessToken:  "jwt",
		RefreshToken: "refresh_token"}
	TEST_TOKEN_REQUEST_BODY cmds.TokenRequestBody = cmds.TokenRequestBody{
		DeviceCode: TEST_DEVICE_CODE}
	TEST_TOKEN_RESPONSE_BODY cmds.CommandResponseBody = cmds.CommandResponseBody{
		Level:   "info",
		Message: TEST_OKTA_TOKEN_RESPONSE_BODY}
	TEST_METHODS                                []string                = []string{"GET", "POST", "DELETE", "PUT", "PATCH"}
	TEST_OKTATOKEN_NODEVICECODE_RESPONSE_BODY   string                  = "{\"error\": \"invalid_grant\",\"error_description\": \"The device code is invalid or has expired.\"}"
	TEST_TOKEN_NODEVICECODE_RESPONSE_BODY       string                  = "{\"message\":\"invalid_grant\",\"level\":\"error\"}"
	TEST_COMMAND_BAD_REQUEST_BODY_NO_COMMAND    string                  = "{}"
	TEST_COMMAND_BAD_REQUEST_BODY_EMPTY_COMMAND cmds.CommandRequestBody = cmds.CommandRequestBody{Commands: []string{""}}
	TEST_COMMAND_REQUEST_BODY_SUCCESS           cmds.CommandRequestBody = cmds.CommandRequestBody{Commands: []string{"Get-ChildItem | Select-Object Name"}}
)

type RouteReponseTestCase struct {
	Route              string
	ExpectedStatusCode int
	Description        string
}

func TestRoutes_RouteResponse_Success(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api", 200, "Test Root Route Status Code is 200"}

	app.Get(test.Route, routes.RootRoute)

	req := httptest.NewRequest("GET", test.Route, nil)

	res, _ := app.Test(req, 1)

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, "✋ App", string(body))
}

func TestRoutes_RouteResponse_MethodNotAllowed(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api", 405, "Test Root Route Status Code is 405 when not a GET request is made"}

	app.Get(test.Route, routes.RootRoute)

	allowedMethod := "GET"

	for _, method := range TEST_METHODS {
		if allowedMethod != method {
			log.Print(fmt.Sprintf("[%s] %s", method, test.Description))
			req := httptest.NewRequest(method, test.Route, nil)

			res, _ := app.Test(req, 1)

			assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, fmt.Sprintf("[%s] %s", method, test.Description))
		}
	}

}

func TestRoutes_AuthorizeResponse_MethodNotAllowed(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/auth/authorize", 405, "Test Authorize Route Status Code is 405 when not a GET request is made"}

	app.Get(test.Route, routes.RootRoute)

	allowedMethod := "GET"

	for _, method := range TEST_METHODS {
		if allowedMethod != method {
			log.Print(fmt.Sprintf("[%s] %s", method, test.Description))
			req := httptest.NewRequest(method, test.Route, nil)

			res, _ := app.Test(req, 1)

			assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, fmt.Sprintf("[%s] %s", method, test.Description))
		}
	}

}

func TestRoutes_TokeneResponse_MethodNotAllowed(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/auth/token", 405, "Test Token Route Status Code is 405 when not a GET request is made"}

	app.Post(test.Route, routes.RootRoute)

	allowedMethod := "POST"

	for _, method := range TEST_METHODS {
		if allowedMethod != method {
			log.Print(fmt.Sprintf("[%s] %s", method, test.Description))
			req := httptest.NewRequest(method, test.Route, nil)

			res, _ := app.Test(req, 1)

			assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, fmt.Sprintf("[%s] %s", method, test.Description))
		}
	}

}

func TestRoutes_CommandResponse_MethodNotAllowed(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 405, "Test Command Route Status Code is 405 when not a GET request is made"}

	app.Post(test.Route, routes.Command)

	allowedMethod := "POST"

	for _, method := range TEST_METHODS {
		if allowedMethod != method {
			log.Print(fmt.Sprintf("[%s] %s", method, test.Description))
			req := httptest.NewRequest(method, test.Route, nil)

			res, _ := app.Test(req, 1)

			assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, fmt.Sprintf("[%s] %s", method, test.Description))
		}
	}

}

func TestRoutes_CommandResponse_NoAuth_NotAuthorized(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 401, "Test Command Route Status Code is 401 no auth header is passed in"}

	app.Post(test.Route, middleware.SecureRoute, routes.Command)

	req := httptest.NewRequest("POST", test.Route, nil)

	res, _ := app.Test(req, 1)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)

}

func TestRoutes_CommandResponse_BadToken_Unauthorized(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 401, "Test Command Route Status Code is 401 auth header is not valid"}

	app.Post(test.Route, middleware.SecureRoute, routes.Command)

	req := httptest.NewRequest("POST", test.Route, nil)

	req.Header.Add("Auhorization", fmt.Sprintf("%s %s", TEST_OKTA_TOKEN_RESPONSE_BODY.TokenType, TEST_OKTA_TOKEN_RESPONSE_BODY.AccessToken))

	res, _ := app.Test(req, 1)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)

}

func TestRoutes_CommandResponse_NoRequestBody(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 400, "Test Command Route Status Code is 400 no request body is passed in"}

	app.Post(test.Route, routes.Command)

	req := httptest.NewRequest("POST", test.Route, nil)

	res, _ := app.Test(req, 1)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)

}

func TestRoutes_CommandResponse_NoCommandInBody(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 400, "Test Command Route Status Code is 400 no request body is passed in"}

	app.Post(test.Route, routes.Command)

	req := httptest.NewRequest("POST", test.Route, bytes.NewBuffer([]byte(TEST_COMMAND_BAD_REQUEST_BODY_NO_COMMAND)))

	res, _ := app.Test(req, 1)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)

}

func TestRoutes_CommandResponse_EmptyCommandInBody(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 400, "Test Command Route Status Code is 400 no request body is passed in"}

	app.Post(test.Route, routes.Command)

	reqBody, _ := json.Marshal(TEST_COMMAND_BAD_REQUEST_BODY_EMPTY_COMMAND)
	req := httptest.NewRequest("POST", test.Route, bytes.NewBuffer(reqBody))

	res, _ := app.Test(req, 1)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)

}

func TestExecRun_CommandResponse_Success(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	_, err := json.Marshal(TEST_COMMAND_REQUEST_SUCCESS_STDOUT)
	if err != nil {
		log.Print(err.Error())
	}

	fmt.Fprint(os.Stdout, "Test")
	os.Exit(0)
}

func Mock_ExecCommand_CommandResponse_Success(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecRun_CommandResponse_Success", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}
func TestRoutes_CallOktaAuthorizeEndpoint_Success(t *testing.T) {

	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "https://tenant.okta.com/oauth2/default")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "api://default")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	err := cmds.ValidateConfig()
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/v1/device/authorize", cmds.OktaIssuer),
		httpmock.NewStringResponder(200,
			fmt.Sprintf(`{\r\n  \"device_code\": \"%s\",\r\n  
			\"user_code\": \"VNKBHRSV\",\r\n  \"verification_uri\": \"https:\/\/tenant.okta.com\/activate\",\r\n  
			\"verification_uri_complete\": \"https:\/\/tenant.okta.com\/activate?user_code=VNKBHRSV\",\r\n  
			\"expires_in\": 600,\r\n  \"interval\": 5\r\n}`, TEST_DEVICE_CODE)))

	res, err := routes.CallOktaAuthorizeEndpoint()
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestRoutes_CallOktaTokenEndpoint_Success(t *testing.T) {

	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "https://tenant.okta.com/oauth2/default")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "api://default")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")
	resBody, _ := json.Marshal(TEST_TOKEN_RESPONSE_BODY)
	err := cmds.ValidateConfig()
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/v1/token", cmds.OktaIssuer),
		httpmock.NewBytesResponder(200,
			resBody))

	res, err := routes.CallOktaTokenEndpoint([]byte(TEST_DEVICE_CODE))
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestRoutes_AuthorizeRoute_Success(t *testing.T) {

	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "https://tenant.okta.com/oauth2/default")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "api://default")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/v1/device/authorize", cmds.OktaIssuer),
		httpmock.NewStringResponder(200, TEST_AUTHORIZE_RESPONSE_BODY))

	err := cmds.ValidateConfig()
	assert.Nil(t, err)

	app := fiber.New()

	test := RouteReponseTestCase{"/api/auth/authorize", 200, "Test Authorize Route Status Code is 200 and Body is Correct"}

	app.Get(test.Route, routes.AuthorizeRoute)

	req := httptest.NewRequest("GET", test.Route, nil)

	res, _ := app.Test(req, 1)

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, TEST_AUTHORIZE_RESPONSE_BODY, string(body))
}

func TestRoute_TokenRoute_Success(t *testing.T) {
	test := RouteReponseTestCase{"/api/auth/token", 201, "Test Token Route Status Code is 201 and Body is Correct"}

	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "https://tenant.okta.com/oauth2/default")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "api://default")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	err := cmds.ValidateConfig()
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	oktaReqBody, _ := json.Marshal(TEST_OKTA_TOKEN_RESPONSE_BODY)
	jsonBody := string(oktaReqBody)
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/v1/token", cmds.OktaIssuer),
		httpmock.NewStringResponder(200, jsonBody))

	app := fiber.New()

	app.Get("/api/auth/token", routes.TokenRoute)

	reqBody, _ := json.Marshal(TEST_TOKEN_REQUEST_BODY)
	req := httptest.NewRequest("GET", "/api/auth/token", bytes.NewBuffer(reqBody))

	res, err := app.Test(req, -1)
	assert.Nil(t, err)

	actualResBody, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	expectedResBody, err := json.Marshal(TEST_TOKEN_RESPONSE_BODY)
	assert.Nil(t, err)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, string(expectedResBody), string(actualResBody))

}

func TestRoute_TokenRoute_NoDeviceCode(t *testing.T) {
	test := RouteReponseTestCase{"/api/auth/token", 400, "Test Token Route Status Code is 400 when body is missing device code"}

	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "https://tenant.okta.com/oauth2/default")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "api://default")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	err := cmds.ValidateConfig()
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/v1/token", cmds.OktaIssuer),
		httpmock.NewStringResponder(400, TEST_OKTATOKEN_NODEVICECODE_RESPONSE_BODY))

	app := fiber.New()

	app.Get("/api/auth/token", routes.TokenRoute)

	req := httptest.NewRequest("GET", "/api/auth/token", nil)

	res, err := app.Test(req, -1)
	assert.Nil(t, err)

	actualResBody, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, string(TEST_TOKEN_NODEVICECODE_RESPONSE_BODY), string(actualResBody))

}
