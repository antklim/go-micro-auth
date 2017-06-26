package main

import (
	"fmt"
	"strconv"
	"time"

	config "../../pkg/config"
	proto "../../pkg/proto/auth"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

// Auth structure, contains different authentification methods
type Auth struct{}

// CreateJwt method implementation
func (auth *Auth) CreateJwt(ctx context.Context, req *proto.CreateJwtRequest, rsp *proto.CreateJwtResponse) error {
	serviceConfig, err := config.Get()
	if err != nil {
		return err
	}

	secret, err := serviceConfig.GetKVPair("jwssecret")
	if err != nil {
		return err
	}

	ttl, err := serviceConfig.GetKVPair("jwtttl")
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

		serviceConfig, err := config.Get()
		if err != nil {
			return nil, err
		}

		secret, err := serviceConfig.GetKVPair("jwssecret")
		return secret, err
	})

	rsp.Valid = err == nil
	if err != nil {
		rsp.Error = err.Error()
	}

	return nil
}
