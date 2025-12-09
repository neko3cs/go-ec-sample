package service

import (
	"go-ec-sample/command"
	"go-ec-sample/domain"
	"go-ec-sample/query"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAllProducts() ([]domain.Product, error) {
	q := query.NewGetAllProductsQuery()
	h := query.NewGetAllProductsQueryHandler()
	return h.Handle(q)
}

func (s *ProductService) GetProduct(id uint) (*domain.Product, error) {
	q := query.NewGetProductQuery(id)
	h := query.NewGetProductQueryHandler()
	return h.Handle(q)
}

func (s *ProductService) CreateProduct(name string, price int) error {
	c := command.NewCreateProductCommand(name, price)
	h := command.NewCreateProductCommandHandler()
	return h.Handle(c)
}

func (s *ProductService) UpdateProduct(id uint, name string, price int) error {
	c := command.NewUpdateProductCommand(id, name, price)
	h := command.NewUpdateProductCommandHandler()
	return h.Handle(c)
}

func (s *ProductService) DeleteProduct(id uint) error {
	c := command.NewDeleteProductCommand(id)
	h := command.NewDeleteProductCommandHandler()
	return h.Handle(c)
}
