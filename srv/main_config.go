package auth

type ConfigHandler interface {
	GetKVPair(key string) ([]byte, error)
}

type Config struct {
	ConfigHandler
}

func InitConfig(hdlr ConfigHandler) *Config {
	return &Config{hdlr}
}

func (c *Config) GetKVPair(key string) ([]byte, error) {
	return c.ConfigHandler.GetKVPair(key)
}
