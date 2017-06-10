package auth

type ConfigHandler interface {
	GetKVPair(key string)
}

type Config struct {
	ConfigHandler
}

func InitConfig(hdlr ConfigHandler) *Config {
	return &Config{hdlr}
}
