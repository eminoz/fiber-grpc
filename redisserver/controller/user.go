package controller

import (
	"context"

	redisproto "gitbub.com/eminoz/graceful-fiber/proto/redis"
	"gitbub.com/eminoz/graceful-fiber/redisserver/service"
)

type redisUserApi struct {
	userService service.UserService
}

func NewRedisUserApi(u service.UserService) redisUserApi {

	return redisUserApi{
		userService: u,
	}
}
func (r redisUserApi) InsertUser(ctx context.Context, user *redisproto.User) (*redisproto.InsertedUserRes, error) {
	res, err := r.userService.InserUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r redisUserApi) GetUser(ctx context.Context, id *redisproto.UserId) (*redisproto.ResUser, error) {
	res, err := r.userService.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r redisUserApi) DeleteUser(ctx context.Context, id *redisproto.UserId) (*redisproto.DeleteUserRes, error) {
	res, err := r.userService.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
