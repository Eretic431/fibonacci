package grpc

import (
	"context"
	fibonacciService "github.com/Eretic431/fibonacci/internal/fibonacci/proto"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
)

type FibonacciService struct {
	fibonacciUC usecase.FibonacciUseCase
}

func (fs *FibonacciService) Get(ctx context.Context, r *fibonacciService.GetRequest) (*fibonacciService.GetResponse, error) {
	numbers, err := fs.fibonacciUC.GetSlice(int(r.GetX()), int(r.GetY()))
	if err != nil {
		return nil, err
	}

	return &fibonacciService.GetResponse{Numbers: numbers}, nil
}
