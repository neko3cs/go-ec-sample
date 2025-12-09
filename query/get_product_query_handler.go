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
	var product domain.Product
	err := db.GetDB().First(&product, q.id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
