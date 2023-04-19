package service

import (
	"context"
	"fmt"

	"gitbub.com/eminoz/graceful-fiber/client/config"
	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	redisproto "gitbub.com/eminoz/graceful-fiber/proto/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type UserService interface {
	CreateUser(user *api.User) (*api.ResUser, error)
	GetUserById(id *api.UserId) (*api.ResUser, error)
	DeleteUserById(id *api.UserId) (*api.DeleteUserRes, error)
	UpdateUserById(user *api.UpdateUser) (*api.ResUser, error)
}

type userService struct {
	server api.UserServiceClient
	redis  redisproto.UserServiceClient
}

func NewUserService() UserService {
	s, r := config.GetConnection()

	return &userService{
		server: api.NewUserServiceClient(s),
		redis:  redisproto.NewUserServiceClient(r),
	}
}
func (u userService) CreateUser(user *api.User) (*api.ResUser, error) {

	ctx := context.Background()

	// md := metadata.New(map[string]string{
	// 	"my-key": "emin",
	// })

	// // Set the metadata to send with the request
	// ctx = metadata.NewOutgoingContext(ctx, md)

	var header, trailer metadata.MD // variable to store header and trailer
	user2, err := u.server.CreateUser(ctx, user, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}
	newUser := redisproto.User{Name: user2.Name, Surname: user2.Surname, Id: user2.Id}
	u.redis.InsertUser(ctx, &newUser)
	a, ok := header["msg"]

	if ok {

		for _, v := range a {
			fmt.Print(v)

		}
	}
	// if header.Get("msg") != nil {

	// 	fmt.Print(header.Get("msg")[0]) // Read metadata sent by server
	// }

	return user2, nil
}
func (u userService) GetUserById(id *api.UserId) (*api.ResUser, error) {
	ctx := context.Background()
	user, err := u.server.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u userService) DeleteUserById(id *api.UserId) (*api.DeleteUserRes, error) {
	res, err := u.server.DeleteUser(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u userService) UpdateUserById(user *api.UpdateUser) (*api.ResUser, error) {
	res, err := u.server.UpdateUserById(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
