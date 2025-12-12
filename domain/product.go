package domain

type Product struct {
	id    uint
	name  string
	price int
	stock int
}

func NewProduct(id uint, name string, price int, stock int) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: price,
		stock: stock,
	}
}

func (p *Product) Id() uint { return p.id }

func (p *Product) Name() string { return p.name }

func (p *Product) Price() int { return p.price }

func (p *Product) Stock() int { return p.stock }
