package api

import (
	"gitbub.com/eminoz/graceful-fiber/client/service"
	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
}
type userApi struct {
	userService service.UserService
}

func NewUserApi(us service.UserService) UserApi {

	return &userApi{userService: us}
}
func (u userApi) CreateUser(ctx *fiber.Ctx) error {
	var usr api.User
	ctx.BodyParser(&usr)
	a := u.userService.CreateUser(&usr)
	return ctx.JSON(a)
}
func (u userApi) GetUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	userRes := u.userService.GetUserById(&api.UserId{Id: userID})
	return ctx.JSON(userRes)
}
