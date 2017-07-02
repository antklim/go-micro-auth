package file

import "fmt"

func getConfig(key, value string) Config {
	kvPairs := make(map[string][]byte)
	kvPairs[key] = []byte(value)

	return Config{
		KVPairs: kvPairs,
	}
}

var getKVPairTestCases = []struct {
	config Config
	key    string
	value  []byte
	err    error
}{
	{
		config: getConfig("blah", "blah"),
		key:    "foo",
		value:  nil,
		err:    fmt.Errorf("Value for key \"foo\" not found"),
	}, {
		config: getConfig("foo", "bar"),
		key:    "foo",
		value:  []byte("bar"),
		err:    nil,
	},
}

var parseConfigLineTestCases = []struct {
	configLine string
	key        string
	value      string
	err        error
}{
	{
		configLine: "",
		key:        "",
		value:      "",
		err:        fmt.Errorf("Cannot parse config key value pair"),
	},
	{
		configLine: "foo = bar   ",
		key:        "foo",
		value:      "bar",
		err:        nil,
	},
}
