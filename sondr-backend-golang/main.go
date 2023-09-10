package main

import (
	"log"
	"sondr-backend/route"
	"sondr-backend/src/controllers"
	"sondr-backend/src/repository"
	"sondr-backend/src/service"
	"sondr-backend/utils/cron"
	"sondr-backend/utils/database"
	"sondr-backend/utils/s3"
	"sondr-backend/utils/validator"

	config "sondr-backend/config"
	"sondr-backend/migration"
	logger "sondr-backend/utils/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.Default()
	logger.NewLogger("./log/debug.log")
	database.GetInstancemysql()

	repository.MySqlInit()
	migration.Migration()
	validator.Init()
	err := s3.ConnectS3()
	if err != nil {
		log.Println("error in s3", err)
	}

	controllers.LoadSocialMediaCreditional()
	//router.Use(middleware.TracingMiddleware())
	service.ChannelStart = make(chan bool)
	service.ChannelEnd = make(chan bool)
	go cron.Init()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("authorization")
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
