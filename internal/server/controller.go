package server

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dstan05/auth/pkg/auth"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Routes struct {
	auth.UnimplementedAuthServer
}

func (r *Routes) Get(_ context.Context, req *auth.GetRequest) (*auth.GetResponse, error) {
	u := auth.GetResponse{
		Id:        req.Id,
		Name:      wrapperspb.String(gofakeit.Name()),
		Email:     wrapperspb.String(gofakeit.Email()),
		Role:      auth.Role(gofakeit.RandomInt([]int{1, 2})),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	return &u, nil
}

func (r *Routes) Create(_ context.Context, _ *auth.CreateRequest) (*auth.CreateResponse, error) {
	return &auth.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (r *Routes) Update(_ context.Context, _ *auth.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (r *Routes) Delete(_ context.Context, _ *auth.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
