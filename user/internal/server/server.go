package server

import (
	"context"
	users "github.com/MaksKazantsev/grpc_service/proto/gen/protos"
	"github.com/MaksKazantsev/grpc_service/user/internal/log"
	"github.com/MaksKazantsev/grpc_service/user/internal/utils/validator"
	"google.golang.org/grpc"
)

type Params struct {
}

func RegisterGRPCServer(s *grpc.Server, p Params) {
	users.RegisterUserServer(s, newServer(p))
}

func newServer(p Params) users.UserServer {
	return &server{
		log:       log.GetLogger(),
		validator: validator.NewValidator(),
	}
}

type server struct {
	users.UnimplementedUserServer

	log log.Logger

	validator validator.Validator
}

func (s *server) Register(ctx context.Context, req *users.RegisterReq) (*users.RegisterRes, error) {
	if err := s.validator.ValidateRegisterReq(req); err != nil {
		return nil, err
	}

	return &users.RegisterRes{}, nil
}
