package server

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	auth "github.com/dstan05/auth/pkg/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Routes struct {
	auth.UnimplementedAuthServer
}

func (s *Routes) Get(ctx context.Context, req *auth.GetRequest) (*auth.GetResponse, error) {
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

func (s *Routes) Create(ctx context.Context, req *auth.CreateRequest) (*auth.CreateResponse, error) {
	return &auth.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *Routes) Update(ctx context.Context, req *auth.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Routes) Delete(ctx context.Context, req *auth.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
