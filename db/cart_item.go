package db

type CartItem struct {
	CartItemId uint `gorm:"primarykey"`
	CartId     uint
	ProductId  uint
	Quantity   int

	Product Product `gorm:"foreignKey:ProductId"`
}
