package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	auth "github.com/dstan05/auth/pkg/v1"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net"
)

type server struct {
	auth.UnimplementedAuthServer
}

func (s *server) Get(ctx context.Context, req *auth.GetRequest) (*auth.GetResponse, error) {
	u := auth.GetResponse{
		Id:        req.Id,
		Name:      wrapperspb.String(gofakeit.Name()),
		Email:     wrapperspb.String(gofakeit.Email()),
		Role:      auth.Role_ADMIN,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}
	return &u, nil
}

func (s *server) Update(ctx context.Context, req *auth.UpdateRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:2231")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	auth.RegisterAuthServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println(color.GreenString("Hello, world!"))
}
