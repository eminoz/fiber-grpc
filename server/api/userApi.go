package usrapi

import (
	"context"

	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	"gitbub.com/eminoz/graceful-fiber/server/service"
)

type UserApi struct {
	userController service.UserService
}

func NewUserApi(c service.UserService) UserApi {
	return UserApi{userController: c}
}
func (u UserApi) CreateUser(ctx context.Context, usr *api.User) (*api.ResUser, error) {
	res, err := u.userController.InsertUser(ctx, usr)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u UserApi) GetUser(ctx context.Context, userId *api.UserId) (*api.ResUser, error) {

	res := u.userController.GetUserById(ctx, userId.Id)
	return res, nil
}
func (u UserApi) DeleteUser(ctx context.Context, userId *api.UserId) (*api.DeleteUserRes, error) {
	res, err := u.userController.DeleteUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
