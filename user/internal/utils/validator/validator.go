package validator

import (
	users "github.com/MaksKazantsev/grpc_service/proto/gen/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

const (
	ERR_INVALID_EMAIL    string = "invalid email"
	ERR_INVALID_PASSWORD string = "invalid password"
	ERR_INVALID_PHONE    string = "invalid phone"
	ERR_INVALID_USERNAME string = "invalid user"
)

type Validator interface {
	ValidateRegisterReq(req *users.RegisterReq) error
	ValidateLoginReq(req *users.LoginReq) error
	ValidateSwitchNotificationsStatusReq(req *users.SwitchNotificationsStatusReq) error
	ValidateCheckIfAuthorizedReq(req *users.CheckIfAuthorizedReq) error
	ValidateResetPasswordReq(req *users.ResetPasswordReq) error
}

func NewValidator() Validator {
	return &validator{
		regExpEmail: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
		regExpPhone: regexp.MustCompile(`^[1-9]\d{9}$`),
	}
}

type validator struct {
	regExpEmail *regexp.Regexp
	regExpPhone *regexp.Regexp
}

func (v validator) ValidateLoginReq(req *users.LoginReq) error {
	//TODO implement me
	panic("implement me")
}

func (v validator) ValidateSwitchNotificationsStatusReq(req *users.SwitchNotificationsStatusReq) error {
	//TODO implement me
	panic("implement me")
}

func (v validator) ValidateCheckIfAuthorizedReq(req *users.CheckIfAuthorizedReq) error {
	//TODO implement me
	panic("implement me")
}

func (v validator) ValidateResetPasswordReq(req *users.ResetPasswordReq) error {
	//TODO implement me
	panic("implement me")
}

func (v validator) ValidateRegisterReq(req *users.RegisterReq) error {
	if ok := v.regExpEmail.MatchString(req.Email); !ok {
		return status.Error(codes.InvalidArgument, ERR_INVALID_EMAIL)
	}
	if ok := v.regExpPhone.MatchString(req.PhoneNumber); !ok {
		return status.Error(codes.InvalidArgument, ERR_INVALID_PHONE)
	}
	if len(req.Username) < 3 {
		return status.Error(codes.InvalidArgument, ERR_INVALID_USERNAME)
	}
	if err := validatePassword(req.Password); err != nil {
		return err
	}
	return nil
}

func validatePassword(pass string) error {
	if len(pass) < 7 || len(pass) > 40 {
		return status.Error(codes.InvalidArgument, ERR_INVALID_PASSWORD)
	}
	return nil
}
