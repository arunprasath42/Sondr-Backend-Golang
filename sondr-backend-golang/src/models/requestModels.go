package models

type Request struct {
	ID           uint   `json:"id,omitempty" form:"id"`
	Email        string `json:"email,omitempty" validate:"email"`
	Password     string `json:"password,omitempty"`
	PageNo       int    `json:"pageno,omitempty" form:"pageNo"`
	PageSize     int    `json:"pageSize,omitempty" form:"pageSize"`
	SearchFilter string `json:"searchfilter,omitempty" form:"searchFilter"`
	From         string `json:"from,omitempty" form:"from"`
	To           string `json:"to,omitempty" form:"to"`
	Status       string `json:"status,omitempty" form:"status"`
	UserId       uint   `json:"userId,omitempty" form:"userId"`
	EventId      uint   `json:"eventId,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	OTP          string `json:"otp,omitempty"`
	NewUser      bool   `json:"newUser,omitempty"`
	Latitude     string `json:"latitude,omitempty" form:"latitude"`
	Longitude    string `json:"longitude,omitempty" form:"longitude"`
}

type RequestLogin struct {
	ID  uint   `json:"id,omitempty"`
	OTP string `json:"otp,omitempty"`
}

/***USER REQUEST FOR SETTINGUP USER PROFILE***/
type UserRequest struct {
	ID                       uint    `json:"id,omitempty"`
	FirstName                string  `json:"firstName,omitempty" valid:"length(3|255)" validate:"required"`
	LastName                 string  `json:"lastName,omitempty" valid:"length(2|255)" validate:"required"`
	PhoneNumber              string  `json:"phoneNumber,omitempty" valid:"length(6|15)"`
	Email                    string  `json:"email,omitempty" valid:"length(5|255)" validate:"required,email"`
	Gender                   string  `json:"gender,omitempty" validate:"required,oneof='Male''Female''Non-Binary'"`
	DOB                      string  `json:"dob,omitempty" validate:"required,datetime=2006-01-02"`
	Occupation               string  `json:"occupation,omitempty"`
	Height                   float64 `json:"height,omitempty"`
	Country                  string  `json:"country,omitempty"`
	City                     string  `json:"city,omitempty"`
	GenderVisibility         bool    `json:"genderVisibility,omitempty"`
	GenderCategoryVisibility bool    `json:"genderCategoryVisibility,omitempty"`
	HeightVisibility         bool    `json:"heightVisibility,omitempty"`
	GroupCategory            string  `json:"groupCategory,omitempty" validate:"required_if=Gender Non-Binary"`
	GroupBy                  string  `json:"groupBy,omitempty" validate:"required_if=Gender Non-Binary"`
	ReferenceID              string  `json:"referenceID,omitempty"`
	ProfileCompletion        int     `json:"profileCompletion,omitempty"`
}

type ResponseString struct {
	Message string `json:"message,omitempty"`
}

/**GENDER INTEREST REQUEST**/
type UserInterest struct {
	ID                 uint   `json:"id,omitempty" validate:"required"`
	Gender             string `json:"gender,omitempty" validate:"required"`
	InterestVisibility bool   `json:"interestVisibility,omitempty"`
}

/**MEETING PURPOSE REQUEST**/
type Purpose struct {
	ID                  uint   `json:"id,omitempty" validate:"required"`
	MeetingPurpose      string `json:"meetingPurpose,omitempty" validate:"required"`
	Gender              string `json:"gender,omitempty"`
	Verified_visibility bool   `json:"verifiedVisibility,omitempty"`
	PurposeVisibility   bool   `json:"purposeVisibility,omitempty"`
}

type UpdateUserProfile struct {
	ID                   uint    `json:"id,omitempty" validate:"required"`
	Occupation           string  `json:"occupation,omitempty"`
	Height               float64 `json:"height,omitempty"`
	Country              string  `json:"country,omitempty"`
	City                 string  `json:"city,omitempty"`
	HeightVisibility     bool    `json:"heightVisibility,omitempty"`
	OccupationVisibility bool    `json:"occupationVisibility,omitempty"`
	Interests            string  `json:"interests,omitempty"`
}

type UploadPhotos struct {
	ID             uint
	Photo1         string `json:"photo1,omitempty"`
	Photo2         string `json:"photo2,omitempty"`
	Photo3         string `json:"photo3,omitempty"`
	Photo4         string `json:"photo4,omitempty"`
	Photo5         string `json:"photo5,omitempty"`
	ProfilePicture bool   `json:"profilePicture,omitempty"`
	IsRegistration bool   `json:"isRegistration,omitempty"`
}

type LocationRequest struct {
	UserId                  uint   `json:"userId,omitempty" validate:"required"`
	CurrentLatitude         string `json:"currentLatitude,omitempty" validate:"required"`
	CurrentLongitude        string `json:"currentLongitude,omitempty" validate:"required"`
	DestinationLatitude     string `json:"destinationLatitude,omitempty" validate:"required"`
	DestinationLongitude    string `json:"destinationLongitude,omitempty" validate:"required"`
	DestinationLocationName string `json:"destinationLocationName,omitempty" validate:"required"`
}

type AnswerRequest struct {
	QuestionId uint   `json:"questionId,omitempty" validate:"required"`
	Answer     string `json:"answer,omitempty" validate:"required"`
}

type CountryRequest struct {
	Country string `json:"country,omitempty"`
}

type Country struct {
	Name     string `json:"name,omitempty"`
	Code     string `json:"code,omitempty"`
	DialCode string `json:"dial_code,omitempty"`
}
type LikeRequest struct {
	SenderUserId   uint `json:"senderUserId,omitempty" validate:"required"`
	ReceiverUserId uint `json:"receiverUserId,omitempty" validate:"required"`
}

type MatchUserInfo struct {
	Id                   uint   `json:"id,omitempty"`
	SenderUserId         uint   `json:"sender_user_id,omitempty"`
	SenderUserName       string `json:"sender_user_name,omitempty"`
	SenderProfilePhoto   string `json:"sender_profile_photo,omitempty"`
	ReceiverUserId       uint   `json:"receiver_user_id,omitempty"`
	ReceiverUserName     string `json:"receiver_user_name,omitempty"`
	ReceiverProfilePhoto string `json:"receiver_profile_photo,omitempty"`
	Status               string `json:"status,omitempty"`
}

type ReportRequest struct {
	ReporteeUserId uint   `json:"reporteeUserId,omitempty" validate:"required"`
	Reason         string `json:"reason,omitempty" validate:"required"`
	Comment        string `json:"comment,omitempty"`
}

/**DASHBOARD REQUEST**/
type Countrequest struct {
	FromTime string `json:"from,omitempty" form:"FromTime"`
	ToTime   string `json:"to,omitempty" form:"ToTime"`
}
type DailyMatchesReportreq struct {
	FromTime string `json:"from,omitempty" form:"FromTime"`
	ToTime   string `json:"to,omitempty" form:"ToTime"`
}
type InterestRequest struct {
	GenreType string   `json:"genreType,omitempty" validate:"required"`
	Interests []string `json:"interests,omitempty" validate:"required"`
}

type GetUserDetails struct {
	ID uint `json:"id,omitempty" validate:"required"`
}

type DropTable struct {
	TableName string `json:"tableName,omitempty" validate:"required"`
}

type UserPhoneno struct {
	ID      uint   `json:"id,omitempty" validate:"required"`
	PhoneNo string `json:"phoneNumber,omitempty"`
	GAuth   string `json:"gauth,omitempty" `
	Email   string `json:"email,omitempty"`
}

type DeactivateAccount struct {
	ID               uint   `json:"id,omitempty" validate:"required"`
	DeactivateReason string `json:"deactivateReason,omitempty" validate:"required"`
}
