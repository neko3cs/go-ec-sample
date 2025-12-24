package service

import (
	"errors"
	"go-ec-sample/command"
	"go-ec-sample/db"
	"go-ec-sample/query"
	"go-ec-sample/querymodel"

	"gorm.io/gorm"
)

type CartService struct{}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) GetCart(userId uint) (*querymodel.Cart, error) {
	var cart *querymodel.Cart
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

func (s *CartService) AddToCart(userId uint, productId uint, quantity int) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		loadCartCmd := command.NewLoadCartCommand(userId)
		loadCartCmdHandler := command.NewLoadCartCommandHandler(tx)
		cart, err := loadCartCmdHandler.Handle(loadCartCmd)
		if err != nil {
			return err
		}

		loadProductCmd := command.NewLoadProductCommand(productId)
		loadProductCmdHandler := command.NewLoadProductCommandHandler(tx)
		product, err := loadProductCmdHandler.Handle(loadProductCmd)
		if err != nil {
			return err
		}

		cart.AddItem(*product, quantity)

		updateCartCmd := command.NewUpdateCartCommand(cart.CartId(), command.NewUpdateCartItems(cart.Items()))
		updateCartCmdHandler := command.NewUpdateCartCommandHandler(tx)
		return updateCartCmdHandler.Handle(updateCartCmd)
	})
}

func (s *CartService) RemoveFromCart(userId uint, productId uint) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		loadCartCmd := command.NewLoadCartCommand(userId)
		loadCartCmdHandler := command.NewLoadCartCommandHandler(tx)
		cart, err := loadCartCmdHandler.Handle(loadCartCmd)
		if err != nil {
			return err
		}

		cart.RemoveItem(productId)

		updateCartCmd := command.NewUpdateCartCommand(cart.CartId(), command.NewUpdateCartItems(cart.Items()))
		updateCartCmdHandler := command.NewUpdateCartCommandHandler(tx)
		return updateCartCmdHandler.Handle(updateCartCmd)
	})
}

func (s *CartService) Checkout(userId uint) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		loadCartCmd := command.NewLoadCartCommand(userId)
		loadCartCmdHandler := command.NewLoadCartCommandHandler(tx)
		cart, err := loadCartCmdHandler.Handle(loadCartCmd)
		if err != nil {
			return err
		}

		updateProductCommandHandler := command.NewUpdateProductCommandHandler(tx)
		for _, item := range cart.Items() {
			loadProductCmd := command.NewLoadProductCommand(item.ProductId())
			loadProductCmdHandler := command.NewLoadProductCommandHandler(tx)
			product, err := loadProductCmdHandler.Handle(loadProductCmd)
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
		err = deleteCartCommandHandler.Handle(deleteCartCommand)
		if err != nil {
			return err
		}

		return nil
	})
}
