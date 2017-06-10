package auth

import consulapi "github.com/hashicorp/consul/api"

type ConsulConfig struct {
	consul *consulapi.Client
	kv     *consulapi.KV
	err    error
}

func InitConsulConfig() *ConsulConfig {
	consulConfig := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(consulConfig)

	if err != nil {
		return &ConsulConfig{consul, nil, err}
	}

	return &ConsulConfig{consul, consul.KV(), err}
}

func (c *ConsulConfig) GetKVPair(key string) ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}

	kvp, _, err := c.kv.Get(key, nil)
	if err != nil {
		return nil, err
	}

	return kvp.Value, err
}
