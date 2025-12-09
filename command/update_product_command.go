package command

type UpdateProductCommand struct {
	Id    uint
	Name  string
	Price int
}

func NewUpdateProductCommand(id uint, name string, price int) *UpdateProductCommand {
	return &UpdateProductCommand{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
