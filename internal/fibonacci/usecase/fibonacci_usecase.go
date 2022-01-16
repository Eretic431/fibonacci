package usecase

import (
	"context"
	"github.com/Eretic431/fibonacci/internal/fibonacci/repository/redis"
	"github.com/Eretic431/fibonacci/internal/models"
)

type FibonacciUseCase struct {
	fr *redis.FibonacciRepository
}

func NewFibonacciUseCase(fr *redis.FibonacciRepository) *FibonacciUseCase {
	return &FibonacciUseCase{fr: fr}
}

func (f *FibonacciUseCase) GetSlice(ctx context.Context, from, to int) ([]int64, error) {
	if from < 1 || to < from {
		return nil, models.ErrInvalidArguments
	}

	prev, err := f.fr.GetLastTwoNumbers(ctx)
	if err != nil {
		return nil, err
	}

	// if true then means that we got all needed numbers in cache already calculated
	if prev[3] >= int64(to) {
		output, err := f.fr.GetInterval(ctx, from, to)
		if err != nil {
			return nil, err
		}

		return output, nil
	}

	result := make([]int64, 0, to-from+1)
	var prev1, prev2 int64
	// if true then means that we need to calculate all numbers and put in cache
	if prev[3] <= int64(from) {
		prev1, prev2 = prev[0], prev[2]
		for i := int(prev[3]) + 1; i <= to; i++ {
			prev1, prev2 = prev2, prev1+prev2
			err = f.fr.Set(ctx, i, prev2)
			if err != nil {
				return nil, err
			}
			if i >= from {
				result = append(result, prev2)
			}
		}
	}

	// if true then means we got some numbers calculated in cache and need to calculate others and put in cache
	if prev[3] < int64(to) && prev[3] > int64(from) {
		// get calculated numbers
		cache, err := f.fr.GetInterval(ctx, from, int(prev[3]))
		if err != nil {
			return nil, err
		}
		result = append(result, cache...)

		prev1, prev2 = prev[0], prev[2]
		for i := int(prev[3]) + 1; i <= to; i++ {
			prev1, prev2 = prev2, prev1+prev2
			err = f.fr.Set(ctx, i, prev2)
			if err != nil {
				return nil, err
			}
			result = append(result, prev2)
		}
	}

	return result, nil
}
