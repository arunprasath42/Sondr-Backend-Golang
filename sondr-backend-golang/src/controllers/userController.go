package controllers

import (
	"errors"
	"net/http"
	"sondr-backend/migration"
	"sondr-backend/src/models"
	service "sondr-backend/src/service"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/logging"
	"sondr-backend/utils/response"
	val "sondr-backend/utils/validator"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/****************************************************API to find all users from the database*******************************************************/

func FindAllUsers(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUsers, err := service.FindAllUsers(reqModel.SearchFilter, reqModel.PageNo, reqModel.PageSize, reqModel.From, reqModel.To)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All users record found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))
}

/***********************************************API to find all reported users from the database***************************************************/

func FindAllReportedUsers(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUsers, err := service.FindAllReportedUsers(reqModel.SearchFilter, reqModel.PageNo, reqModel.PageSize, reqModel.From, reqModel.To)

	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All reported users record found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))
}

// /**************************************************API to find all blocked users from the database**************************************************/

func FindAllBlockedUsers(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUsers, err := service.FindAllBlockedUsers(reqModel.SearchFilter, reqModel.PageNo, reqModel.PageSize, reqModel.From, reqModel.To)

	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All blocked users record found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))
}

// /*************************************************API to find all hosted events of a user from the database*******************************************/

func FindAllHostedEvents(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readEvents, err := service.FindAllHostedEvents(reqModel.UserId, reqModel.PageNo, reqModel.PageSize, reqModel.From, reqModel.To)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All hosted events of user found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readEvents))
}

/*************************************************API to find all reports of a user from the database***********************************************/

func FindAllReports(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readReports, err := service.FindAllReports(reqModel.UserId, reqModel.PageNo, reqModel.PageSize, reqModel.From, reqModel.To)

	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All reports of user found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readReports))
}

/*********************************************************API to block and unblock a user************************************************/

func BlockAndUnblockUser(c *gin.Context) {
	id := c.GetString("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	request := &models.Request{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(request.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	block, err := service.BlockAndUnblockUser(uint(userId), request.UserId)

	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User is blocked successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(block))
}

/**************************************************API to get all info of user from the database**************************************************/

func AboutUser(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUsers, err := service.AboutUser(reqModel.UserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User info found succesfully:", readUsers)
	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))

}

/*************************************************API to get all uploaded photos of user from the database********************************************/

func GetUploadedPhotos(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUsers, err := service.GetUploadedPhotos(reqModel.UserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User uploaded photos fetched successfully:", readUsers)
	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))

}

/*************************************************API to get all info of user data in the database************************************************/

func GetUsersMetadata(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUsers, err := service.GetUsersMetadata(reqModel.UserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User metadata found succesfully:", readUsers)
	c.JSON(http.StatusOK, response.SuccessResponse(readUsers))

}

/************************************************************************************************************************************************
*                                                         USER CONTROLLER-IOS                                                                   *
*************************************************************************************************************************************************/

/**************************************************API to insert Prompt Questions**************************************************************/

func InsertPromptQuestions(c *gin.Context) {
	var service = service.UserService{}
	question, err := service.InsertPromptQuestion()
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All prompt questions inserted successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(question))
}

/***********************************************API to get prompt questions for user**************************************************************/

func GetPromptQuestion(c *gin.Context) {
	var service = service.UserService{}
	question, err := service.GetPromptQuestion()
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All prompt questions record found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(question))
}

/***********************************************API to answer the prompt questions******************************************************/
func AddAnswerOfPromptQuestion(c *gin.Context) {
	id := c.GetString("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := []*models.AnswerRequest{}
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel[0]); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	saved, err := service.AddAnswerOfPromptQuestion(uint(userId), reqModel)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Prompt question answer saved successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}

/***********************************************API to fetch all prompts of a user************************************************/
func GetAllPromptOfUser(c *gin.Context) {
	id := c.GetString("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	var service = service.UserService{}
	readPrompts, err := service.GetAllPromptOfUser(uint(userId))
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("All prompt questions record found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readPrompts))
}

/*************************************************API to edit Prompt Answers******************************************************/
func EditPromptAnswers(c *gin.Context) {
	id := c.GetString("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := []*models.AnswerRequest{}
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel[0]); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	updatePrompts, err := service.EditPromptAnswers(userId, reqModel)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Prompt answer edited successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(updatePrompts))
}

/*************************************************API to get the Matched Profiles of user**************************************************/

func GetMatchedProfiles(c *gin.Context) {
	id := c.GetString("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readMatchProfile, err := service.GetMatchedProfiles(uint(userId), reqModel.PageNo, reqModel.PageSize, reqModel.SearchFilter)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Matched profiles found successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(readMatchProfile))
}

/*************************************************************Get User Profile by Id*******************************************/
func GetUserProfileById(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	UserID := c.Query("userId")
	UserId, err := strconv.ParseInt(UserID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	if err := val.ValidateVariable(UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUser, err := service.GetUserProfileById(Id, UserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User info found succesfully.", readUser)
	c.JSON(http.StatusOK, response.SuccessResponse(readUser))
}

/*******************************************************Location Checkin Api***********************************************************/
func LocationCheckIn(c *gin.Context) {
	reqModel := models.LocationRequest{}
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readUser, err := service.LocationCheckIn(reqModel.UserId, reqModel.CurrentLatitude, reqModel.CurrentLongitude, reqModel.DestinationLatitude, reqModel.DestinationLongitude, reqModel.DestinationLocationName)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User checked in location succesfully.", readUser)
	c.JSON(http.StatusOK, response.SuccessResponse(readUser))
}

/**************************************************List of Profiles Location check in****************************************************/
func ListOfProfilesLocationCheckIn(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if reqModel.Latitude == "" && reqModel.Longitude == "" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("latitude and longitude is invalid or missing")))
		return
	}
	var service = service.UserService{}
	readUser, err := service.ListOfProfilesLocationCheckIn(uint(Id), reqModel.Latitude, reqModel.Longitude, reqModel.PageNo, reqModel.PageSize)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("List of checked in profiles fetched succesfully.", readUser)
	c.JSON(http.StatusOK, response.SuccessResponse(readUser))
}

/************************************************List of Events near by locations**********************************************/
func ListOfEventsAndUserLocations(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if reqModel.Latitude == "" && reqModel.Longitude == "" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("latitude and longitude is invalid or missing")))
		return
	}
	var service = service.UserService{}
	readUser, err := service.ListOfEventsAndUserLocations(Id, reqModel.Latitude, reqModel.Longitude)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Events fetched succesfully.", readUser)
	c.JSON(http.StatusOK, response.SuccessResponse(readUser))
}

/*************************************************Location Checkout Api***********************************************************/
func LocationCheckOut(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	var service = service.UserService{}
	readUser, err := service.LocationCheckOut(Id)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User checked out location succesfully.", readUser)
	c.JSON(http.StatusOK, response.SuccessResponse(readUser))
}

/***********************************************Like User Profile******************************************************/
func LikeAndUnlikeUserProfile(c *gin.Context) {
	reqModel := &models.LikeRequest{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	like, err := service.LikeAndUnlikeUserProfile(reqModel.SenderUserId, reqModel.ReceiverUserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("You liked another user.", like)
	c.JSON(http.StatusOK, response.SuccessResponse(like))
}

/************************************************List Notifications Of User*****************************************************/
func GetAllNotifications(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readNotifications, err := service.GetAllNotifications(Id, int64(reqModel.PageNo), int64(reqModel.PageSize))
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("List of Notifications fetched succesfully.", readNotifications)
	c.JSON(http.StatusOK, response.SuccessResponse(readNotifications))
}

/*****************************************************Read Notifications*******************************************************/

func ReadNotification(c *gin.Context) {
	reqModel := models.Request{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	readNotification, err := service.ReadNotification(reqModel.ID, reqModel.UserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Notifications read succesfully.", readNotification)
	c.JSON(http.StatusOK, response.SuccessResponse(readNotification))
}

/*****************************************************Report Profile*****************************************************/
func ReportProfile(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := models.ReportRequest{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	report, err := service.ReportProfile(uint(Id), reqModel.ReporteeUserId, reqModel.Reason, reqModel.Comment)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Profile is reported.", report)
	c.JSON(http.StatusOK, response.SuccessResponse(report))
}

/*******************************************************Insert User Interests*************************************************************/
func InsertUserInterests(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := []*models.InterestRequest{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel[0]); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	interest, err := service.InsertUserInterests(uint(Id), reqModel)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Interest is inserted.", interest)
	c.JSON(http.StatusOK, response.SuccessResponse(interest))
}

/*******************************************************UpdateLastVisitedOfUser*************************************************************/
func UpdateLastVisitedOfUser(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	var service = service.UserService{}
	lastVisited, err := service.UpdateLastVisitedOfUser(uint(Id))
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("LastVisited is updated.", lastVisited)
	c.JSON(http.StatusOK, response.SuccessResponse(lastVisited))
}

/*************************************************Create and Delete Table**********************************************************/
func DropTable(c *gin.Context) {
	reqModel := models.DropTable{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	create, err := service.DropTable(reqModel.TableName)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Table is created.", create)
	c.JSON(http.StatusOK, response.SuccessResponse(create))
}

/************************************************* Logout **********************************************************/
func Logout(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	var service = service.UserService{}
	logout, err := service.Logout(uint(Id))
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Logout successfull.", logout)
	c.JSON(http.StatusOK, response.SuccessResponse(logout))
}

/************************************************AutoMigrateTableFunc*****************************************************************/
func AutoMigrateTables(c *gin.Context) {
	migration.Migration()
	c.JSON(http.StatusOK, response.SuccessResponse("Table automigrated successfully."))
}

/************************************************RejectProfile*********************************************************************/
func RejectProfile(c *gin.Context) {
	id := c.GetString("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	reqModel := models.Request{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.UserId, "required", "user_id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	rejectProfile, err := service.RejectProfile(uint(Id), reqModel.UserId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Profile Rejected..", rejectProfile)
	c.JSON(http.StatusOK, response.SuccessResponse(rejectProfile))
}
