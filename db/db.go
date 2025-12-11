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

	db.AutoMigrate(&Product{})

	db.Create(&Product{Name: "Apple", Price: 120})
	db.Create(&Product{Name: "Banana", Price: 80})
	db.Create(&Product{Name: "Strawberry", Price: 300})

	database = db
}

func GetDB() *gorm.DB {
	return database
}
