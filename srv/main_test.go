package main

import "testing"
import proto "./proto/auth"
import "golang.org/x/net/context"

// TODO: stub consul KeyValue
// TODO: stub time library
// due to iat and exp field addition token will depend on current timestamp

func TestJwtGeneration(t *testing.T) {
	for _, test := range testCases {
		auth := new(Auth)
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
