package config

import "github.com/caarlos0/env"

type Config struct {
	ServerCfg *ServerConfig
}

type ServerConfig struct {
	Production bool   `env:"PRODUCTION" envDefault:"false"`
	Port       string `env:"PORT" envDefault:":8080"`
}

func GetConfig() (*Config, error) {
	sCfg := &ServerConfig{}
	if err := env.Parse(sCfg); err != nil {
		return nil, err
	}

	return &Config{ServerCfg: sCfg}, nil
}
