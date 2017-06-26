package config

import "fmt"

var initConfigTestCases = []struct {
	source       string
	configInited bool
	err          error
}{
	{"blah", false, fmt.Errorf("Unsupported config source: 'blah'")},
	{"consul", true, nil},
	{"file", true, nil},
}

var getTestCases = []struct {
	config *Config
	err    error
}{
	{nil, fmt.Errorf("Config is not inited")},
	{&Config{nil}, nil},
}

type testConfigHandler struct{}

func (c *testConfigHandler) GetKVPair(key string) ([]byte, error) {
	switch key {
	case "blah":
		return []byte("blah"), nil
	default:
		return nil, fmt.Errorf("Test error")
	}
}

var getKVPairTestCases = []struct {
	config   *Config
	key      string
	expected []byte
	err      error
}{
	{&Config{new(testConfigHandler)}, "blah", []byte("blah"), nil},
	{&Config{new(testConfigHandler)}, "foo", nil, fmt.Errorf("Test error")},
}
