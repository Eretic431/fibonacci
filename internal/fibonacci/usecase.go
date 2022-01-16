//go:generate mockgen -source usecase.go -destination mock/usecase.go -package mock
package fibonacci

import "context"

type FibonacciUseCase interface {
	GetSlice(ctx context.Context, from, to int) ([]int64, error)
}
