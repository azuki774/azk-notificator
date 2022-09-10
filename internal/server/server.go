package server

import (
	"azk-notificator/internal/model"
	"azk-notificator/internal/telemetry"
	"context"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	Logger      *zap.Logger
	QueueClient QueueClient
	Host        string
	Port        string
}

func (s *Server) NewHTTPServer() (srv *http.Server, err error) {
	r := NewHandler()
	srv = &http.Server{
		Handler:      r,
		Addr:         net.JoinHostPort(s.Host, s.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv, nil
}

func (s *Server) Start() (err error) {
	s.Logger.Info("server start")
	ServerForHandler = s

	httpSrv, err := s.NewHTTPServer()
	if err != nil {
		s.Logger.Error("failed to create HTTP Server", zap.Error(err))
	}

	err = httpSrv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.Logger.Error("failed to shutdown HTTP server", zap.Error(err))
	}

	s.Logger.Info("HTTP server closed")
	return nil
}

func (s *Server) Enqueue(ctx context.Context, q model.Queue) (err error) {
	l := telemetry.LoggerWithSpanID(ctx, s.Logger)
	err = s.QueueClient.Push(ctx, q)
	if err != nil {
		l.Error("failed to enqueue", zap.Error(err))
		return err
	}

	l.Info("enqueue")
	return nil
}
