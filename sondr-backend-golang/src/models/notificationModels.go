package models

import (
	db "sondr-backend/utils/database"
	"time"

	"github.com/jinzhu/gorm"
)

type Notifications struct {
	gorm.Model
	SenderUserId   uint   `json:"sender_user_id,omitempty"`
	ReceiverUserId uint   `json:"receiver_user_id,omitempty"`
	Message        string `json:"message,omitempty"`
	Type           string `json:"type,omitempty"`
	IsRead         bool   `json:"is_read,omitempty"`
}
type NotificationResponse struct {
	Id                 uint      `json:"id,omitempty"`
	NotificationTime   time.Time `json:"notificationTime,omitempty"`
	SenderUserId       int       `json:"senderUserId,omitempty"`
	ReceiverUserId     int       `json:"receiverUserId,omitempty"`
	Message            string    `json:"message,omitempty"`
	Type               string    `json:"type,omitempty"`
	IsRead             bool      `json:"isRead"`
	SenderProfilePhoto string    `json:"senderProfilePhoto,omitempty"`
}

func NotificationMigrate() {
	db.DB.Debug().AutoMigrate(&Notifications{}).AddForeignKey("sender_user_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("receiver_user_id", "users(id)", "CASCADE", "CASCADE")
}
