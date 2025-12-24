package query

import (
	"errors"
	"go-ec-sample/db"
	"go-ec-sample/querymodel"

	"gorm.io/gorm"
)

type GetCartQueryHandler struct {
	db *gorm.DB
}

func NewGetCartQueryHandler(db *gorm.DB) *GetCartQueryHandler {
	return &GetCartQueryHandler{db: db}
}

func (h *GetCartQueryHandler) Handle(query *GetCartQuery) (*querymodel.Cart, error) {
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

	cartItems := make([]*querymodel.CartItem, 0, len(c.Items))
	for _, item := range c.Items {
		cartItem := querymodel.CartItem{
			ProductId:    item.ProductId,
			ProductName:  item.Product.Name,
			Price:        item.Product.Price,
			Quantity:     item.Quantity,
			SubTotalCost: item.Product.Price * item.Quantity,
		}
		cartItems = append(cartItems, &cartItem)
	}

	totalCost := 0
	for _, item := range cartItems {
		totalCost += item.SubTotalCost
	}
	cart := querymodel.Cart{
		Items:     cartItems,
		TotalCost: totalCost,
	}
	return &cart, nil
}
