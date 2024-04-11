package db

import (
	"context"
	"github.com/MaksKazantsev/grpc_service/user/internal/models"
)

type Repository interface {
	Register(ctx context.Context, req models.RegisterReq) error
	Login(ctx context.Context, email string) (LoginInfo, error)
	GetUserInfo(ctx context.Context, uuid string) (models.UserInfoRes, error)
	SwitchNotificationsStatus(ctx context.Context, uuid string) error
	GetUserByUUID(ctx context.Context, uuids string) (EmailNotificationsInfo, error)
	CheckIfAuthorized(ctx context.Context, uuid string, role string) error
	ResetPassword(ctx context.Context, uuid string, password string) error
	GetPassword(ctx context.Context, uuid string) (string, error)
}

type LoginInfo struct {
	UUID          string
	Email         string
	Password      string
	PermissionLvl string
}

type EmailNotificationsInfo struct {
	Email    string
	Username string
}
