package service

import (
	"context"
	"fmt"

	"gitbub.com/eminoz/graceful-fiber/client/config"
	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
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
	client api.UserServiceClient
}

func NewUserService() UserService {
	cnn := config.GetConnection()

	return &userService{
		client: api.NewUserServiceClient(cnn),
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
	user2, err := u.client.CreateUser(ctx, user, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return nil, err
	}

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
	user, err := u.client.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u userService) DeleteUserById(id *api.UserId) (*api.DeleteUserRes, error) {
	res, err := u.client.DeleteUser(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u userService) UpdateUserById(user *api.UpdateUser) (*api.ResUser, error) {
	res, err := u.client.UpdateUserById(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
