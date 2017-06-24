package config

import "fmt"
import (
	consulConfig "../../pkg/config/consul"
	fileConfig "../../pkg/config/file"
)

var (
	Source   string
	FilePath string
	config   *Config
)

type ConfigHandler interface {
	GetKVPair(key string) ([]byte, error)
}

type Config struct {
	ConfigHandler
}

func Init() error {
	var err error
	config, err = initConfig()
	return err
}

func initConfig() (*Config, error) {
	var configHandler ConfigHandler
	var err error

	switch Source {
	case "consul":
		configHandler, err = consulConfig.Init("auth/config/")
		break
	case "file":
		configHandler, err = fileConfig.Init(FilePath)
		break
	default:
		err = fmt.Errorf("Unsupported config source: '%s'", Source)
		break
	}

	if err != nil {
		return nil, err
	}

	return &Config{configHandler}, nil
}

func Get() (*Config, error) {
	var err error
	if config == nil {
		err = fmt.Errorf("Config is not inited")
	}
	return config, err
}

func (c *Config) GetKVPair(key string) ([]byte, error) {
	return c.ConfigHandler.GetKVPair(key)
}
