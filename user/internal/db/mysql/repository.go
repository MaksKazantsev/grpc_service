package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaksKazantsev/grpc_service/user/internal/db"
	"github.com/MaksKazantsev/grpc_service/user/internal/models"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func NewRepository(db *sqlx.DB) db.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sqlx.DB
}

const ROLE_USER = "user"

func (r repository) Register(ctx context.Context, req models.RegisterReq) error {
	q := `INSERT INTO users (uuid, username, password, permlvl, email, phone_number) VALUES(?,?,?,?,?,?)`

	if err := r.db.QueryRowx(q, req.UUID, req.Username, req.Password, ROLE_USER, req.Email, req.PhoneNumber).Err(); err != nil {
		return &models.Error{
			Message: fmt.Sprintf("failed to query raw register: %v", err),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (r repository) Login(ctx context.Context, email string) (db.LoginInfo, error) {
	q := `SELECT uuid, email, password, permlvl FROM users WHERE email = ? LIMIT 1`

	var info db.LoginInfo

	err := r.db.QueryRowx(q, email).StructScan(&info)
	{
		if errors.Is(sql.ErrNoRows, err) {
			return db.LoginInfo{}, &models.Error{
				Message: fmt.Sprintf("user not found: %v", err),
				Status:  http.StatusNotFound,
			}
		}
		if err != nil {
			return db.LoginInfo{}, &models.Error{
				Message: fmt.Sprintf("failed to query: %v", err),
				Status:  http.StatusInternalServerError,
			}
		}

	}
	return info, nil
}

func (r repository) GetUserInfo(ctx context.Context, uuid string) (models.UserInfoRes, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) SwitchNotificationsStatus(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetUserByUUID(ctx context.Context, uuids string) (db.EmailNotificationsInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) CheckIfAuthorized(ctx context.Context, uuid string, role string) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) ResetPassword(ctx context.Context, uuid string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) GetPassword(ctx context.Context, uuid string) (string, error) {
	//TODO implement me
	panic("implement me")
}
