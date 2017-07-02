package config

import "fmt"
import (
	consulConfig "github.com/antklim/go-micro-auth/config/consul"
	fileConfig "github.com/antklim/go-micro-auth/config/file"
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

func Init() (*Config, error) {
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

func (c *Config) GetKVPair(key string) ([]byte, error) {
	return c.ConfigHandler.GetKVPair(key)
}
