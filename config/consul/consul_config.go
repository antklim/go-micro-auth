package consul

import consulapi "github.com/hashicorp/consul/api"

// TODO: Interaction with the consul should be tested in integration tests

type Config struct {
	consul    *consulapi.Client
	kv        *consulapi.KV
	keyPrefix string
}

func Init(keyPrefix string) (*Config, error) {
	consul, err := consulapi.NewClient(consulapi.DefaultConfig())
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
