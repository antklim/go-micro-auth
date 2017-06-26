package auth

import "fmt"
import proto "../../../pkg/proto/auth"

func getCreateJwtRequest(username, password string) *proto.CreateJwtRequest {
	return &proto.CreateJwtRequest{
		Username: username,
		Password: password,
	}
}

var createJwtTestCases = []struct {
	request     *proto.CreateJwtRequest
	header      string
	claims      string
	description string
}{
	{getCreateJwtRequest("username", "password"), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", "eyJleHAiOjE4MDAwMCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ", "should successfuly generate JWS"},
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
	request          *proto.ValidateJwtRequest
	expectedResponse *proto.ValidateJwtResponse
}{
	{getValidateJwtRequest(""), getValidateJwtResponse(fmt.Errorf("token contains an invalid number of segments"))},
	{getValidateJwtRequest("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc5NzI1MDgsImlhdCI6MTQ5Nzc5MjUwOCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.3AUFTX--FZtFbQwFiMHpKxUsicou-xhbjGcymunPxpG-vTxs1xxixDNy_pl7xNxQMCWxAwUxJC-WHMITDiroTUP_F-cSv4CPDShbhxxmcLqfd7BLrtRGBeoDKs5gHfmr80cKdEt23Il3EWD_6f1c4ItNMilLLL_d00bPrPg7wck"), getValidateJwtResponse(fmt.Errorf("Unexpected signing method: RS256"))},
	{getValidateJwtRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc5NzI1MDgsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.fW6vIfgYjANXPpOkwFc6gI5PIxCCvH1KVfWkqOD-huY"), getValidateJwtResponse(fmt.Errorf("Required field 'iat' not found"))},
	{getValidateJwtRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0OTc3OTI1MDgsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.jQQEgdCopSBd7ivRI-Q9t-F2KomDIqyKOzq69GScvS4"), getValidateJwtResponse(fmt.Errorf("Required field 'exp' not found"))},
	{getValidateJwtRequest("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MTMzMzIzMDQsImlhdCI6MTQ5Nzc5MjMwNCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.TtGPJzgyK_Ybiw0-4KqLu-kOe-oW9N1A_dzTdcdMzZ8"), getValidateJwtResponse(nil)},
}
