package cron

import (
	"sondr-backend/src/repository"
	"sondr-backend/src/service"

	"github.com/robfig/cron"
)

func Init() {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", service.EventServiceCronJob)
	c.AddFunc("0 0 0 * * *", repository.UpdateActiveStatusOfUser)
	c.Start()

}
