package service

import (
	"context"
	"encoding/json"

	redisproto "gitbub.com/eminoz/graceful-fiber/proto/redis"
	"github.com/go-redis/redis/v8"
)

type UserService interface {
	InserUser(ctx context.Context, user *redisproto.User) (redisproto.InsertedUserRes, error)
	GetUserById(ctx context.Context, userId *redisproto.UserId) (redisproto.ResUser, error)
	DeleteUser(ctx context.Context, userId *redisproto.UserId) (redisproto.DeleteUserRes, error)
}

type userService struct {
	Redis *redis.Client
}

func NewUserService() UserService {
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use the default database
	})
	return &userService{
		Redis: r,
	}
}
func (u userService) InserUser(ctx context.Context, user *redisproto.User) (redisproto.InsertedUserRes, error) {
	jsonuser, err := json.Marshal(user)
	if err != nil {
		return redisproto.InsertedUserRes{}, err
	}
	res := u.Redis.HSet(ctx, "redisproto", user.Id, jsonuser)
	if res.Val() != 1 {
		return redisproto.InsertedUserRes{Msg: "could not insert", IsInsert: false}, err
	}
	return redisproto.InsertedUserRes{Msg: "inserted", IsInsert: true}, nil

}

func (u userService) GetUserById(ctx context.Context, userId *redisproto.UserId) (redisproto.ResUser, error) {
	panic("not implemented") // TODO: Implement
}

func (u userService) DeleteUser(ctx context.Context, userId *redisproto.UserId) (redisproto.DeleteUserRes, error) {
	panic("not implemented") // TODO: Implement
}
