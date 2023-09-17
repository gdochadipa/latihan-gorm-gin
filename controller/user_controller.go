package controller

import (
	"fmt"
	"latihan-api-startup/helper"
	"latihan-api-startup/model/domain"
	"latihan-api-startup/model/web"
	"latihan-api-startup/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImp struct {
	userService service.UserService
	authService service.AuthService
}

type UserController interface {
	RegisterUser(context *gin.Context)
	Login(context *gin.Context)
	CheckEmailAvailbility(context *gin.Context)
	UploadAvatar(context *gin.Context)
	FetchUser(context *gin.Context)
	GetUser(context *gin.Context)
}

func NewUserController(userService service.UserService, authService service.AuthService) UserController {
	return &UserControllerImp{userService: userService, authService: authService}
}

func (s *UserControllerImp) RegisterUser(context *gin.Context) {
	var input web.RegisterUserInput

	err := context.ShouldBindJSON(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := map[string]any{"errors": errors}
		response := helper.APIResponse("Register account Failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUsers, err := s.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := s.authService.GenerateToken(newUsers.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := domain.FormatUser(newUsers, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	context.JSON(http.StatusOK, response)

}

func (s *UserControllerImp) Login(context *gin.Context) {
	var input web.LoginInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]any{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	logginUser, err := s.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := s.authService.GenerateToken(logginUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := domain.FormatUser(logginUser, token)

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)

	context.JSON(http.StatusOK, response)
}

func (s UserControllerImp) CheckEmailAvailbility(context *gin.Context) {
	var input web.CheckEmailInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		context.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable := s.userService.IsEmailAvailable(input)

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	context.JSON(http.StatusOK, response)
}

func (s UserControllerImp) UploadAvatar(context *gin.Context) {
	file, err := context.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := context.MustGet("currentUser").(domain.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = context.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = s.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	context.JSON(http.StatusOK, response)
}

func (s *UserControllerImp) FetchUser(context *gin.Context) {
	currentUser := context.MustGet("currentUser").(domain.User)

	formatter := web.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	context.JSON(http.StatusOK, response)
}

func (s *UserControllerImp) GetUser(context *gin.Context) {
	getUser, err := s.userService.GetUserByID(1)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		context.JSON(http.StatusBadRequest, response)
		return
	}

	newUser := domain.FormatUser(getUser, "")

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", newUser)

	context.JSON(http.StatusOK, response)

}
