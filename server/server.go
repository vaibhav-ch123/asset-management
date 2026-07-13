package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/vaibhav-ch123/asset-management/handlers"
	"github.com/vaibhav-ch123/asset-management/middlewares"
)

type Server struct {
	router http.Handler
	server *http.Server
}

const (
	ReadTimeout       = 5 * time.Minute
	ReadHeaderTimeout = 30 * time.Second
	WriteTimeout      = 5 * time.Minute
)

func SetupRoute() *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "server healthy!")
	})

	mux.HandleFunc("POST /v1/register", handlers.RegisterEmployee)
	mux.HandleFunc("POST /v1/login", handlers.LoginEmployee)

	mux.Handle("GET /v1/employee", middlewares.AuthMiddleWare(middlewares.ShouldHaveAdmin(http.HandlerFunc(handlers.GetEmployees))))

	return &Server{
		router: middlewares.ContentTypeMiddleware(middlewares.PanicRecoveryMiddleware(mux)),
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
