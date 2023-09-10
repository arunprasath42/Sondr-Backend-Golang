package models

import (
	db "sondr-backend/utils/database"

	"gorm.io/gorm"
)

type UserSocialMediaDetails struct {
	gorm.Model
	UserID       uint   `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	FacebookURL  string `json:"facebook_url,omitempty"`
	InstagramURL string `json:"instagram_url,omitempty"`
}

func SocialMediaMigrate() {
	db.DB.Debug().AutoMigrate(&UserSocialMediaDetails{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
