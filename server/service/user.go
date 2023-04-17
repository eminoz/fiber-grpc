package service

import (
	"context"

	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	"gitbub.com/eminoz/graceful-fiber/server/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type UserService interface {
	InsertUser(ctx context.Context, user *api.User) (*api.ResUser, error)
	GetUserById(ctx context.Context, id string) *api.ResUser
	DeleteUser(context.Context, *api.UserId) (*api.DeleteUserRes, error)
}
type userService struct {
	userservice repo.UserRepo
}

func NewUserService(s repo.UserRepo) UserService {
	return &userService{
		userservice: s,
	}
}
func (u userService) InsertUser(ctx context.Context, user *api.User) (*api.ResUser, error) {
	// Read metadata from the incoming context
	// incomingMD, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	panic("No metadata was sent with the request")
	// }
	// // Access a specific value in the metadata
	// myValue := incomingMD.Get("my-key")
	// if myValue != nil {
	// 	a := myValue[0]
	// 	fmt.Print(a)
	// }
	if user.GetPassword() == "" {
		// Create metadata to send with the response
		md := metadata.New(map[string]string{"msg": "user password emty"})

		// Attach metadata to response
		grpc.SetHeader(ctx, md)
		return &api.ResUser{}, nil
	}
	insertone, err := u.userservice.InsertUser(ctx, user)
	if err != nil {
		return &api.ResUser{}, nil
	}

	return insertone, nil

}
func (u userService) GetUserById(ctx context.Context, id string) *api.ResUser {
	res := u.userservice.GetUserById(id)
	return res
}
func (u userService) DeleteUser(ctx context.Context, userID *api.UserId) (*api.DeleteUserRes, error) {
	res, err := u.userservice.DeleteUserById(userID.Id)

	if err != nil {
		return &api.DeleteUserRes{Msg: "user could not delete", IsDeleted: res}, nil
	}
	return &api.DeleteUserRes{Msg: "user was deleted", IsDeleted: res}, nil
}
