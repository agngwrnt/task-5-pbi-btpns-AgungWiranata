package router

import (
	"btpn/controllers"
	"btpn/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rute untuk pengguna
	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.PUT("/users/:userId", controllers.UpdateUser)
	r.DELETE("/users/:userId", controllers.DeleteUser)

	// Rute untuk foto
	r.POST("/photos", controllers.UploadPhoto)
	r.GET("/photos", controllers.GetPhotos)
	r.PUT("/photos/:photoId", controllers.UpdatePhoto)
	r.DELETE("/photos/:photoId", controllers.DeletePhoto)

	// Rute Authentation
	r.POST("/photos", middlewares.AuthMiddleware(), controllers.UploadPhoto)
	r.GET("/photos", middlewares.AuthMiddleware(), controllers.GetPhotos)
	r.PUT("/photos/:photoId", middlewares.AuthMiddleware(), controllers.UpdatePhoto)
	r.DELETE("/photos/:photoId", middlewares.AuthMiddleware(), controllers.DeletePhoto)

	//func UploadPhoto(c *gin.Context) {
	//	userID, _ := c.Get("user_id")
	//	router.POST("/upload-photo", controllers.UploadPhoto)
	//	// Gunakan userID untuk operasi terkait pengguna yang terotentikasi
	//	// ...
	//}

	return r
}
