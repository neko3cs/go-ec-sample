package command

import (
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type CreateCartCommandHandler struct {
	db *gorm.DB
}

func NewCreateCartCommandHandler(db *gorm.DB) *CreateCartCommandHandler {
	return &CreateCartCommandHandler{db: db}
}

func (h *CreateCartCommandHandler) Handle(cmd *CreateCartCommand) error {
	cart := db.Cart{
		UserId: cmd.UserId,
	}
	return h.db.Create(&cart).Error
}
