package controllers

import (
	"btpn/database"
	"btpn/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Implementasikan fungsi-fungsi sesuai dengan kebutuhan (UploadPhoto, GetPhotos, UpdatePhoto, DeletePhoto)

func UploadPhoto(c *gin.Context) {
	// Mendapatkan data dari request
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mendapatkan ID pengguna dari konteks
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Mengatur UserID dengan ID pengguna yang terotentikasi
	photo.UserID = userID.(uint)

	// Simpan data foto ke database
	database.DB.Create(&photo)

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil diunggah", "photo": photo})
}

func DeletePhoto(c *gin.Context) {
	userID, _ := c.Get("user_id")
	photoID := c.Params.ByName("photoId")

	var photo models.Photo

	// Cari foto berdasarkan ID
	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
		return
	}

	// Periksa apakah pengguna yang terotentikasi adalah pemilik foto
	if photo.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki izin untuk menghapus foto ini"})
		return
	}

	// Hapus foto dari database
	database.DB.Delete(&photo)

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil dihapus"})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	// Dapatkan daftar foto dari database
	result := database.DB.Find(&photos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar foto"})
		return
	}

	// Kirim respons dengan daftar foto
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Params.ByName("photoId")

	var photo models.Photo

	// Cari foto berdasarkan ID
	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
		return
	}

	// Handle logika pembaruan foto di sini

	// Simpan pembaruan ke database
	database.DB.Save(&photo)

	// Kirim respons
	c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil diperbarui"})
}
