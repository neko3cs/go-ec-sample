package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type GetAllProductsQueryHandler struct{}

func NewGetAllProductsQueryHandler() *GetAllProductsQueryHandler {
	return &GetAllProductsQueryHandler{}
}

func (h *GetAllProductsQueryHandler) Handle(query *GetAllProductsQuery) ([]domain.Product, error) {
	var products []domain.Product
	err := db.GetDB().Find(&products).Error
	return products, err
}
