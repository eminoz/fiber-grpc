package router

import (
	"gitbub.com/eminoz/graceful-fiber/client/api"
	"gitbub.com/eminoz/graceful-fiber/client/service"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	Route fiber.Router
}

func NewUserRouter(r fiber.Router) *UserRouter {
	return &UserRouter{
		Route: r,
	}
}

func (u UserRouter) UserRouters() {
	us := service.NewUserService()
	usr := api.NewUserApi(us)
	u.Route.Post("/create", usr.CreateUser)
	u.Route.Get("/getuser/:id", usr.GetUserById)
	u.Route.Delete("/deleteuser/:id", usr.DeleteUserById)
}
