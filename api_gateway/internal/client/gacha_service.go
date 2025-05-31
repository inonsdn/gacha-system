package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"local.dev/api_gateway/proto_file/gacha"
)

type GachaServiceClient struct {
	client gacha.GachaServiceClient
}

func NewGachaServiceClient() *GachaServiceClient {
	// conn, err := grpc.Dial("gacha-service:50052", grpc.WithInsecure()) // Port depends on Docker setup
	// TODO: what is DialOptions
	// insecure is for dev, for production will use others
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Sprintf("Could not connect to GachaService: %v", err))
	}

	client := gacha.NewGachaServiceClient(conn)
	return &GachaServiceClient{client: client}
}

func (c *GachaServiceClient) GetServiceName() string {
	return "Gacha"
}

func (c *GachaServiceClient) Draw(userID string) (*gacha.DrawResponse, error) {
	req := &gacha.DrawRequest{UserId: userID}
	return c.client.Draw(context.Background(), req)
}

func (c *GachaServiceClient) GetGachaInfo(userId string, gachaType string) (*gacha.GachaResponse, error) {
	req := &gacha.GachaRequest{GachaType: gachaType}
	fmt.Println("GetGachaInfo")
	return c.client.GetGachaInfo(context.Background(), req)
}
