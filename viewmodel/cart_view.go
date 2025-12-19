package viewmodel

import "go-ec-sample/domain"

type CartView struct {
	Items []CartItemView `json:"items"`
}

type CartItemView struct {
	ProductId uint   `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
}

func NewCartView(cart *domain.Cart) *CartView {
	items := make([]CartItemView, 0)
	for _, item := range cart.Items() {
		items = append(items, CartItemView{
			ProductId: item.ProductId(),
			Name:      item.Name(),
			Price:     item.Price(),
			Quantity:  item.Quantity(),
		})
	}
	return &CartView{
		Items: items,
	}
}

func (cv *CartView) ToDomainModel() *domain.Cart {
	cartItems := make([]*domain.CartItem, 0)
	for _, item := range cv.Items {
		cartItem := domain.NewCartItem(item.ProductId, item.Name, item.Price, item.Quantity)
		cartItems = append(cartItems, cartItem)
	}
	cart := domain.NewCartWithItems(cartItems)
	return cart
}
