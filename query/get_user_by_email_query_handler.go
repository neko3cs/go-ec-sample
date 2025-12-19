package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"

	"gorm.io/gorm"
)

type GetUserByEmailQueryHandler struct {
	db *gorm.DB
}

func NewGetUserByEmailQueryHandler(db *gorm.DB) *GetUserByEmailQueryHandler {
	return &GetUserByEmailQueryHandler{db: db}
}

func (h *GetUserByEmailQueryHandler) Handle(query *GetUserByEmailQuery) (*domain.User, error) {
	u := &db.User{}
	if err := h.db.Where("email = ?", query.Email).First(u).Error; err != nil {
		return nil, err
	}

	user := domain.NewUser(u.Id, u.Name, u.Email, u.Password, u.IsAdmin)
	return user, nil
}
