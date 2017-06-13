package file

// TODO: Implement file config initialisation and GetKVPair

type Config struct {
	file interface{}
	Err  error
}

func InitConfig(configFilePath string) *Config {
	return &Config{nil, nil}
}

func (c *Config) GetKVPair(key string) ([]byte, error) {
	if c.Err != nil {
		return nil, c.Err
	}

	return nil, nil
}
