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

// func createJwt()

// Auth structure, contains different authentification methods
type Auth struct {
	jwssecret []byte
	jwtttl    int32
	// TODO: replace jwssecret and jwtttl with config interface
	// config ConfigHandler
}

// CreateJwt method implementation
func (auth *Auth) CreateJwt(ctx context.Context, req *proto.CreateJwtRequest, rsp *proto.CreateJwtResponse) error {
	iat := int32(time.Now().Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      iat,
		"exp":      iat + auth.jwtttl,
		"username": req.GetUsername(),
		"password": req.GetPassword(),
	})

	tokenString, err := token.SignedString(auth.jwssecret)
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

	// TODO: Move the KeyValue pair getters into CreateJwt method
	secret, err := config.GetKVPair("jwssecret")
	if err != nil {
		log.Fatal(err)
		return
	}

	ttl, err := config.GetKVPair("jwtttl")
	if err != nil {
		log.Fatal(err)
		return
	}

	exp64, err := strconv.ParseInt(string(ttl), 10, 32)
	if err != nil {
		log.Fatal(err)
		return
	}

	proto.RegisterAuthHandler(service.Server(), &Auth{jwssecret: secret, jwtttl: int32(exp64)})

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
