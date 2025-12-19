package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"

	"gorm.io/gorm"
)

type GetProductQueryHandler struct {
	db *gorm.DB
}

func NewGetProductQueryHandler(db *gorm.DB) *GetProductQueryHandler {
	return &GetProductQueryHandler{db: db}
}

func (h *GetProductQueryHandler) Handle(q *GetProductQuery) (*domain.Product, error) {
	var dbProduct db.Product
	err := h.db.First(&dbProduct, q.id).Error
	if err != nil {
		return nil, err
	}

	product := domain.NewProduct(dbProduct.Id, dbProduct.Name, dbProduct.Price, dbProduct.Stock)
	return product, nil
}
