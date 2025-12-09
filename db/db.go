package db

import (
	"log"

	"go-ec-sample/domain"

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

	db.AutoMigrate(&domain.Product{})

	db.Create(&domain.Product{Name: "Apple", Price: 120})
	db.Create(&domain.Product{Name: "Banana", Price: 80})
	db.Create(&domain.Product{Name: "Strawberry", Price: 300})

	database = db
}

func GetDB() *gorm.DB {
	return database
}
