package query

import (
	"errors"
	"go-ec-sample/db"
	"go-ec-sample/domain"

	"gorm.io/gorm"
)

type GetCartQueryHandler struct {
	db *gorm.DB
}

func NewGetCartQueryHandler(db *gorm.DB) *GetCartQueryHandler {
	return &GetCartQueryHandler{db: db}
}

func (h *GetCartQueryHandler) Handle(query *GetCartQuery) (*domain.Cart, error) {
	var c db.Cart
	err := h.db.
		Preload("Items.Product").
		Where("user_id = ?", query.userId).
		First(&c).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	cartItems := make([]*domain.CartItem, 0, len(c.Items))
	for _, item := range c.Items {
		cartItem := domain.NewCartItem(
			item.CartItemId,
			item.CartId,
			item.ProductId,
			domain.NewProduct(
				item.Product.Id,
				item.Product.Name,
				item.Product.Price,
				item.Product.Stock,
			),
			item.Quantity,
		)
		cartItems = append(cartItems, cartItem)
	}

	cart := domain.NewCart(c.CartId, c.UserId, cartItems)
	return cart, nil
}
