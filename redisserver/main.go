package main

import (
	"net"

	redisproto "gitbub.com/eminoz/graceful-fiber/proto/redis"
	"gitbub.com/eminoz/graceful-fiber/redisserver/controller"
	"gitbub.com/eminoz/graceful-fiber/redisserver/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":4043")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	rs := service.NewUserService()
	r := controller.NewRedisUserApi(rs)
	redisproto.RegisterUserServiceServer(srv, r)
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}
