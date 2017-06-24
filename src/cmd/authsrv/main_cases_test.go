package main

var createJwtTestCases = []struct {
	username    string
	password    string
	header      string
	claims      string
	description string
}{
	{"username", "password", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", "eyJleHAiOjE4MDAwMCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ", "should successfuly generate JWS"},
}

var validateJwtTestCases = []struct {
	token    string
	isError  bool
	expected string
}{
	{"", true, "token contains an invalid number of segments"},
	{"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc5NzI1MDgsImlhdCI6MTQ5Nzc5MjUwOCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.3AUFTX--FZtFbQwFiMHpKxUsicou-xhbjGcymunPxpG-vTxs1xxixDNy_pl7xNxQMCWxAwUxJC-WHMITDiroTUP_F-cSv4CPDShbhxxmcLqfd7BLrtRGBeoDKs5gHfmr80cKdEt23Il3EWD_6f1c4ItNMilLLL_d00bPrPg7wck", true, "Unexpected signing method: RS256"},
	{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTc5NzI1MDgsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.fW6vIfgYjANXPpOkwFc6gI5PIxCCvH1KVfWkqOD-huY", true, "Required field 'iat' not found"},
	{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0OTc3OTI1MDgsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.jQQEgdCopSBd7ivRI-Q9t-F2KomDIqyKOzq69GScvS4", true, "Required field 'exp' not found"},
	{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MTMzMzIzMDQsImlhdCI6MTQ5Nzc5MjMwNCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.TtGPJzgyK_Ybiw0-4KqLu-kOe-oW9N1A_dzTdcdMzZ8", false, ""},
}
