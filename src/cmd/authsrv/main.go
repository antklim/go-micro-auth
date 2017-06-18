package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	config "../../pkg/config"
	consulConfig "../../pkg/config/consul"
	proto "../../pkg/proto/auth"
	jwt "github.com/dgrijalva/jwt-go"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Auth structure, contains different authentification methods
type Auth struct {
	config config.ConfigHandler
}

// CreateJwt method implementation
func (auth *Auth) CreateJwt(ctx context.Context, req *proto.CreateJwtRequest, rsp *proto.CreateJwtResponse) error {
	secret, err := auth.config.GetKVPair("jwssecret")
	if err != nil {
		return err
	}

	ttl, err := auth.config.GetKVPair("jwtttl")
	if err != nil {
		return err
	}

	exp, err := strconv.ParseInt(string(ttl), 10, 32)
	if err != nil {
		return err
	}

	iat := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      iat,
		"exp":      iat + exp,
		"username": req.GetUsername(),
		"password": req.GetPassword(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return err
	}

	rsp.Token = tokenString

	return nil
}

// ValidateJwt method implementation
func (auth *Auth) ValidateJwt(ctx context.Context, req *proto.ValidateJwtRequest, rsp *proto.ValidateJwtResponse) error {
	tokenString := req.GetToken()
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		claims := token.Claims.(jwt.MapClaims)
		if claims["iat"] == nil {
			return nil, fmt.Errorf("Required field 'iat' not found")
		}

		if claims["exp"] == nil {
			return nil, fmt.Errorf("Required field 'exp' not found")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	// TODO: config should be switchable based on environment variables or CLI options
	serviceConfig := consulConfig.InitConfig("auth/config/")
	if serviceConfig.Err != nil {
		log.Fatal(serviceConfig.Err)
		return
	}

	proto.RegisterAuthHandler(service.Server(), &Auth{config.InitConfig(serviceConfig)})

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
