package cmds

import (
	"fmt"
	cmds "powershell-proxy/cmds"

	"github.com/gofiber/fiber/v2"
)

func RootRoute(c *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", cmds.AppName)
	return c.SendString(msg)
}
