package command

type CreateProductCommand struct {
	Name  string
	Price int
}

func NewCreateProductCommand(name string, price int) *CreateProductCommand {
	return &CreateProductCommand{
		Name:  name,
		Price: price,
	}
}
