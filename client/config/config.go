package config

import (
	"google.golang.org/grpc"
)

var redisConnection *grpc.ClientConn
var serverConnection *grpc.ClientConn

func SetupConfig() {
	sconn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	rconn, err := grpc.Dial("localhost:4043", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	redisConnection = rconn
	serverConnection = sconn
}
func GetConnection() (*grpc.ClientConn, *grpc.ClientConn) {
	return serverConnection, redisConnection
}
