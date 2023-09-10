package models

import (
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

type ReportedUsers struct {
	gorm.Model
	ReporterUserId uint   `json:"reporter_user_id,omitempty"`
	ReporteeUserId uint   `json:"reportee_user_id,omitempty"`
	Reason         string `json:"reason,omitempty"`
	Comment        string `json:"comment,omitempty"`
}

type ReportsResponse struct {
	UniqueId     uint      `json:"uniqueId,omitempty"`
	FirstName    string    `json:"firstName,omitempty"`
	LastName     string    `json:"lastName,omitempty"`
	Status       bool      `json:"status,omitempty"`
	ReportedDate time.Time `json:"reportedDate,omitempty"`
	Reason       string    `json:"reason,omitempty"`
	Comment      string    `json:"comment,omitempty"`
}
type ReportInfo struct {
	Count           int64 `json:"count,omitempty"`
	ReporteResponse []*ReportsResponse
}

func ReportedUserMigrate() {
	db.DB.Debug().AutoMigrate(&ReportedUsers{}).
		AddForeignKey("reporter_user_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("reportee_user_id", "users(id)", "CASCADE", "CASCADE")
}
