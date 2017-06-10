package auth

import consulapi "github.com/hashicorp/consul/api"

type ConsulConfig struct {
	consul    *consulapi.Client
	kv        *consulapi.KV
	keyPrefix string
	err       error
}

func InitConsulConfig(keyPrefix string) *ConsulConfig {
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)

	if err != nil {
		return &ConsulConfig{consul, nil, keyPrefix, err}
	}

	return &ConsulConfig{consul, consul.KV(), keyPrefix, err}
}

func (c *ConsulConfig) GetKVPair(key string) ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}

	kvp, _, err := c.kv.Get(c.keyPrefix+key, nil)
	if err != nil {
		return nil, err
	}

	return kvp.Value, err
}
