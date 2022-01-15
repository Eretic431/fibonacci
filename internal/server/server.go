package server

import (
	"github.com/Eretic431/fibonacci/config"
	grpcS "github.com/Eretic431/fibonacci/internal/fibonacci/delivery/grpc"
	httpS "github.com/Eretic431/fibonacci/internal/fibonacci/delivery/http"
	"github.com/cockroachdb/cmux"
	"go.uber.org/zap"
	"net"
)

type Server struct {
	Log         *zap.SugaredLogger
	Cfg         *config.Config
	GrpcService *grpcS.FibonacciService
	HttpService *httpS.FibonacciService
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", s.Cfg.ServerCfg.Port)
	if err != nil {
		s.Log.Errorw("could not listen", "error", err)
		return
	}

	m := cmux.New(l)
	grpcL := m.Match(cmux.HTTP2()) // consider all http2 traffic is grpc requests
	httpL := m.Match(cmux.HTTP1())

	go s.GrpcService.Serve(grpcL)
	go s.HttpService.Serve(httpL)

	if err := m.Serve(); err != nil {
		s.Log.Errorw("serving error", "error", err)
		return
	}
}
