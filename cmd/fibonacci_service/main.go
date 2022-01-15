package main

import (
	"github.com/Eretic431/fibonacci/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {
	s, closer, err := initServer()
	if err != nil {
		log.Fatalf("could not init server: %v", err)
	}
	closer()
	s.Run()
}

func newLogger(c *config.Config) (*zap.SugaredLogger, func(), error) {
	var logger *zap.Logger
	var err error

	if c.ServerCfg.Production {
		logger, err = zap.NewProduction()
	} else {
		conf := zap.NewDevelopmentConfig()
		conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = conf.Build()
	}

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = logger.Sync()
	}

	return logger.Sugar(), cleanup, nil
}
