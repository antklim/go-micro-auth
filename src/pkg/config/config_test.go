package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_initConfig(t *testing.T) {
	for _, test := range initConfigTestCases {
		actualResult, err := initConfig(test.source)

		if test.err == nil {
			require.NoError(t, err, "Expected no error")
			// TODO: validate cofig type (file, consul)
			require.NotNil(t, actualResult, "Expected config not nil")
		} else {
			require.Nil(t, actualResult, "Expected no config inited")

			if !reflect.DeepEqual(test.err, err) {
				t.Fatalf("Expected %v, but got %v", test.err, err)
			}
		}
	}
}

func TestGet(t *testing.T) {
	for _, test := range getTestCases {
		config = test.config
		actualResult, err := Get()

		if test.err == nil {
			require.NoError(t, err, "Expected no error")

			if !reflect.DeepEqual(test.config, actualResult) {
				t.Fatalf("Expected %v, but got %v", test.config, actualResult)
			}
		} else {
			require.Nil(t, actualResult, "Expected no config returned")

			if !reflect.DeepEqual(test.err, err) {
				t.Fatalf("Expected %v, but got %v", test.err, err)
			}
		}
	}
}
