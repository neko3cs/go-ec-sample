package service

import (
	"go-ec-sample/domain"
	"go-ec-sample/query"
)

type LoginService struct{}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) Authenticate(email, password string) (bool, *domain.User) {
	q := query.NewGetUserByEmailQuery(email)
	h := query.NewGetUserByEmailQueryHandler()
	user, err := h.Handle(q)
	if err != nil || user.Password() != password {
		return false, nil
	}
	return true, user
}
