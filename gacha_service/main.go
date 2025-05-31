package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"local.dev/gacha_service/internal"
	"local.dev/gacha_service/proto_file/gacha"
)

func run() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Println("Got error while running", err)
	}

	server := grpc.NewServer()

	gacha.RegisterGachaServiceServer(server, internal.GachaService{})

	err = server.Serve(lis)

	if err != nil {
		fmt.Println("Found Error when running service", err)
	}

}

func main() {
	run()
}
