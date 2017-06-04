package main

var testCases = []struct {
	username    string
	password    string
	expected    string
	description string
}{
	{"username", "password", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4MDAwMCwicGFzc3dvcmQiOiJwYXNzd29yZCIsInVzZXJuYW1lIjoidXNlcm5hbWUifQ.GSPMU93RPoPu6cfoMqqh1FmuMF3Cz0VOBRbvZuxYfhk", "should succssfuly generate JWT"},
}
