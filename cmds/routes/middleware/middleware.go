package cmds

import (
	"log"
	cmds "powershell-proxy/cmds"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SecureRoute(c *fiber.Ctx) (err error) {
	jwt := strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)
	token, err := cmds.ValidateJwt(jwt)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return c.SendStatus(401)
	}
	cmds.Auth = token.Claims
	log.Printf("INFO: Token found for user: %s", cmds.Auth["sub"])
	return c.Next()

}
