package domain

type Product struct {
	id    uint
	name  string
	price int
}

func NewProduct(id uint, name string, price int) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: price,
	}
}

func (p *Product) Id() uint { return p.id }

func (p *Product) Name() string { return p.name }

func (p *Product) Price() int { return p.price }
