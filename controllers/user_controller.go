package controllers

import (
	"btpn/database"
	"btpn/helpers"
	"btpn/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	// Mendapatkan data dari request
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alamat email sudah terdaftar"})
		return
	}

	// Memeriksa validitas alamat email
	if !govalidator.IsEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alamat email tidak valid"})
		return
	}

	// Memeriksa validitas username
	if err := helpers.ValidateUsername(user.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi menggunakan GoValidator
	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password sebelum menyimpan di database
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses kata sandi"})
		return
	}
	user.Password = hashedPassword

	// Simpan data pengguna ke database
	database.DB.Create(&user)

	// Generate JWT token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token"})
		return
	}

	// Kirim respons dengan token
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func LoginUser(c *gin.Context) {
	// Mendapatkan data dari request
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	// Cari pengguna berdasarkan email
	result := database.DB.Where("email = ?", inputUser.Email).First(&existingUser)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	// Memeriksa validitas alamat email
	if !govalidator.IsEmail(inputUser.Email) { // Menambahkan validasi alamat email
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alamat email tidak valid"})
		return
	}

	// Periksa kecocokan kata sandi
	if !helpers.CheckPasswordHash(inputUser.Password, existingUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kata sandi salah"})
		return
	}

	// Generate JWT token
	token, err := helpers.GenerateToken(existingUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token"})
		return
	}

	// Kirim respons dengan token
	c.JSON(http.StatusOK, gin.H{"token": token, "user": existingUser})
}

func DeleteUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	inputUserID := c.Params.ByName("userId")

	var user models.User

	// Cari pengguna berdasarkan ID
	result := database.DB.First(&user, inputUserID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	// Periksa apakah pengguna yang terotentikasi adalah pemilik pengguna yang akan dihapus
	if user.ID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki izin untuk menghapus pengguna ini"})
		return
	}

	// Hapus pengguna dari database
	database.DB.Delete(&user)

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{"message": "Akun pengguna berhasil dihapus"})
}

func UpdateUser(c *gin.Context) {
	//userID, _ := c.Get("user_id")
	inputUserID := c.Params.ByName("userId")

	// Pastikan userID dan inputUserID adalah tipe data uint

	// Periksa apakah pengguna yang terotentikasi adalah pemilik pengguna yang akan diupdate
	//if userID.(uint) != inputUserID {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki izin untuk memperbarui pengguna ini"})
	//	return
	//}

	var user models.User

	// Cari pengguna berdasarkan ID
	result := database.DB.First(&user, inputUserID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	// Handle logika pembaruan pengguna di sini

	// Simpan pembaruan ke database
	database.DB.Save(&user)

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil diperbarui"})
}
