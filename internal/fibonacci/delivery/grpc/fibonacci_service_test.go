package grpc

import (
	"context"
	"github.com/Eretic431/fibonacci/config"
	"github.com/Eretic431/fibonacci/internal/fibonacci/mock"
	fibonacciService "github.com/Eretic431/fibonacci/internal/fibonacci/proto"
	"github.com/Eretic431/fibonacci/internal/models"
	"github.com/Eretic431/fibonacci/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibonacciService_GetInvalidArguments(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fibonacciUC := mock.NewMockFibonacciUseCase(ctrl)

	cfg := &config.Config{ServerCfg: &config.ServerConfig{Production: true}}
	log, _, err := logger.NewLogger(cfg)
	require.NoError(t, err)

	fibonacciS := NewGrpcFibonacciService(fibonacciUC, log)
	req := &fibonacciService.GetRequest{
		X: 0,
		Y: 0,
	}

	ctx := context.Background()
	fibonacciUC.EXPECT().GetSlice(ctx, gomock.Any(), gomock.Any()).Return(nil, models.ErrInvalidArguments)

	resp, err := fibonacciS.Get(ctx, req)
	require.Nil(t, resp)
	require.EqualError(t, err, models.ErrInvalidArguments.Error())
}

func TestFibonacciService_GetValidArguments(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fibonacciUC := mock.NewMockFibonacciUseCase(ctrl)

	cfg := &config.Config{ServerCfg: &config.ServerConfig{Production: true}}
	log, _, err := logger.NewLogger(cfg)
	require.NoError(t, err)

	fibonacciS := NewGrpcFibonacciService(fibonacciUC, log)
	req := &fibonacciService.GetRequest{
		X: 0,
		Y: 0,
	}

	ctx := context.Background()
	fibonacciUC.EXPECT().GetSlice(ctx, gomock.Any(), gomock.Any()).Return([]int64{}, nil)

	resp, err := fibonacciS.Get(ctx, req)
	require.NotNil(t, resp)
	require.NoError(t, err)
}
