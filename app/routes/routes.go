package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	c_auth "main/app/controller/auth"
	_ "main/docs"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/doc/*", swagger.Handler)
	// 127.0.0.1/auth/register
	auth := app.Group("/auth")
	auth.Post("/register", c_auth.Register)
	auth.Post("/login", c_auth.Login) //handler
	auth.Post("/forget", c_auth.ForgetPassword)
	auth.Post("/verify", c_auth.Verification)

}
