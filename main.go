package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "main/adapters/config"
	router "main/app/routes"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "APP name",
		ServerHeader:  "No header",
	})

	app.Use(logger.New())

	router.SetupRoutes(app)

	if os.Getenv("app_env") == "dev" {
		app.Listen(":3000") //h2 https
	} else {
		app.ListenTLS(":3000", "cert", "key") //https SSL
	}

}
