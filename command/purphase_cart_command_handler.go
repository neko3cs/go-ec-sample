package command

import (
	"errors"
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type PurchaseCartCommandHandler struct {
	db *gorm.DB
}

func NewPurchaseCartCommandHandler(db *gorm.DB) *PurchaseCartCommandHandler {
	return &PurchaseCartCommandHandler{db: db}
}

func (h *PurchaseCartCommandHandler) Handle(cmd *PurchaseCartCommand) error {
	var product db.Product
	err := h.db.First(&product, cmd.ProductId).Error
	if err != nil {
		return err
	}
	if product.Stock < cmd.Quantity {
		return errors.New("在庫不足")
	}
	product.Stock -= cmd.Quantity
	return h.db.Save(&product).Error
}
