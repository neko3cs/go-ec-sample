package command

type PurchaseCartCommand struct {
	ProductId uint
	Quantity  int
}

func NewPurchaseCartCommand(productId uint, quantity int) *PurchaseCartCommand {
	return &PurchaseCartCommand{
		ProductId: productId,
		Quantity:  quantity,
	}
}
