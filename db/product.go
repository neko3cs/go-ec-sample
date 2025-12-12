package db

type Product struct {
	Id    uint `gorm:"primarykey"`
	Name  string
	Price int
	Stock int
}
