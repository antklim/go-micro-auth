package consul

import (
	"errors"
	"testing"
)

func TestGetKVPair(t *testing.T) {
	config := Config{nil, nil, "test", errors.New("Test error")}
	_, err := config.GetKVPair("key")
	if err == nil {
		t.Error("GetKVPair should return error")
	}
}
