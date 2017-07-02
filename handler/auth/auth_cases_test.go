package auth

import "fmt"
import proto "github.com/antklim/go-micro-auth/proto/auth"

type testConfigHandler struct {
	KVPairs map[string][]byte
	ErrKey  string
	Err     error
}

func (c testConfigHandler) GetKVPair(key string) ([]byte, error) {
	if c.Err != nil && key == c.ErrKey {
		return nil, c.Err
	}

	return c.KVPairs[key], nil
}

func getTestConfigHandler(err error, errKey string) testConfigHandler {
	if err != nil {
		return testConfigHandler{Err: err, ErrKey: errKey}
	}

	kvPairs := make(map[string][]byte, 2)
	kvPairs["jwtttl"] = []byte("180000")
	kvPairs["jwssecret"] = []byte("secret")
	return testConfigHandler{KVPairs: kvPairs, ErrKey: "", Err: nil}
}

func getCreateJwtRequest(username, password string) *proto.CreateJwtRequest {
	return &proto.CreateJwtRequest{
		Username: username,
		Password: password,
	}
}

var createJwtTestCases = []struct {
	configHandler testConfigHandler
	request       *proto.CreateJwtRequest
	header        string
	claims        string
	err           error
}{
	{
		configHandler: getTestConfigHandler(nil, ""),
		request:       getCreateJwtRequest("username", "password"),
		header:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		claims:        "eyJleHAiOjE4MDAwMCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ",
		err:           nil,
	}, {
		configHandler: getTestConfigHandler(fmt.Errorf("Key not found"), "jwtttl"),
		request:       getCreateJwtRequest("username", "password"),
		header:        "",
		claims:        "",
		err:           fmt.Errorf("Key not found"),
	}, {
		configHandler: getTestConfigHandler(fmt.Errorf("Key not found"), "jwssecret"),
		request:       getCreateJwtRequest("username", "password"),
		header:        "",
		claims:        "",
		err:           fmt.Errorf("Key not found"),
	},
}

func getValidateJwtRequest(token string) *proto.ValidateJwtRequest {
	return &proto.ValidateJwtRequest{
		Token: token,
	}
}

func getValidateJwtResponse(err error) *proto.ValidateJwtResponse {
	if err != nil {
		return &proto.ValidateJwtResponse{
			Valid: false,
			Error: err.Error(),
		}
	}

	return &proto.ValidateJwtResponse{
		Valid: true,
	}
}

var validateJwtTestCases = []struct {
	configHandler    testConfigHandler
	request          *proto.ValidateJwtRequest
	expectedResponse *proto.ValidateJwtResponse
}{
	{
		configHandler:    getTestConfigHandler(nil, ""),
		request:          getValidateJwtRequest(""),
		expectedResponse: getValidateJwtResponse(fmt.Errorf("token contains an invalid number of segments")),
	}, {
		configHandler:    getTestConfigHandler(nil, ""),
		request:          getValidateJwtRequest("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc5NzI1MDgsImlhdCI6MTQ5Nzc5MjUwOCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.3AUFTX--FZtFbQwFiMHpKxUsicou-xhbjGcymunPxpG-vTxs1xxixDNy_pl7xNxQMCWxAwUxJC-WHMITDiroTUP_F-cSv4CPDShbhxxmcLqfd7BLrtRGBeoDKs5gHfmr80cKdEt23Il3EWD_6f1c4ItNMilLLL_d00bPrPg7wck"),
		expectedResponse: getValidateJwtResponse(fmt.Errorf("Unexpected signing method: RS256")),
	}, {
		configHandler:    getTestConfigHandler(nil, ""),
		request:          getValidateJwtRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc5NzI1MDgsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.fW6vIfgYjANXPpOkwFc6gI5PIxCCvH1KVfWkqOD-huY"),
		expectedResponse: getValidateJwtResponse(fmt.Errorf("Required field 'iat' not found")),
	}, {
		configHandler:    getTestConfigHandler(nil, ""),
		request:          getValidateJwtRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0OTc3OTI1MDgsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.jQQEgdCopSBd7ivRI-Q9t-F2KomDIqyKOzq69GScvS4"),
		expectedResponse: getValidateJwtResponse(fmt.Errorf("Required field 'exp' not found")),
	}, {
		configHandler:    getTestConfigHandler(nil, ""),
		request:          getValidateJwtRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MTMzMzIzMDQsImlhdCI6MTQ5Nzc5MjMwNCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.TtGPJzgyK_Ybiw0-4KqLu-kOe-oW9N1A_dzTdcdMzZ8"),
		expectedResponse: getValidateJwtResponse(nil),
	},
}
