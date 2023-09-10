package models

import (
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

type Kycs struct {
	gorm.Model
	UserId     uint   `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	Status     string `json:"status,omitempty"`
	VerifiedBy string `json:"verified_by,omitempty"`
	KycPhoto1  string `json:"kyc_photo_1,omitempty"`
	KycPhoto2  string `json:"kyc_photo_2,omitempty"`
}

type KycResponse struct {
	KycStatusReq    []*ListKycstatusReq    `json:"kycStatusReq,omitempty"`
	ReqCount        int                    `json:"reqCount,omitempty"`
	KycStatusVerify []*ListKycStatusVerify `json:"kycStatusVerify,omitempty"`
	VerifyCount     int                    `json:"verifyCount,omitempty"`
}
type ListKycstatusReq struct {
	Id          uint       `json:"id,omitempty"`
	FirstName   string     `json:"firstName,omitempty"`
	LastName    string     `json:"lastName,omitempty"`
	CreatedAt   time.Time  `json:"verificationRequested,omitempty"`
	LastVisited *time.Time `json:"lastVisited,omitempty" gorm:"last_visited"`
}
type ListKycStatusVerify struct {
	Id         uint      `json:"id,omitempty"`
	FirstName  string    `json:"firstName,omitempty"`
	LastName   string    `json:"lastName,omitempty"`
	VerifiedBy string    `json:"verifiedBy,omitempty"`
	UpdatedAt  time.Time `json:"verifiedDataAndTime,omitempty"`
}

func KYCMigrate() {
	db.DB.Debug().AutoMigrate(&Kycs{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
