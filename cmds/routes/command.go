package cmds

import (
	"encoding/json"
	"strings"

	cmds "powershell-proxy/cmds"

	"github.com/gofiber/fiber/v2"
)

func Command(c *fiber.Ctx) error {

	depth, err := cmds.ConvertDepthString(c.Query("depth"))
	if err != nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: err.Error(), Level: "error"})
	}

	var body cmds.CommandRequestBody
	json.Unmarshal(c.Body(), &body)

	if body.Commands == nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: strings.ToUpper("You must specify the 'commands' property in the request body."),
			Level: "error"})
	}

	if body.Commands[0] == "" {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: strings.ToUpper("'commands' property in the request body cannot be empty."),
			Level: "error"})
	}

	output, err := cmds.ExecuteCommand(body, depth)
	if err != nil {
		c.Status(400)
		return c.JSON(cmds.CommandResponseBody{Message: err.Error(), Level: "error"})
	}

	c.Status(200)
	c.Set("content-type", "application/json")
	return c.Send(output.Bytes())
}
