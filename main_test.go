package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIPAddressNotValid(t *testing.T) {

	ip := "0.0.0"

	assert.False(t, checkIPAddress(ip))

}

func TestCheckIPAddressValid(t *testing.T) {

	ip := "0.0.0.0"

	assert.True(t, checkIPAddress(ip))

}

func TestConvertDepthStringValid(t *testing.T) {

	depth := "5"

	expected := 5

	actual, _ := convertDepthString(depth)

	assert.Equal(t, expected, actual)

}

func TestConvertDepthStringNotValid_NotInt(t *testing.T) {

	depth := "eeewrwe"

	expected := -1

	actual, _ := convertDepthString(depth)

	assert.Equal(t, expected, actual)

}

func TestConvertDepthStringNotValid_ToBig(t *testing.T) {

	depth := "7"

	expected := 4

	actual, _ := convertDepthString(depth)

	assert.Equal(t, expected, actual)
}

func TestConvertDepthStringNotValid_ToSmall(t *testing.T) {

	depth := "0"

	expected := 4

	actual, _ := convertDepthString(depth)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "Powershell Proxy API", AppNameDefault)
	assert.Nil(t, validateConfig())
	assert.Equal(t, AppNameDefault, AppName)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "Powershell Proxy API", AppNameDefault)
	assert.Nil(t, validateConfig())
	assert.Equal(t, "App", AppName)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "0.0.0.0", ListenAddressDefault)
	assert.Nil(t, validateConfig())
	assert.Equal(t, ListenAddressDefault, ListenAddress)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, ip, os.Getenv("PWSHPRXY_LISTEN_ADDR"))
	assert.Nil(t, validateConfig())
	assert.Equal(t, ip, ListenAddress)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, port, os.Getenv("PWSHPRXY_LISTEN_PORT"))
	assert.Nil(t, validateConfig())
	assert.Equal(t, port, ListenPort)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "8000", ListenPortDefault)
	assert.Nil(t, validateConfig())
	assert.Equal(t, ListenPortDefault, ListenPort)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Nil(t, validateConfig())
	assert.Equal(t, "1234", OktaAudience)
	assert.Equal(t, "1234", OktaClientId)
	assert.Equal(t, "1234", OktaIssuer)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Nil(t, validateConfig())
	assert.Equal(t, "0.0.0.0", ListenAddress)
	assert.Equal(t, "App", AppName)
	assert.Equal(t, "8000", ListenPort)
	assert.Equal(t, "core", PowerShellType)
	assert.Equal(t, "1234", OktaAudience)
	assert.Equal(t, "1234", OktaClientId)
	assert.Equal(t, "1234", OktaIssuer)
	assert.Equal(t, "pwsh", Shell)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "core", PowerShellType)
	assert.Nil(t, validateConfig())
	assert.Equal(t, "pwsh", Shell)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	assert.Equal(t, "windows", PowerShellType)
	assert.Nil(t, validateConfig())
	assert.Equal(t, "powershell", Shell)

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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
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

	ListenAddress = os.Getenv("PWSHPRXY_LISTEN_ADDR")
	ListenPort = os.Getenv("PWSHPRXY_LISTEN_PORT")
	AppName = os.Getenv("PWSHPRXY_APP_NAME")
	PowerShellType = os.Getenv("PWSHPRXY_TYPE")
	OktaIssuer = os.Getenv("PWSHPRXY_OKTA_ISSUER")
	OktaClientId = os.Getenv("PWSHPRXY_OKTA_CLIENT_ID")
	OktaAudience = os.Getenv("PWSHPRXY_OKTA_AUDIENCE")

	actual := validateConfig()
	assert.Error(t, actual)
	assert.EqualError(t, actual, "FATAL: Env Variable 'PWSHPRXY_TYPE' must be set to either 'core' or 'powershell'")
}
