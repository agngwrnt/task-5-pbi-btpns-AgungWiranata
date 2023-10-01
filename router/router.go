package router

import (
	"btpn/controllers"
	"btpn/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rute untuk pengguna
	r.POST("/api/v1/users/register", controllers.RegisterUser)
	r.POST("/api/v1/users/login", controllers.LoginUser)
	r.PUT("/api/v1/users/:userId", controllers.UpdateUser)
	r.DELETE("/api/v1/users/:userId", controllers.DeleteUser)

	r.POST("/api/v1/photos", controllers.UploadPhoto)
	r.GET("/api/v1/photos", controllers.GetPhotos)
	r.PUT("/api/v1/photos/:photoId", controllers.UpdatePhoto)
	r.DELETE("/api/v1/photos/:photoId", controllers.DeletePhoto)

	authGroup := r.Group("/api/v1/photos")
	authGroup.Use(middlewares.AuthMiddleware())
	{
		authGroup.POST("/auth", controllers.UploadPhoto)
		authGroup.GET("/auth", controllers.GetPhotos)
		authGroup.PUT("/auth/:photoId", controllers.UpdatePhoto)
		authGroup.DELETE("/auth/:photoId", controllers.DeletePhoto)
	}

	return r
}
