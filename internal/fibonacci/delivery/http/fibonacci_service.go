package http

import (
	"errors"
	"github.com/Eretic431/fibonacci/internal/fibonacci/usecase"
	"github.com/Eretic431/fibonacci/internal/models"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type FibonacciService struct {
	fibonacciUC usecase.FibonacciUseCase
	log         *zap.SugaredLogger
}

func (fs *FibonacciService) GetHandler(w http.ResponseWriter, r *http.Request) {
	from, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		fs.badRequestResponse(w, r, err)
		return
	}

	to, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		fs.badRequestResponse(w, r, err)
		return
	}

	numbers, err := fs.fibonacciUC.GetSlice(from, to)
	if err != nil {
		if errors.Is(err, models.ErrInvalidArguments) {
			fs.badRequestResponse(w, r, err)
			return
		}
		fs.serverErrorResponse(w, r, err)
		return
	}

	output := &struct {
		Numbers []int64 `json:"numbers"`
	}{numbers}

	if err := fs.writeJSON(w, http.StatusOK, output); err != nil {
		fs.serverErrorResponse(w, r, err)
		return
	}
}
