package repository

import (
	"sondr-backend/src/models"
)

/***This interface is common for all the repository files****/
type MysqlRepository interface {
	Insert(req interface{}) error
	Insert4SubAdmins(obj *models.Admins) error
	FindById(obj interface{}, id int) error
	Update(obj interface{}, id int, update interface{}) error
	Delete(obj interface{}, id int) error
	GetAdmin(obj interface{}, email string) error
	ReadSubAdmin(obj interface{}, value ...interface{}) error //Read Subadmin data from database
	Find(obj interface{}, tableName string, selectQuery string, whereQuery string, value ...interface{}) error
	FindAdminLogin(obj interface{}, value ...interface{}) error
	ForgetPassword(obj interface{}, value interface{}) error
	ListAllAdmins(obj interface{}, searchFilter string, pageNo int, pageSize int) (int, error)
	UpdateSubAdmin(obj interface{}, id int, update interface{}) error
	LastTenUsers(obj interface{}) error
	TopTenLocationUsers(obj interface{}) error
	KycVerifiedUsers(obj interface{}, fromtime, totime string) error
	DailyMatchesReport(obj interface{}) error
	SiteVisitorAnalytics(obj interface{}, fromtime, totime string) error
	Profilesetup(obj interface{}, id int, update interface{}) error
	EmailValidation(obj interface{}, email string) error
	UpdateUserProfile(obj interface{}, id int, update map[string]interface{}) error
	UpdateProfilePicture(obj interface{}, id int, update interface{}) error
	UpdatewithUserID(obj interface{}, id int, update map[string]interface{}) error
	VerifyPassword(obj interface{}, value ...interface{}) error
	DeleteByID(obj interface{}, id int) error
	BlockSubAdmin(obj interface{}, id int, update interface{}) error
	GetDashboardCount(obj interface{}, fromtime, totime string) error
	FirstOrCreate(obj interface{}, user_id uint) error
	UpdateUserPhotos(obj interface{}, user_id uint) error
	UpdateifExist(obj interface{}, user_id int, update interface{}) error
	InsertPromptQuestion(obj interface{}) error
	ListAllCountries(obj interface{}) error
	GetUserProfile(obj interface{}, id int) error
	DropTable(tableName string) error
	GetUserByID(obj interface{}, id int) error
	GetUserInterests(obj interface{}, id int) error
	SoftDeleteUser(obj interface{}, id int) error

	/**** Repository for event ***/
	FindAllEvents(obj interface{}, pageNo, pageSize int, fromDate, toDate, search string) (int64, error)
	FindEventById(obj interface{}, id uint) error
	FetchInvitedUsers(obj interface{}, id uint) (int64, error)
	FetchAttendiesUsers(obj interface{}, id uint) (int64, error)
	CancelEvent(obj interface{}, reason string, id uint) error
	GetUserEmailId(obj interface{}, id uint) error

	/*****Repository for KYC ***/
	FindAllKycStatusRequest(obj interface{}, pageNo, pageSize int) (int, error)
	FindAllKycStatusVerify(obj interface{}, pageNo, pageSize int) (int, error)
	KycApproveAndDisApprove(id uint, status string, email string) error

	/*************KYC USer Module****************/
	InsertOrUpdateKYC(obj *models.Kycs) error

	/*******USER MODULE*********/
	UpdateBlockUser(obj interface{}, id int, update interface{}) error
	FindAllUsers(obj interface{}, pageNo, pageSize int, from, to, search string) (int64, error)
	FindAllReportedUsers(obj interface{}, pageNo, pageSize int, from, to, search string) (int64, error)
	FindAllBlockedUsers(obj interface{}, pageNo, pageSize int, from, to, search string) (int64, error)
	FindAllHostedEvents(obj interface{}, userId uint, pageNo, pageSize int, from, to string) (int64, error)
	CountAttendedEventsOfUser(obj interface{}, userId uint) error
	FindAllReports(obj interface{}, userId uint, pageNo, pageSize int, from, to string) (int64, error)
	CountReportsOfUser(obj interface{}, userId uint) error
	AboutUser(obj *models.UserInfo, userId int) error
	GetUploadedPhotos(obj interface{}, userId int) error
	GetUsersMetadata(obj *models.UsersMetaData, userId int) error
	GetPromptQuestion(obj interface{}) (int64, error)
	AddAnswerOfPromptQuestion(obj *models.UserAnswers) error
	GetAllPromptOfUser(obj interface{}, userId uint) (int64, error)
	EditPromptAnswers(obj *models.UserAnswers) error
	DeletePromptAnswers(userId int64) error
	GetMatchedProfiles(obj interface{}, userId uint, pageNo, pageSize int, search string) (int64, error)
	FindUserById(obj interface{}, id, userId int) error
	FindUserInterestsById(obj interface{}, userId int) error
	FindUserPromptQuestionAnswersById(obj interface{}, userId int) error
	LocationCheckIn(obj interface{}, userId int) error
	FindUserLocationCoordinates(obj interface{}, userId uint) error
	FindUserProfilesOnLocationCoordinates(obj interface{}, coordinates string, userId uint, pageNo, pageSize int) (int64, error)
	FindEventsbyCoordinates(obj interface{}) error
	LocationCheckOut(obj interface{}, userId uint) error
	LikeUserProfile(obj *models.MatchUserInfo, id, userId uint) error
	UpdateLikeStatus(obj *models.Match, id, userId uint, status string) error
	UpdateMatchedProfileLikeStatus(obj *models.Match, id, userId uint, status string) error
	GetAllNotifications(obj interface{}, userId, pageNo, pageSize int64) (int64, error)
	ReadNotification(obj interface{}, id, userId uint) error
	DeleteLikeRequest(obj interface{}, id int) error
	ReportProfile(report *models.ReportedUsers) error
	InsertUserInterests(interests *models.UserInterests) error
	GetUserVisibility(obj *models.UserVisibility, senderUserId, receiverUserId uint) error
	FindLocationCheckedInOfUser(event *models.EventCheckInInfo, location *models.LocationCheckInInfo, id uint)
	UpdateLastVisitedOfUser(id uint) error
	Logout(id uint) error
	RejectProfile(id, userId uint) error
	DeleteUserInterests(userId int64) error
	/********IOS Create OTP ***********/
	FindUserWithPhoneNo(obj interface{}, phno string) (bool, error)
	FindUserWithReferenceID(obj interface{}, refId string) (bool, error)

	/***********EVENT MODULE USER MODULE **********/
	UpdateEvent(obj interface{}, id uint, update interface{}) error
	DeleteEventMetadatas(id uint) error
	ListInvitedUserId(obj interface{}, id uint) (int64, error)
	InvitedEvents(obj interface{}, id uint) (int64, error)
	HostedEvents(obj interface{}, id uint) (int64, error)
	VerifyEventCheckin(eventid uint, userid uint) error
	GetEventMetadata(obj interface{}, eventid uint, userid uint) (bool, error)
	UpdateEventMetadata(obj interface{}, eventid uint, userid uint, update interface{}) error
	EventCheckout(obj interface{}, eventid uint, userid uint) error
	FindEventByDate(obj interface{}, date string) error
	FindEventByStartTime(obj interface{}, date string) error
	FindEventByEndTime(obj interface{}, date string) error
	UpdateEventExpiryStatus(obj interface{}, eventid uint) error
	UpdateEventStartingStatus(obj interface{}, eventid uint) error
	ListProfilesEventCheckIn(obj interface{}, eventId, userId uint) (int64, error)
}
