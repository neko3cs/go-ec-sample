package db

type User struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	IsAdmin  bool
}
