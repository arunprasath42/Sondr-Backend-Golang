package models

import (
	db "sondr-backend/utils/database"

	"gorm.io/gorm"
)

//table for blocked users
type Blocked struct {
	gorm.Model
	BlockerId uint `gorm:"foreignKey:id" json:"blocker_id,omitempty"`
	BlockeeId uint `gorm:"foreignKey:id" json:"blockee_id,omitempty"`
}

//struct for request for blocked users
type BlockRequest struct {
	UserId  uint `json:"userId,omitempty"`
	AdminId uint `json:"adminId,omitempty"`
}

//Migrate the table to automatically create the table.
func BlockedMigrate() {
	db.DB.Debug().AutoMigrate(&Blocked{}).
		AddForeignKey("blocker_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("blockee_id", "users(id)", "CASCADE", "CASCADE")
}
