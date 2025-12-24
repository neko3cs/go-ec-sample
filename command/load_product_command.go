package command

type LoadProductCommand struct {
	ProductId uint
}

func NewLoadProductCommand(productId uint) *LoadProductCommand {
	return &LoadProductCommand{ProductId: productId}
}
