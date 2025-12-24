package query

import (
	"go-ec-sample/db"
	"go-ec-sample/querymodel"

	"gorm.io/gorm"
)

type GetProductQueryHandler struct {
	db *gorm.DB
}

func NewGetProductQueryHandler(db *gorm.DB) *GetProductQueryHandler {
	return &GetProductQueryHandler{db: db}
}

func (h *GetProductQueryHandler) Handle(q *GetProductQuery) (*querymodel.ProductDetail, error) {
	var dbProduct db.Product
	err := h.db.First(&dbProduct, q.id).Error
	if err != nil {
		return nil, err
	}

	product := querymodel.ProductDetail{
		Id:    dbProduct.Id,
		Name:  dbProduct.Name,
		Price: dbProduct.Price,
		Stock: dbProduct.Stock,
	}
	return &product, nil
}
