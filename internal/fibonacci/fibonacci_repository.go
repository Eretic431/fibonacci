//go:generate mockgen -source fibonacci_repository.go -destination mock/fibonacci_repository.go -package mock
package fibonacci

import "context"

type FibonacciRepository interface {
	Get(ctx context.Context, key int) (int64, error)
	Set(ctx context.Context, key int, value int64) error
	GetLastTwoNumbers(ctx context.Context) ([]int64, error)
	GetInterval(ctx context.Context, from, to int) ([]int64, error)
}
