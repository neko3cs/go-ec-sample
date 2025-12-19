package service

import (
	"errors"
	"go-ec-sample/command"
	"go-ec-sample/db"
	"go-ec-sample/domain"
	"go-ec-sample/query"

	"gorm.io/gorm"
)

type CartService struct{}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) GetCart(userId uint) (*domain.Cart, error) {
	var cart *domain.Cart
	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		q := query.NewGetCartQuery(userId)
		qh := query.NewGetCartQueryHandler(tx)
		c, err := qh.Handle(q)
		if err != nil {
			return err
		}
		if c == nil {
			cmd := command.NewCreateCartCommand(userId)
			ch := command.NewCreateCartCommandHandler(tx)
			err := ch.Handle(cmd)
			if err != nil {
				return err
			}
			c, err = qh.Handle(q)
			if err != nil {
				return err
			}
		}
		cart = c
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *CartService) SaveCart(cart *domain.Cart) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		cmd := command.NewUpdateCartCommand(cart.CartId(), command.NewUpdateCartItems(cart.Items()))
		h := command.NewUpdateCartCommandHandler(tx)
		return h.Handle(cmd)
	})
}

func (s *CartService) Checkout(cart *domain.Cart) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		updateProductCommandHandler := command.NewUpdateProductCommandHandler(tx)
		for _, item := range cart.Items() {
			getProductQuery := query.NewGetProductQuery(item.ProductId())
			getProductQueryHandler := query.NewGetProductQueryHandler(tx)
			product, err := getProductQueryHandler.Handle(getProductQuery)
			if err != nil {
				return err
			}
			if product.Stock() < item.Quantity() {
				return errors.New("在庫不足")
			}
			product.AddStock(-item.Quantity())
			updateProductCommand := command.NewUpdateProductCommand(product.Id(), product.Name(), product.Price(), product.Stock())
			err = updateProductCommandHandler.Handle(updateProductCommand)
			if err != nil {
				return err
			}
		}
		deleteCartCommandHandler := command.NewDeleteCartCommandHandler(tx)
		deleteCartCommand := command.NewDeleteCartCommand(cart.CartId())
		err := deleteCartCommandHandler.Handle(deleteCartCommand)
		if err != nil {
			return err
		}

		return nil
	})
}
