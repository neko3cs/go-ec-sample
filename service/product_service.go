package service

import (
	"go-ec-sample/command"
	"go-ec-sample/domain"
	"go-ec-sample/query"
)

type ProductService struct {
	query   *query.ProductQuery
	command *command.ProductCommand
}

func NewProductService(q *query.ProductQuery, c *command.ProductCommand) *ProductService {
	return &ProductService{
		query:   q,
		command: c,
	}
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	return s.query.FindAll()
}

func (s *ProductService) GetProduct(id uint) (*domain.Product, error) {
	return s.query.FindByID(id)
}

func (s *ProductService) CreateProduct(name string, price int) error {
	p := &domain.Product{
		Name:  name,
		Price: price,
	}
	return s.command.InsertProduct(p)
}

func (s *ProductService) UpdateProduct(id uint, name string, price int) error {
	p, err := s.query.FindByID(id)
	if err != nil {
		return err
	}

	p.Name = name
	p.Price = price

	return s.command.UpdateProduct(p)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.command.DeleteProduct(id)
}
