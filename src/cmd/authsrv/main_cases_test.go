package main

var testCases = []struct {
	username    string
	password    string
	header      string
	claims      string
	description string
}{
	{"username", "password", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", "eyJleHAiOjE4MDAwMCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ", "should successfuly generate JWS"},
}
