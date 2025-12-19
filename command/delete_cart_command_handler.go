package command

import (
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type DeleteCartCommandHandler struct {
	db *gorm.DB
}

func NewDeleteCartCommandHandler(db *gorm.DB) *DeleteCartCommandHandler {
	return &DeleteCartCommandHandler{db: db}
}

func (h *DeleteCartCommandHandler) Handle(cmd *DeleteCartCommand) error {
	err := h.db.Delete(&db.CartItem{}, "cart_id = ?", cmd.CartId).Error
	if err != nil {
		return err
	}
	err = h.db.Delete(&db.Cart{}, "cart_id = ?", cmd.CartId).Error
	if err != nil {
		return err
	}
	return nil
}
