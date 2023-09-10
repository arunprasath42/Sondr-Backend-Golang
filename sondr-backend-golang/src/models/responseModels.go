package models

import "time"

type AllUserResponse struct {
	Id              uint    `json:"id,omitempty"`
	FirstName       string  `json:"firstName,omitempty"`
	LastName        string  `json:"lastName,omitempty"`
	LastVisited     *string `json:"lastVisited,omitempty"`
	LastEventHosted string  `json:"lastEventHosted,omitempty"`
	IsBlocked       bool    `json:"isBlocked"`
	Type            string  `json:"type,omitempty"`
}

type UserInfo struct {
	Email                 string    `json:"email,omitempty"`
	PhoneNo               string    `json:"phoneNo,omitempty"`
	ProfilePhoto          string    `json:"profilePhoto,omitempty"`
	Age                   string    `json:"age,omitempty"`
	Location              string    `json:"location,omitempty"`
	CreatedDate           time.Time `json:"createdDate,omitempty"`
	LastVisited           *string   `json:"lastVisited,omitempty"`
	IsBlocked             bool      `json:"isBlocked,omitempty"`
	KycPhoto1             string    `json:"kycPhoto1,omitempty"`
	KycPhoto2             string    `json:"kycPhoto2,omitempty"`
	KycVerificationStatus string    `json:"kycVerification_status,omitempty"`
	VerifiedBy            string    `json:"verifiedBy,omitempty"`
	VerifiedDate          time.Time `json:"verifiedDate,omitempty"`
	FacebookURL           string    `json:"facebookURL,omitempty"`
	InstagramURL          string    `json:"instagramURL,omitempty"`
}
type UsersMetaData struct {
	MatchCount        int64 `json:"matchCount"`
	LikeSentCount     int64 `json:"likeSentCount"`
	LikeReceivedCount int64 `json:"likeReceivedCount"`
	RejectionsCount   int64 `json:"rejectionsCount"`
	TotalCheckIns     int64 `json:"totalCheckIns"`
}
type KycPhotosresponse struct {
	KycPhotos []string
}
type UploadedPhotosresponse struct {
	Photo1 string `json:"photo1,omitempty"`
	Photo2 string `json:"photo2,omitempty"`
	Photo3 string `json:"photo3,omitempty"`
	Photo4 string `json:"photo4,omitempty"`
	Photo5 string `json:"photo5,omitempty"`
}
type UserResponse struct {
	Count                     int64                     `json:"count,omitempty"`
	CurrentPage               int64                     `json:"currentPage,omitempty"`
	TotalPages                int64                     `json:"totalPages,omitempty"`
	Limit                     int64                     `json:"limit,omitempty"`
	Message                   string                    `json:"message,omitempty"`
	AllUserResponse           []*AllUserResponse        `json:"users,omitempty"`
	MatchedProfiles           []*MatchedProfileResponse `json:"matchedProfiles,omitempty"`
	LocationCheckedInProfiles []*MatchedProfileResponse `json:"locationCheckedInProfiles,omitempty"`
	UserProfileResponse       *UserProfileResponse      `json:"userProfileResponse,omitempty"`
	QuestionResponse          []*Questions              `json:"promptQuestions,omitempty"`
	UserUploadedPhotos        []*string                 `json:"userUploadedPhotos,omitempty"`
	Notifications             []*NotificationResponse   `json:"notifications,omitempty"`
}
type LocationProfile struct {
	UserId       uint    `json:"userId,omitempty"`
	Coordinates  string  `json:"coordinates,omitempty"`
	LocationName string  `json:"locationName,omitempty"`
	Distance     float64 `json:"distance"`
}
type MapResponse struct {
	EventLocationResponse []*EventLocationResponse `json:"eventLocationResponse"`
	UserLocationResponse  []*LocationProfile       `json:"userLocationResponse"`
	EventCheckInInfo      EventCheckInInfo         `json:"eventCheckInInfo"`
	LocationCheckInInfo   LocationCheckInInfo      `json:"locationCheckInInfo"`
}
type QuestionResponse struct {
	Count          int64             `json:"count,omitempty"`
	Questions      []*Questions      `json:"questions,omitempty"`
	PromptResponse []*PromptResponse `json:"promptResponse,omitempty"`
}

type PromptResponse struct {
	ID       int64  `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}
type UserInterestsResponse struct {
	Interests string `json:"interests,omitempty"`
}
type MatchedProfileResponse struct {
	UserId       uint   `json:"userId,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	ProfilePhoto string `json:"profilePhoto,omitempty"`
	Age          string `json:"age,omitempty"`
}
type UserProfileResponse struct {
	FirstName          string            `json:"firstName,omitempty"`
	LastName           string            `json:"lastName,omitempty"`
	ProfilePhoto       string            `json:"profilePhoto,omitempty"`
	Age                string            `json:"age,omitempty"`
	Occupation         string            `json:"occupation,omitempty"`
	Country            string            `json:"country,omitempty"`
	City               string            `json:"city,omitempty"`
	Height             float64           `json:"height,omitempty"`
	GenderVisibility   bool              `json:"genderVisibility,omitempty"`
	InterestVisibility bool              `json:"interestVisibility,omitempty"`
	PurposeVisibility  bool              `json:"purposeVisibility,omitempty"`
	HeightVisibility   bool              `json:"heightVisibility,omitempty"`
	SenderUserId       int64             `json:"senderUserId,omitempty"`
	MatchStatus        string            `json:"matchStatus,omitempty"`
	FriendCount        int64             `json:"friendCount,omitempty"`
	UserPhotos         []interface{}     `json:"userPhotos,omitempty"`
	Interests          []*string         `json:"interests,omitempty"`
	PromptQuestions    []*PromptResponse `json:"promptQuestions,omitempty"`
}

type CountryResponse struct {
	Error   bool      `json:"error,omitempty"`
	Message string    `json:"msg,omitempty"`
	Data    []Country `json:"data,omitempty"`
}

type CityResponse struct {
	Error   bool        `json:"error,omitempty"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

/***************DASHBOARD RESPONSE***************/
type TotalCount struct {
	TotalUsers             int `json:"total_users"`
	ActiveUsers            int `json:"active_users"`
	InactiveUsers          int `json:"inactive_users"`
	BlockedUsers           int `json:"blocked_users"`
	HostedEvents           int `json:"hosted_events"`
	Event_total_check_ins  int `json:"event_total_check_ins"`
	User_total_check_ins   int `json:"user_total_check_ins"`
	Totalcheckins          int `json:"total_check_ins"`
	Event_total_check_outs int `json:"event_total_check_outs"`
	User_total_check_outs  int `json:"user_total_check_outs"`
	Totalcheckouts         int `json:"total_check_outs"`
	Location_check_ins     int `json:"location_check_ins"`
}

//LAST REGISTERED USERS
type LastTenUsers struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Active    bool   `json:"active,omitempty"`
}

type TopTenLocationUsers struct {
	Position      int    `json:"position,omitempty"`
	LocationName  string `json:"locationName,omitempty"`
	NumberOfUsers int    `json:"numberOfUsers,omitempty"`
}

type Kycverification struct {
	Status                   bool    `json:"status,omitempty"`
	TotalUsers               int     `json:"total_users"`
	KYCVerified              int     `json:"kyc_verified"`
	KYCUnverified            int     `json:"kyc_unverified"`
	KYCVerifiedPercentage    float64 `json:"kycVerifiedPercentage"`
	KYCNotVerifiedPercentage float64 `json:"kycNotVerifiedPercentage"`
}

type DailyMatchesReport struct {
	Day            string `json:"day,omitempty"`
	MatchedMatches int    `json:"matched_matches,omitempty"`
}

type SiteVisitorAnalytics struct {
	TotalVisitors int    `json:"total_visitors"`
	Month         string `json:"month"`
}

type Userdetails struct {
	ID                          uint              `json:"id,omitempty" valid:"required"`
	FirstName                   string            `json:"firstName,omitempty" valid:"required"`
	LastName                    string            `json:"lastName,omitempty" valid:"required"`
	Email                       string            `json:"email,omitempty"`
	Age                         string            `json:"age,omitempty"`
	ProfilePhoto                string            `json:"profilePhoto,omitempty" valid:"required"`
	Photo2                      string            `json:"photo2,omitempty"`
	Photo3                      string            `json:"photo3,omitempty" `
	Photo4                      string            `json:"photo4,omitempty" `
	Photo5                      string            `json:"photo5,omitempty" `
	PhoneNo                     string            `json:"phoneNo,omitempty" `
	DOB                         string            `json:"dob,omitempty" valid:"required"`
	Country                     string            `json:"country" valid:"required"`
	City                        string            `json:"city" valid:"required"`
	Height                      float64           `json:"height" valid:"required"`
	Status                      string            `json:"status,omitempty" valid:"required"`
	FriendsCount                int               `json:"friendsCount"`
	NotificationsCount          int               `json:"notificationsCount"`
	MeetingPurpose              string            `json:"lookingFor,omitempty"`
	Gender                      string            `json:"interestedIn,omitempty"`
	HideDetailsVisibility       bool              `json:"hideDetails" valid:"required"`
	AllowprofileslikeVisibility bool              `json:"allowProfilestolike" valid:"required"`
	EnableVisibility            bool              `json:"enableProfile" valid:"required"`
	VerifiedVisibility          bool              `json:"verifiedVisibility"`
	ProfileStatus               string            `json:"profileStatus,omitempty"`
	ProfileCompletion           int               `json:"profileCompletion"`
	PromptQuestions             []*PromptResponse `json:"promptQuestions,omitempty"`
	Interests                   []string          `json:"interests"`
}

type LikeResponse struct {
	IsProfileLiked bool   `json:"isProfileLiked"`
	Message        string `json:"message,omitempty"`
}
type VerifiedVisibility struct {
	VerifiedVisibility bool `json:"verified_visibility,omitempty"`
}
type Enabled struct {
	Visible bool `json:"visible,omitempty"`
}
type UserVisibility struct {
	HideDetailsVisibility       bool `json:"hideDetailsVisibility,omitempty"`
	AllowprofileslikeVisibility bool `json:"allowprofileslikeVisibility,omitempty"`
}

type LocationCheckInInfo struct {
	IsLocationCheckedIn bool   `json:"isLocationCheckedIn"`
	LocationCoordinates string `json:"locationCoordinates,omitempty"`
}
type EventCheckInInfo struct {
	IsEventCheckedIn bool   `json:"isEventCheckedIn"`
	EventId          uint   `json:"eventId,omitempty"`
	EventCoordinates string `json:"eventCoordinates,omitempty"`
}
