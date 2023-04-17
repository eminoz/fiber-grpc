package main

import (
	"net"

	api "gitbub.com/eminoz/graceful-fiber/proto/pb"
	usrapi "gitbub.com/eminoz/graceful-fiber/server/api"
	"gitbub.com/eminoz/graceful-fiber/server/config"
	"gitbub.com/eminoz/graceful-fiber/server/db"
	"gitbub.com/eminoz/graceful-fiber/server/repo"
	"gitbub.com/eminoz/graceful-fiber/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	config.SetupConfig()
	db.SetDatabase()
	s := repo.NewUserController()
	c := service.NewUserService(s)
	u := usrapi.NewUserApi(c)
	api.RegisterUserServiceServer(srv, u)
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}
