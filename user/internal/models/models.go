package models

import "github.com/golang-jwt/jwt/v4"

type RegisterReq struct {
	UUID        string
	Username    string
	Password    string
	Email       string
	PhoneNumber string
}

type RegisterRes struct {
	UUID  string
	Token string
}

type LoginReq struct {
	Email    string
	Password string
}

type Claims struct {
	UUID            string `json:"uuid"`
	PermissionLevel string `json:"role"`
	*jwt.RegisteredClaims
}

type UserInfoRes struct {
	Username        string
	NotificationsOn bool
	Email           string
}

type ResetPasswordReq struct {
	OldPassword string
	NewPassword string
	Token       string
}
