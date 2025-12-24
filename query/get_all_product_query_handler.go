package query

import (
	"go-ec-sample/db"
	"go-ec-sample/querymodel"

	"gorm.io/gorm"
)

type GetAllProductsQueryHandler struct {
	db *gorm.DB
}

func NewGetAllProductsQueryHandler(db *gorm.DB) *GetAllProductsQueryHandler {
	return &GetAllProductsQueryHandler{db: db}
}

func (h *GetAllProductsQueryHandler) Handle(query *GetAllProductsQuery) ([]querymodel.ProductListItem, error) {
	var dbProducts []db.Product
	err := h.db.Find(&dbProducts).Error
	if err != nil {
		return nil, err
	}

	var products []querymodel.ProductListItem
	for _, p := range dbProducts {
		products = append(products, querymodel.ProductListItem{
			Id:    p.Id,
			Name:  p.Name,
			Price: p.Price,
			Stock: p.Stock,
		})
	}
	return products, err
}
