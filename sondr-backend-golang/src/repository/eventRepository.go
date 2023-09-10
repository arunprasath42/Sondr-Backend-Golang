package repository

import (
	"errors"
	"fmt"
	"sondr-backend/src/models"
	"sondr-backend/utils/database"
	"time"
)

func (r *MySqlRepositoryRepo) FindAllEvents(obj interface{}, pageNo, pageSize int, fromDate, toDate, search string) (int64, error) {

	offset := (pageNo - 1) * pageSize
	var count int64

	db := database.DB.Debug().Table("events")
	if search != "" {
		db = db.Where("concat_ws('',events.event_name,events.id)like ? ", "%"+search+"%")
	}
	if fromDate != "" && toDate != "" {
		db = db.Where("date(events.created_at) BETWEEN ? AND ?", fromDate, toDate)
	}
	if fromDate != "" && toDate != "" && search != "" {
		db = db.Where("concat_ws('',events.event_name,events.id)like ?  AND date(events.created_at) BETWEEN ? AND ?", "%"+search+"%", fromDate, toDate)
	}
	err := db.Select("events.id,events.location,events.event_name,events.event_mode,events.date,events.start_time,events.end_time,events.status,users.first_name,users.last_name").
		Joins("INNER JOIN users ON events.host_user_id=users.id").Count(&count).Limit(pageSize).Offset(offset).Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil
}
func (r *MySqlRepositoryRepo) FindEventById(obj interface{}, id uint) error {

	err := database.DB.Debug().Table("events").Where("events.id=?", id).Select("events.id,events.created_at,events.host_user_id,events.location,events.event_name,events.event_mode,events.date,events.coordinates,events.start_time,events.end_time,events.status,events.password,users.first_name,users.last_name").
		Joins("INNER JOIN users ON events.host_user_id=users.id").Find(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *MySqlRepositoryRepo) FetchInvitedUsers(obj interface{}, id uint) (int64, error) {
	var count int64

	err := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id = ?", id).Select("event_metadatas.invited_user_id, users.first_name, users.last_name").
		Joins("INNER JOIN users ON event_metadatas.invited_user_id=users.id").Find(obj).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *MySqlRepositoryRepo) FetchAttendiesUsers(obj interface{}, id uint) (int64, error) {
	var count int64
	var event models.Events
	err := database.DB.Debug().Table("events").Where("id = ? && status = 'Completed'", id).Find(&event).Error
	if err != nil {
		fmt.Println("error", err)
		return count, err
	}
	err = database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id= ? && is_attended = 1", id).Select("event_metadatas.check_out,event_metadatas.check_in,event_metadatas.invited_user_id,users.first_name,users.last_name").
		Joins("INNER JOIN users ON event_metadatas.invited_user_id=users.id").Count(&count).Scan(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *MySqlRepositoryRepo) CancelEvent(obj interface{}, reason string, id uint) error {
	var update models.Events
	update.Reason = reason
	update.Status = "Cancelled"

	err := database.DB.Debug().Table("events").Where("(status ='Planned' ||  status ='On-Going') && id = ?", id).Find(obj).Update(&update).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) GetUserEmailId(obj interface{}, id uint) error {
	err := database.DB.Debug().Table("event_metadatas").Where("event_id = ?", id).Select("users.email").
		Joins("INNER JOIN users ON event_metadatas.invited_user_id=users.id").Find(obj).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) ListInvitedUserId(obj interface{}, id uint) (int64, error) {
	var count int64
	err := database.DB.Debug().Table("event_metadatas").Select("event_metadatas.invited_user_id").Where("event_id = ?", id).Scan(obj).Count(&count).Error
	if err != nil {
		return count, nil
	}
	return count, nil
}
func (r *MySqlRepositoryRepo) UpdateEvent(obj interface{}, id uint, update interface{}) error {
	err := database.DB.Debug().Table("events").Where("status ='Planned' ||  status ='On-Going' && id = ?", id).Find(obj).Update(update).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) DeleteEventMetadatas(id uint) error {

	err := database.DB.Debug().Unscoped().Table("event_metadatas").Where("event_id = ?", id).Delete(&models.EventMetadatas{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) InvitedEvents(obj interface{}, id uint) (int64, error) {
	var count int64
	err := database.DB.Debug().Table("events").Select("events.id,events.event_name,events.date,events.start_time,events.end_time,events.location,events.coordinates,users.first_name,users.last_name,users.profile_photo").
		Where("(events.status ='Planned' ||  events.status ='On-Going') && event_metadatas.invited_user_id = ?", id).Joins("INNER JOIN users ON events.host_user_id=users.id").Joins("INNER JOIN event_metadatas ON events.id=event_metadatas.event_id").Count(&count).Scan(obj).Error
	if err != nil {
		return count, nil
	}

	return count, nil

}

func (r *MySqlRepositoryRepo) HostedEvents(obj interface{}, id uint) (int64, error) {
	var count int64
	err := database.DB.Debug().Table("events").Select("events.id,events.event_name,events.date,events.start_time,events.end_time,events.location,events.coordinates,users.first_name,users.last_name,users.profile_photo").
		Where("(events.status ='Planned' ||  events.status ='On-Going') && events.host_user_id = ?", id).Joins("INNER JOIN users ON events.host_user_id=users.id").Scan(obj).Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *MySqlRepositoryRepo) VerifyEventCheckin(eventid uint, userid uint) error {
	var result models.EventMetadatas
	dbs := database.DB.Debug().Table("events").Where("events.id = ? && events.status = 'On-Going'", eventid).Find(&models.Events{})
	if dbs.RowsAffected == 0 {
		return errors.New("please check in correct time")
	}
	d := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.invited_user_id = ? && event_metadatas.check_in is NOT NULL && event_metadatas.check_out is NULL", userid).Find(&models.EventMetadatas{})
	if d.RowsAffected > 0 {
		return errors.New("please check out from another event")
	}
	db := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id =? && event_metadatas.invited_user_id = ? && is_attended = ? && check_out is NULL ", eventid, userid, true).First(&result)
	if db.RowsAffected != 0 {
		return errors.New("you have already check in in please checkout first")
	}
	return nil
}

func (r *MySqlRepositoryRepo) GetEventMetadata(obj interface{}, eventid uint, userid uint) (bool, error) {
	db := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id = ? && event_metadatas.invited_user_id = ?", eventid, userid).First(obj)
	if db.RowsAffected == 0 {
		return false, nil
	}
	if err := db.Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *MySqlRepositoryRepo) UpdateEventMetadata(obj interface{}, eventid uint, userid uint, update interface{}) error {

	err := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id = ? && event_metadatas.invited_user_id = ?  ", eventid, userid).Find(obj).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *MySqlRepositoryRepo) EventCheckout(obj interface{}, eventid uint, userid uint) error {

	dbs := database.DB.Debug().Table("events").Where("events.id = ? && (events.status = 'On-Going' || events.status = 'Completed')", eventid).Find(&models.Events{})
	if dbs.RowsAffected == 0 {
		return errors.New("Internal server error")
	}
	if err := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id = ? && event_metadatas.invited_user_id = ? && is_attended = ? && event_metadatas.check_in is NOT NULL ", eventid, userid, true).Find(obj).Update("check_out", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
func (r *MySqlRepositoryRepo) FindEventByDate(obj interface{}, date string) error {

	if err := database.DB.Debug().Table("events").Where("events.date = ?", date).Find(obj).Error; err != nil {
		return err
	}
	return nil
}
func (r *MySqlRepositoryRepo) FindEventByStartTime(obj interface{}, date string) error {
	t1 := time.Now().UTC()
	endDate := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 59, 59, 0, time.UTC)
	if err := database.DB.Debug().Table("events").Where("events.start_time  BETWEEN ? AND ? ", date, endDate).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) UpdateEventExpiryStatus(obj interface{}, eventid uint) error {
	var update models.Events
	update.Status = "Completed"
	if err := database.DB.Debug().Table("events").Where("events.id = ? && events.status = 'On-Going'", eventid).Find(obj).Update(&update).Error; err != nil {
		return err
	}
	return nil

}
func (r *MySqlRepositoryRepo) FindEventByEndTime(obj interface{}, date string) error {
	t1 := time.Now().UTC()
	endDate := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 59, 59, 0, time.UTC)

	if err := database.DB.Debug().Table("events").Where("events.end_time BETWEEN ? AND ? ", date, endDate).Find(obj).Error; err != nil {
		return err
	}
	return nil
}
func (r *MySqlRepositoryRepo) UpdateEventStartingStatus(obj interface{}, eventid uint) error {
	var update models.Events
	update.Status = "On-Going"
	if err := database.DB.Debug().Table("events").Where("events.id = ? && events.status = 'Planned'", eventid).Find(obj).Update(&update).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) ListProfilesEventCheckIn(obj interface{}, eventId, userId uint) (int64, error) {
	var count int64
	visibility := models.VerifiedVisibility{}
	if err := database.DB.Debug().Table("visibilities").Select("verified_visibility").Where("user_id = ?", userId).First(&visibility).Error; err != nil {
		return count, err
	}
	db1 := database.DB
	if visibility.VerifiedVisibility {
		db1 = database.DB.Where("(users.id != ?) and (event_metadatas.is_attended = true) and (event_metadatas.event_id = ?)and (check_out is NULL) and (k.status = 'Approved') and (users.id not in(select reportee_user_id from reported_users where reporter_user_id=? union select reporter_user_id from reported_users where reportee_user_id=?))", userId, eventId, userId, userId)
	} else {
		db1 = database.DB.Where("(users.id != ?) and (event_metadatas.is_attended = true) and (event_metadatas.event_id = ?)and (check_out is NULL) and (users.id not in(select reportee_user_id from reported_users where reporter_user_id=? union select reporter_user_id from reported_users where reportee_user_id=?))", userId, eventId, userId, userId)
	}
	db := database.DB.Debug().Table("event_metadatas").Where("event_metadatas.event_id = ? &&event_metadatas.invited_user_id = ?&& is_attended = true && check_in is NOT NULL  && check_out is NULL ", eventId, userId).Find(&models.EventMetadatas{})
	if db.RowsAffected == 0 {
		return count, errors.New("First check in")
	}
	if err := db1.Debug().Table("users").Select("users.id as user_id,first_name,last_name,profile_photo,DATE_FORMAT(FROM_DAYS(DATEDIFF(now(),dob)), '%Y')+0 AS age").
		Joins("left join event_metadatas  on event_metadatas.invited_user_id=users.id").
		Joins("left join kycs as k on users.id=k.user_id").
		Count(&count).Find(obj).Error; err != nil {
		return count, errors.New("Error while fetching list of logged in user for this location.")
	}
	return count, nil
}

/*************Repository for IOS Log in ********************/

func (r *MySqlRepositoryRepo) FindUserWithPhoneNo(obj interface{}, phno string) (bool, error) {
	if err := database.DB.Table("users").Where("phone_no = ?", phno).Find(obj).Error; err != nil {
		fmt.Println("error in find", err)
		var found bool
		if err.Error() == "record not found" {
			err = nil
			found = true
		}
		return found, err
	}
	return false, nil

}
func (r *MySqlRepositoryRepo) FindUserWithReferenceID(obj interface{}, refId string) (bool, error) {
	if err := database.DB.Table("users").Where("reference_id = ?", refId).Find(obj).Error; err != nil {
		fmt.Println("error in find", err)
		var found bool
		if err.Error() == "record not found" {
			err = nil
			found = true
		}
		return found, err
	}
	return false, nil
}
