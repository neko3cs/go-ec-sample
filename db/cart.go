package db

type Cart struct {
	CartId uint `gorm:"primarykey"`
	UserId uint
	Items  []CartItem
}
