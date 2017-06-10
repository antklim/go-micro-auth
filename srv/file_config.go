package auth

// TODO: Implement file config initialisation and GetKVPair

type FileConfig struct {
	file interface{}
	err  error
}

func InitFileConfig(configFilePath string) *FileConfig {
	return &FileConfig{nil, nil}
}

func (c *FileConfig) GetKVPair(key string) ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}

	return nil, nil
}
