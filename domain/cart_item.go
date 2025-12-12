package domain

type CartItem struct {
	productId uint
	name      string
	price     int
	quantity  int
}

func NewCartItem(productId uint, name string, price int, quantity int) *CartItem {
	return &CartItem{
		productId: productId,
		name:      name,
		price:     price,
		quantity:  quantity,
	}
}

func (ci *CartItem) ProductId() uint { return ci.productId }

func (ci *CartItem) Name() string { return ci.name }

func (ci *CartItem) Price() int { return ci.price }

func (ci *CartItem) Quantity() int { return ci.quantity }

func (ci *CartItem) AddQuantity(count int) {
	ci.quantity += count
	if ci.quantity < 0 {
		ci.quantity = 0
	}
}
