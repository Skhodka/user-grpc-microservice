package grpcserv

import (
	"fmt"
	"log/slog"
	"net"
	"time"
	"usermic/internal/transport/grpc/handlers"
	userv1 "usermic/internal/transport/grpc/pb"
	"usermic/internal/usecase/registration"

	"google.golang.org/grpc"
)

type Server struct {
	log    *slog.Logger
	server *grpc.Server
}

func NewServer(log *slog.Logger, regUc *registration.RegistrationUC, timeout time.Duration) *Server {
	serv := grpc.NewServer()
	handl := handlers.NewUserHandler(log, regUc, timeout)
	userv1.RegisterUserServiceServer(serv, handl)

	return &Server{
		log:    log,
		server: serv,
	}
}

func (s *Server) MustStart(port int) {
	s.log.Info("starting server", slog.Int("port", port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic("failed to start listening")
	}

	defer l.Close()

	if err := s.server.Serve(l); err != nil {
		panic("failed serve")
	}
}

func (s *Server) Stop() {
	s.log.Info("starting graceful stop")
	s.server.GracefulStop()
}
