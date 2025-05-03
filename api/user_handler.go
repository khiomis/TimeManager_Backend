package api

import (
	"backend_time_manager/database"
	"backend_time_manager/dto"
	"backend_time_manager/entity"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigureUserApiRoutes(router *gin.Engine) {
	router.POST("/user", handleCreateUser)
	router.PUT("/user", handleUpdateUser)
	router.DELETE("/user/remove", handleDeactivateUser)
	router.GET("/user/:id", handleGetUser)
	router.PUT("/user/activate", handleActivateUser)
	router.PUT("/user/forgot-password", handleForgotPassword)
	router.PUT("/user/reset-password", handleResetPassword)
}

func handleCreateUser(context *gin.Context) {
	var accDto dto.CreateUserDto

	if err := context.BindJSON(&accDto); err != nil {
		context.String(400, err.Error())
		return
	}

	if _, err := database.FindUserByEmail(accDto.Email); err != nil {
		context.String(400, "Email already used")
		return
	}

	user := entity.User{
		Email:     accDto.Email,
		Name:      accDto.Name,
		Password:  accDto.Password,
		Status:    entity.UserPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := database.SaveUser(user)

	if err != nil {
		context.String(400, err.Error())
		return
	}

	userResponse := dto.UserDTO{Id: user.Id, Name: user.Name, Email: user.Email, Status: user.Status, UpdatedAt: user.UpdatedAt}

	context.IndentedJSON(200, userResponse)
}

func handleActivateUser(context *gin.Context) {
	activationCode := context.DefaultQuery("code", "")

	if activationCode == "" {
		context.String(400, "Code not provided")
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
	id := context.Param("id")

	user, err := database.FindUserById(id)

	if err != nil {
		context.String(400, "User not found")
		return
	}

	userResponse := dto.UserDTO{Id: user.Id, Name: user.Name, Email: user.Email, Status: user.Status, UpdatedAt: user.UpdatedAt}

	context.IndentedJSON(200, userResponse)
}

func handleForgotPassword(context *gin.Context) {

}

func handleResetPassword(context *gin.Context) {

}
