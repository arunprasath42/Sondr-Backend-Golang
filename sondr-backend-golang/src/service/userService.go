package service

import (
	"errors"
	"math"
	"sondr-backend/src/models"
	"sondr-backend/src/repository"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/distance"
	"sondr-backend/utils/logging"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
)

type UserService struct{}

/********************************************************API to find all users from database********************************************/

func (c *UserService) FindAllUsers(search string, pageNo, pageSize int, from, to string) (*models.UserResponse, error) {
	var user []*models.AllUserResponse
	var resp models.UserResponse
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllUsers(&user, pageNo, pageSize, from, to, search)
	if err != nil {
		logging.Logger.WithError(err).WithField("error", err).Error("error in fetching all the users", err)
		return nil, err
	}

	resp.AllUserResponse = user
	resp.Count = count
	return &resp, nil
}

/***************************************************API to find all reported users from database**********************************************/

func (c *UserService) FindAllReportedUsers(search string, pageNo, pageSize int, from, to string) (*models.UserResponse, error) {
	var user []*models.AllUserResponse
	var resp models.UserResponse
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllReportedUsers(&user, pageNo, pageSize, from, to, search)
	if err != nil {
		logging.Logger.Error("error in fetching all the reported users", err)
		return nil, err
	}
	resp.AllUserResponse = user
	resp.Count = count
	return &resp, nil
}

/*****************************************************API to find all reported users from database*****************************************/

func (c *UserService) FindAllBlockedUsers(search string, pageNo, pageSize int, from, to string) (*models.UserResponse, error) {
	var user []*models.AllUserResponse
	var resp models.UserResponse
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllBlockedUsers(&user, pageNo, pageSize, from, to, search)
	if err != nil {
		logging.Logger.Error("error in fetching all the blocked users", err)
		return nil, err
	}
	resp.AllUserResponse = user
	resp.Count = count
	return &resp, nil
}

/***************************************************API to find all hosted events of users from database****************************************/

func (c *UserService) FindAllHostedEvents(userId uint, pageNo, pageSize int, from, to string) (*models.UserEventResponse, error) {
	var user []*models.Eventresponse
	var total models.TotalEventsRecord
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllHostedEvents(&user, userId, pageNo, pageSize, from, to)
	if err != nil {
		logging.Logger.Error("error in fetching all the hosted events of user", err)
		return nil, err
	}
	err = repository.Repo.CountAttendedEventsOfUser(&total, userId)
	if err != nil {
		return nil, err
	}
	response := models.UserEventResponse{
		Count:          count,
		EventInfo:      user,
		AttendedEvents: total.AttendedEvents,
	}
	return &response, nil
}

/********************************************API to find all reports of user from database******************************************************/

func (c *UserService) FindAllReports(userId uint, pageNo, pageSize int, from, to string) (*models.ReportInfo, error) {
	var user []*models.ReportsResponse
	var response models.ReportInfo
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllReports(&user, userId, pageNo, pageSize, from, to)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while finding all reports of user:", err)
		return nil, err
	}
	response = models.ReportInfo{
		Count:           count,
		ReporteResponse: user,
	}
	return &response, nil
}

/************************************************API to block and unblock users from database***************************************************/

func (c *UserService) BlockAndUnblockUser(id, userId uint) (string, error) {
	user := models.Users{}
	if err := repository.Repo.FindById(&user, int(userId)); err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while blocking and unblocking user:", err)
		return "Unable to find user", err
	}
	update := map[string]interface{}{"active": false, "blocked": true, "blocker_id": id, "blocker_type": constant.ADMINBLOCKED}
	if user.Blocked {
		update = map[string]interface{}{"active": true, "blocked": false, "blocker_id": nil, "blocker_type": nil}
	}
	if err := repository.Repo.UpdateBlockUser(&user, int(userId), update); err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while blocking and unblocking user:", err)
		return "Unable to find user", err
	}
	if user.Blocked {
		return "User is blocked", nil
	} else {
		return "User is unblocked", nil
	}
}

/********************************************************API to fetch about users from database*****************************************************/

func (c *UserService) AboutUser(userId uint) (*models.UserInfo, error) {
	user := models.UserInfo{}
	if err := repository.Repo.AboutUser(&user, int(userId)); err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while getting user info:", err)
		return nil, err
	}
	return &user, nil
}

/***********************************************API to get uploaded photos of users from database*************************************************/

func (c *UserService) GetUploadedPhotos(userId uint) (*models.UserResponse, error) {
	resp := &models.UserResponse{}
	user := &models.UploadedPhotosresponse{}
	if err := repository.Repo.GetUploadedPhotos(&user, int(userId)); err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while getting uploded photos:", err)
		return nil, err
	}
	resp.UserUploadedPhotos = []*string{&user.Photo1, &user.Photo2, &user.Photo3, &user.Photo4, &user.Photo5}
	return resp, nil
}

/*****************************************************API to get user metadata from database*********************************************************/

func (c *UserService) GetUsersMetadata(userId uint) (*models.UsersMetaData, error) {
	user := models.UsersMetaData{}
	if err := repository.Repo.GetUsersMetadata(&user, int(userId)); err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while finding metadata of user:", err)
		return nil, err
	}
	return &user, nil
}

/*********************************************************************************************************************************************
*                                                              IOS USER MODULE                                                                *
**********************************************************************************************************************************************/

/*************************************************Insert Prompt Questions************************************************/
func (c *UserService) InsertPromptQuestion() (*models.UserResponse, error) {
	user := []models.Questions{
		{Id: 1,
			Question: "Fact about me that surprises people"},
		{Id: 2,
			Question: "The most spontaneous thing i have done"},
		{Id: 3,
			Question: "Dating me is like..."},
		{Id: 4,
			Question: "I want someone who.."},
		{Id: 5,
			Question: "A shower thought i recently had.."},
		{Id: 6,
			Question: "Green flag i look for..."},
		{Id: 7,
			Question: "All i ask is that you..."},
		{Id: 8,
			Question: "I'm a regular at"},
		{Id: 9,
			Question: "My most irrational fear..."},
		{Id: 10,
			Question: "Give me travel tips for..."},
		{Id: 11,
			Question: "Two truths and a lie..."},
		{Id: 12,
			Question: "I geek out on..."},
		{Id: 13,
			Question: "We're the same type of weird if..."},
		{Id: 14,
			Question: "Worst fad I participated in..."},
		{Id: 15,
			Question: "Unusual skills..."},
		{Id: 16,
			Question: "This year I really want to..."},
		{Id: 17,
			Question: "My most irrational fear..."},
		{Id: 18,
			Question: "If you saw the targeted ads I get, you'd think I'm a..."},
	}
	resp := &models.UserResponse{}
	for k := range user {
		if err := repository.Repo.InsertPromptQuestion(&models.Questions{Id: user[k].Id, Question: user[k].Question}); err != nil {
			logging.Logger.WithField("error", err).WithError(err).Error("Error while inserting prompt questions:", err)
			return nil, err
		}
	}
	resp.Message = "Questions inserted successfully."
	return resp, nil
}

/***************************************************API to get prompt questions from database************************************************/
func (c *UserService) GetPromptQuestion() (*models.UserResponse, error) {
	user := []*models.Questions{}
	resp := &models.UserResponse{}
	count, err := repository.Repo.GetPromptQuestion(&user)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while finding prompt questions:", err)
		return nil, err
	}
	resp.Count = count
	resp.QuestionResponse = user
	return resp, nil
}

/************************************************API to save prompt Answer********************************************************/
func (c *UserService) AddAnswerOfPromptQuestion(userId uint, req []*models.AnswerRequest) (*models.UserResponse, error) {
	var resp models.UserResponse
	for k := range req {
		err := repository.Repo.AddAnswerOfPromptQuestion(&models.UserAnswers{UserId: uint(userId), QuestionId: req[k].QuestionId, Answer: req[k].Answer})
		if err != nil {
			logging.Logger.WithField("error", err).WithError(err).Error("Error while finding prompt questions:", err)
			return nil, err
		}
	}
	resp.Message = "Answer added successfully."
	return &resp, nil
}

/***********************************************API to get all Prompt of User***************************************************/
func (c *UserService) GetAllPromptOfUser(userId uint) (*models.QuestionResponse, error) {
	var answer []*models.PromptResponse
	var resp models.QuestionResponse
	count, err := repository.Repo.GetAllPromptOfUser(&answer, userId)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while finding prompt questions:", err)
		return nil, errors.New("error while fetching prompt answers")
	}
	resp.Count = count
	resp.PromptResponse = answer
	return &resp, nil
}

/********************************************Edit Peompts Api*******************************************************/
func (c *UserService) EditPromptAnswers(userId int64, req []*models.AnswerRequest) (*models.UserResponse, error) {
	var resp models.UserResponse
	if err := repository.Repo.DeletePromptAnswers(userId); err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while deleting prompt answer:", err)
		return nil, errors.New("error while deleting prompt answers")
	}
	for k := range req {
		err := repository.Repo.EditPromptAnswers(&models.UserAnswers{UserId: uint(userId), QuestionId: req[k].QuestionId, Answer: req[k].Answer})
		if err != nil {
			logging.Logger.WithField("error", err).WithError(err).Error("Error while updating prompt answer:", err)
			return nil, errors.New("error while updating prompt answers")
		}
	}
	resp.Message = "Prompt answer updated successfully."
	return &resp, nil
}

/********************************************Get Matched Profiles Api***************************************************/
func (c *UserService) GetMatchedProfiles(userId uint, pageNo, pageSize int, search string) (*models.UserResponse, error) {
	if pageSize == 0 {
		pageSize = 10
	}
	resp := models.UserResponse{}
	match := []*models.MatchedProfileResponse{}
	count, err := repository.Repo.GetMatchedProfiles(&match, userId, pageNo, pageSize, search)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error("Error while fetching matched profiles:", err)
		return nil, err
	}
	if count == 0 {
		resp.Message = "No matched profiles found."
		return &resp, nil
	}
	remainder := count % int64(pageSize)
	var totalPages int64
	if remainder == 0 {
		totalPages = count / int64(pageSize)
	} else {
		totalPages = count/int64(pageSize) + 1
	}
	resp.Count = count
	resp.Limit = int64(pageSize)
	resp.TotalPages = totalPages
	resp.CurrentPage = int64(pageNo)
	resp.MatchedProfiles = match
	return &resp, nil
}

/************************************Get User Profile By Id****************************************************/
func (c *UserService) GetUserProfileById(id, userId int64) (*models.UserResponse, error) {
	resp := &models.UserResponse{}
	profile := &models.UserProfileResponse{}
	photos := models.UploadedPhotosresponse{}
	prompts := []*models.PromptResponse{}
	interests := []models.UserInterestsResponse{}
	var interestsResponse []*string
	if err := repository.Repo.FindUserById(&profile, int(id), int(userId)); err != nil {
		return nil, err
	}
	if profile.SenderUserId != id || profile.MatchStatus == "Requested" {
		profile.MatchStatus = ""
	}
	if err := repository.Repo.GetUploadedPhotos(&photos, int(userId)); err != nil {
		return nil, err
	}
	if err := repository.Repo.FindUserInterestsById(&interests, int(userId)); err != nil {
		return nil, err
	}
	if err := repository.Repo.FindUserPromptQuestionAnswersById(&prompts, int(userId)); err != nil {
		return nil, err
	}
	for k := range interests {
		interestsResponse = append(interestsResponse, &interests[k].Interests)
	}
	vals := structs.Values(photos)
	photoResponse := []interface{}{}
	for k := range vals {
		if vals[k] != "" {
			photoResponse = append(photoResponse, vals[k])
		}
	}
	profile.PromptQuestions = prompts
	profile.Interests = interestsResponse
	resp.UserProfileResponse = profile
	resp.UserProfileResponse.UserPhotos = photoResponse
	return resp, nil
}

/*******************************************************Location Checkin Api************************************************************/
func (c *UserService) LocationCheckIn(id uint, currentLatitude, currentLongitude, destinationLatitude, destinationLongitude, locationName string) (*models.UserResponse, error) {
	var resp models.UserResponse
	lat1, err := strconv.ParseFloat(currentLatitude, 64)
	if err != nil {
		return nil, err
	}
	long1, err := strconv.ParseFloat(currentLongitude, 64)
	if err != nil {
		return nil, err
	}
	lat2, err := strconv.ParseFloat(destinationLatitude, 64)
	if err != nil {
		return nil, err
	}
	long2, err := strconv.ParseFloat(destinationLongitude, 64)
	if err != nil {
		return nil, err
	}
	profileEnabled := models.Enabled{}
	if err := repository.Repo.Find(&profileEnabled, "users", "visible", "id=?", id); err != nil {
		return nil, err
	}
	if !profileEnabled.Visible {
		return nil, errors.New("please enable your profile before checkIn")
	}
	if distance.Distance(lat1, long1, lat2, long2) > 1 {
		return nil, errors.New("you can't check in to the location which is more than 1 mile away")
	}
	req := models.UserLocations{
		UserId:       id,
		LocationName: locationName,
		Coordinates:  destinationLatitude + "," + destinationLongitude,
		CheckIn:      time.Now(),
	}
	err = repository.Repo.LocationCheckIn(&req, int(id))
	if err != nil {
		return nil, err
	}
	resp.Message = "You have checked in this location successfully"
	return &resp, nil
}

/***************************************************Get list of LoggedIn Profiles*******************************************************/
func (c *UserService) ListOfProfilesLocationCheckIn(userId uint, latitude, longitude string, pageNo, pageSize int) (*models.UserResponse, error) {
	if pageSize == 0 {
		pageSize = 10
	}
	coordinates := latitude + "," + longitude
	user := []*models.MatchedProfileResponse{}
	resp := models.UserResponse{}
	count, err := repository.Repo.FindUserProfilesOnLocationCoordinates(&user, coordinates, userId, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &resp, nil
	}
	remainder := count % int64(pageSize)
	var totalPages int64
	if remainder == 0 {
		totalPages = count / int64(pageSize)
	} else {
		totalPages = count/int64(pageSize) + 1
	}
	resp.Count = count
	resp.Limit = int64(pageSize)
	resp.TotalPages = totalPages
	resp.CurrentPage = int64(pageNo)
	resp.LocationCheckedInProfiles = user
	return &resp, nil
}

/*******************************************************List of Events near by Locations*******************************************************/
func (c *UserService) ListOfEventsAndUserLocations(userId int64, latitude, longitude string) (*models.MapResponse, error) {
	resp := models.MapResponse{}

	repository.Repo.FindLocationCheckedInOfUser(&resp.EventCheckInInfo, &resp.LocationCheckInInfo, uint(userId))
	if (resp.EventCheckInInfo != models.EventCheckInInfo{}) {
		resp.EventCheckInInfo.IsEventCheckedIn = true
	}
	if (resp.LocationCheckInInfo != models.LocationCheckInInfo{}) {
		resp.LocationCheckInInfo.IsLocationCheckedIn = true
	}

	lat1, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return nil, err
	}
	long1, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil, err
	}
	var lat2, long2 float64
	var Event []*models.EventLocationResponse
	EventsResponse := []*models.EventLocationResponse{}
	profiles := []*models.LocationProfile{}
	LocationResponse := []*models.LocationProfile{}
	if err := repository.Repo.FindUserLocationCoordinates(&profiles, uint(userId)); err != nil {
		return nil, err
	}
	if err := repository.Repo.FindEventsbyCoordinates(&Event); err != nil {
		return nil, err
	}
	if len(Event) == 0 && len(profiles) == 0 {
		return &resp, nil
	}
	for k := range Event {
		lat := strings.Split(Event[k].Coordinates, ",")
		lat2, err = strconv.ParseFloat(lat[0], 64)
		if err != nil {
			return nil, err
		}
		long2, err = strconv.ParseFloat(lat[1], 64)
		if err != nil {
			return nil, err
		}
		dist := distance.Distance(lat1, long1, lat2, long2)
		if dist <= 5 {
			Event[k].Distance = math.Round(dist*1.60934*100000) / 100
			EventsResponse = append(EventsResponse, Event[k])
		}
	}
	for k := range profiles {
		lat := strings.Split(profiles[k].Coordinates, ",")
		lat2, err = strconv.ParseFloat(lat[0], 64)
		if err != nil {
			return nil, err
		}
		long2, err = strconv.ParseFloat(lat[1], 64)
		if err != nil {
			return nil, err
		}
		dist := distance.Distance(lat1, long1, lat2, long2)
		if dist <= 5 {
			profiles[k].Distance = math.Round(dist*1.60934*100000) / 100
			LocationResponse = append(LocationResponse, profiles[k])
		}
	}
	resp.EventLocationResponse = EventsResponse
	resp.UserLocationResponse = LocationResponse
	return &resp, nil
}

/******************************************************Location CheckedOut Api******************************************************/
func (c *UserService) LocationCheckOut(userId int64) (*models.UserResponse, error) {
	resp := &models.UserResponse{}
	userLocation := &models.UserLocations{}
	if err := repository.Repo.LocationCheckOut(userLocation, uint(userId)); err != nil {
		return nil, err
	}
	resp.Message = "You have checked out this location successfully."
	return resp, nil
}

/****************************************************Like User Profile***********************************************************/
func (c *UserService) LikeAndUnlikeUserProfile(senderUserId, receiverUserId uint) (*models.LikeResponse, error) {
	resp := &models.LikeResponse{}
	user := models.Users{}
	match := &models.MatchUserInfo{}
	visibility := &models.UserVisibility{}
	if err := repository.Repo.LikeUserProfile(match, senderUserId, receiverUserId); err != nil {
		return nil, err
	}
	//when both user liked each other
	if match.SenderUserId != senderUserId && match.Status == constant.MATCHSTATUSREQUESTED {

		if err := repository.Repo.UpdateLikeStatus(&models.Match{}, senderUserId, receiverUserId, constant.MATCHSTATUSMATCHED); err != nil {
			return nil, err
		}
		notification := []models.Notifications{
			{
				SenderUserId:   match.ReceiverUserId,
				ReceiverUserId: match.SenderUserId,
				Message:        match.ReceiverUserName + " and your profile matched.",
				Type:           constant.NOTIFICATIONTYPELIKE,
				IsRead:         false,
			},
			{
				SenderUserId:   match.SenderUserId,
				ReceiverUserId: match.ReceiverUserId,
				Message:        match.SenderUserName + " and your profile matched.",
				Type:           constant.NOTIFICATIONTYPELIKE,
				IsRead:         false,
			},
		}

		for k := range notification {
			err := repository.Repo.Insert(&notification[k])
			if err != nil {
				return nil, err
			}
		}
		resp.IsProfileLiked = true
		resp.Message = "Congratulations it's a match!"
		return resp, nil
	}

	//when sender user unlike the profile
	if match.SenderUserId == senderUserId && match.Status == constant.MATCHSTATUSREQUESTED {
		if err := repository.Repo.DeleteLikeRequest(&models.Match{}, int(match.Id)); err != nil {
			return nil, err
		}
		resp.Message = "You have unlike this profile."
		return resp, nil
	}

	//when receiver user unlike the profile.
	if match.SenderUserId != senderUserId && match.Status == constant.MATCHSTATUSMATCHED {
		if err := repository.Repo.UpdateLikeStatus(&models.Match{}, senderUserId, receiverUserId, constant.MATCHSTATUSREQUESTED); err != nil {
			return nil, err
		}
		resp.Message = "You have unlike this profile."
		return resp, nil
	}

	//When sender user unlike the profile after matching
	if match.SenderUserId == senderUserId && match.Status == constant.MATCHSTATUSMATCHED {
		if err := repository.Repo.UpdateMatchedProfileLikeStatus(&models.Match{}, senderUserId, receiverUserId, constant.MATCHSTATUSREQUESTED); err != nil {
			return nil, err
		}
		resp.Message = "You have unlike this profile."
		return resp, nil
	}

	//default like case
	if err := repository.Repo.GetUserVisibility(visibility, senderUserId, receiverUserId); err != nil {
		return nil, err
	}
	var message string
	err := repository.Repo.FindById(&user, int(senderUserId))
	if err != nil {
		return nil, errors.New("sender user name is not found")
	}
	if visibility.HideDetailsVisibility && visibility.AllowprofileslikeVisibility {
		message = user.FirstName + " " + user.LastName + " liked your profile."
	} else {
		message = "Someone liked your profile."
	}

	notification := models.Notifications{
		SenderUserId:   senderUserId,
		ReceiverUserId: receiverUserId,
		Message:        message,
		Type:           constant.NOTIFICATIONTYPELIKE,
		IsRead:         false,
	}
	err = repository.Repo.Insert(&notification)
	if err != nil {
		return nil, err
	}
	resp.IsProfileLiked = true
	resp.Message = "You have liked this profile."
	return resp, nil
}

/*********************************************************List of notifications Of user**********************************************/
func (c *UserService) GetAllNotifications(userId, pageNo, pageSize int64) (*models.UserResponse, error) {
	if pageSize == 0 {
		pageSize = 10
	}
	resp := &models.UserResponse{}
	notification := []*models.NotificationResponse{}
	count, err := repository.Repo.GetAllNotifications(&notification, userId, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		resp.Message = "No notifications found."
		return resp, nil
	}
	remainder := count % int64(pageSize)
	var totalPages int64
	if remainder == 0 {
		totalPages = count / int64(pageSize)
	} else {
		totalPages = count/int64(pageSize) + 1
	}
	resp.Message = "Notifications have found successfully"
	resp.Count = count
	resp.Limit = int64(pageSize)
	resp.TotalPages = totalPages
	resp.CurrentPage = int64(pageNo)
	resp.Notifications = notification
	return resp, nil
}

/******************************************************Read Single Notification**************************************************/
func (c *UserService) ReadNotification(id, userId uint) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	notification := []*models.Notifications{}
	if err := repository.Repo.ReadNotification(&notification, id, userId); err != nil {
		return nil, err
	}
	if userId == 0 {
		resp.Message = "You have read this notification."
	} else {
		resp.Message = "You have read all the notifications."
	}
	return &resp, nil
}

/*******************************************************Report Profile*************************************************************/
func (c *UserService) ReportProfile(reporterUserId, reporteeUserId uint, reason, comment string) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	report := models.ReportedUsers{
		ReporterUserId: reporterUserId,
		ReporteeUserId: reporteeUserId,
		Reason:         reason,
		Comment:        comment,
	}
	if err := repository.Repo.ReportProfile(&report); err != nil {
		return nil, err
	}
	resp.Message = "Profile is reported."
	return &resp, nil
}

/******************************************************Insert User Interests*****************************************************/
func (c *UserService) InsertUserInterests(userId uint, req []*models.InterestRequest) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	if err := repository.Repo.DeleteUserInterests(int64(userId)); err != nil {
		return nil, err
	}
	for k := range req {
		for l := range req[k].Interests {
			err := repository.Repo.InsertUserInterests(&models.UserInterests{UserID: userId, GenreType: req[k].GenreType, Interests: req[k].Interests[l]})
			if err != nil {
				logging.Logger.WithField("error", err).WithError(err).Error("Error while inserting interests:", err)
				return nil, errors.New("error while inserting interests")
			}
		}

	}
	resp.Message = "Interests inserted successfully."
	return &resp, nil
}

/**************************************************CreateAndDropTable****************************************************/
func (c *UserService) DropTable(tableName string) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	if err := repository.Repo.DropTable(tableName); err != nil {
		return nil, err
	}
	resp.Message = "Table deleted successfully."
	return &resp, nil
}

/*************************************************UpdateLastVisitedOfUser***********************************************/
func (c *UserService) UpdateLastVisitedOfUser(id uint) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	if err := repository.Repo.UpdateLastVisitedOfUser(id); err != nil {
		return nil, err
	}
	resp.Message = "Last visited updated successfully."
	return &resp, nil
}

/*************************************************Logout***********************************************/
func (c *UserService) Logout(id uint) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	if err := repository.Repo.Logout(id); err != nil {
		return nil, err
	}
	resp.Message = "You have logged out successfully."
	return &resp, nil
}

/*************************************************Reject Profile*****************************************/
func (c *UserService) RejectProfile(id, userId uint) (*models.UserResponse, error) {
	resp := models.UserResponse{}
	if err := repository.Repo.RejectProfile(id, userId); err != nil {
		return nil, err
	}
	resp.Message = "You have rejected this profile."
	return &resp, nil
}
