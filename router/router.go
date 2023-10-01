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

	r.POST("/photos", controllers.UploadPhoto)
	r.GET("/photos", controllers.GetPhotos)
	r.PUT("/photos/:photoId", controllers.UpdatePhoto)
	r.DELETE("/photos/:photoId", controllers.DeletePhoto)

	authGroup := r.Group("/photos")
	authGroup.Use(middlewares.AuthMiddleware())
	{
		authGroup.POST("/auth", controllers.UploadPhoto)
		authGroup.GET("/auth", controllers.GetPhotos)
		authGroup.PUT("/auth/:photoId", controllers.UpdatePhoto)
		authGroup.DELETE("/auth/:photoId", controllers.DeletePhoto)
	}

	return r
}
