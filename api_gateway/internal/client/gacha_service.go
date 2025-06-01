package client

import (
	"context"
	"fmt"

	gachapb "github.com/inonsdn/gacha-system/proto/gacha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GachaServiceClient struct {
	client gachapb.GachaServiceClient
}

func NewGachaServiceClient() *GachaServiceClient {
	// conn, err := grpc.Dial("gacha-service:50052", grpc.WithInsecure()) // Port depends on Docker setup
	// TODO: what is DialOptions
	// insecure is for dev, for production will use others
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Sprintf("Could not connect to GachaService: %v", err))
	}

	client := gachapb.NewGachaServiceClient(conn)
	return &GachaServiceClient{client: client}
}

func (c *GachaServiceClient) GetServiceName() string {
	return "Gacha"
}

func (c *GachaServiceClient) Draw(userID string, gachaId string, amount int32) (*gachapb.DrawResponse, error) {
	req := &gachapb.DrawRequest{
		UserId:     userID,
		GachaId:    gachaId,
		DrawAmount: amount,
	}
	return c.client.Draw(context.Background(), req)
}

func (c *GachaServiceClient) GetGachaInfo(userId string, gachaType string) (*gachapb.GachaResponse, error) {
	req := &gachapb.GachaRequest{GachaType: gachaType}
	fmt.Println("GetGachaInfo")
	return c.client.GetGachaInfo(context.Background(), req)
}
