package consul

import consulapi "github.com/hashicorp/consul/api"

type Config struct {
	consul    *consulapi.Client
	kv        *consulapi.KV
	keyPrefix string
}

func Init(keyPrefix string) (*Config, error) {
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)

	if err != nil {
		return nil, err
	}

	return &Config{consul, consul.KV(), keyPrefix}, nil
}

func (c *Config) GetKVPair(key string) ([]byte, error) {
	kvp, _, err := c.kv.Get(c.keyPrefix+key, nil)
	if err != nil {
		return nil, err
	}

	return kvp.Value, err
}
