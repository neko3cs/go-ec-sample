package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	dsn := "file::memory:?cache=shared"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect SQLite:", err)
	}

	db.AutoMigrate(
		&Product{},
		&User{})

	db.Create(&Product{Name: "Apple", Price: 120})
	db.Create(&Product{Name: "Banana", Price: 80})
	db.Create(&Product{Name: "Strawberry", Price: 300})
	db.Create(&User{Name: "Admin", Email: "admin@example.com", Password: "password", IsAdmin: true})
	db.Create(&User{Name: "User", Email: "user@example.com", Password: "password", IsAdmin: false})

	database = db
}

func GetDB() *gorm.DB {
	return database
}
