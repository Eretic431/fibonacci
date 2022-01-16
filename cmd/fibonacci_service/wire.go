//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Eretic431/fibonacci/config"
	"github.com/Eretic431/fibonacci/internal/fibonacci"
	grpcS "github.com/Eretic431/fibonacci/internal/fibonacci/delivery/grpc"
	httpS "github.com/Eretic431/fibonacci/internal/fibonacci/delivery/http"
	"github.com/Eretic431/fibonacci/internal/fibonacci/repository/redis"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
	"github.com/Eretic431/fibonacci/internal/server"
	"github.com/Eretic431/fibonacci/pkg/logger"
	pRedis "github.com/Eretic431/fibonacci/pkg/redis"
	"github.com/Eretic431/fibonacci/pkg/utils"
	"github.com/google/wire"
)

func initServer() (*server.Server, func(), error) {
	wire.Build(
		config.GetConfig,
		logger.NewLogger,
		pRedis.NewRedisPool,
		redis.NewFibonacciRepository,
		wire.Bind(new(fibonacci.FibonacciRepository), new(*redis.FibonacciRepository)),
		usecase.NewFibonacciUseCase,
		wire.Bind(new(fibonacci.FibonacciUseCase), new(*usecase.FibonacciUseCase)),
		grpcS.NewGrpcFibonacciService,
		utils.NewHttpHelper,
		httpS.NewHttpFibonacciService,
		wire.Struct(new(server.Server), "*"),
	)

	return nil, nil, nil
}
