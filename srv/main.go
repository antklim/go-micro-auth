package auth

import (
	"log"
	"strconv"
	"time"

	proto "./proto/auth"
	"github.com/dgrijalva/jwt-go"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Auth structure, contains different authentification methods
type Auth struct {
	config ConfigHandler
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

	exp64, err := strconv.ParseInt(string(ttl), 10, 32)
	if err != nil {
		return err
	}

	iat := int32(time.Now().Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      iat,
		"exp":      iat + int32(exp64),
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
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	// TODO: config should be switchable based on environment variables or CLI options
	config := InitConsulConfig("auth/config/")
	if config.err != nil {
		log.Fatal(config.err)
		return
	}

	proto.RegisterAuthHandler(service.Server(), &Auth{config})

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
