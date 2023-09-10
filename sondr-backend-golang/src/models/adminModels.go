package models

import (
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

type Admins struct {
	gorm.Model
	Name     string    `gorm:"column:name;type:varchar(255)" json:"name,omitempty" validate:"required"`
	Photo    string    `gorm:"column:photo;type:varchar(255)" json:"photo,omitempty"`
	Email    string    `gorm:"column:email;type:varchar(255);unique" json:"email,omitempty" validate:"required,email"`
	Password string    `gorm:"column:password;type:varchar(255)" json:"password,omitempty" validate:"required" validatePasswordString:"password"`
	LastSeen time.Time `gorm:"default:NULL" json:"last_seen,omitempty"`
	Role     string    `gorm:"column:role;type:enum('Admin','SubAdmin')" json:"role,omitempty"`
	Blocked  bool      `gorm:"column:blocked;type:bool" json:"blocked,omitempty"`
}

type AdminForgotPassword struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type AdminBlock struct {
	ID      uint `json:"id,omitempty"`
	Blocked bool `json:"blocked,omitempty"`
}

/***Struct for forget password*/
type Admin struct {
	ID          int
	Email       string
	NewPassword string `json:"newpassword,omitempty" validatePasswordString:"password"`
	Password    string `json:"password" validatePasswordString:"password"`
}

type Login struct {
	UniqueId    uint
	AccessToken string
}

/***For Listing admin table ***/
type ListAdmin struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Blocked   bool      `json:"blocked"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type AdminResponse struct {
	ListAdmin []*ListAdmin `json:"listAdmin,omitempty"`
	Count     int          `json:"count,omitempty"`
	Admin     *ListAdmin   `json:"admin,omitempty"`
}

func AdminMigrate() {
	db.DB.Debug().AutoMigrate(&Admins{})
}
