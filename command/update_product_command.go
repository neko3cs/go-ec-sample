package command

type UpdateProductCommand struct {
	Id    uint
	Name  string
	Price int
	Stock int
}

func NewUpdateProductCommand(id uint, name string, price int, stock int) *UpdateProductCommand {
	return &UpdateProductCommand{
		Id:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
}
