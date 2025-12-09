package service

import (
	"go-ec-sample/domain"
	"go-ec-sample/query"
)

type ProductService struct {
	query *query.ProductQuery
}

func NewProductService() *ProductService {
	return &ProductService{
		query: query.NewProductQuery(),
	}
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.query.FindAll()
}

func (s *ProductService) GetProduct(id uint) (*domain.Product, error) {
	return s.query.FindByID(id)
}
