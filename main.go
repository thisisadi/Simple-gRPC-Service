package main

import (
	"log"
	"net"

	"github.com/thisisadi/Total-Corp-Task/user"
	"github.com/thisisadi/Total-Corp-Task/utils"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("unable to listen : %s", err)
	}
	println("server running on port 3000")
	serverRegistrar := grpc.NewServer()
	service := &utils.UserServiceServer{}

	user.RegisterUserServiceServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("server stopped : %s", err)
	}
}
