package file

// TODO: Implement file config initialisation and GetKVPair
// TODO: Interaction with the file should be tested in integration tests

type Config struct {
	file interface{}
}

func Init(configFilePath string) (*Config, error) {
	return &Config{nil}, nil
}

func (c *Config) GetKVPair(key string) ([]byte, error) {
	return nil, nil
}
