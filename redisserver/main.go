package main

import (
	"net"

	redisproto "gitbub.com/eminoz/graceful-fiber/proto/redis"
	"gitbub.com/eminoz/graceful-fiber/redisserver/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	r := controller.NewRedisUserApi()
	redisproto.RegisterUserServiceServer(srv, r)
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}
