package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type GetProductQueryHandler struct{}

func NewGetProductQueryHandler() *GetProductQueryHandler {
	return &GetProductQueryHandler{}
}

func (h *GetProductQueryHandler) Handle(q *GetProductQuery) (*domain.Product, error) {
	var dbProduct db.Product
	err := db.GetDB().First(&dbProduct, q.id).Error
	if err != nil {
		return nil, err
	}

	product := domain.NewProduct(dbProduct.Id, dbProduct.Name, dbProduct.Price)
	return product, nil
}
