package domain

type CartItem struct {
	cartItemId uint
	cartId     uint
	productId  uint
	quantity   int
	product    *Product
}

func NewCartItem(cartItemId uint, cartId uint, productId uint, product *Product, quantity int) *CartItem {
	return &CartItem{
		cartItemId: cartItemId,
		cartId:     cartId,
		productId:  productId,
		product:    product,
		quantity:   quantity,
	}
}

func (ci *CartItem) CartItemId() uint { return ci.cartItemId }
func (ci *CartItem) CartId() uint     { return ci.cartId }
func (ci *CartItem) ProductId() uint  { return ci.productId }
func (ci *CartItem) Name() string     { return ci.product.Name() }
func (ci *CartItem) Price() int       { return ci.product.Price() }
func (ci *CartItem) Quantity() int    { return ci.quantity }

func (ci *CartItem) AddQuantity(count int) {
	ci.quantity += count
	if ci.quantity < 0 {
		ci.quantity = 0
	}
}
