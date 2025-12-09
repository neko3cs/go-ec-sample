package query

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type ProductQuery struct{}

func NewProductQuery() *ProductQuery {
	return &ProductQuery{}
}

func (q *ProductQuery) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := db.GetDB().Find(&products).Error
	return products, err
}

func (q *ProductQuery) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := db.GetDB().First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
