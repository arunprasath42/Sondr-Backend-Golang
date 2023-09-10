package models

import (
	"database/sql"
	db "sondr-backend/utils/database"
	"time"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	HostUserID  uint      `json:"hostUserId,omitempty"`
	EventName   string    `json:"eventName,omitempty"`
	Location    string    `json:"location,omitempty"`
	EventMode   string    `json:"eventMode,omitempty"`
	Password    string    `json:"password,omitempty"`
	Date        string    `json:"date,omitempty"`
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	Status      string    `json:"status,omitempty"`
	Coordinates string    `json:"coordinates,omitempty"`
}

type EventMetadatas struct {
	gorm.Model
	EventId       uint         `json:"eventId,omitempty"`
	InvitedUserId uint         `gorm:"foreignKey:id" json:"invitedUserId,omitempty"`
	IsAttended    bool         `json:"isAttended,omitempty"`
	CheckIn       time.Time    `json:"checkIn,omitempty" gorm:"default:NULL"`
	CheckOut      sql.NullTime `json:"checkOut,omitempty" gorm:"default:NULL"`
	NoOfCheckIn   int64        `json:"noOfCheckIn,omitempty"`
}

type EventRequest struct {
	EventId       uint   `json:"eventId,omitempty"`
	HostUserID    uint   `json:"hostUserId,omitempty" validate:"required"`
	EventName     string `json:"eventName,omitempty" validate:"required"`
	Location      string `json:"location,omitempty" validate:"required"`
	EventMode     string `json:"eventMode,omitempty" validate:"required,oneof='Private''Open'"`
	Password      string `json:"password,omitempty" validate:"required_if=EventMode Private"`
	Date          string `json:"date,omitempty" validate:"required"`
	StartTime     string `json:"startTime,omitempty" validate:"required,timeFormat=2006-01-02 15:04:05"  `
	EndTime       string `json:"endTime,omitempty" validate:"required,timeFormat=2006-01-02 15:04:05"`
	Coordinates   string `json:"coordinates,omitempty" validate:"required"`
	InvitedUserId []uint `json:"invitedUserId,omitempty"`
	Status        string `json:"status,omitempty" validate:"omitempty,oneof='Planned' 'On-Going' 'Cancelled'"`
}

type EventCheckInRequest struct {
	UserID    uint   `json:"userID,omitempty" validate:"required"`
	Latitude  string `json:"latitude,omitempty" validate:"required"`
	Longitude string `json:"longitude,omitempty" validate:"required"`
	EventId   uint   `json:"eventID,omitempty" validate:"required"`
	Password  string `json:"password,omitempty"`
}

type EventResponse struct {
	Events                   []*ListEvent              `json:"events,omitempty"`
	Count                    int64                     `json:"count,omitempty"`
	Event                    *Event                    `json:"event,omitempty"`
	InvitedGuest             []*InvitedGuest           `json:"invitedGuest,omitempty"`
	AttendedGuest            []*AttendedGuest          `json:"attendedGuest,omitempty"`
	InvitedGuestCount        int64                     `json:"invitedGuestCount,omitempty"`
	AttendedGuestCount       int64                     `json:"attendedGuestCount,omitempty"`
	Message                  string                    `json:"message,omitempty"`
	EventID                  uint                      `json:"eventID,omitempty"`
	InivitedEvents           []*ListInvitedEvents      `json:"inivitedEvents,omitempty"`
	InivitedEventCount       int64                     `json:"inivitedEventCount,omitempty"`
	HostedEvents             []*ListHostedEvents       `json:"hostedEvents,omitempty"`
	HostedEventCount         int64                     `json:"hostedEventCount,omitempty"`
	ListProfilesEventCheckIn []*MatchedProfileResponse `json:"listProfilesEventCheckIn,omitempty"`
	InvitedUserId            []*uint                   `json:"invitedUserId,omitempty"`
}
type InviUsersID struct {
	InvitedUserId uint
}
type StartTimeQueue struct {
	EventID   uint
	StartedAt time.Time
}
type EndTimeQueue struct {
	EventID uint
	EndAt   time.Time
}
type ListHostedEvents struct {
	ID               uint      `json:"eventId,omitempty"`
	FirstName        string    `json:"firstName,omitempty"`
	LastName         string    `json:"lastName,omitempty"`
	Location         string    `json:"location,omitempty"`
	Date             string    `json:"date,omitempty"`
	StartTime        time.Time `json:"startTime,omitempty"`
	EndTime          time.Time `json:"endTime,omitempty"`
	ProfilePhoto     string    `json:"profilePic,omitempty"`
	Coordinates      string    `json:"coordinates,omitempty"`
	EventName        string    `json:"eventName,omitempty"`
	InvitedUserCount int64     `json:"invitedUserCount,omitempty"`
}
type AttendedGuest struct {
	InvitedUserId uint      `json:"invitedUserID,omitempty"`
	FirstName     string    `json:"firstName,omitempty"`
	LastName      string    `json:"lastName,omitempty"`
	CheckIn       time.Time `json:"checkIn,omitempty"`
	CheckOut      time.Time `json:"checkOut,omitempty"`
}
type InvitedGuest struct {
	InvitedUserId uint   `json:"invitedUserID,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
}
type ListEvent struct {
	ID        uint      `json:"id,omitempty"`
	EventName string    `json:"eventName,omitempty"`
	Location  string    `json:"location,omitempty"`
	Status    string    `json:"status,omitempty"`
	EventMode string    `json:"eventMode,omitempty"`
	Date      string    `json:"date,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
}
type Event struct {
	Id          uint      `json:"id,omitempty"`
	HostUserID  uint      `json:"hostUserId,omitempty"`
	EventName   string    `json:"eventName,omitempty"`
	Location    string    `json:"location,omitempty"`
	EventMode   string    `json:"eventMode,omitempty"`
	Password    string    `json:"password,omitempty"`
	Date        string    `json:"date,omitempty"`
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
	Status      string    `json:"status,omitempty"`
	FirstName   string    `json:"firstName,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	CreatedAt   time.Time `json:"createdDate,omitempty"`
	Coordinates string    `json:"coordinates,omitempty"`
}

type EmailId struct {
	Email string
}
type Eventresponse struct {
	ID                 uint   `json:"id,omitempty"`
	EventName          string `json:"eventName,omitempty"`
	Location           string `json:"location,omitempty"`
	EventMode          string `json:"eventMode,omitempty"`
	Date               string `json:"date,omitempty"`
	Status             string `json:"status,omitempty"`
	ParticipationCount int    `json:"participationCount,omitempty"`
}
type TotalEventsRecord struct {
	AttendedEvents int64 `json:"attendedEvents,omitempty"`
}
type UserEventResponse struct {
	Count          int64            `json:"count,omitempty"`
	AttendedEvents int64            `json:"attendedEvents,omitempty"`
	EventInfo      []*Eventresponse `json:"eventInfo,omitempty"`
}
type UserLoginResponse struct {
	Message       string `json:"message,omitempty"`
	UserId        uint   `json:"userID,omitempty"`
	Token         string `json:"token,omitempty"`
	Otp           string `json:"otp,omitempty"`
	ProfileStatus string `json:"profileStatus,omitempty"`
}
type ListInvitedEvents struct {
	ID               uint      `json:"eventId,omitempty"`
	EventName        string    `json:"eventName,omitempty"`
	Location         string    `json:"location,omitempty"`
	FirstName        string    `json:"firstName,omitempty"`
	LastName         string    `json:"lastName,omitempty"`
	ProfilePhoto     string    `json:"profilePic,omitempty"`
	Date             string    `json:"date,omitempty"`
	StartTime        time.Time `json:"startTime,omitempty"`
	EndTime          time.Time `json:"endTime,omitempty"`
	Coordinates      string    `json:"coordinates,omitempty"`
	InvitedUserCount int64     `json:"invitedUserCount,omitempty"`
}

type EventLocationResponse struct {
	Id          uint      `json:"id,omitempty"`
	HostUserID  uint      `json:"hostUserId,omitempty"`
	EventName   string    `json:"eventName,omitempty"`
	Location    string    `json:"location,omitempty"`
	EventMode   string    `json:"eventMode,omitempty"`
	Date        string    `json:"date,omitempty"`
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
	Status      string    `json:"status,omitempty"`
	Coordinates string    `json:"coordinates,omitempty"`
	Distance    float64   `json:"distance,omitempty"`
}
type UserDetails struct {
	PhoneNumber string
	GAuth       string
}

type UserEmailDetails struct {
	Email string
	GAuth string
}

func EventMigrate() {
	db.DB.AutoMigrate(&Events{}).AddForeignKey("host_user_id", "users(id)", "CASCADE", "CASCADE")
	db.DB.AutoMigrate(&EventMetadatas{}).AddForeignKey("event_id", "events(id)", "CASCADE", "CASCADE").AddForeignKey("invited_user_id", "users(id)", "CASCADE", "CASCADE")
}
