package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"

	"gorm.io/gorm"
)

type GetAllProductsQueryHandler struct {
	db *gorm.DB
}

func NewGetAllProductsQueryHandler(db *gorm.DB) *GetAllProductsQueryHandler {
	return &GetAllProductsQueryHandler{db: db}
}

func (h *GetAllProductsQueryHandler) Handle(query *GetAllProductsQuery) ([]domain.Product, error) {
	var dbProducts []db.Product
	err := h.db.Find(&dbProducts).Error
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, p := range dbProducts {
		products = append(products, *domain.NewProduct(p.Id, p.Name, p.Price, p.Stock))
	}
	return products, err
}
