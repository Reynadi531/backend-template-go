package auth

import (
	"backend-template-go/internal/entities/web"
)

type AuthService interface {
	Register(name string, email string, password string) web.Response
	Login(email string, password string) web.Response
	Refresh(token string, userId string) web.Response
}
