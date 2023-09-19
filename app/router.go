package app

import (
	"latihan-api-startup/controller"
	"latihan-api-startup/helper"
	"latihan-api-startup/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController controller.UserController,
	campaignController controller.CampaignController,
	transactionController controller.TransactionController,
	userService service.UserService,
	authService service.AuthService) *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("api/v1")

	api.GET("/users", userController.GetUser)
	api.POST("/user/register", userController.RegisterUser)
	api.POST("/user/login", userController.Login)
	api.POST("/user/email-check", userController.CheckEmailAvailbility)
	api.POST("/user/avatars", userController.UploadAvatar)
	api.GET("/user/fetch", userController.FetchUser)
	api.GET("/user/test", userController.GetUser)

	api.GET("/campaigns", campaignController.GetCampaigns)
	api.GET("/campaigns/:id", campaignController.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignController.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignController.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignController.UploadImage)

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionController.GetCampaignTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionController.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionController.CreateTransaction)
	api.POST("/transactions/notification", transactionController.GetNotification)

	return router
}

func authMiddleware(authService service.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set to cookie current user in gin
		context.Set("currentUser", user)
	}

}
