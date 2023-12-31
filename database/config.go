package database

import (
	"btpn/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	_ = "user=silahkan masukan host sendiri password=1 host=localhost port=5432 database=postgres"
	database, err := gorm.Open(postgres.Open("user=silahkan masukan host sendiri password=1 host=localhost port=5432 database=postgres"), &gorm.Config{})
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

// :)
