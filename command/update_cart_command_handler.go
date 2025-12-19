package command

import (
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type UpdateCartCommandHandler struct {
	db *gorm.DB
}

func NewUpdateCartCommandHandler(db *gorm.DB) *UpdateCartCommandHandler {
	return &UpdateCartCommandHandler{db: db}
}

func (h *UpdateCartCommandHandler) Handle(cmd *UpdateCartCommand) error {
	err := h.db.
		Where("cart_id = ?", cmd.CartId).
		Delete(&db.CartItem{}).
		Error
	if err != nil {
		return err
	}

	for _, item := range cmd.Items {
		cartItem := db.CartItem{
			CartId:    cmd.CartId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
		err := h.db.Create(&cartItem).Error
		if err != nil {
			return err
		}
	}

	return nil
}
