package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	userpb "github.com/inonsdn/gacha-system/proto/user"
	"github.com/inonsdn/gacha-system/user_service/internal"
)

func run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("Got error while running", err)
	}

	server := grpc.NewServer()

	userpb.RegisterUserServiceServer(server, internal.UserService{})
	fmt.Println("Run User Service at 50051")
	err = server.Serve(lis)

	if err != nil {
		fmt.Println("Found Error when running service", err)
	}

}

func main() {
	run()
}
