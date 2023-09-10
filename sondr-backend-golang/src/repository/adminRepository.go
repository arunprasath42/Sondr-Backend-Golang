package repository

import (
	//"sondr-backend/src/models"

	"sondr-backend/src/models"
	"sondr-backend/utils/database"
)

/***************************** CREATING SUB-ADMINS ********************************/
func (r *MySqlRepositoryRepo) Insert(req interface{}) error {
	if err := database.DB.Debug().Create(req).Error; err != nil {
		return err
	}
	return nil
}

/**********************INSERTING 4 SUB ADMINS*********************************/
func (r *MySqlRepositoryRepo) Insert4SubAdmins(obj *models.Admins) error {

	admins := []*models.Admins{{Name: "Harshil", Email: "harshil@gmail.com", Password: "123456", Blocked: false, Role: "Admin"}, {Name: "Arun", Email: "arun.ps@accubit.com", Password: "123456", Blocked: false, Role: "Admin"}}

	for _, admin := range admins {
		if err := database.DB.Debug().Create(admin).Error; err != nil {
			return err
		}
	}
	return nil
}

/***************** LOGIN -FINDIND THE ADMIN DETAILS FROM DATABASE *****************/
func (r *MySqlRepositoryRepo) FindAdminLogin(obj interface{}, value ...interface{}) error {
	sqlDB := database.DB.Debug().Table("admins")

	if err := sqlDB.Where("email = ?", value...).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***********************READING ADMIN DATA FROM DATABASE**************************/
func (r *MySqlRepositoryRepo) GetAdmin(obj interface{}, email string) error {
	if err := database.DB.Debug().Where("email = ? ", email).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/**************************************LISTING ALL THE SUB ADMINS******************/
func (r *MySqlRepositoryRepo) ListAllAdmins(obj interface{}, searchFilter string, pageNo int, pageSize int) (int, error) {

	offset := (pageNo - 1) * pageSize
	var count int

	sqlDB := database.DB.Debug().Table("admins").Order("created_at desc")
	if searchFilter != "" {
		sqlDB = sqlDB.Where("concat_ws('',admins.id,admins.name)like ? ", "%"+searchFilter+"%")
	}
	err := sqlDB.Select("admins.id, admins.name, admins.email, admins.password,admins.blocked, admins.created_at").Count(&count).Limit(pageSize).Offset(offset).Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil

}

/****************************** FORGET PASSWORD ***********************************/
func (r *MySqlRepositoryRepo) ForgetPassword(obj interface{}, value interface{}) error {
	sqlDB := database.DB.Debug().Table("admins")
	if err := sqlDB.Where("email = ?", value).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***************************READ ADMIN DATA FROM DATABASE**************************/
func (r *MySqlRepositoryRepo) ReadSubAdmin(obj interface{}, value ...interface{}) error {
	sqlDB := database.DB.Debug().Table("admins")
	if err := sqlDB.Select("admins.name,admins.email,admins.id,admins.photo").Where("id = ? ", value...).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/**************************** UPDATING SUBADMIN DATA *****************************/
func (r *MySqlRepositoryRepo) UpdateSubAdmin(obj interface{}, id int, update interface{}) error {
	if err := database.DB.Debug().Model(obj).Where("id = ?", id).Find(obj).Update(update).Error; err != nil {
		return err
	}
	return nil

}

/*****************************VERIFY PASSWORD ************************************/
func (r *MySqlRepositoryRepo) VerifyPassword(obj interface{}, value ...interface{}) error {
	sqlDB := database.DB.Debug().Table("admins")
	if err := sqlDB.Where("id = ? AND password = ?", value...).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***************************DELETING SUBADMIN*************************************/

func (r *MySqlRepositoryRepo) DeleteByID(obj interface{}, id int) error {
	if err := database.DB.Unscoped().Debug().Where("id IN (?) ", id).First(obj).Delete(obj).Error; err != nil {
		return err
	}
	return nil
}

/***********************BLOCKING SUBADMIN***************************************/
func (r *MySqlRepositoryRepo) BlockSubAdmin(obj interface{}, id int, update interface{}) error {
	if err := database.DB.Debug().Model(obj).Where("id = ?", id).Find(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

/***********************Dashboard Module(Total counts)****************************/
func (r *MySqlRepositoryRepo) GetDashboardCount(obj interface{}, fromtime, totime string) error {

	sqlDB := database.DB.Debug().Table("users")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(*) as total_users, count(case when blocked = 0 then 1 end) as active_users,count(case when active = 0 then 1 end) as inactive_users, count(case when blocked = 1 then 1 end) as blocked_users").Find(obj).Error; err != nil {
		return err
	}

	//HOSTED EVENTS
	sqlDB = database.DB.Debug().Table("events")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(*) as hosted_events").Find(obj).Error; err != nil {
		return err
	}

	//TOTAL CHECKINS - EVENTMETADATA
	sqlDB = database.DB.Debug().Table("event_metadatas")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(case when check_in is NOT NULL AND check_out is NULL then 1 end) as event_total_check_ins").Find(obj).Error; err != nil {
		return err
	}

	//TOTAL CHECKINS - USERLOCATIONS
	sqlDB = database.DB.Debug().Table("user_locations")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(case when check_in is NOT NULL AND check_out is NULL then 1 end) as user_total_check_ins").Find(obj).Error; err != nil {
		return err
	}

	//TOTAL CHECKOUTS - EVENTMETADATAS
	sqlDB = database.DB.Debug().Table("event_metadatas")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(case when check_in is NOT NULL AND check_out is NOT NULL then 1 end) as event_total_check_outs").Find(obj).Error; err != nil {
		return nil
	}

	//TOTAL CHECKOUTS - USERLOCATIONS
	sqlDB = database.DB.Debug().Table("user_locations")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(case when check_in is NOT NULL AND check_out is NOT NULL then 1 end) as user_total_check_outs").Find(obj).Error; err != nil {
		return nil
	}

	//LOCATIONS CHECKINS
	sqlDB = database.DB.Debug().Table("user_locations")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("check_in IS NOT NULL AND check_out IS NULL AND created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(distinct location_name) as location_check_ins").Find(obj).Error; err != nil {
		return nil

	}
	return nil
}

/*************************LAST REGISTERED USERS********************************/
func (r *MySqlRepositoryRepo) LastTenUsers(obj interface{}) error {
	if err := database.DB.Debug().Table("users").Order("created_at desc").Limit(10).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/*************************LIST TOP 10 LOCATION ******************************/
func (r *MySqlRepositoryRepo) TopTenLocationUsers(obj interface{}) error {
	if err := database.DB.Debug().Table("user_locations").Select("location_name, count(*) as number_of_users").Group("location_name").Order("number_of_users desc").Limit(10).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/************************KYC VERIFIED & UNVERIFIED***************************/
func (r *MySqlRepositoryRepo) KycVerifiedUsers(obj interface{}, fromtime, totime string) error {

	sqlDB := database.DB.Debug().Table("kycs")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(*) as total_users, count(case when status = 'Approved' then 1 end) as kyc_verified, count(case when status = 'Requested' OR status= 'DisApproved' then 1 end) as kyc_unverified").Find(obj).Error; err != nil {
		return err
	}

	return nil
}

/************************DAILY MATCHES REPORT********************************/
func (r *MySqlRepositoryRepo) DailyMatchesReport(obj interface{}) error {

	sqlDB := database.DB.Debug().Table("matches")

	if err := sqlDB.Select("count(*) as total_matches, count(case when status = 'Matched' then 'Matched' end) as matched_matches, count(case when status = 0 then 1 end) as unmatched_matches, DATE_FORMAT(created_at,'%w') as day").Group("day").Find(obj).Error; err != nil {
		return err
	}

	return nil
}

/************************SITE VISITOR ANALYTICS******************************/
func (r *MySqlRepositoryRepo) SiteVisitorAnalytics(obj interface{}, fromtime, totime string) error {
	sqlDB := database.DB.Debug().Table("users")
	if fromtime != "" && totime != "" {
		sqlDB = sqlDB.Where("created_at BETWEEN ? AND ?", fromtime, totime)
	}
	if err := sqlDB.Select("count(*) as total_visitors, DATE_FORMAT(created_at,'%m') as month").Group("month").Find(obj).Error; err != nil {
		return err
	}

	return nil
}

/****************************  PROFILE SETUP *****************************/
func (r *MySqlRepositoryRepo) Profilesetup(obj interface{}, id int, update interface{}) error {
	if err := database.DB.Debug().Model(obj).Where("id = ?", id).Find(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil

}

/******************************UPDATING USERPROFILE*****************************/
func (r *MySqlRepositoryRepo) UpdateUserProfile(obj interface{}, id int, update map[string]interface{}) error {
	if err := database.DB.Debug().Model(obj).Where("id = ?", id).Find(obj).Updates(update).Error; err != nil {
		return err
	}

	if database.DB.RowsAffected == 0 {
		return database.DB.Error
	}
	return nil

}

/****************************** UPDATE PROFILEPIC *****************************/
func (r *MySqlRepositoryRepo) UpdateProfilePicture(obj interface{}, id int, update interface{}) error {
	if err := database.DB.Debug().Model(obj).Where("user_id = ?", id).Find(obj).Updates(update).Error; err != nil {
		return err
	}

	if database.DB.RowsAffected == 0 {
		return database.DB.Error
	}
	return nil

}

/******************************UPDATING USERPROFILE*****************************/
func (r *MySqlRepositoryRepo) UpdatewithUserID(obj interface{}, id int, update map[string]interface{}) error {
	if err := database.DB.Debug().Model(obj).Where("user_id = ?", id).Find(obj).Updates(update).Error; err != nil {
		return err
	}

	if database.DB.RowsAffected == 0 {
		return database.DB.Error
	}
	return nil

}

/*****************************UPDATE USER PHOTOS*****************************/
func (r *MySqlRepositoryRepo) UpdateUserPhotos(obj interface{}, user_id uint) error {

	db := database.DB.Debug().Where("user_id = ?", user_id).Find(&models.UserPhotos{})
	if db.RowsAffected == 0 {
		if err := database.DB.Debug().Create(obj).Error; err != nil {
			return err
		}
	}
	return nil

}

/********************************FIRST OR CREATE***************************/
func (r *MySqlRepositoryRepo) FirstOrCreate(obj interface{}, user_id uint) error {

	if err := database.DB.Debug().FirstOrCreate(obj, "user_id = ?", user_id).Error; err != nil {
		return err
	}

	return nil

}

/******************************UPDATING USERPROFILE*****************************/
func (r *MySqlRepositoryRepo) UpdateifExist(obj interface{}, user_id int, update interface{}) error {

	if err := database.DB.Debug().Where("user_id = ? ", user_id).First(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

/*****************************INSERTING INTO LIST ALL COUNTRIES*****************************/
func (r *MySqlRepositoryRepo) ListAllCountries(obj interface{}) error {
	if err := database.DB.Debug().Table("countries").Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/******************************GET USER PROFILE*****************************/
func (r *MySqlRepositoryRepo) GetUserProfile(obj interface{}, id int) error {
	if err := database.DB.Debug().Table("users").Select("users.first_name,users.last_name, users.profile_photo, users.profile_status, DATE_FORMAT(FROM_DAYS(DATEDIFF(now(),dob)), '%Y')+0 AS age, user_photos.photo2, user_photos.photo3, user_photos.photo4, user_photos.photo5, users.phone_no, users.dob, users.country, users.city, users.height, kycs.status, visibilities.allowprofileslike_visibility, visibilities.verified_visibility, visibilities.hide_details_visibility, visibilities.enable_visibility, meeting_purposes.meeting_purpose, user_interested_ins.gender").
		Joins("LEFT JOIN kycs ON users.id = kycs.user_id").
		Joins("LEFT JOIN user_photos on users.id = user_photos.user_id").
		Joins("LEFT JOIN visibilities on users.id = visibilities.user_id").
		Joins("LEFT JOIN meeting_purposes on users.id = meeting_purposes.user_id").
		Joins("LEFT JOIN user_interested_ins on users.id = user_interested_ins.user_id").
		Where("users.id = ?", id).
		Find(obj).Error; err != nil {
		return err
	}

	if err := database.DB.Debug().Table("notifications").Select("count(id) as notifications_count").Where("is_read=false AND receiver_user_id=?", id).Find(obj).Error; err != nil {
		return err
	}

	if err := database.DB.Debug().Table("matches").Select("count(id) as friends_count").Where("status='Matched' AND sender_user_id = ? OR receiver_user_id=?", id, id).Find(obj).Error; err != nil {
		return err
	}

	return nil

}

//GetUserInterests
func (r *MySqlRepositoryRepo) GetUserInterests(obj interface{}, id int) error {
	if err := database.DB.Debug().Table("user_interests").Select("interests").Where("user_id = ?", id).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/*****************************EMAIL VALIDATION *****************************/
func (r *MySqlRepositoryRepo) EmailValidation(obj interface{}, email string) error {
	if err := database.DB.Debug().Table("users").Where("email = ?", email).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/****************************GET USER DETAILS BASED ON ID******************/
func (r *MySqlRepositoryRepo) GetUserByID(obj interface{}, id int) error {
	if err := database.DB.Debug().Table("users").Where("id = ?", id).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

//SoftDeleteUser
func (r *MySqlRepositoryRepo) SoftDeleteUser(obj interface{}, id int) error {

	if err := database.DB.Debug().Table("users").Where("id = ?", id).Delete(obj).Error; err != nil {
		return err
	}

	return nil
}
