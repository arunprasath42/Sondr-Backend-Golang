package repository

import (
	"errors"
	"fmt"
	"sondr-backend/src/models"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/database"
	"sondr-backend/utils/logging"
	"time"
)

/******************************************************Fetching all users data from database*************************************************/

func (r *MySqlRepositoryRepo) FindAllUsers(obj interface{}, pageNo, pageSize int, from, to, search string) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64

	db := database.DB.Debug().Table("users as u")
	if search != "" {
		db = db.Where("concat_ws(' ',first_name,last_name,u.id) LIKE ?", "%"+search+"%")
	}
	if from != "" && to != "" {
		db = db.Where("date(u.created_at) BETWEEN ? AND ?", from, to)
	}
	if from != "" && to != "" && search != "" {
		db = db.Where("(date(u.created_at)  BETWEEN ? AND ?) AND (concat_ws(' ',u.first_name,u.last_name,u.id) LIKE ? )", from, to, "%"+search+"%")
	}
	err := db.Select("u.id,u.first_name, u.last_name,u.last_visited,Max(e.date) last_event_hosted,blocked as is_blocked").
		Group("u.id").
		Joins("left join events as e on e.host_user_id=u.id").Count(&count).Limit(pageSize).Offset(offset).Order("u.created_at DESC").Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil

}

/********************************************************Fetching all reported users data from database****************************************/

func (r *MySqlRepositoryRepo) FindAllReportedUsers(obj interface{}, pageNo, pageSize int, from, to, search string) (int64, error) {

	offset := (pageNo - 1) * pageSize
	var count int64
	db := database.DB.Debug().Table("reported_users as r")
	if search != "" {
		db = db.Where("concat_ws(' ',first_name,last_name,u.id) LIKE ?", "%"+search+"%")
	}
	if from != "" && to != "" {
		db = db.Where("date(r.created_at) BETWEEN ? AND ?", from, to)
	}
	if from != "" && to != "" && search != "" {
		db = db.Where("(date(r.created_at)  BETWEEN ? AND ?) AND (concat_ws(' ',u.first_name,u.last_name,u.id) LIKE ? )", from, to, "%"+search+"%")
	}
	err := db.Select("u.id,u.first_name,u.last_name,u.last_visited,Max(e.date) last_event_hosted, blocked as is_blocked").
		Group("u.id").
		Joins("left join users as u on u.id=r.reportee_user_id").
		Joins("left join events as e on e.host_user_id=r.reportee_user_id").Count(&count).Limit(pageSize).Offset(offset).Order("MAX(r.created_at) DESC").Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

/********************************************************Fetching all the blocked users data from database****************************************/

func (r *MySqlRepositoryRepo) FindAllBlockedUsers(obj interface{}, pageNo, pageSize int, from, to, search string) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64
	var where string
	db := database.DB.Debug().Table("users")
	where = "blocked is true"
	if search != "" {
		fmt.Println("enter")
		where = "blocked is true and concat_ws(' ',first_name,last_name,id) LIKE '%" + search + "%'"
	}
	if from != "" && to != "" {
		where = "blocked is true and date(updated_at) BETWEEN '" + from + "' AND '" + to + "'"
	}
	if from != "" && to != "" && search != "" {
		where = "blocked is true AND (date(updated_at)  BETWEEN '" + from + "' AND '" + to + "') AND (concat_ws(' ',first_name,last_name,id) LIKE '%" + search + "%')"
	}
	err := db.Select("id,first_name,last_name, last_visited, blocked as is_blocked, blocker_type as type").
		Where(where).
		Count(&count).Limit(pageSize).Offset(offset).Order("updated_at DESC").Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil

}

/************************************************Fetching all hosted events data of user from database**********************************************/

func (r *MySqlRepositoryRepo) FindAllHostedEvents(obj interface{}, userId uint, pageNo, pageSize int, from, to string) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64
	db := database.DB.Debug().Table("events")
	db = db.Where("host_user_id = ?", userId)
	if from != "" && to != "" {
		db = db.Where("host_user_id =? AND (date(events.created_at) BETWEEN ? AND ?)", userId, from, to)
	}
	err := db.Select("events.id,event_name,location,event_mode,date,start_time,end_time,status,sum(is_attended=true) as participation_count").
		Group("events.id").
		Joins("left join event_metadatas on events.id=event_metadatas.event_id").
		Count(&count).Limit(pageSize).Offset(offset).Order("events.created_at DESC").Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil

}

/*****************************************************Count the no of reports of a user*************************************************/
func (r *MySqlRepositoryRepo) CountAttendedEventsOfUser(obj interface{}, userId uint) error {
	err := database.DB.Table("event_metadatas").Select("sum(is_attended=true) attended_events").
		Where("invited_user_id = ?", userId).Find(obj).Error
	if err != nil {
		return err
	}
	return nil

}

/*********************************************************Fetching all reports data of user from database******************************************/

func (r *MySqlRepositoryRepo) FindAllReports(obj interface{}, userId uint, pageNo, pageSize int, from, to string) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64
	db := database.DB.Debug().Table("users as u")
	db = db.Where("r.reportee_user_id = ?", userId)
	if from != "" && to != "" {
		db = db.Where("r.reportee_user_id =? AND (date(r.created_at) BETWEEN ? AND ?)", userId, from, to)
	}
	err := db.Select("u.id unique_id,u.first_name,u.last_name,u.active as status,r.created_at as reported_date,reason,comment").
		Joins("left join reported_users as r on r.reporter_user_id=u.id").
		Count(&count).Limit(pageSize).Offset(offset).Order("r.created_at DESC").Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil

}

/*****************************************************Count the no of reports of a user*************************************************/
func (r *MySqlRepositoryRepo) CountReportsOfUser(obj interface{}, userId uint) error {
	err := database.DB.Table("reported_users").Select("count(*) total_reports").
		Where("reportee_user_id = ?", userId).Find(obj).Error
	if err != nil {
		return err
	}
	return nil

}

/********************************************************Fetching all info of user from database*************************************************/

func (r *MySqlRepositoryRepo) AboutUser(obj *models.UserInfo, userId int) error {

	if err := database.DB.Debug().Table("users as u").Select("u.email,u.phone_no,u.profile_photo,u.dob as age, concat(u.country,' ',u.city) as location,u.created_at as created_date,u.last_visited,u.blocked as is_blocked,k.kyc_photo1,k.kyc_photo2,k.status as kyc_verification_status,k.verified_by,k.updated_at as verified_date,s.facebook_url,s.instagram_url").
		Joins("left join kycs as k on k.user_id=u.id").
		Joins("left join user_social_media_details as s on s.user_id=u.id").
		Where("u.id = ? ", userId).First(obj).Error; err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.NOTFOUND, err)
		return err
	}
	logging.Logger.Info("All info of user found successfully.")
	return nil
}

/****************************************************Fetching all uploaded photos of user from database*********************************************/

func (r *MySqlRepositoryRepo) GetUploadedPhotos(obj interface{}, userId int) error {
	if err := database.DB.Debug().Table("user_photos").Where("user_id = ? ", userId).Find(obj).Error; err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.NOTFOUND, err)
		return err
	}
	logging.Logger.Info("All uploaded photos of user found successfully.")
	return nil
}

/******************************************************Fetching user metadata from database***************************************************/

func (r *MySqlRepositoryRepo) GetUsersMetadata(obj *models.UsersMetaData, userId int) error {

	if err := database.DB.Debug().Raw("select "+
		"(SELECT COUNT(*) FROM matches WHERE ? in(sender_user_id,receiver_user_id) and status = 'Matched') as match_count,"+
		"(SELECT COUNT(*) FROM matches WHERE sender_user_id = ? and status = 'Requested') as like_sent_count,"+
		"(SELECT COUNT(*) FROM matches WHERE receiver_user_id = ? and status = 'Requested') as like_received_count,"+
		"(SELECT SUM(no_of_rejections) from matches where receiver_user_id = ?  and status = 'Rejected') as rejections_count,"+
		"((SELECT COALESCE(SUM(no_of_check_in),0) FROM event_metadatas WHERE invited_user_id = ?)+ (select count(*) from user_locations where user_id = ?)) as total_check_ins",
		userId, userId, userId, userId, userId, userId).Find(obj).Error; err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(constant.NOTFOUND, err)
		return err
	}
	logging.Logger.Info("All metadata of user found successfully.")
	return nil
}

/*********************************************************************************************************************************************
*                                                              IOS USER MODULE                                                                *
**********************************************************************************************************************************************/

/****************************************************Get Prompt Questions from database*************************************************/
func (r *MySqlRepositoryRepo) GetPromptQuestion(obj interface{}) (int64, error) {
	var count int64
	if err := database.DB.Table("questions").Find(obj).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

/*****************************************************Add answer of prompt question*************************************************/
func (r *MySqlRepositoryRepo) AddAnswerOfPromptQuestion(obj *models.UserAnswers) error {
	db := database.DB.Debug().Where("user_id = ? and question_id = ?", obj.UserId, obj.QuestionId).First(obj)
	if db.RowsAffected != 0 {
		return errors.New("you have already answered this question")
	}
	if err := database.DB.Debug().Create(obj).Error; err != nil {
		return errors.New("error while saving the answer")
	}
	if err := database.DB.Debug().Where("id = ?", obj.UserId).First(&models.Users{}).Update("profile_status", "KYC").Error; err != nil {
		return errors.New("error while updating the user profile status")
	}
	return nil
}

/*****************************************************API to get all Prompt of user***********************************************/
func (r *MySqlRepositoryRepo) GetAllPromptOfUser(obj interface{}, userId uint) (int64, error) {
	var count int64
	if err := database.DB.Table("user_answers").Select("question,answer").
		Joins("left join questions on questions.id=user_answers.question_id").
		Where("user_id = ?", userId).Count(&count).Find(obj).Error; err != nil {
		return count, err
	}
	return count, nil
}

/*************************************************API edit Prompt answers*********************************************************/
func (r *MySqlRepositoryRepo) EditPromptAnswers(obj *models.UserAnswers) error {
	if err := database.DB.Debug().Create(obj).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************API edit Prompt answers*********************************************************/
func (r *MySqlRepositoryRepo) DeletePromptAnswers(userId int64) error {
	if err := database.DB.Debug().Table("user_answers").
		Where("user_id = ?", userId).Unscoped().Delete(&models.UserAnswers{}).Error; err != nil {
		return err
	}
	return nil
}

/************************************************API to get Matched Profiles*******************************************************/

func (r *MySqlRepositoryRepo) GetMatchedProfiles(obj interface{}, userId uint, pageNo, pageSize int, search string) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64
	if err := database.DB.Debug().Table("users").Select("users.id as user_id,first_name,last_name,profile_photo,DATE_FORMAT(FROM_DAYS(DATEDIFF(now(),dob)), '%Y')+0 AS age").
		Joins("left join reported_users as r on users.id=r.reporter_user_id").
		Where("users.id in(select sender_user_id from matches where receiver_user_id = ? and status = 'Matched' union select receiver_user_id from matches where sender_user_id = ? and status = 'Matched') and blocked is false and users.id not in(select reporter_user_id from reported_users where reportee_user_id = ? union select reportee_user_id from reported_users where reporter_user_id = ?) and concat_ws(' ',first_name,last_name) LIKE ?", userId, userId, userId, userId, "%"+search+"%").
		Count(&count).
		Find(obj).Limit(pageSize).Offset(offset).Error; err != nil {
		return count, errors.New("error while fetching matched profiles")
	}
	return count, nil
}

/**********************************************Find Unblocked user details by id**************************************************/
func (r *MySqlRepositoryRepo) FindUserById(obj interface{}, id, userId int) error {
	if err := database.DB.Debug().Raw("select first_name,last_name,profile_photo,DATE_FORMAT(FROM_DAYS(DATEDIFF(now(),dob)), '%Y')+0 AS age,occupation,country,city,height,gender_visibility,interest_visibility,purpose_visibility,height_visibility, (select sender_user_id from matches where (sender_user_id = ? and receiver_user_id = ?) or (sender_user_id= ? and receiver_user_id= ?)) as sender_user_id,(select status from matches where (sender_user_id = ? and receiver_user_id = ?) or (sender_user_id= ? and receiver_user_id= ?)) as match_status,(select count(*) from matches where (sender_user_id= ? or receiver_user_id = ?) and status = 'Matched') as friend_count from users left join visibilities on visibilities.user_id=users.id where users.id= ?", id, userId, userId, id, id, userId, userId, id, id, id, userId).
		Find(obj).Error; err != nil {
		return errors.New("error while fetching user data")
	}
	return nil
}

/**********************************************Find Unblocked user details by id**************************************************/
func (r *MySqlRepositoryRepo) FindUserInterestsById(obj interface{}, userId int) error {
	if err := database.DB.Debug().Table("user_interests").Select("interests").Where("user_id = ?", userId).Find(obj).Error; err != nil {
		return errors.New("error while fetching interests of user")
	}
	return nil
}

/**********************************************Find Unblocked user details by id**************************************************/
func (r *MySqlRepositoryRepo) FindUserPromptQuestionAnswersById(obj interface{}, userId int) error {
	if err := database.DB.Debug().Table("user_answers").Select("question_id as id,question,answer").
		Joins("left join questions on questions.id=user_answers.question_id").
		Where("user_id = ?", userId).Find(obj).Error; err != nil {
		return errors.New("error while fetching prompt questions and answers")
	}
	return nil
}

/*********************************************Location Checkin Api*************************************************************/
func (r *MySqlRepositoryRepo) LocationCheckIn(obj interface{}, userId int) error {
	db := database.DB.Debug().Table("user_locations").Where("user_id = ? and check_out is NULL", userId).Find(obj)
	if db.RowsAffected != 0 {
		return errors.New("you have already logged in please checkout first")
	}
	if err := database.DB.Debug().Create(obj).Error; err != nil {
		return errors.New("unable to check in the location")
	}
	return nil
}

/*******************************************FindUserLocationCoordinates*****************************************************/
func (r *MySqlRepositoryRepo) FindUserLocationCoordinates(obj interface{}, userId uint) error {
	if err := database.DB.Debug().Table("user_locations").Select("user_id,coordinates,location_name").Where("user_id != ? and (check_in is not null and check_out is null)", userId).Find(obj).Error; err != nil {
		return errors.New("error while fetching coordinates of users")
	}
	return nil
}

/******************************************Fetch Users checkin same location************************************************/
func (r *MySqlRepositoryRepo) FindUserProfilesOnLocationCoordinates(obj interface{}, coordinates string, userId uint, pageNo, pageSize int) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64
	visibility := models.VerifiedVisibility{}
	if err := database.DB.Debug().Table("visibilities").Select("verified_visibility").Where("user_id = ?", userId).First(&visibility).Error; err != nil {
		return count, err
	}
	db := database.DB
	if visibility.VerifiedVisibility {
		db = database.DB.Where("(users.id != ?) and (coordinates = ?) and (check_out is NULL) and (k.status = 'Approved') and (users.id not in(select reportee_user_id from reported_users where reporter_user_id=? union select reporter_user_id from reported_users where reportee_user_id=?))", userId, coordinates, userId, userId)
	} else {
		db = database.DB.Where("(users.id != ?) and (coordinates = ?) and (check_out is NULL) and (users.id not in(select reportee_user_id from reported_users where reporter_user_id=? union select reporter_user_id from reported_users where reportee_user_id=?))", userId, coordinates, userId, userId)
	}

	err := db.Debug().Table("users").Select("users.id as user_id,first_name,last_name,profile_photo,DATE_FORMAT(FROM_DAYS(DATEDIFF(now(),dob)), '%Y')+0 AS age").
		Joins("left join user_locations as ul on ul.user_id=users.id").
		Joins("left join kycs as k on users.id=k.user_id").
		Limit(pageSize).Offset(offset).Count(&count).Find(obj).Error

	if err != nil {
		return count, errors.New("error while fetching list of logged in user for this location")
	}
	return count, nil
}

/************************************************Find Events by coordinates*************************************************************/
func (r *MySqlRepositoryRepo) FindEventsbyCoordinates(obj interface{}) error {
	if err := database.DB.Debug().Table("events").Select("id,host_user_id,event_name,location,event_mode,start_time,end_time,status,coordinates").
		Where("status in('Planned','On-Going') and DATE(date)=CURDATE() and start_time <= CURTIME()+ interval 1 hour").
		Find(obj).Error; err != nil {
		return errors.New("error while fetching list of logged in user for this location")
	}
	return nil
}

/************************************************Location Checkout Api****************************************************************/
func (r *MySqlRepositoryRepo) LocationCheckOut(obj interface{}, userId uint) error {
	if err := database.DB.Debug().Where("user_id= ?", userId).Find(obj).Update("check_out", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************Like User Profile*******************************************************************/
func (r *MySqlRepositoryRepo) LikeUserProfile(obj *models.MatchUserInfo, id, userid uint) error {
	db := database.DB.Debug().Table("matches as m").Select("m.id,u.id as sender_user_id,concat(u.first_name,' ',u.last_name) as sender_user_name,u.profile_photo as sender_profile_photo, us.id as receiver_user_id,concat(us.first_name,' ',us.last_name) as receiver_user_name,us.profile_photo as receiver_profile_photo, m.status").
		Joins("left join users as u on u.id=m.sender_user_id").
		Joins("left join users as us on us.id=m.receiver_user_id").
		Where("((sender_user_id = ? and receiver_user_id = ?) or (sender_user_id = ? and receiver_user_id = ?)) and status in('Requested','Matched')", id, userid, userid, id).Find(obj)
	if db.RowsAffected == 0 {
		obj := &models.Match{SenderUserId: id, ReceiverUserId: userid, Status: constant.MATCHSTATUSREQUESTED}
		if err := database.DB.Debug().Create(&obj).Error; err != nil {
			return err
		}
	}
	return nil
}

/*************************************************Update Like Status*******************************************************************/
func (r *MySqlRepositoryRepo) UpdateLikeStatus(obj *models.Match, id, userid uint, status string) error {
	if err := database.DB.Debug().Where("(sender_user_id = ? and receiver_user_id = ?) or (sender_user_id = ? and receiver_user_id = ?)", id, userid, userid, id).Find(obj).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************Update Like Status*******************************************************************/
func (r *MySqlRepositoryRepo) UpdateMatchedProfileLikeStatus(obj *models.Match, id, userid uint, status string) error {
	if err := database.DB.Debug().Where("sender_user_id = ? and receiver_user_id = ? ", id, userid).Find(obj).Updates(map[string]interface{}{"status": status, "sender_user_id": userid, "receiver_user_id": id}).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************List Of notifications of user*****************************************************/
func (r *MySqlRepositoryRepo) GetAllNotifications(obj interface{}, userId, pageNo, pageSize int64) (int64, error) {
	offset := (pageNo - 1) * pageSize
	var count int64
	if err := database.DB.Debug().Table("notifications as n").
		Select("n.id,n.created_at as notification_time,sender_user_id,receiver_user_id,message,type,is_read,profile_photo as sender_profile_photo").
		Joins("left join users on users.id=n.sender_user_id").
		Where("receiver_user_id = ?", userId).
		Count(&count).Limit(pageSize).Offset(offset).Order("n.created_at desc").Find(obj).Error; err != nil {
		return count, err
	}
	return count, nil
}

/**************************************************Read Single Notification*****************************************************/
func (r *MySqlRepositoryRepo) ReadNotification(obj interface{}, id, userId uint) error {
	db := database.DB.Debug().Table("notifications")
	if userId != 0 {
		db = db.Where("receiver_user_id = ?", userId)
	} else {
		db = db.Where("id = ?", id)
	}
	row := db.Find(obj).Update("is_read", true)
	if row.RowsAffected == 0 {
		return errors.New("no notifications found")
	}
	return nil
}

/*****************************************************Report Profile***********************************************************/
func (r *MySqlRepositoryRepo) ReportProfile(report *models.ReportedUsers) error {
	db := database.DB.Debug().Where("reporter_user_id = ? and reportee_user_id= ?", report.ReporterUserId, report.ReporteeUserId).First(report)
	if db.RowsAffected == 0 {
		if err := database.DB.Debug().Create(report).Error; err != nil {
			return errors.New("error while reporting user")
		}
		return nil
	}
	return nil
}

/*****************************************************Insert User Interests*****************************************************/
func (r *MySqlRepositoryRepo) InsertUserInterests(interests *models.UserInterests) error {
	if err := database.DB.Debug().Create(interests).Error; err != nil {
		return errors.New("error while inserting user interests")
	}
	return nil
}

/*******************************************Delete User Interests*********************************************************/
func (r *MySqlRepositoryRepo) DeleteUserInterests(userId int64) error {
	if err := database.DB.Debug().Where("user_id = ?", userId).Delete(&models.UserInterests{}).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************Get user visibilities for notifications*****************************************************/
func (r *MySqlRepositoryRepo) GetUserVisibility(obj *models.UserVisibility, senderUserId, receiverUserId uint) error {
	if err := database.DB.Debug().Raw("select (select hide_details_visibility from visibilities where user_id =?) hide_details_visibility, (select allowprofileslike_visibility from visibilities where user_id=?) allowprofileslike_visibility", senderUserId, receiverUserId).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************updateLastVisitedOfUser*****************************************************/
func (r *MySqlRepositoryRepo) UpdateLastVisitedOfUser(id uint) error {
	if err := database.DB.Debug().Where("id = ?", id).Find(&models.Users{}).Update("last_visited", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************updateLastVisitedOfUser*****************************************************/
func UpdateActiveStatusOfUser() {
	if err := database.DB.Debug().Where("last_visited+interval 2 month<now() and active is true").Find(&[]models.Users{}).Update("active", false).Error; err != nil {
		logging.Logger.Error("************************Active status is not Updated**************************)")
	}
	logging.Logger.Info("************************Active Status is Updated**************************)")
}

/*************************************************FindLocationCheckedInOfUser*****************************************************/
func (r *MySqlRepositoryRepo) FindLocationCheckedInOfUser(event *models.EventCheckInInfo, location *models.LocationCheckInInfo, id uint) {
	database.DB.Debug().Table("user_locations").Select("coordinates as location_coordinates").Where("user_id = ? and check_in is not null and check_out is null", id).Find(location)
	database.DB.Debug().Table("events as e").Select("e.id as event_id,coordinates as event_coordinates").
		Joins("left join event_metadatas on e.id=event_metadatas.event_id").
		Where("invited_user_id = ? and check_in is not null and check_out is null", id).Find(event)
}

/*************************************************Logout*****************************************************/
func (r *MySqlRepositoryRepo) Logout(id uint) error {
	if err := database.DB.Debug().Where("id = ?", id).Find(&models.Users{}).Updates(map[string]interface{}{"last_visited": time.Now(), "is_logged_in": false}).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************RejectProfile*****************************************************/
func (r *MySqlRepositoryRepo) RejectProfile(id, userId uint) error {
	matchRecord := models.Match{}
	db := database.DB.Debug().Table("matches").Where("sender_user_id = ? and receiver_user_id = ? and status ='Rejected'", id, userId).Find(&matchRecord).Update("no_of_rejections", matchRecord.NoOfRejections+1)
	if db.RowsAffected != 0 {
		return nil
	}
	if err := database.DB.Debug().Create(&models.Match{SenderUserId: id, ReceiverUserId: userId, Status: constant.MATCHSTATUSREJECTED, NoOfRejections: 1}).Error; err != nil {
		return err
	}
	return nil
}
