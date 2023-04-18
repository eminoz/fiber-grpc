package api

import (
	"gitbub.com/eminoz/graceful-fiber/client/service"
	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUserById(ctx *fiber.Ctx) error
	DeleteUserById(ctx *fiber.Ctx) error
	UpdateUserById(ctx *fiber.Ctx) error
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
	a, err := u.userService.CreateUser(&usr)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(a)
}
func (u userApi) GetUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	userRes, err := u.userService.GetUserById(&api.UserId{Id: userID})
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(userRes)
}

func (u userApi) DeleteUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	res, err := u.userService.DeleteUserById(&api.UserId{Id: userID})
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(res)
}
func (u userApi) UpdateUserById(ctx *fiber.Ctx) error {
	var user *api.UpdateUser

	ctx.BodyParser(&user)
	res, err := u.userService.UpdateUserById(user)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(res)
}
