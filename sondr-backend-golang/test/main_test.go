package test

import (
	"fmt"
	"log"
	"sondr-backend/src/repository"
	"sondr-backend/src/service"
	"sondr-backend/utils/database"
	"sondr-backend/utils/logging"
	"sondr-backend/utils/s3"

	"github.com/spf13/viper"
)

func init() {
	fmt.Println("called")
	//config.LoadConfig()
	logging.NewLogger("../log/debug.log")
	LoadConfig()
	database.GetInstancemysql()
	s3.ConnectS3()
	repository.MySqlInit()
	service.ChannelStart = make(chan bool)
	service.ChannelEnd = make(chan bool)
	go service.EventDurationStarting()
	go service.EventDurationEnding()

}

func LoadConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("..")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Config not found...")
	}
}
