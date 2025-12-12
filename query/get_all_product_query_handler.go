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
	var dbProducts []db.Product
	err := db.GetDB().Find(&dbProducts).Error
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, p := range dbProducts {
		products = append(products, *domain.NewProduct(p.Id, p.Name, p.Price, p.Stock))
	}
	return products, err
}
