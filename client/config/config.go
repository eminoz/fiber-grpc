package config

import (
	"google.golang.org/grpc"
)

var connection *grpc.ClientConn

func SetupConfig() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	connection = conn
}
func GetConnection() *grpc.ClientConn {
	return connection
}
