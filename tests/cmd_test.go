package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	cmds "powershell-proxy/cmds"

	"github.com/stretchr/testify/assert"
)

var (
	TEST_COMMAND_SUCCESS_STDOUT string = "This is the test standard out from the mock function"
	TEST_COMMAND_FAIL_STDERR    string = "This is the test standard error from the mock function"
)

func TestCmds_CheckIPAddressNotValid(t *testing.T) {

	ip := "0.0.0"

	assert.False(t, cmds.CheckIPAddress(ip))

}

func TestCmds_CheckIPAddressValid(t *testing.T) {

	ip := "0.0.0.0"

	assert.True(t, cmds.CheckIPAddress(ip))

}

func TestCmds_ConvertDepthStringValid(t *testing.T) {

	depth := "5"

	expected := 5

	actual, _ := cmds.ConvertDepthString(depth)

	assert.Equal(t, expected, actual)

}

func TestCmds_ConvertDepthStringNotValid_NotInt(t *testing.T) {

	depth := "eeewrwe"

	expected := -1

	actual, _ := cmds.ConvertDepthString(depth)

	assert.Equal(t, expected, actual)

}

func TestCmds_ConvertDepthStringNotValid_ToBig(t *testing.T) {

	depth := "7"

	expected := 4

	actual, _ := cmds.ConvertDepthString(depth)

	assert.Equal(t, expected, actual)
}

func TestCmds_ConvertDepthStringNotValid_ToSmall(t *testing.T) {

	depth := "0"

	expected := 4

	actual, _ := cmds.ConvertDepthString(depth)

	assert.Equal(t, expected, actual)

}

func TestValidateConfigs_EnvVarNotSet(t *testing.T) {

	os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Unsetenv("PWSHPRXY_TYPE")
	os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_OKTA_CLIENT_ID' not set")

}

func TestValidateConfigs_EnvAppNameNotSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "Powershell Proxy API", cmds.AppNameDefault)
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, cmds.AppNameDefault, cmds.AppName)

}

func TestValidateConfigs_EnvAppNameSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "Powershell Proxy API", cmds.AppNameDefault)
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, "App", cmds.AppName)

}

func TestValidateConfigs_EnvListenAddrNotSet(t *testing.T) {

	os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "0.0.0.0", cmds.ListenAddressDefault)
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, cmds.ListenAddressDefault, cmds.ListenAddress)

}

func TestValidateConfigs_EnvListenAddrNotValid(t *testing.T) {
	ip := "0.0.0"
	os.Setenv("PWSHPRXY_LISTEN_ADDR", ip)
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_LISTEN_ADDR' is set, but '0.0.0' is not a valid ipv4 address")

}

func TestValidateConfigs_EnvListenAddrSetValid(t *testing.T) {

	ip := "1.1.1.1"

	os.Setenv("PWSHPRXY_LISTEN_ADDR", ip)
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, ip, os.Getenv("PWSHPRXY_LISTEN_ADDR"))
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, ip, cmds.ListenAddress)

}

func TestValidateConfigs_EnvListenPortSetValid(t *testing.T) {

	port := "3000"

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", port)
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, port, os.Getenv("PWSHPRXY_LISTEN_PORT"))
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, port, cmds.ListenPort)

}

func TestValidateConfigs_EnvListenPortNotValid(t *testing.T) {
	port := "notaport"
	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", port)
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_LISTEN_PORT' is set, but 'notaport' is not a number")
}

func TestValidateConfigs_NoEnvListenPort(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "8000", cmds.ListenPortDefault)
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, cmds.ListenPortDefault, cmds.ListenPort)

}

func TestValidateConfigs_EnvOktaClientIdNotSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_OKTA_CLIENT_ID' not set")

}

func TestValidateConfigs_EnvOktaIssuerNotSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_OKTA_ISSUER' not set")

}

func TestValidateConfigs_EnvOktaAudienceNotSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_OKTA_AUDIENCE' not set")

}

func TestValidateConfigs_EnvOktaNotSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_OKTA_CLIENT_ID' not set")
}

func TestValidateConfigs_EnvOktaSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, "1234", cmds.OktaAudience)
	assert.Equal(t, "1234", cmds.OktaClientId)
	assert.Equal(t, "1234", cmds.OktaIssuer)

}

func TestValidateConfigs_EnvAllSetValid(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, "0.0.0.0", cmds.ListenAddress)
	assert.Equal(t, "App", cmds.AppName)
	assert.Equal(t, "8000", cmds.ListenPort)
	assert.Equal(t, "core", cmds.PowerShellType)
	assert.Equal(t, "1234", cmds.OktaAudience)
	assert.Equal(t, "1234", cmds.OktaClientId)
	assert.Equal(t, "1234", cmds.OktaIssuer)
	assert.Equal(t, "pwsh", cmds.Shell)

}

func TestValidateConfigs_EnvTypeSetToCore(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "core")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "core", cmds.PowerShellType)
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, "pwsh", cmds.Shell)

}

func TestValidateConfigs_EnvTypeSetToWindows(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Setenv("PWSHPRXY_TYPE", "windows")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "windows", cmds.PowerShellType)
	assert.Nil(t, cmds.ValidateConfig())
	assert.Equal(t, "powershell", cmds.Shell)

}

func TestValidateConfigs_EnvTypeNotSet(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_TYPE' not set, cannot start webserver")

}

func TestValidateConfigs_EnvTypeSetNotValid(t *testing.T) {

	os.Setenv("PWSHPRXY_LISTEN_ADDR", "0.0.0.0")
	defer os.Unsetenv("PWSHPRXY_LISTEN_ADDR")
	os.Setenv("PWSHPRXY_LISTEN_PORT", "8000")
	defer os.Unsetenv("PWSHPRXY_LISTEN_PORT")
	os.Setenv("PWSHPRXY_APP_NAME", "App")
	defer os.Unsetenv("PWSHPRXY_APP_NAME")
	// Set to notvalid
	os.Setenv("PWSHPRXY_TYPE", "notvalid")
	defer os.Unsetenv("PWSHPRXY_TYPE")
	os.Setenv("PWSHPRXY_OKTA_CLIENT_ID", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_CLIENT_ID")
	os.Setenv("PWSHPRXY_OKTA_ISSUER", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_ISSUER")
	os.Setenv("PWSHPRXY_OKTA_AUDIENCE", "1234")
	defer os.Unsetenv("PWSHPRXY_OKTA_AUDIENCE")

	cmds.ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	cmds.ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	cmds.AppName = os.Getenv("PWSHPRXY_APP_NAME")
	cmds.PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	cmds.OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	cmds.OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	cmds.OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := cmds.ValidateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_TYPE' must be set to either 'core' or 'powershell'")
}

func TestExecRun_Success(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprint(os.Stdout, TEST_COMMAND_SUCCESS_STDOUT)
	os.Exit(0)
}

func Mock_ExecCommand_Success(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecRun_Success", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestExecCommand_Success(t *testing.T) {
	cmds.ExecCommand = Mock_ExecCommand_Success
	defer func() { cmds.ExecCommand = exec.Command }()
	output, err := cmds.ExecuteCommand(cmds.CommandRequestBody{Commands: []string{"test"}}, 1)
	assert.Nil(t, err)
	assert.Equal(t, output.String(), TEST_COMMAND_SUCCESS_STDOUT)
}

func TestExecRun_Fail(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprint(os.Stderr, TEST_COMMAND_FAIL_STDERR)
	os.Exit(0)
}

func Mock_ExecCommand_Fail(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecRun_Fail", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}

func TestExecCommand_Fail(t *testing.T) {
	cmds.ExecCommand = Mock_ExecCommand_Fail
	defer func() { cmds.ExecCommand = exec.Command }()
	output, err := cmds.ExecuteCommand(cmds.CommandRequestBody{Commands: []string{"test"}}, 1)
	assert.Equal(t, output.Len(), 0)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), TEST_COMMAND_FAIL_STDERR)
}
