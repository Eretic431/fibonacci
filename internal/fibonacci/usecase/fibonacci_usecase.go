package usecase

import "github.com/Eretic431/fibonacci/internal/models"

type FibonacciUseCase struct {
}

func NewFibonacciUseCase() *FibonacciUseCase {
	return &FibonacciUseCase{}
}

func (f *FibonacciUseCase) GetSlice(from, to int) ([]int64, error) {
	if from < 1 || to < from {
		return nil, models.ErrInvalidArguments
	}

	result := make([]int64, 0, to-from+1)
	var prev1, prev2 int64 = 0, 1

	if from == 1 {
		result = append(result, prev1)
	}
	if to >= 2 && from <= 2 {
		result = append(result, prev2)
	}

	for i := 3; i <= to; i++ {
		prev1, prev2 = prev2, prev1+prev2
		if i >= from {
			result = append(result, prev2)
		}
	}

	return result, nil
}
