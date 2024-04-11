package server

import (
	"context"
	users "github.com/MaksKazantsev/grpc_service/proto/gen/protos"
	"github.com/MaksKazantsev/grpc_service/user/internal/log"
	"github.com/MaksKazantsev/grpc_service/user/internal/service"
	"github.com/MaksKazantsev/grpc_service/user/internal/utils/converter"
	"github.com/MaksKazantsev/grpc_service/user/internal/utils/validator"
	"google.golang.org/grpc"
)

type Params struct {
	service service.Service
}

func RegisterGRPCServer(s *grpc.Server, p Params) {
	users.RegisterUserServer(s, newServer(p))
}

func newServer(p Params) users.UserServer {
	return &server{
		log:       log.GetLogger(),
		validator: validator.NewValidator(),
		converter: converter.NewConverter(),
		service:   p.service,
	}
}

type server struct {
	users.UnimplementedUserServer

	log log.Logger

	service service.Service

	validator validator.Validator

	converter converter.Converter
}

func (s *server) Register(ctx context.Context, req *users.RegisterReq) (*users.RegisterRes, error) {
	if err := s.validator.ValidateRegisterReq(req); err != nil {
		return nil, err
	}

	res, err := s.service.Register(ctx, s.converter.RegisterReqToService(req))
	if err != nil {
		return nil, s.handleError(err)
	}

	return s.converter.RegisterResToPb(res), nil
}

func (s *server) Login(ctx context.Context, req *users.LoginReq) (*users.LoginRes, error) {
	if err := s.validator.ValidateLoginReq(req); err != nil {
		return nil, err
	}

	token, err := s.service.Login(ctx, s.converter.LoginReqToService(req))
	if err != nil {
		return nil, s.handleError(err)
	}

	return s.converter.LoginResToPb(token), nil
}
