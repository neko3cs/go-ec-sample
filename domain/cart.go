package domain

type Cart struct {
	cartId uint
	userId uint
	items  []*CartItem
}

func NewCart(cartId uint, userId uint, items []*CartItem) *Cart {
	return &Cart{
		cartId: cartId,
		userId: userId,
		items:  items,
	}
}

func (c *Cart) CartId() uint       { return c.cartId }
func (c *Cart) UserId() uint       { return c.userId }
func (c *Cart) Items() []*CartItem { return c.items }

func (c *Cart) TotalCost() int {
	total := 0
	for _, item := range c.items {
		total += item.Price() * item.Quantity()
	}
	return total
}

func (c *Cart) AddItem(p Product, quantity int) {
	for i, item := range c.items {
		if item.productId == p.Id() {
			c.items[i].AddQuantity(quantity)
			return
		}
	}
	c.items = append(c.items, NewCartItem(0, 0, p.Id(), &p, quantity))
}

func (c *Cart) RemoveItem(productId uint) {
	for i, item := range c.items {
		if item.ProductId() == productId {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return
		}
	}
}
