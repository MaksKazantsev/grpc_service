package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/grpc_service/user/internal/db"
	"github.com/MaksKazantsev/grpc_service/user/internal/models"
	"github.com/google/uuid"
	"net/http"
)

type Service interface {
	Register(ctx context.Context, req models.RegisterReq) (models.RegisterRes, error)
	Login(ctx context.Context, req models.LoginReq) (string, error)
	SwitchNotificationsStatus(ctx context.Context, uuid string) error
	CheckIfAuthorized(ctx context.Context, token string) (string, error)
	ResetPassword(ctx context.Context, req models.ResetPasswordReq) error
}

type Params struct {
	repo db.Repository
}

func NewService(p Params) Service {
	return &service{
		repo: p.repo,
	}
}

type service struct {
	repo db.Repository
}

const ROLE_USER string = "user"

func (s service) Register(ctx context.Context, req models.RegisterReq) (models.RegisterRes, error) {
	req.UUID = uuid.New().String()

	_, err := s.repo.Login(ctx, req.Email)
	if err == nil {
		return models.RegisterRes{}, &models.Error{
			Message: fmt.Sprintf("user with email: %s already exists", req.Email),
			Status:  http.StatusBadRequest,
		}
	}

	passHash, err := hashPass(req.Password)

	req.Password = passHash
	req.PhoneNumber = encrypt(req.PhoneNumber)

	if err != nil {
		return models.RegisterRes{}, fmt.Errorf("failed to hash password: %w", err)
	}

	if err = s.repo.Register(ctx, req); err != nil {
		return models.RegisterRes{}, err
	}

	token, err := generateToken(req.UUID, ROLE_USER)
	if err != nil {
		return models.RegisterRes{}, fmt.Errorf("failed to generate new token: %w", err)
	}

	return models.RegisterRes{UUID: req.UUID, Token: token}, nil
}

func (s service) Login(ctx context.Context, req models.LoginReq) (string, error) {
	info, err := s.repo.Login(ctx, req.Email)
	if err != nil {
		return "", nil
	}
	if err = comparePass(info.Password, req.Password); err != nil {
		return "", err
	}
	token, err := generateToken(info.UUID, info.PermissionLvl)
	if err != nil {
		return "", fmt.Errorf("failed to generate new token: %w", err)
	}
	return token, nil
}

func (s service) SwitchNotificationsStatus(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s service) CheckIfAuthorized(ctx context.Context, token string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) ResetPassword(ctx context.Context, req models.ResetPasswordReq) error {
	//TODO implement me
	panic("implement me")
}
