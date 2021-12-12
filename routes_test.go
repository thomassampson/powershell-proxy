package main

import (
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	cmds "powershell-proxy/cmds"
	routes "powershell-proxy/cmds/routes"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type RouteReponseTestCase struct {
	Route              string
	ExpectedStatusCode int
	Description        string
}

func TestRoutes_RouteResponse(t *testing.T) {

	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")

	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")

	cmds.ValidateConfig()

	log.Print(cmds.AppNameDefault)

	app := fiber.New()

	test := RouteReponseTestCase{"/api", 200, "Test Root Route Status Code is 200"}

	app.Get(test.Route, routes.RootRoute)

	req := httptest.NewRequest("GET", test.Route, nil)

	res, _ := app.Test(req, 1)

	body, _ := ioutil.ReadAll(res.Body)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, "✋ App", string(body))
}
