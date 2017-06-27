package auth

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"

	proto "../../../pkg/proto/auth"
	"golang.org/x/net/context"
)

type testClaims struct {
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestCreateJwt(t *testing.T) {
	for _, test := range createJwtTestCases {
		auth := &Auth{ConfigHandler: test.configHandler}
		req := test.request
		rsp := &proto.CreateJwtResponse{}
		err := auth.CreateJwt(context.TODO(), req, rsp)

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected error to be %v, but got %v", test.err, err)
		}

		if test.err != nil {
			continue
		}

		tokenParts := strings.Split(rsp.GetToken(), ".")

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
		auth := &Auth{ConfigHandler: test.configHandler}
		rsp := &proto.ValidateJwtResponse{}
		err := auth.ValidateJwt(context.TODO(), test.request, rsp)

		require.NoError(t, err, "Expected no error")

		if !reflect.DeepEqual(test.expectedResponse, rsp) {
			t.Fatalf("Expected %v, but got %v", test.expectedResponse, rsp)
		}
	}
}
