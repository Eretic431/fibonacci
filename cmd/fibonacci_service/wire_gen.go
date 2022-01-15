// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/Eretic431/fibonacci/config"
	"github.com/Eretic431/fibonacci/internal/fibonacci/delivery/grpc"
	"github.com/Eretic431/fibonacci/internal/fibonacci/delivery/http"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
	"github.com/Eretic431/fibonacci/internal/server"
)

// Injectors from wire.go:

func initServer() (*server.Server, func(), error) {
	configConfig, err := config.GetConfig()
	if err != nil {
		return nil, nil, err
	}
	sugaredLogger, cleanup, err := newLogger(configConfig)
	if err != nil {
		return nil, nil, err
	}
	fibonacciUseCase := usecase.NewFibonacciUseCase()
	fibonacciService := grpc.NewGrpcFibonacciService(fibonacciUseCase, sugaredLogger)
	httpFibonacciService := http.NewGrpcFibonacciService(fibonacciUseCase, sugaredLogger)
	serverServer := &server.Server{
		Log:         sugaredLogger,
		Cfg:         configConfig,
		GrpcService: fibonacciService,
		HttpService: httpFibonacciService,
	}
	return serverServer, func() {
		cleanup()
	}, nil
}