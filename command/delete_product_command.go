package command

type DeleteProductCommand struct {
	Id uint
}

func NewDeleteProductCommand(id uint) *DeleteProductCommand {
	return &DeleteProductCommand{
		Id: id,
	}
}
