package server

import (
	grpcS "github.com/Eretic431/fibonacci/internal/fibonacci/delivery/grpc"
	httpS "github.com/Eretic431/fibonacci/internal/fibonacci/delivery/http"
	"github.com/cockroachdb/cmux"
	"go.uber.org/zap"
	"net"
)

type Server struct {
	log         *zap.SugaredLogger
	grpcService *grpcS.FibonacciService
	httpService *httpS.FibonacciService
}

func (s *Server) Run() {
	// TODO: get port from config
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		s.log.Errorw("could not listen", "error", err)
		return
	}

	m := cmux.New(l)

	if err := m.Serve(); err != nil {
		s.log.Errorw("serving error", "error", err)
		return
	}
}
