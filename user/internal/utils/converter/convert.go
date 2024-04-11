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
	return models.RegisterReq{
		Username:    req.Username,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
}

func (c converter) RegisterResToPb(res models.RegisterRes) *users.RegisterRes {
	return &users.RegisterRes{
		UUID:  res.UUID,
		Token: res.Token,
	}
}

func (c converter) LoginReqToService(req *users.LoginReq) models.LoginReq {
	return models.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c converter) LoginResToPb(token string) *users.LoginRes {
	return &users.LoginRes{
		Token: token,
	}
}

func (c converter) ResetPasswordReqToService(req *users.ResetPasswordReq) models.ResetPasswordReq {
	panic("implement me")
}
func (c converter) CheckIfAuthorizedResToPb(role string) *users.CheckIfAuthorizedRes {
	//TODO implement me
	panic("implement me")
}
