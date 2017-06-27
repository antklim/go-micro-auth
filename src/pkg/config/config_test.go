package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_initConfig(t *testing.T) {
	for _, test := range initConfigTestCases {
		actualResult, err := initConfig(test.source)

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected error to be %v, but got %v", test.err, err)
		}

		if test.configInited {
			// TODO: validate cofig type (file, consul)
			require.NotNil(t, actualResult, "Expected config not nil")
		} else {
			require.Nil(t, actualResult, "Expected no config inited")
		}
	}
}

func TestGet(t *testing.T) {
	for _, test := range getTestCases {
		config = test.config
		actualResult, err := Get()

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected error to be %v, but got %v", test.err, err)
		}

		if !reflect.DeepEqual(test.config, actualResult) {
			t.Fatalf("Expected %v, but got %v", test.config, actualResult)
		}
	}
}

func TestGetKVPair(t *testing.T) {
	for _, test := range getKVPairTestCases {
		config = test.config
		actualResult, err := config.GetKVPair(test.key)

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected %v, but got %v", test.err, err)
		}

		if !reflect.DeepEqual(test.expected, actualResult) {
			t.Fatalf("Expected %v, but got %v", test.expected, actualResult)
		}
	}
}
