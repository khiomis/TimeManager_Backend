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

	router := gin.Default()

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return
	}

	api.ConfigureAuthorizationApiRoutes(router)
	api.ConfigureUserApiRoutes(router)
	api.ConfigureProjectApiRoutes(router)
	api.ConfigureEntryApiRoutes(router)
	api.ConfigureTagApiRoutes(router)
	api.ConfigureTaskApiRoutes(router)

	database.ConnectDatabase()

	if err := router.Run("localhost:8080"); err != nil {
		return
	}
}
