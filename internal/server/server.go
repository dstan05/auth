package server

import (
	"fmt"
	"github.com/dstan05/auth/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const port = 1212

type Server struct {
	grps     *grpc.Server
	listener net.Listener
}

func Init() (Server, error) {
	s := Server{}
	s.grps = grpc.NewServer()

	reflection.Register(s.grps)
	auth.RegisterAuthServer(s.grps, &Routes{})

	return s, nil
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	s.listener = lis
	if err = s.grps.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() (Server, error) {
	s.grps.Stop()
	if err := s.listener.Close(); err != nil {
		return *s, err
	}

	return *s, nil
}
