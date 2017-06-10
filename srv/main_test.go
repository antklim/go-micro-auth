package auth

import (
	"errors"
	"testing"

	proto "./proto/auth"
)

import "golang.org/x/net/context"

// TODO: stub time library
// due to iat and exp field addition token will depend on current timestamp

// Config mock
type testConfig struct{}

var _ ConfigHandler = (*testConfig)(nil)

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

func TestJwtGeneration(t *testing.T) {
	for _, test := range testCases {
		auth := &Auth{new(testConfig)}
		req := proto.CreateJwtRequest{
			Username: test.username,
			Password: test.password,
		}
		rsp := proto.CreateJwtResponse{}
		err := auth.CreateJwt(context.TODO(), &req, &rsp)

		if err != nil {
			t.Fatalf("Error happened")
		}

		if rsp.GetToken() != test.expected {
			// t.Fatalf("Jwt(ctx, req, rsp) = %s, want %s (%s)",
			// 	rsp.GetToken(), test.expected, test.description)
		}
	}
}
