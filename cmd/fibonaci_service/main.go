package main

import (
	"fmt"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
)

func main() {
	uc := &usecase.FibonacciUseCase{}
	fmt.Println(uc.GetSlice(1, 10))
}
