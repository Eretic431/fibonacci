package http

import (
	"context"
	"fmt"
	"github.com/Eretic431/fibonacci/config"
	"github.com/Eretic431/fibonacci/internal/fibonacci/mock"
	"github.com/Eretic431/fibonacci/internal/models"
	"github.com/Eretic431/fibonacci/pkg/logger"
	"github.com/Eretic431/fibonacci/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestFibonacciService_GetHandlerInvalidQueries(t *testing.T) {
	queries := []struct {
		x string
		y string
	}{
		{
			"asd", "asd",
		},
		{
			"asd", "2",
		},
		{
			"1", "asd",
		},
		{
			"", "",
		},
		{
			"1", "",
		},
		{
			"", "2",
		},
	}

	for i, query := range queries {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fibonacciUC := mock.NewMockFibonacciUseCase(ctrl)

			cfg := &config.Config{ServerCfg: &config.ServerConfig{Production: true}}
			log, _, err := logger.NewLogger(cfg)
			require.NoError(t, err)

			httpS := NewHttpFibonacciService(fibonacciUC, log, utils.NewHttpHelper(log))

			values := url.Values{}
			values.Add("x", query.x)
			values.Add("y", query.y)

			require.HTTPStatusCode(t, httpS.GetHandler, http.MethodGet, "/fibonacci", values, http.StatusBadRequest)
			require.HTTPBodyContains(t, httpS.GetHandler, http.MethodGet, "/fibonacci", values, "invalid syntax")
		})
	}
}

func TestFibonacciService_GetHandlerInvalidArguments(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fibonacciUC := mock.NewMockFibonacciUseCase(ctrl)

	cfg := &config.Config{ServerCfg: &config.ServerConfig{Production: true}}
	log, _, err := logger.NewLogger(cfg)
	require.NoError(t, err)

	httpS := NewHttpFibonacciService(fibonacciUC, log, utils.NewHttpHelper(log))

	endpoint := "/fibonacci" // valid query params to pass atoi
	var values = url.Values{}
	values.Set("x", "1")
	values.Set("y", "2")

	// we're doing two requests with same input data so we expect two calls of GetSlice()
	fibonacciUC.EXPECT().GetSlice(context.Background(), gomock.Any(), gomock.Any()).Return(nil, models.ErrInvalidArguments)
	fibonacciUC.EXPECT().GetSlice(context.Background(), gomock.Any(), gomock.Any()).Return(nil, models.ErrInvalidArguments)

	require.HTTPStatusCode(t, httpS.GetHandler, http.MethodGet, endpoint, values, http.StatusBadRequest)
	require.HTTPBodyContains(t, httpS.GetHandler, http.MethodGet, endpoint, values, models.ErrInvalidArguments.Error())
}

func TestFibonacciService_GetHandlerValidArguments(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fibonacciUC := mock.NewMockFibonacciUseCase(ctrl)

	cfg := &config.Config{ServerCfg: &config.ServerConfig{Production: true}}
	log, _, err := logger.NewLogger(cfg)
	require.NoError(t, err)

	httpS := NewHttpFibonacciService(fibonacciUC, log, utils.NewHttpHelper(log))

	endpoint := "/fibonacci" // valid query params to pass atoi
	var values = url.Values{}
	values.Set("x", "1")
	values.Set("y", "2")

	// we're doing two requests with same input data so we expect two calls of GetSlice()
	fibonacciUC.EXPECT().GetSlice(context.Background(), gomock.Any(), gomock.Any()).Return([]int64{}, nil)
	fibonacciUC.EXPECT().GetSlice(context.Background(), gomock.Any(), gomock.Any()).Return([]int64{}, nil)

	require.HTTPStatusCode(t, httpS.GetHandler, http.MethodGet, endpoint, values, http.StatusOK)
	require.HTTPBodyContains(t, httpS.GetHandler, http.MethodGet, endpoint, values, "numbers")
}
