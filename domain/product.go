package domain

type Product struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Price int
}
