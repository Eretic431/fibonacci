package main

import (
	"fmt"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
)

func main() {
	uc := &usecase.FibonacciUseCase{}
	f, err := uc.GetSlice(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
