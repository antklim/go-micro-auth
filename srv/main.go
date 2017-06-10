package auth

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

// func createJwt()

// Auth structure, contains different authentification methods
type Auth struct {
	jwssecret []byte
	jwtttl    int32
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

	// TODO: Move the following code to config initialisation part
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	kv := consul.KV()

	// TODO: Move the KeyValue pair getter into config interface
	secret, err := getKVPair(kv, "auth/config/jwssecret")
	if err != nil {
		log.Fatal(err)
		return
	}

	ttl, err := getKVPair(kv, "auth/config/jwtttl")
	if err != nil {
		log.Fatal(err)
		return
	}

	exp64, err := strconv.ParseInt(string(ttl), 10, 32)
	if err != nil {
		log.Fatal(err)
		return
	}

	////////////////////////
	proto.RegisterAuthHandler(service.Server(), &Auth{jwssecret: secret, jwtttl: int32(exp64)})

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
