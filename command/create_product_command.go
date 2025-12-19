package command

type CreateProductCommand struct {
	Name  string
	Price int
	Stock int
}

func NewCreateProductCommand(name string, price int, stock int) *CreateProductCommand {
	return &CreateProductCommand{
		Name:  name,
		Price: price,
		Stock: stock,
	}
}
