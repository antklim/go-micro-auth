package main

import (
	"log"
	"strconv"
	"time"

	proto "./proto/auth"
	"github.com/dgrijalva/jwt-go"
	consulapi "github.com/hashicorp/consul/api"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func getKVPair(kv *consulapi.KV, key string) ([]byte, error) {
	kvp, _, err := kv.Get(key, nil)
	return kvp.Value, err
}

// Auth structure, contains different authentification methods
type Auth struct{}

// Jwt method implementation
func (auth *Auth) Jwt(ctx context.Context, req *proto.JwtRequest, rsp *proto.JwtResponse) error {
	// TODO: move consul client initiation out of the scope of the method
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return err
	}

	kv := consul.KV()

	secret, err := getKVPair(kv, "auth/config/jwssecret")
	if err != nil {
		return err
	}

	ttl, err := getKVPair(kv, "auth/config/jwtttl")
	if err != nil {
		return err
	}

	iat := int32(time.Now().Unix())
	exp64, err := strconv.ParseInt(string(ttl), 10, 32)
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      iat,
		"exp":      iat + int32(exp64),
		"username": req.GetUsername(),
		"password": req.GetPassword(),
	})
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
