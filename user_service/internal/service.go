package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	userpb "github.com/inonsdn/gacha-system/proto/user"
)

var jwtSecret = []byte("ebcc8244-5629-4633-bbdd-f5f2253a13bd")

type UserService struct {
	userpb.UnimplementedUserServiceServer
}

func (u UserService) Login(c context.Context, request *userpb.UserLoginRequest) (*userpb.UserLoginResponse, error) {
	var err error

	fmt.Println("Login", request.LoginName, request.Password)

	// TODO: verify login

	claims := jwt.MapClaims{
		"user_id": request.LoginName + request.Password,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret)
	fmt.Println("signed", signed)

	return &userpb.UserLoginResponse{
		JwtToken: signed,
		Error:    "",
	}, err
}
