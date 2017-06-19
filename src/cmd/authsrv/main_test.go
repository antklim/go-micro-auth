package main

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	config "../../pkg/config"
	proto "../../pkg/proto/auth"
	"golang.org/x/net/context"
)

// Config mock
type testConfig struct{}

var _ config.ConfigHandler = (*testConfig)(nil)

func (c *testConfig) GetKVPair(key string) ([]byte, error) {
	switch key {
	case "jwssecret":
		return []byte("secret"), nil
	case "jwtttl":
		return []byte("180000"), nil
	default:
		return nil, errors.New("Test error")
	}
}

type testClaims struct {
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestCreateJwt(t *testing.T) {
	for _, test := range createJwtTestCases {
		auth := &Auth{new(testConfig)}
		req := proto.CreateJwtRequest{
			Username: test.username,
			Password: test.password,
		}
		rsp := proto.CreateJwtResponse{}

		if err := auth.CreateJwt(context.TODO(), &req, &rsp); err != nil {
			t.Fatalf("Error happened: %v", err)
		}

		tokenParts := strings.Split(rsp.GetToken(), ".")

		if l := len(tokenParts); l != 3 {
			t.Fatalf("JWS token should contain %d segments, but found %d segments", 3, l)
		}

		if tokenParts[0] != test.header {
			t.Fatalf("JWS header expected: %s, but found: %s", test.header, tokenParts[0])
		}

		claimBytes, err := jwt.DecodeSegment(tokenParts[1])

		if err != nil {
			t.Fatalf("Error happened: %v", err)
		}

		claims := testClaims{}
		if err = json.Unmarshal(claimBytes, &claims); err != nil {
			t.Fatalf("Error happened: %v", err)
		}

		if test.username != claims.Username {
			t.Fatalf("Username expected: %s, but found: %s", test.username, claims.Username)
		}

		if test.password != claims.Password {
			t.Fatalf("Password expected: %s, but found: %s", test.password, claims.Password)
		}

		if claims.Iat > time.Now().Unix() {
			t.Fatal("'iat' field should be before current time")
		}

		if claims.Exp != claims.Iat+180000 {
			t.Fatalf("'exp' field value expected: %d, but found: %d", claims.Iat+180000, claims.Exp)
		}
	}
}

func TestValidateJwt(t *testing.T) {
	for _, test := range validateJwtTestCases {
		auth := &Auth{new(testConfig)}
		req := proto.ValidateJwtRequest{
			Token: test.token,
		}
		rsp := proto.ValidateJwtResponse{}
		err := auth.ValidateJwt(context.TODO(), &req, &rsp)

		if test.isError == true {
			if err == nil {
				t.Fatal("Validation should fail")
			}

			if test.expected != err.Error() {
				t.Fatalf("Expected error: %s, but found: %s", test.expected, err.Error())
			}
		} else {
			if err != nil {
				t.Fatalf("Validation should pass, but found error: %v", err)
			}
		}
	}
}
