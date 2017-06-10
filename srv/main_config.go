package auth

type ConfigHandler interface {
	GetKVPair(key string)
}

type Config struct {
	ConfigHandler
}

func Init(hdlr ConfigHandler) *Config {
	return &Config{hdlr}
}
