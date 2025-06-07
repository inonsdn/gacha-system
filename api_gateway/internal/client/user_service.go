package client

import (
	"context"
	"fmt"

	userpb "github.com/inonsdn/gacha-system/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	client userpb.UserServiceClient
}

func NewUserServiceClient() *UserServiceClient {
	// conn, err := grpc.Dial("gacha-service:50051", grpc.WithInsecure()) // Port depends on Docker setup
	// TODO: what is DialOptions
	// insecure is for dev, for production will use others
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Sprintf("Could not connect to UserService: %v", err))
	}

	client := userpb.NewUserServiceClient(conn)
	return &UserServiceClient{client: client}
}

func (c *UserServiceClient) Login(loginName string, password string) (*userpb.UserLoginResponse, error) {
	req := &userpb.UserLoginRequest{
		LoginName: loginName,
		Password:  password,
	}
	return c.client.Login(context.Background(), req)
}

func (c *UserServiceClient) Register(loginName string, password string) (*userpb.RegistrationResponse, error) {
	req := &userpb.RegistrationRequest{
		LoginName: loginName,
		Password:  password,
		TelNo:     "",
		Address:   "",
		Birthday:  "",
	}
	return c.client.Register(context.Background(), req)
}
