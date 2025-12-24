package command

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"

	"gorm.io/gorm"
)

type LoadProductCommandHandler struct {
	db *gorm.DB
}

func NewLoadProductCommandHandler(db *gorm.DB) *LoadProductCommandHandler {
	return &LoadProductCommandHandler{db: db}
}

func (h *LoadProductCommandHandler) Handle(cmd *LoadProductCommand) (*domain.Product, error) {
	var dbProduct db.Product
	err := h.db.First(&dbProduct, cmd.ProductId).Error
	if err != nil {
		return nil, err
	}
	product := domain.NewProduct(dbProduct.Id, dbProduct.Name, dbProduct.Price, dbProduct.Stock)
	return product, nil
}
