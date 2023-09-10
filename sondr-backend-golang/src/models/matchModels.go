package models

import (
	db "sondr-backend/utils/database"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	SenderUserId   uint   `gorm:"foreignKey:id" json:"sender_user_id,omitempty"`
	ReceiverUserId uint   `gorm:"foreignKey:id" json:"receiver_user_id,omitempty"`
	Status         string `gorm:"column:status;type:enum('Requested','Matched')" json:"status,omitempty"`
	NoOfRejections int64  `json:"no_of_rejections,omitempty" gorm:"default:NULL"`
}

func MatchMigrate() {
	db.DB.AutoMigrate(&Match{}).
		AddForeignKey("sender_user_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("receiver_user_id", "users(id)", "CASCADE", "CASCADE")

}
