package middlewares

import (
	"btpn/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan token dari header Authorization
		tokenString, err := helpers.ExtractToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Memeriksa dan memverifikasi token
		claims, err := helpers.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Periksa apakah klaim (claims) valid
		if err := claims.Valid(); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}

		// Menambahkan ID pengguna ke konteks untuk digunakan di handler akhir
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
