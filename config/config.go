package config

import "github.com/caarlos0/env"

type Config struct {
	ServerCfg *ServerConfig
	RedisCfg  *RedisConfig
}

type ServerConfig struct {
	Production bool   `env:"PRODUCTION" envDefault:"false"`
	Port       string `env:"PORT" envDefault:":8080"`
}

type RedisConfig struct {
	Url string `env:"REDIS_URL" envDefault:"localhost:6379"`
}

func GetConfig() (*Config, error) {
	sCfg := &ServerConfig{}
	if err := env.Parse(sCfg); err != nil {
		return nil, err
	}

	rCfg := &RedisConfig{}
	if err := env.Parse(rCfg); err != nil {
		return nil, err
	}

	return &Config{
		ServerCfg: sCfg,
		RedisCfg:  rCfg,
	}, nil
}
