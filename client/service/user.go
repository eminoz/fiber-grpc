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
	CreateUser(user *api.User) *api.ResUser
	GetUserById(id *api.UserId) *api.ResUser
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
func (u userService) CreateUser(user *api.User) *api.ResUser {

	ctx := context.Background()

	// md := metadata.New(map[string]string{
	// 	"my-key": "emin",
	// })

	// // Set the metadata to send with the request
	// ctx = metadata.NewOutgoingContext(ctx, md)

	var header, trailer metadata.MD // variable to store header and trailer
	user2, _ := u.client.CreateUser(ctx, user, grpc.Header(&header), grpc.Trailer(&trailer))
	a, ok := header["msg"]

	if ok {

		for i, v := range a {
			fmt.Print(v)
			fmt.Print(i)

		}
	}
	// if header.Get("msg") != nil {

	// 	fmt.Print(header.Get("msg")[0]) // Read metadata sent by server
	// }

	return user2
}
func (u userService) GetUserById(id *api.UserId) *api.ResUser {
	ctx := context.Background()
	fmt.Print(id)
	user, _ := u.client.GetUser(ctx, id)
	return user
}
