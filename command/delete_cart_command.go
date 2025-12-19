package command

type DeleteCartCommand struct {
	CartId uint
}

func NewDeleteCartCommand(cartId uint) *DeleteCartCommand {
	return &DeleteCartCommand{CartId: cartId}
}
