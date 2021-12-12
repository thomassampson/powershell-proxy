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
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var (
	TEST_COMMAND_REQUEST_SUCCESS_STDOUT = []struct {
		Name string `json:"name"`
	}{struct {
		Name string "json:\"name\""
	}{Name: "main.go"}, struct {
		Name string "json:\"name\""
	}{Name: "go.mod"}}
	TEST_COMMAND_REQUEST_FAIL_STDERR = "exit status 1"
)

func TestExecRun_CommandRoute_Success(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	data, err := json.Marshal(TEST_COMMAND_REQUEST_SUCCESS_STDOUT)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, string(data))
	os.Exit(0)
}

func Mock_ExecCommand_CommandRoute_Success(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecRun_CommandRoute_Success", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestExecCommand_CommandRoute_Success(t *testing.T) {
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")

	cmds.Shell = "pwsh"
	defer func() { cmds.Shell = "" }()
	cmds.ExecCommand = Mock_ExecCommand_CommandRoute_Success
	defer func() { cmds.ExecCommand = exec.Command }()

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 200, "Test Command Route Status Code is 200"}

	app.Post(test.Route, routes.Command)

	reqBody, _ := json.Marshal(cmds.CommandRequestBody{Commands: []string{"Get-ChildItem"}})
	req := httptest.NewRequest("POST", test.Route, bytes.NewBuffer(reqBody))

	res, err := app.Test(req, -1)
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, err)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, err)

	expected, _ := json.Marshal(TEST_COMMAND_REQUEST_SUCCESS_STDOUT)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, string(expected), string(resBody))

}

func TestExecRun_CommandRoute_Fail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	fmt.Fprint(os.Stderr, TEST_COMMAND_REQUEST_FAIL_STDERR)
	os.Exit(0)
}

func Mock_ExecCommand_CommandRoute_Fail(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecRun_CommandRoute_Fail", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestExecCommand_CommandRoute_Fail(t *testing.T) {
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")

	cmds.Shell = "pwsh"
	defer func() { cmds.Shell = "" }()
	cmds.ExecCommand = Mock_ExecCommand_CommandRoute_Fail
	defer func() { cmds.ExecCommand = exec.Command }()

	app := fiber.New()

	test := RouteReponseTestCase{"/api/command", 400, "Test Command Route Status Code is 400 due to error"}

	app.Post(test.Route, routes.Command)

	reqBody, _ := json.Marshal(cmds.CommandRequestBody{Commands: []string{"Get-FakeCommand"}})
	req := httptest.NewRequest("POST", test.Route, bytes.NewBuffer(reqBody))

	res, err := app.Test(req, -1)
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, err)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}
	assert.Nil(t, err)

	var resObj cmds.CommandResponseBody
	err = json.Unmarshal(resBody, &resObj)
	assert.Nil(t, err)

	assert.Equalf(t, test.ExpectedStatusCode, res.StatusCode, test.Description)
	assert.Equal(t, cmds.CommandResponseBody{Message: TEST_COMMAND_REQUEST_FAIL_STDERR, Level: "error"}, resObj)

}
