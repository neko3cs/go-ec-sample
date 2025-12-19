package service

import (
	"errors"
	"go-ec-sample/db"
	"go-ec-sample/domain"
	"go-ec-sample/query"

	"gorm.io/gorm"
)

type LoginService struct{}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) Authenticate(email, password string) (bool, *domain.User) {
	var user *domain.User

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		q := query.NewGetUserByEmailQuery(email)
		h := query.NewGetUserByEmailQueryHandler(tx)
		u, err := h.Handle(q)
		if err != nil {
			return err
		}
		if u.Password() != password {
			return errors.New("invalid password")
		}
		user = u
		return nil
	})

	if err != nil {
		return false, nil
	}
	return true, user
}
