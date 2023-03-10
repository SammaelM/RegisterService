package server

import (
	"context"
	"mainmod/api/grpc/proto"
	"mainmod/internal/response"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	httpServer *http.Server
}

//http

func (s *Server) Run(port string, handler http.HandlerFunc) error {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 25, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

//grpc

func RegServ() *grpc.Server {
	s := grpc.NewServer()
	srv := response.GRPServer{}
	proto.RegisterRegistrServer(s, srv)
	return s
}
