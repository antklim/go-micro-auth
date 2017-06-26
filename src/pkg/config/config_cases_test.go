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
