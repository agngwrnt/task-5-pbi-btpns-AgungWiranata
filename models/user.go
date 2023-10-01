package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" valid:"required,length(6|50),uniquePassword"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   `json:"photos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func init() {
	govalidator.TagMap["uniquePassword"] = govalidator.Validator(func(str string) bool {
		// Implementasi logika untuk memeriksa apakah kata sandi sudah digunakan
		// Return false jika kata sandi sudah digunakan
		// Return true jika kata sandi unik
		return true // Ganti dengan implementasi Anda sendiri
	})
}
