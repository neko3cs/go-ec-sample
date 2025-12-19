package command

import "go-ec-sample/domain"

type UpdateCartCommand struct {
	CartId uint
	Items  []UpdateCartItem
}

type UpdateCartItem struct {
	CartItemId uint
	ProductId  uint
	Quantity   int
}

func NewUpdateCartCommand(cartId uint, items []UpdateCartItem) *UpdateCartCommand {
	return &UpdateCartCommand{
		CartId: cartId,
		Items:  items,
	}
}

func NewUpdateCartItems(items []*domain.CartItem) []UpdateCartItem {
	result := make([]UpdateCartItem, len(items))
	for i, item := range items {
		result[i] = UpdateCartItem{
			CartItemId: item.CartItemId(),
			ProductId:  item.ProductId(),
			Quantity:   item.Quantity(),
		}
	}
	return result
}
