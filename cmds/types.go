package cmds

import (
	"io"
	"net/http"
)

type Http interface {
	Post(url string, contentType string, body io.Reader) (res *http.Response, err error)
}

type CommandResponseBody struct {
	Message interface{} `json:"message"`
	Level   string      `json:"level"`
}

type CommandRequestBody struct {
	Commands []string `json:"commands"`
}

type TokenRequestBody struct {
	DeviceCode string `json:"device_code"`
}

type TokenResponseBody struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type TokenErrorResponseBody struct {
	Error string `json:"error"`
}
