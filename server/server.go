package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	router *http.ServeMux
	server *http.Server
}

const (
	ReadTimeout       = 5 * time.Minute
	ReadHeaderTimeout = 30 * time.Second
	WriteTimeout      = 5 * time.Minute
)

func SetupRoute() *Server {

	r := http.NewServeMux()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "server healthy!")
	})

	return &Server{
		router: r,
	}
}

func (s *Server) Run(port string) error {

	s.server = &http.Server{
		Addr:              port,
		Handler:           s.router,
		ReadTimeout:       ReadTimeout,
		ReadHeaderTimeout: ReadHeaderTimeout,
		WriteTimeout:      WriteTimeout,
	}

	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *Server) ShutDownServer(timeOut time.Duration) error {

	ctx, cancelCtx := context.WithTimeout(context.Background(), timeOut)

	defer cancelCtx()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
