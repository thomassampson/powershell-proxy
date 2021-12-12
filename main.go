package main

import (
	"fmt"
	"log"
	"os/exec"

	cmds "powershell-proxy/cmds"
	routes "powershell-proxy/cmds/routes"
	middleware "powershell-proxy/cmds/routes/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	log.Printf("🔵 Starting %s", cmds.AppNameDefault)

	cmds.ExecCommand = exec.Command

	err := cmds.ValidateConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// GET /api
	app.Get("/api", routes.RootRoute)

	// GET /api/auth/authorize
	app.Get("/api/auth/authorize", routes.AuthorizeRoute)

	// POST /api/auth/token
	app.Post("/api/auth/token", routes.TokenRoute)

	// POST /api/command
	app.Post("/api/command", middleware.SecureRoute, routes.Command)

	log.Printf("🟢 Started %s", cmds.AppName)
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", cmds.ListenAddress, cmds.ListenPort)))

}
