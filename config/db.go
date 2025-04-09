package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"olxkz/models" // ✅ добавь этот импорт
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=olx port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Не удалось подключиться к базе данных!")
	}

	DB = db
	fmt.Println("База данных подключена!")

	// ✅ Используем models.Product
	db.AutoMigrate(&models.Product{}, &models.Category{})
}
