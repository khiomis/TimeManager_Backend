package api

import (
	"backend_time_manager/database"
	"backend_time_manager/dto"
	"backend_time_manager/entity"
	"backend_time_manager/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func ConfigurePublicAuthorizationApiRoutes(router *gin.RouterGroup) {
	router.POST("/auth/sign-in", handleSignIn)
	router.POST("/auth/sign-in/validate-token", handleConfirmSignIn)
}

func ConfigurePrivateAuthorizationApiRoutes(router *gin.RouterGroup) {
	router.POST("/auth/refresh", handleRefreshToken)
	router.DELETE("/auth/logout", handleSignOut)
}

func ValidateAndLoadToken(context *gin.Context) {
	token := context.GetHeader("Authorization")

	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	token = strings.Replace(token, "bearer ", "", 1)

	claim, err := utils.ParseToken(token)
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := database.FindUserById(claim.UserID)

	if err != nil {
		_ = context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	session, err := database.FindSessionByUuid(claim.SessionId)

	if err != nil {
		_ = context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	context.Set("logged_user", user)
	context.Set("logged_session", session)
}

func handleSignIn(context *gin.Context) {
	var signInBody dto.SignInDto
	if err := context.BindJSON(&signInBody); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.FindUserByEmail(signInBody.Email)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if !utils.CheckPasswordHash(signInBody.Password, user.Password) {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
		return
	}

	validationToken := entity.ValidationToken{
		ExpireAt: time.Now().Add(time.Minute * 15),
		Code:     utils.GenerateCharToken(6),
		Type:     entity.ValidationTokenTypeSignIn,
		IdUser:   user.Id,
	}

	validationToken, err = database.InsertToken(validationToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.TokenValidationDto{
		Id:    validationToken.Id,
		Email: user.Email,
		Token: "",
	}

	context.JSON(http.StatusOK, response)
}

func handleConfirmSignIn(context *gin.Context) {
	var data dto.TokenValidationDto
	if err := context.BindJSON(&data); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationToken, err := database.FindToken(data.Id, data.Token, entity.ValidationTokenTypeSignIn)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := database.FindUserByEmail(data.Email)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if validationToken.IdUser != user.Id {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token does not match"})
		return
	}

	if validationToken.ExpireAt.After(time.Now()) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	session := entity.Session{
		ExpireAt: time.Now().Add(7 * 24 * time.Hour),
		IdUser:   user.Id,
	}
	session, err = database.CreateSession(session)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authToken, err := utils.GenerateAccessToken(user.Id, session)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Id, session)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.AuthUserDto{
		AuthToken:          authToken,
		RefreshToken:       refreshToken,
		NeedSetNewPassword: false,
		User:               dto.UserDTO{}.From(user),
	}

	err = database.RemoveToken(validationToken.Id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

func handleRefreshToken(context *gin.Context) {
	session := getLoggedSession(context)
	user := getLoggedUser(context)

	var refreshToken string
	if err := context.BindJSON(refreshToken); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No refresh token provided"})
		return
	}

	refreshClaims, err := utils.ParseToken(refreshToken)

	if err != nil {
		_ = context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if refreshClaims.UserID != user.Id || refreshClaims.SessionId != session.Id {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token"})
		return
	}

	newAuthToken, err := utils.GenerateAccessToken(user.Id, session)

	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": newAuthToken})
}

func handleSignOut(context *gin.Context) {
	session := getLoggedSession(context)

	err := database.DeleteSession(session.Id)

	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.Status(http.StatusOK)
}

func getLoggedUser(context *gin.Context) entity.User {
	return context.MustGet("logged_user").(entity.User)
}

func getLoggedSession(context *gin.Context) entity.Session {
	return context.MustGet("logged_session").(entity.Session)
}
