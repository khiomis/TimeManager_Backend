package main

import (
	"backend_time_manager/api"
	"backend_time_manager/database"
	"backend_time_manager/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon utils="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error is occurred  on .env file please check")
	}

	utils.InitJwt()

	database.ConnectDatabase()

	router := gin.Default()

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return
	}

	defaultRouter := router.Group("/api")

	api.ConfigurePublicAuthorizationApiRoutes(defaultRouter)
	api.ConfigurePublicUserApiRoutes(defaultRouter)

	privateRouter := defaultRouter.Group("/", api.ValidateAndLoadToken)
	api.ConfigurePrivateAuthorizationApiRoutes(privateRouter)
	api.ConfigurePrivateUserApiRoutes(privateRouter)
	api.ConfigureProjectApiRoutes(privateRouter)
	api.ConfigureEntryApiRoutes(privateRouter)
	api.ConfigureTagApiRoutes(privateRouter)
	api.ConfigureTaskApiRoutes(privateRouter)

	if err := router.Run("localhost:8080"); err != nil {
		return
	}
}
