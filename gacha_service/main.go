package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/inonsdn/gacha-system/gacha_service/internal"
	"github.com/inonsdn/gacha-system/gacha_service/internal/dbhandler"
	gachapb "github.com/inonsdn/gacha-system/proto/gacha"
)

func run() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		fmt.Println("Got error while running", err)
	}

	server := grpc.NewServer()

	gachapb.RegisterGachaServiceServer(server, internal.GachaService{
		DBHandler: dbhandler.NewDBHandler(),
	})
	fmt.Println("Run Gacha Service at 50052")
	err = server.Serve(lis)

	if err != nil {
		fmt.Println("Found Error when running service", err)
	}

}

func main() {
	run()
}
