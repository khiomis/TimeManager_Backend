package api

import (
	"backend_time_manager/database"
	"backend_time_manager/dto"
	"backend_time_manager/entity"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigureAccountApiRoutes(router *gin.Engine) {
	router.POST("/account", HandleCreateAccount)
	router.GET("/account/:id", HandleGetAccount)
	router.PUT("/account/:id", HandleUpdateAccount)
	router.PUT("/account/activate", HandleActivateAccount)
	router.DELETE("/account/:id/remove", HandleDeactivateAccount)
	router.PUT("/account/forgot-password", HandleForgotPassword)
	router.PUT("/account/reset-password", HandleResetPassword)
}

func HandleCreateAccount(context *gin.Context) {
	var accDto dto.CreateAccountDto

	if err := context.BindJSON(&accDto); err != nil {
		context.String(400, err.Error())
		return
	}

	if _, err := database.FindAccountByEmail(accDto.Email); err != nil {
		context.String(400, "Email already used")
		return
	}

	account := entity.Account{
		Email:     accDto.Email,
		Name:      accDto.Name,
		Password:  accDto.Password,
		Status:    entity.AccountPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	account, err := database.SaveAccount(account)

	if err != nil {
		context.String(400, err.Error())
		return
	}

	accountResponse := dto.AccountDTO{Id: account.Id, Name: account.Name, Email: account.Email, Status: account.Status, UpdatedAt: account.UpdatedAt}

	context.IndentedJSON(200, accountResponse)
}

func HandleActivateAccount(context *gin.Context) {
	activationCode := context.DefaultQuery("code", "")

	if activationCode == "" {
		context.String(400, "Code not provided")
	}
	//id := context.Param("id")
	//
	//account, err := database.FindAccountById(id)
	//
	//if err != nil {
	//	context.String(400, "Account not found")
	//	return
	//}
	//
	//if account.Status != entity.AccountPending {
	//	context.String(http.StatusNotAcceptable, "Not acceptable")
	//	return
	//}
	//
	//account.Status = entity.AccountActivated
	//
	//account, err = database.SaveAccount(account)
	//
	//if err != nil {
	//	context.String(400, err.Error())
	//	return
	//}
	//
	//accountResponse := dto.AccountDTO{Id: account.Id, Name: account.Name, Email: account.Email, Status: account.Status}
	//
	//context.IndentedJSON(200, accountResponse)
}

func HandleDeactivateAccount(context *gin.Context) {

}

func HandleUpdateAccount(context *gin.Context) {
	id := context.Param("id")

	context.JSON(200, gin.H{"id": id})
}

func HandleGetAccount(context *gin.Context) {
	id := context.Param("id")

	account, err := database.FindAccountById(id)

	if err != nil {
		context.String(400, "Account not found")
		return
	}

	accountResponse := dto.AccountDTO{Id: account.Id, Name: account.Name, Email: account.Email, Status: account.Status, UpdatedAt: account.UpdatedAt}

	context.IndentedJSON(200, accountResponse)
}

func HandleForgotPassword(context *gin.Context) {

}

func HandleResetPassword(context *gin.Context) {

}
