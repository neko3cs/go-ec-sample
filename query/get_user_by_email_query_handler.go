package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type GetUserByEmailQueryHandler struct{}

func NewGetUserByEmailQueryHandler() *GetUserByEmailQueryHandler {
	return &GetUserByEmailQueryHandler{}
}

func (h *GetUserByEmailQueryHandler) Handle(query *GetUserByEmailQuery) (*domain.User, error) {
	u := &db.User{}
	if err := db.GetDB().Where("email = ?", query.Email).First(u).Error; err != nil {
		return nil, err
	}

	user := domain.NewUser(u.Id, u.Name, u.Email, u.Password, u.IsAdmin)
	return user, nil
}
