package app

import (
	"latihan-api-startup/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController controller.UserController) *gin.Engine {
	router := gin.Default()

	api := router.Group("api/v1")

	api.GET("/users", userController.GetUser)
	api.POST("/user/register", userController.RegisterUser)
	api.POST("/user/login", userController.Login)
	api.POST("/user/email-check", userController.CheckEmailAvailbility)
	api.POST("/user/avatars", userController.UploadAvatar)
	api.GET("/user/fetch", userController.FetchUser)

	return router
}
