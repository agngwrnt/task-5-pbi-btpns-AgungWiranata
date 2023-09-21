package database

import (
	"btpn/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "<koneksi_database>" // Ganti <koneksi_database> dengan string koneksi database Anda (contoh: "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	database, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database

	// Auto Migrate the models
	MigrateDB()
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
