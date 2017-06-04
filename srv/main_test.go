package main

import "testing"
import proto "./proto/auth"
import "golang.org/x/net/context"

func TestJwtGeneration(t *testing.T) {
	for _, test := range testCases {
		auth := new(Auth)
		req := proto.JwtRequest{
			Username: test.username,
			Password: test.password,
		}
		rsp := proto.JwtResponse{}
		err := auth.Jwt(context.TODO(), &req, &rsp)

		if err != nil {
			t.Fatalf("Error happened")
		}

		if rsp.GetToken() != test.expected {
			t.Fatalf("Jwt(ctx, req, rsp) = %s, want %s (%s)",
				rsp.GetToken(), test.expected, test.description)
		}
	}
}
