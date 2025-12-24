package service

import (
	"go-ec-sample/command"
	"go-ec-sample/db"
	"go-ec-sample/query"
	"go-ec-sample/querymodel"

	"gorm.io/gorm"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAllProducts() ([]querymodel.ProductListItem, error) {
	var products []querymodel.ProductListItem

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		q := query.NewGetAllProductsQuery()
		h := query.NewGetAllProductsQueryHandler(tx)
		ps, err := h.Handle(q)
		if err != nil {
			return err
		}
		products = ps
		return nil
	})

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProduct(id uint) (*querymodel.ProductDetail, error) {
	var product *querymodel.ProductDetail

	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		q := query.NewGetProductQuery(id)
		h := query.NewGetProductQueryHandler(tx)
		p, err := h.Handle(q)
		if err != nil {
			return err
		}
		product = p
		return nil
	})

	if err != nil {
		return nil, err
	}
	return product, err
}

func (s *ProductService) CreateProduct(name string, price int, stock int) error {
	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		c := command.NewCreateProductCommand(name, price, stock)
		h := command.NewCreateProductCommandHandler(tx)
		return h.Handle(c)
	})
	return err
}

func (s *ProductService) UpdateProduct(id uint, name string, price int, stock int) error {
	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		c := command.NewUpdateProductCommand(id, name, price, stock)
		h := command.NewUpdateProductCommandHandler(tx)
		return h.Handle(c)
	})
	return err
}

func (s *ProductService) DeleteProduct(id uint) error {
	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		c := command.NewDeleteProductCommand(id)
		h := command.NewDeleteProductCommandHandler(tx)
		return h.Handle(c)
	})
	return err
}
