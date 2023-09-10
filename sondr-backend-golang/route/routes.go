package route

import (
	controllers "sondr-backend/src/controllers"
	"sondr-backend/src/service"
	"sondr-backend/utils/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/api/v1")

	/***ADD THE ADMIN ROUTES***/
	api.POST("/login", controllers.Login)
	api.POST("/forgotPassword", controllers.ForgetPassword)
	api.POST("/createSubadmin", middleware.Middleware(), controllers.CreateSubadmin)
	api.POST("/generatePassword", middleware.Middleware(), controllers.GeneratePassword)
	api.GET("/listSubadmin", middleware.Middleware(), controllers.ListSubAdmins)
	api.GET("/fetchAdmindetails", middleware.Middleware(), controllers.ReadSubAdmin)
	api.PUT("/updateSubadmin", middleware.Middleware(), controllers.UpdateSubAdmin)
	api.GET("/verifyPassword", middleware.Middleware(), controllers.VerifyPassword)
	api.PUT("/changePassword", middleware.Middleware(), controllers.ChangePassword)
	api.GET("/sharetoGmail", middleware.Middleware(), controllers.SharetoGmail)
	api.DELETE("/deleteSubadmin", middleware.Middleware(), controllers.DeleteSubAdmin)
	api.POST("/blockSubAdmin", middleware.Middleware(), controllers.BlockSubAdmin)
	api.POST("/dashboardcount", middleware.Middleware(), controllers.DashboardCount)
	api.GET("/last10registeredusers", middleware.Middleware(), controllers.LastTenUsers)
	api.GET("/top10CheckedInLocations", middleware.Middleware(), controllers.TopTenCheckedInLocations)
	api.POST("/KycVerifiedUsers", middleware.Middleware(), controllers.KycVerifiedUsers)
	api.POST("/DailyMatchesReport", middleware.Middleware(), controllers.DailyMatchesReport)
	api.POST("/SiteVisitorAnalytics", middleware.Middleware(), controllers.SiteVisitorAnalytics)

	/*******COUNTRIES & CITY ROUTES********/
	api.GET("/listallcountries", controllers.ListAllCountries)
	api.POST("/cities", controllers.Cities)

	api.GET("/listevents", middleware.Middleware(), controllers.ListAllEvents)
	api.GET("/getevent", middleware.Middleware(), controllers.GetEventById)
	api.GET("/getinvitedusers", middleware.Middleware(), controllers.GetInvitedUsers)
	api.GET("/getAttendedUsers", middleware.Middleware(), controllers.GetAttendedUsers)
	api.PUT("/cancelevent", middleware.Middleware(), controllers.CancelEvent)

	api.POST("/uploadKycPhotos", middleware.UserMiddleware(), controllers.UploadKycPhotos)
	api.GET("/listKycStatusRequest", middleware.Middleware(), controllers.ListAllKycRequestStatus)
	api.GET("/listKycStatusVerified", middleware.Middleware(), controllers.ListAllKycVerifiedStatus)
	api.PUT("/kycApproveAndDisApprove", middleware.Middleware(), controllers.KycApprove)

	/******USER ROUTES******/
	api.GET("/findAllUsers", middleware.Middleware(), controllers.FindAllUsers)
	api.GET("/findAllReportedUsers", middleware.Middleware(), controllers.FindAllReportedUsers)
	api.GET("/findAllBlockedUsers", middleware.Middleware(), controllers.FindAllBlockedUsers)
	api.GET("/findAllHostedEvents", middleware.Middleware(), controllers.FindAllHostedEvents)
	api.GET("/findAllReports", middleware.Middleware(), controllers.FindAllReports)
	api.GET("/findUserInfo", middleware.Middleware(), controllers.AboutUser)
	api.GET("/uploadedPhotos", middleware.Middleware(), controllers.GetUploadedPhotos)
	api.PUT("/blockAndUnblockUser", middleware.Middleware(), controllers.BlockAndUnblockUser)
	api.GET("/getUserMetaData", middleware.Middleware(), controllers.GetUsersMetadata)

	/******USER ROUTES-USER PANEL******/
	api.POST("/createOtp", controllers.CreateOTP)
	api.POST("/verfiyOtp", controllers.VerifyOtp)
	api.GET("/emailValidation", controllers.EmailValidation)
	api.POST("/profileSetup", controllers.UserProfileSetup)
	api.POST("/genderInterests", middleware.UserMiddleware(), controllers.UserInterests)
	api.PUT("/meetingPurpose", middleware.UserMiddleware(), controllers.MeetingPurpose)
	api.POST("/updateUserProfile", middleware.UserMiddleware(), controllers.UpdateUserProfile)
	api.POST("/uploadProfilePic", middleware.UserMiddleware(), controllers.UploadProfilePic)
	api.POST("/phonenumberGenerateotp", controllers.PhoneNumberGenerateOTP)
	api.POST("/UpdatePhoneNumberByID", controllers.UpdatePhoneNumberByID)
	api.POST("/VerifyEmailOtp", controllers.VerifyEmailOtp)
	api.POST("/EmailIdGenerateOtp", controllers.EmailIdGenerateOtp)

	api.POST("/insertPromptQuestions", controllers.InsertPromptQuestions)
	api.GET("/getPromptQuestions", middleware.UserMiddleware(), controllers.GetPromptQuestion)
	api.POST("/answerPromptQuestion", middleware.UserMiddleware(), controllers.AddAnswerOfPromptQuestion)
	api.GET("/getAllPromptsOfUser", middleware.UserMiddleware(), controllers.GetAllPromptOfUser)
	api.PUT("/editPromptAnswer", middleware.UserMiddleware(), controllers.EditPromptAnswers)
	api.GET("/getMatchedProfiles", middleware.UserMiddleware(), controllers.GetMatchedProfiles)
	api.GET("/getUserProfile", middleware.UserMiddleware(), controllers.GetUserProfileById)
	api.POST("/locationCheckIn", middleware.UserMiddleware(), controllers.LocationCheckIn)
	api.GET("/listOfLocationCheckedInProfiles", middleware.UserMiddleware(), controllers.ListOfProfilesLocationCheckIn)
	api.GET("/listOfEventsAndUserLocations", middleware.UserMiddleware(), controllers.ListOfEventsAndUserLocations)
	api.GET("/locationCheckOut", middleware.UserMiddleware(), controllers.LocationCheckOut)
	api.POST("/likeUserProfile", middleware.UserMiddleware(), controllers.LikeAndUnlikeUserProfile)
	api.GET("/getAllNotifications", middleware.UserMiddleware(), controllers.GetAllNotifications)
	api.GET("/readNotifications", middleware.UserMiddleware(), controllers.ReadNotification)
	api.POST("/reportProfile", middleware.UserMiddleware(), controllers.ReportProfile)
	api.POST("/addUserInterests", middleware.UserMiddleware(), controllers.InsertUserInterests)
	api.GET("/getUserProfileDetails", middleware.UserMiddleware(), controllers.GetUserProfile)
	api.PUT("/updatePreferences", middleware.UserMiddleware(), controllers.UpdatePreferences)
	api.PUT("/updateUserDetails", middleware.UserMiddleware(), controllers.UpdateUserDetails)
	api.DELETE("/deleteUser", middleware.UserMiddleware(), controllers.DeleteUser)
	api.GET("/updateLastVisited", middleware.UserMiddleware(), controllers.UpdateLastVisitedOfUser)
	api.POST("/dropTable", controllers.DropTable)
	api.GET("/autoMigrateTables", controllers.AutoMigrateTables)
	api.GET("/logout", middleware.UserMiddleware(), controllers.Logout)
	api.POST("/rejectProfile", middleware.UserMiddleware(), controllers.RejectProfile)
	/********EVENT ROUTES USER PANEL *********/

	api.POST("/createEvent", middleware.UserMiddleware(), controllers.CreateEvent)
	api.PUT("/updateEvent", middleware.UserMiddleware(), controllers.UpdateEvent)
	api.GET("/invitedEvents", middleware.UserMiddleware(), controllers.InvitedEvents)
	api.GET("/hostedEvents", middleware.UserMiddleware(), controllers.HostedEvents)
	api.POST("/eventcheckin", middleware.UserMiddleware(), controllers.EventCheckIn)
	api.POST("/eventcheckout", middleware.UserMiddleware(), controllers.EventCheckOut)
	api.POST("/listUsersEventCheckIn", middleware.UserMiddleware(), controllers.ListProfilesEventCheckIn)
	api.GET("/fetchEventByID", middleware.UserMiddleware(), controllers.FetchEventById)

	/***********Fb Login ********/
	api.GET("/login-fb", controllers.FaceBookLogin)
	api.GET("/callback-facebook", controllers.FacebookCallBack)
	api.GET("/login-gl", controllers.GoogleLogin)
	api.GET("/callback-google", controllers.CallBackGoogle)
	api.GET("/login-insta", controllers.InstagramLogin)
	api.GET("/callback-instagram", controllers.CallBackInstagram)
	go service.EventDurationStarting()
	go service.EventDurationEnding()
	service.LoadEvents()

	router.Run(viper.GetString("server.port"))
}
