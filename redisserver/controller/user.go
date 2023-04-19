package controller

import (
	"context"

	redisproto "gitbub.com/eminoz/graceful-fiber/proto/redis"
)

type redisUserApi struct {
}

func NewRedisUserApi() redisUserApi {

	return redisUserApi{}
}
func (r redisUserApi) InsertUser(context.Context, *redisproto.User) (*redisproto.ResUser, error) {
	return &redisproto.ResUser{}, nil
}
func (r redisUserApi) GetUser(context.Context, *redisproto.UserId) (*redisproto.ResUser, error) {
	return &redisproto.ResUser{}, nil
}
func (r redisUserApi) DeleteUser(context.Context, *redisproto.UserId) (*redisproto.DeleteUserRes, error) {
	return &redisproto.DeleteUserRes{}, nil
}
