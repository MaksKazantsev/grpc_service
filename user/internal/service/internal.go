package service

import (
	"encoding/base64"
	"fmt"
	"github.com/MaksKazantsev/grpc_service/user/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

const (
	ENV_SECRET_KEY = "SECRET_KEY"
)

func hashPass(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", &models.Error{
			Message: fmt.Sprintf("failed to generate hash from value: %v", err),
			Status:  http.StatusInternalServerError,
		}
	}

	return string(b), nil
}

func comparePass(hash, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		return &models.Error{
			Message: fmt.Sprintf("invalid password: %v", err),
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func generateToken(uuid, permlvl string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, models.Claims{
		UUID:            uuid,
		PermissionLevel: permlvl,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 21)),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv(ENV_SECRET_KEY)))
	if err != nil {
		return "", &models.Error{
			Message: fmt.Sprintf("failed to sign token: %v", err),
			Status:  http.StatusInternalServerError,
		}
	}

	return tokenString, err
}

func ParseToken(token string) (string, string, error) {
	c := models.Claims{}
	_, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(ENV_SECRET_KEY)), nil
	})
	if err != nil {
		return "", "", &models.Error{
			Message: "invalid token provided",
			Status:  http.StatusBadRequest,
		}
	}
	return c.UUID, c.PermissionLevel, nil
}

func encrypt(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func decrypt(value string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", &models.Error{
			Message: fmt.Sprintf("failed to decrypt value: %v", err),
			Status:  http.StatusInternalServerError,
		}
	}
	return string(data), nil
}
