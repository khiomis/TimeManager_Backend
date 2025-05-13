package api

import (
	"backend_time_manager/database"
	"backend_time_manager/dto"
	"backend_time_manager/entity"
	"backend_time_manager/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ConfigurePublicUserApiRoutes(router *gin.RouterGroup) {
	router.POST("/users", handleCreateUser)
	router.PUT("/users/activate", handleActivateUser)
	router.PUT("/users/forgot-password", handleForgotPassword)
	router.PUT("/users/reset-password", handleResetPassword)
}

func ConfigurePrivateUserApiRoutes(router *gin.RouterGroup) {
	router.PUT("/users", handleUpdateUser)
	router.DELETE("/users/remove", handleDeactivateUser)
	router.GET("/users/:userId", loadUserContext, handleGetUser)
}

func handleCreateUser(context *gin.Context) {
	var accDto dto.CreateUserDto

	if err := context.BindJSON(&accDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailAlreadyInUse, err := database.CheckEmailAlreadyInUseUser(accDto.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if emailAlreadyInUse {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	password, err := utils.HashPassword(accDto.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := entity.User{
		Email:    accDto.Email,
		Name:     accDto.Name,
		Password: password,
		Status:   entity.UserPending,
	}

	savedUser, err := database.SaveUser(user)

	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	userResponse := dto.UserDTO{}.From(savedUser)

	context.IndentedJSON(http.StatusCreated, userResponse)
}

func handleActivateUser(context *gin.Context) {
	activationCode := context.DefaultQuery("code", "")

	if activationCode == "" {
		context.String(http.StatusBadRequest, "Code not provided")
	}
	//id := context.Param("id")
	//
	//user, err := database.FindUserById(id)
	//
	//if err != nil {
	//	context.String(400, "User not found")
	//	return
	//}
	//
	//if user.Status != entity.UserPending {
	//	context.String(http.StatusNotAcceptable, "Not acceptable")
	//	return
	//}
	//
	//user.Status = entity.UserActivated
	//
	//user, err = database.SaveUser(user)
	//
	//if err != nil {
	//	context.String(400, err.Error())
	//	return
	//}
	//
	//userResponse := dto.UserDTO{Id: user.Id, Name: user.Name, Email: user.Email, Status: user.Status}
	//
	//context.IndentedJSON(200, userResponse)
}

func handleDeactivateUser(context *gin.Context) {

}

func handleUpdateUser(context *gin.Context) {
}

func handleGetUser(context *gin.Context) {
	user := getUserFromPath(context)

	userResponse := dto.UserDTO{}.From(user)

	context.IndentedJSON(200, userResponse)
}

func handleForgotPassword(context *gin.Context) {

}

func handleResetPassword(context *gin.Context) {

}

func loadUserContext(context *gin.Context) {
	id := context.Param("userId")

	if err := uuid.Validate(id); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	user, err := database.FindUserByUuid(uuid.MustParse(id))
	if err != nil {
		context.String(http.StatusNotFound, "User not found")
		return
	}

	context.Set("user", user)
}

func getUserFromPath(context *gin.Context) entity.User {
	return context.MustGet("user").(entity.User)
}
