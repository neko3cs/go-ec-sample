package domain

type Cart struct {
	items []*CartItem
}

func NewCart() *Cart {
	return &Cart{
		items: []*CartItem{},
	}
}

func NewCartWithItems(items []*CartItem) *Cart {
	return &Cart{
		items: items,
	}
}

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
	c.items = append(c.items, NewCartItem(p.Id(), p.Name(), p.Price(), quantity))
}

func (c *Cart) RemoveItem(productId uint) {
	for i, item := range c.items {
		if item.ProductId() == productId {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return
		}
	}
}
