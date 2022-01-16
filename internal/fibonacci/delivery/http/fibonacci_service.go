package http

import (
	"errors"
	"github.com/Eretic431/fibonacci/internal/fibonacci"
	"github.com/Eretic431/fibonacci/internal/models"
	"github.com/Eretic431/fibonacci/pkg/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strconv"
)

type FibonacciService struct {
	fibonacciUC fibonacci.FibonacciUseCase
	log         *zap.SugaredLogger
	httpH       *utils.HttpHelper
}

func NewHttpFibonacciService(fuc fibonacci.FibonacciUseCase, log *zap.SugaredLogger, httpH *utils.HttpHelper) *FibonacciService {
	return &FibonacciService{fibonacciUC: fuc, log: log, httpH: httpH}
}

func (fs *FibonacciService) GetHandler(w http.ResponseWriter, r *http.Request) {
	from, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		fs.httpH.BadRequestResponse(w, r, err)
		return
	}

	to, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		fs.httpH.BadRequestResponse(w, r, err)
		return
	}

	numbers, err := fs.fibonacciUC.GetSlice(r.Context(), from, to)
	if err != nil {
		if errors.Is(err, models.ErrInvalidArguments) {
			fs.httpH.BadRequestResponse(w, r, err)
			return
		}
		fs.httpH.ServerErrorResponse(w, r, err)
		return
	}

	output := &struct {
		Numbers []int64 `json:"numbers"`
	}{numbers}

	if err := fs.httpH.WriteJSON(w, http.StatusOK, output); err != nil {
		fs.httpH.ServerErrorResponse(w, r, err)
		return
	}
}

func (fs *FibonacciService) Serve(l net.Listener) {
	server := &http.Server{Handler: fs.router()}
	fs.log.Infof("Starting http service")
	if err := server.Serve(l); err != nil {
		fs.log.Errorw("error serving http", "error", err)
		return
	}
}

func (fs *FibonacciService) router() http.Handler {
	r := chi.NewRouter()
	r.Get("/fibonacci", fs.GetHandler)

	return r
}
