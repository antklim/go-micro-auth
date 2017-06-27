package file

import (
	"reflect"
	"testing"
)

func TestGetKVPair(t *testing.T) {
	for _, test := range getKVPairTestCases {
		value, err := test.config.GetKVPair(test.key)

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected error to be %v, but got %v", test.err, err)
		}

		if !reflect.DeepEqual(test.value, value) {
			t.Fatalf("Expected value to be %v, but got %v", test.value, value)
		}
	}
}

func Test_parseConfigLine(t *testing.T) {
	for _, test := range parseConfigLineTestCases {
		key, value, err := parseConfigLine(test.configLine)

		if !reflect.DeepEqual(test.err, err) {
			t.Fatalf("Expected error to be %v, but got %v", test.err, err)
		}

		if !reflect.DeepEqual(test.key, key) {
			t.Fatalf("Expected key to be %v, but got %v", test.key, key)
		}

		if !reflect.DeepEqual(test.value, value) {
			t.Fatalf("Expected value to be %v, but got %v", test.value, value)
		}
	}
}
