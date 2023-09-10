package migration

import (
	"sondr-backend/src/models"
	"sondr-backend/src/service"
)

/****DB migration to be added***/

func Migration() {
	models.AdminMigrate()
	models.UserMigrate()
	models.EventMigrate()
	models.MatchMigrate()
	models.ReportedUserMigrate()
	models.SocialMediaMigrate()
	models.KYCMigrate()
	models.NotificationMigrate()

	//service.Insert4SubAdmins()
	service.GetAllCountries()
}
