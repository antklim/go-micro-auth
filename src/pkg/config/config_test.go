package config

import "testing"

func TestGet(t *testing.T) {
	config = nil
	_, err := Get()
	if err.Error() != "Config is not inited" {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func Test_initConfig(t *testing.T) {
	_, err := initConfig("blah")
	if err.Error() != "Unsupported config source: 'blah'" {
		t.Fatalf("Unexpected error: %v", err)
	}

}
