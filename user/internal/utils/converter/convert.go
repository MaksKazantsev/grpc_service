package converter

import (
	users "github.com/MaksKazantsev/grpc_service/proto/gen/protos"
	"github.com/MaksKazantsev/grpc_service/user/internal/models"
)

type Converter interface {
	ToService
	ToPb
}

type ToService interface {
	RegisterReqToService(req *users.RegisterReq) models.RegisterReq
	LoginReqToService(req *users.LoginReq) models.LoginReq
	ResetPasswordReqToService(req *users.ResetPasswordReq) models.ResetPasswordReq
}

type ToPb interface {
	RegisterResToPb(res models.RegisterRes) *users.RegisterRes
	LoginResToPb(token string) *users.LoginRes
	CheckIfAuthorizedResToPb(role string) *users.CheckIfAuthorizedRes
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct {
}

func (c converter) RegisterReqToService(req *users.RegisterReq) models.RegisterReq {
	//TODO implement me
	panic("implement me")
}

func (c converter) LoginReqToService(req *users.LoginReq) models.LoginReq {
	//TODO implement me
	panic("implement me")
}

func (c converter) ResetPasswordReqToService(req *users.ResetPasswordReq) models.ResetPasswordReq {
	//TODO implement me
	panic("implement me")
}

func (c converter) RegisterResToPb(res models.RegisterRes) *users.RegisterRes {
	//TODO implement me
	panic("implement me")
}

func (c converter) LoginResToPb(token string) *users.LoginRes {
	//TODO implement me
	panic("implement me")
}

func (c converter) CheckIfAuthorizedResToPb(role string) *users.CheckIfAuthorizedRes {
	//TODO implement me
	panic("implement me")
}
