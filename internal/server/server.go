package server

import (
	"fmt"
	auth "github.com/dstan05/auth/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const port = 1212

type Server struct {
	grps *grpc.Server
}

func Init() (Server, error) {
	server := Server{}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return server, err
	}

	server.grps = grpc.NewServer()
	reflection.Register(server.grps)
	auth.RegisterAuthServer(server.grps, &Routes{})

	if err = server.grps.Serve(lis); err != nil {
		return server, err
	}

	return server, nil
}

func (server *Server) Stop() Server {
	server.grps.Stop()
	return *server
}
