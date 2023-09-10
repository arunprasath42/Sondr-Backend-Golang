package models

import (
	"database/sql"
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

//user table migrate
type Users struct {
	gorm.Model
	FirstName        string     `gorm:"type:varchar(255);" json:"first_name,omitempty" valid:"length(3|255)" validate:"required"`
	LastName         string     `gorm:"type:varchar(255);" json:"last_name,omitempty" valid:"length(2|255)" validate:"required"`
	Email            string     `gorm:"type:varchar(255);unique" json:"email,omitempty" valid:"length(5|255)" validate:"required,email"`
	PhoneNo          string     `gorm:"type:varchar(25);unique;default:NULL" json:"phone_no,omitempty" valid:"length(10|25)"`
	Gender           string     `json:"gender,omitempty" validate:"required"`
	DOB              string     `json:"dob,omitempty" validate:"required,datetime=2006-01-02"`
	ProfilePhoto     string     `json:"profile_photo,omitempty"`
	Occupation       string     `json:"occupation,omitempty"`
	Height           float64    `json:"height,omitempty"`
	Country          string     `gorm:"type:varchar(50);" json:"country,omitempty" valid:"length(10|50)"`
	City             string     `gorm:"type:varchar(50);" json:"city,omitempty" valid:"length(10|50)"`
	LastSeen         time.Time  `json:"last_seen,omitempty" gorm:"default:NULL"`
	Active           bool       `gorm:"type:bool;default:true" json:"active,omitempty"`
	Visible          bool       `gorm:"type:bool;default:true" json:"visible,omitempty"`
	LastVisited      *time.Time `json:"last_visited" gorm:"default:NULL"`
	Blocked          bool       `gorm:"type:bool;default:false" json:"blocked,omitempty"`
	BlockerID        int64      `json:"blocker_id,omitempty"`
	BlockerType      string     `json:"blocker_type,omitempty"`
	GAuth            string     `json:"gAuth,omitempty"`
	ProfileStatus    string     `json:"profileStatus,omitempty"`
	ReferenceID      string     `gorm:"unique;default:NULL" json:"referenceID,omitempty"`
	IsLoggedIn       bool       `json:"isLoggedIn,omitempty" gorm:"type:bool;default:false"`
	DeactivateReason string     `json:"deactivateReason,omitempty"`
}

/*****************Table for Toggle ON/OFF Visibility***************************/
type Visibility struct {
	gorm.Model
	UserID                      uint `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	GenderVisibility            bool `gorm:"type:bool" json:"genderVisibility,omitempty"`
	GenderCategoryVisibility    bool `gorm:"type:bool" json:"genderCategoryVisibility,omitempty"`
	InterestVisibility          bool `gorm:"type:bool" json:"interestVisibility,omitempty"`
	PurposeVisibility           bool `gorm:"type:bool" json:"purposeVisibility,omitempty"`
	HeightVisibility            bool `gorm:"type:bool" json:"heightVisibility,omitempty"`
	VerifiedVisibility          bool `gorm:"type:bool;default:true" json:"verifiedVisibility,omitempty"`
	OccupationVisibility        bool `gorm:"type:bool" json:"occupationVisibility,omitempty"`
	HideDetailsVisibility       bool `gorm:"type:bool" json:"hideDetailsVisibility,omitempty"`
	AllowprofileslikeVisibility bool `gorm:"type:bool;default:false" json:"allowprofileslikeVisibility,omitempty"`
	EnableVisibility            bool `gorm:"type:bool;default:true" json:"enableVisibility,omitempty"`
}

/************* Struct for Gender classification like cisgender, Trangender etc.,************/
type GroupGender struct {
	gorm.Model
	UserID        uint   `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	GroupCategory string `json:"group_category,omitempty"`
	GroupBy       string `json:"group_by,omitempty"`
}

//user photos table migrate
type UserPhotos struct {
	gorm.Model
	UserID uint   `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	Photo1 string `json:"photo1,omitempty"`
	Photo2 string `json:"photo2,omitempty"`
	Photo3 string `json:"photo3,omitempty"`
	Photo4 string `json:"photo4,omitempty"`
	Photo5 string `json:"photo5,omitempty"`
}

//user interested in table migrate
type UserInterestedIn struct {
	gorm.Model
	UserID uint   `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	Gender string `json:"gender,omitempty"`
}

// Meeting purpose table migrate
type MeetingPurpose struct {
	gorm.Model
	UserID         uint   `gorm:"foreignKey:id;unique" json:"user_id,omitempty"`
	MeetingPurpose string `json:"meeting_purpose,omitempty"`
}

// User Interests table migrate
type UserInterests struct {
	gorm.Model
	UserID    uint   `gorm:"foreignKey:id" json:"user_id,omitempty"`
	GenreType string `json:"genre_type,omitempty" type:"enum('Creativity','Music','Travelling')"`
	Interests string `json:"interests,omitempty"`
}

//Question table
type Questions struct {
	Id       uint   `gorm:"primarykey" json:"id,omitempty"`
	Question string `json:"question,omitempty"`
}

//Answers Table
type UserAnswers struct {
	gorm.Model
	UserId     uint   `gorm:"foreignKey:id" json:"user_id,omitempty"`
	QuestionId uint   `gorm:"foreignKey:id" json:"question_id,omitempty"`
	Answer     string `json:"answer,omitempty"`
}

type UserLocations struct {
	gorm.Model
	UserId       uint         `json:"user_id,omitempty"`
	Coordinates  string       `json:"coordinates,omitempty"`
	LocationName string       `json:"location_name,omitempty"`
	CheckIn      time.Time    `json:"check_in,omitempty"`
	CheckOut     sql.NullTime `json:"check_out,omitempty"`
}

type Countries struct {
	ID       uint   `gorm:"primarykey" json:"id,omitempty"`
	Name     string `gorm:"type:varchar(255);unique" json:"name,omitempty"`
	Code     string `gorm:"type:varchar(255)" json:"code,omitempty"`
	DialCode string `gorm:"type:varchar(255)" json:"dial_code,omitempty"`
}

type AllCountries struct {
	ID   uint   `gorm:"primarykey" json:"id,omitempty"`
	Name string `gorm:"type:varchar(255);unique" json:"name,omitempty"`
}

func UserMigrate() {
	db.DB.Debug().AutoMigrate(&Users{})
	db.DB.Debug().AutoMigrate(&UserPhotos{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&UserInterestedIn{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&UserInterests{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&MeetingPurpose{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&Questions{})
	db.DB.Debug().AutoMigrate(&UserAnswers{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").
		AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&Visibility{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&GroupGender{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.Debug().AutoMigrate(&Countries{})
	db.DB.Debug().AutoMigrate(&UserLocations{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}
