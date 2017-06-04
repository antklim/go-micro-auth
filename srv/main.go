package main

import (
	"log"

	proto "./proto/auth"
	"github.com/dgrijalva/jwt-go"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Auth structure, contains different authentification methods
type Auth struct{}

// Jwt method implementation
func (auth *Auth) Jwt(ctx context.Context, req *proto.JwtRequest, rsp *proto.JwtResponse) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      180000,
		"username": req.GetUsername(),
		"password": req.GetPassword(),
	})

	secret := []byte("secret")
	tokenString, err := token.SignedString(secret)

	if err == nil {
		rsp.Token = tokenString
	}

	return err
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	proto.RegisterAuthHandler(service.Server(), new(Auth))

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
