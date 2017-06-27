package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Interaction with the file should be tested in integration tests

const keyValureDelimiter string = "="

type Config struct {
	KVPairs map[string][]byte
}

func Init(configFilePath string) (*Config, error) {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	config := &Config{KVPairs: make(map[string][]byte)}
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		configLine := scanner.Text()

		if len(strings.TrimSpace(configLine)) == 0 {
			continue
		}

		key, value, err := parseConfigLine(configLine)
		if err != nil {
			return nil, err
		}

		config.KVPairs[key] = []byte(value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c Config) GetKVPair(key string) ([]byte, error) {
	value, prs := c.KVPairs[key]

	if prs == false {
		return nil, fmt.Errorf("Value for key %q not found", key)
	}

	return value, nil
}

func parseConfigLine(configLine string) (string, string, error) {
	parts := strings.Split(configLine, keyValureDelimiter)

	if len(parts) != 2 {
		return "", "", fmt.Errorf("Cannot parse config key value pair")
	}

	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
}
