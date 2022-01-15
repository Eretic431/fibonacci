package grpc

import (
	"context"
	"errors"
	fibonacciService "github.com/Eretic431/fibonacci/internal/fibonacci/proto"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
	"github.com/Eretic431/fibonacci/internal/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type FibonacciService struct {
	fibonacciUC *usecase.FibonacciUseCase
	log         *zap.SugaredLogger
}

func NewGrpcFibonacciService(fuc *usecase.FibonacciUseCase, log *zap.SugaredLogger) *FibonacciService {
	return &FibonacciService{fibonacciUC: fuc, log: log}
}

func (fs *FibonacciService) Get(ctx context.Context, r *fibonacciService.GetRequest) (*fibonacciService.GetResponse, error) {
	numbers, err := fs.fibonacciUC.GetSlice(int(r.GetX()), int(r.GetY()))
	if err != nil {
		if errors.Is(err, models.ErrInvalidArguments) {
			fs.log.Debugw("client error", "error", err)
			return nil, err
		}
		fs.log.Errorw("server error", "error", err)

		return nil, err
	}

	return &fibonacciService.GetResponse{Numbers: numbers}, nil
}

func (fs *FibonacciService) Serve(l net.Listener) {
	s := grpc.NewServer()
	fibonacciService.RegisterFibonacciServiceServer(s, fs)
	if err := s.Serve(l); err != nil {
		fs.log.Errorw("error serving grpc", "error", err)
	}
}
