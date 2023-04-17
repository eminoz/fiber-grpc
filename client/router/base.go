package router

import "github.com/gofiber/fiber/v2"

func UserRoute() *fiber.App {
	f := fiber.New()
	usergroup := f.Group("/user")
	NewUserRouter(usergroup).UserRouters()
	return f
}
