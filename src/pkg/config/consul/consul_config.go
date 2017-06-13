package consul

import consulapi "github.com/hashicorp/consul/api"

type Config struct {
	consul    *consulapi.Client
	kv        *consulapi.KV
	keyPrefix string
	Err       error
}

func InitConfig(keyPrefix string) *Config {
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)

	if err != nil {
		return &Config{consul, nil, keyPrefix, err}
	}

	return &Config{consul, consul.KV(), keyPrefix, err}
}

func (c *Config) GetKVPair(key string) ([]byte, error) {
	if c.Err != nil {
		return nil, c.Err
	}

	kvp, _, err := c.kv.Get(c.keyPrefix+key, nil)
	if err != nil {
		return nil, err
	}

	return kvp.Value, err
}
