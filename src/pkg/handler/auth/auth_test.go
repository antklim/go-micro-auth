package auth

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"

	config "../../../pkg/config"
	proto "../../../pkg/proto/auth"
	"golang.org/x/net/context"
)

// Config mock
type testConfig struct{}

func (c testConfig) GetKVPair(key string) ([]byte, error) {
	switch key {
	case "jwssecret":
		return []byte("secret"), nil
	case "jwtttl":
		return []byte("180000"), nil
	default:
		return nil, errors.New("Test error")
	}
}

var _, _ = config.Set(&config.Config{new(testConfig)})

type testClaims struct {
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestCreateJwt(t *testing.T) {
	for _, test := range createJwtTestCases {
		auth := new(Auth)
		req := test.request
		rsp := &proto.CreateJwtResponse{}
		err := auth.CreateJwt(context.TODO(), req, rsp)

		require.NoError(t, err, "Expected no error")

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

		if req.GetUsername() != claims.Username {
			t.Fatalf("Expected username %s, but got %s", req.GetUsername(), claims.Username)
		}

		if req.GetPassword() != claims.Password {
			t.Fatalf("Expected password %s, but got %s", req.GetPassword(), claims.Password)
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
		auth := new(Auth)
		rsp := &proto.ValidateJwtResponse{}
		err := auth.ValidateJwt(context.TODO(), test.request, rsp)

		require.NoError(t, err, "Expected no error")

		if !reflect.DeepEqual(test.expectedResponse, rsp) {
			t.Fatalf("Expected %v, but got %v", test.expectedResponse, rsp)
		}
	}
}
