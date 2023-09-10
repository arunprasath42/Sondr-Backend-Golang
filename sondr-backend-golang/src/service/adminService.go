package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"sondr-backend/src/models"
	"sondr-backend/src/repository"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/logging"
	"sondr-backend/utils/notification"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pquerna/otp/totp"
	"github.com/sethvargo/go-password/password"

	"github.com/google/uuid"

	"github.com/spf13/viper"
)

var emailOtp sync.Map

type AdminService struct{}

/********************************INSERT 4 SUBADMINS**********************************/
func Insert4SubAdmins() (string, error) {
	var obj *models.Admins
	if err := repository.Repo.Insert4SubAdmins(obj); err != nil {
		return "", errors.New("unable to create admin")
	}
	return "Admins inserted sucessfully", nil
}

/*******************************CREATING SUB-ADMINS**********************************/
func (c *AdminService) CreateSubadmin(insert *models.Admins) (*models.Admins, error) {
	if insert.Role == "" {
		insert.Role = "SubAdmin"
	}

	//CHECKING IF EMAIL ALREADY EXISTS
	obj := models.Admins{}
	if err := repository.Repo.FindAdminLogin(&obj, insert.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	if err := repository.Repo.Insert(insert); err != nil {
		return nil, errors.New("unable to create subadmin")
	}

	return insert, nil

}

/********************************GENERATE PASSWORD MANUALLY***********************************************/
func (c *AdminService) GeneratePassword() (string, error) {
	genaratePassword, err := password.Generate(12, 1, 1, false, false)
	if err != nil {
		log.Fatal(err)
	}
	return genaratePassword, nil
}

/******************************************LOGIN *******************************************************/
func (c *AdminService) Login(req *models.Request) (*models.Login, error) {
	var resp models.Login
	obj := models.Admins{}

	if err := repository.Repo.FindAdminLogin(&obj, req.Email); err != nil {
		return nil, errors.New("invalid email")
	}
	if obj.Password == req.Password {
		resp.UniqueId = obj.ID
		resp.AccessToken, _ = GenerateToken(&obj)
		return &resp, nil
	} else {
		return nil, errors.New("invalid password")
	}

}

/*************************************** GENERATING JWT TOKEN *******************************************/
func GenerateToken(admin *models.Admins) (string, error) {
	var mySigningKey = []byte(viper.GetString("secret.Key"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = admin.Email
	claims["role"] = admin.Role
	claims["password"] = admin.Password
	claims["id"] = strconv.FormatUint(uint64(admin.ID), 10)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*************************************** FORGET PASSWORD  ************************************************/
func (c *AdminService) ForgetPassword(req *models.AdminForgotPassword) (string, error) {
	obj := models.Admins{}
	if err := repository.Repo.ForgetPassword(&obj, req.Email); err != nil {
		return "", errors.New("email not found")
	}

	newPassword, err := password.Generate(12, 2, 0, false, false)
	if err != nil {
		log.Fatal(err)
	}

	if err := repository.Repo.UpdateSubAdmin(&models.Admins{}, int(obj.ID), &models.Admins{Password: newPassword}); err != nil {
		return "", errors.New("unable to change password")
	}
	body := fmt.Sprintf("Hi %s,\n\nYour new password is : %s\n\nThanks,\nTeam Sondr", obj.Name, newPassword)

	sendMail, err := notification.SendMail(constant.Subject, body, req.Email)
	if err != nil {
		return "", errors.New("unable to send email")
	}
	fmt.Println(sendMail)
	return "Password has been sent to your email", nil
}

/*********************************LISTING SUBADMINS*********************************************/
func (r *AdminService) ListSubAdmins(pageNo int, pageSize int, searchFilter string, role string) (*models.AdminResponse, error) {
	var admin []*models.ListAdmin
	var resp models.AdminResponse

	/*Assigning default pageSize , if the pageSize is  empty*/
	if pageSize == 0 {
		pageSize = 10
	}

	var count int
	count, err := repository.Repo.ListAllAdmins(&admin, searchFilter, pageNo, pageSize)
	if role != "Admin" {

		for _, v := range admin {
			v.Password = "********"
		}
	}
	if err != nil {
		return nil, err
	}
	resp.ListAdmin = admin
	resp.Count = count
	return &resp, nil
}

/*********************************READ SUBADMIN DETAILS*********************************************/
func (c *AdminService) ReadSubAdmin(obj interface{}) (*models.ListAdmin, error) {
	var admin models.ListAdmin

	if err := repository.Repo.ReadSubAdmin(&admin, obj); err != nil {
		return nil, err
	}
	return &admin, nil
}

/*********************************UPDATE SUBADMIN DETAILS*********************************************/
func (c *AdminService) UpdateSubAdmin(req *models.Admins, profilePicture multipart.File, profilePictureHandler *multipart.FileHeader) (string, error) {
	url := viper.GetString("s3.url")

	var fileName string
	if profilePicture != nil {
		fileName = "profilePicture/" + profilePictureHandler.Filename
		errFile := make(chan error)
		go UploadFileS3(profilePicture, fileName, profilePictureHandler.Size, errFile)
		err := <-errFile
		if err != nil {
			logging.Logger.WithField("error in uploading the s3 image3", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
			return "", err

		}
	}
	req.Photo = url + fileName
	if req.Photo == constant.S3URL {
		req.Photo = ""
	}
	if err := repository.Repo.UpdateSubAdmin(&req, int(req.ID), &models.Admins{Name: req.Name, Email: req.Email, Photo: req.Photo, Password: req.Password}); err != nil {
		return "Unable to update user", err
	}
	return "Update Successfull", nil
}

/***********************************VERIFY PASSWORD*********************************************/
func (c *AdminService) VerifyPassword(req *models.Admins) (string, error) {
	obj := models.Admins{}
	if err := repository.Repo.VerifyPassword(&obj, req.ID, req.Password); err != nil {
		return "", errors.New("Password is incorrect. Please try again")
	}
	if req.Password == obj.Password {
		return "Verification Successfull", nil
	} else if req.Password == "" {
		return "", errors.New("please enter the  current password to verify")
	} else {
		return "", errors.New("please enter the correct password")
	}
}

/*********************************CHANGE PASSWORD*********************************************/
func (c *AdminService) ChangePassword(req *models.Admin) (string, error) {
	obj := models.Admins{}
	if err := repository.Repo.FindById(&obj, req.ID); err != nil {
		return "", errors.New("unable to change password")
	}
	if req.NewPassword != req.Password {
		return "", errors.New("new password and confirm password does not match")
	}
	if req.NewPassword == "" || req.Password == "" {
		return "", errors.New("please enter the new password")
	}

	if req.Password == obj.Password {
		return "", errors.New("current password and new password should not be same")
	}

	if err := repository.Repo.UpdateSubAdmin(&models.Admins{}, int(req.ID), &models.Admins{Password: req.NewPassword}); err != nil {
		return "", errors.New("unable to change password")
	}
	return "Password has been changed successfully", nil
}

/************************************** SHARETOGMAIL *******************************************************/
func (c *AdminService) SharetoGmail(req *models.Admins) (string, error) {
	var admins models.Admins

	if err := repository.Repo.FindById(&admins, int(req.ID)); err != nil {
		return "", errors.New("unable to share the details")
	}

	body := fmt.Sprintf("<html><body><h1>Hi %s,</h1><p>Please find the details below :</p><p>Name: %s</p><p>Email: %s</p><p>Password: %s</p><p>Thanks,</p><p>Team Sondr</p></body></html>", admins.Name, admins.Name, admins.Email, admins.Password)

	sendMail, err := notification.HtmlMail(constant.SubjectforSubadmin, body, req.Email)
	if err != nil {
		return "", errors.New("unable to send email")
	}
	fmt.Println(sendMail)
	return "Email sent sucessfully", nil
}

/****************************************DELETE SUBADMIN*****************************************************/

func (c *AdminService) DeleteSubAdmin(del *models.Admins) (string, error) {
	if err := repository.Repo.DeleteByID(del, int(del.ID)); err != nil {
		return "Unable to delete user", err
	}
	return "Delete Sucessfull", nil
}

/****************************************BLOCKING SUBADMIN*****************************************************/

func (c *AdminService) BlockSubAdmin(req *models.AdminBlock) (string, error) {
	if err := repository.Repo.UpdateSubAdmin(&models.Admins{}, int(req.ID), map[string]interface{}{"Blocked": req.Blocked}); err != nil {
		return "Unable to block user", err
	}
	return "Block Sucessfull", nil
}

/****************************************DASHBOARD COUNT*****************************************************/
func (c *AdminService) DashboardCount(fromtime, totime string) (*models.TotalCount, error) {
	var count models.TotalCount

	if err := repository.Repo.GetDashboardCount(&count, fromtime, totime); err != nil {
		return nil, err
	}
	count.Totalcheckins = count.Event_total_check_ins + count.User_total_check_ins
	count.Totalcheckouts = count.Event_total_check_outs + count.User_total_check_outs

	return &count, nil
}

/**************************************LAST REGISTERED USERS*************************************************/
func (c *UserService) ListLastTenUsers() ([]models.LastTenUsers, error) {

	var resp []models.LastTenUsers

	if err := repository.Repo.LastTenUsers(&resp); err != nil {
		return nil, err
	}

	for i := 0; i < len(resp); i++ {
		resp[i].Name = resp[i].FirstName + " " + resp[i].LastName
	}
	return resp, nil
}

/**************************************TOP 10 LOCATIONS CHECKEDIN******************************************/
func (c *UserService) TopTenLocationUsers() ([]models.TopTenLocationUsers, error) {

	var resp []models.TopTenLocationUsers

	if err := repository.Repo.TopTenLocationUsers(&resp); err != nil {
		return nil, err
	}

	for i := 0; i < len(resp); i++ {
		resp[i].Position = i + 1
	}
	return resp, nil
}

/**************************KYC PIECHART****************************/
func (c *UserService) KYCVerifiedUnverified(fromtime, totime string) ([]models.Kycverification, error) {

	var resp []models.Kycverification

	if err := repository.Repo.KycVerifiedUsers(&resp, fromtime, totime); err != nil {
		return nil, err
	}

	for i := 0; i < len(resp); i++ {

		resp[i].KYCVerifiedPercentage = float64((resp[i].KYCVerified * 100) / resp[i].TotalUsers)
		fmt.Println("Verified", resp[i].KYCVerifiedPercentage)
	}

	for j := 0; j < len(resp); j++ {

		resp[j].KYCNotVerifiedPercentage = float64((resp[j].KYCUnverified * 100) / resp[j].TotalUsers)
		fmt.Println("UnVerified", resp[j].KYCNotVerifiedPercentage)
	}
	return resp, nil
}

/****&****************DAILY MATCHES REPORT****************/
func (c *UserService) DailyMatchesReport() (map[string]interface{}, error) {

	var resp []models.DailyMatchesReport

	days := make(map[string]interface{})

	days["0"] = 0
	days["1"] = 0
	days["2"] = 0
	days["3"] = 0
	days["4"] = 0
	days["5"] = 0
	days["6"] = 0

	if err := repository.Repo.DailyMatchesReport(&resp); err != nil {
		return nil, err
	}

	for i := 0; i < len(resp); i++ {
		switch resp[i].Day {
		case "0":
			days["0"] = resp[i].MatchedMatches
		case "1":
			days["1"] = resp[i].MatchedMatches
		case "2":
			days["2"] = resp[i].MatchedMatches
		case "3":
			days["3"] = resp[i].MatchedMatches
		case "4":
			days["4"] = resp[i].MatchedMatches
		case "5":
			days["5"] = resp[i].MatchedMatches
		case "6":
			days["6"] = resp[i].MatchedMatches
		}
	}

	return days, nil

}

/*****************************SITE VISITOR ANALYTICS********************/
func (c *UserService) SiteVisitorAnalytics(fromtime, totime string) (map[string]interface{}, error) {

	var resp []models.SiteVisitorAnalytics
	month := make(map[string]interface{})

	month["01"] = 0
	month["02"] = 0
	month["03"] = 0
	month["04"] = 0
	month["05"] = 0
	month["06"] = 0
	month["07"] = 0
	month["08"] = 0
	month["09"] = 0
	month["10"] = 0
	month["11"] = 0
	month["12"] = 0

	if err := repository.Repo.SiteVisitorAnalytics(&resp, fromtime, totime); err != nil {
		return nil, err
	}

	for i := 0; i < len(resp); i++ {
		switch resp[i].Month {
		case "01":
			month["01"] = resp[i].TotalVisitors
		case "02":
			month["02"] = resp[i].TotalVisitors
		case "03":
			month["03"] = resp[i].TotalVisitors
		case "04":
			month["04"] = resp[i].TotalVisitors
		case "05":
			month["05"] = resp[i].TotalVisitors
		case "06":
			month["06"] = resp[i].TotalVisitors
		case "07":
			month["07"] = resp[i].TotalVisitors
		case "08":
			month["08"] = resp[i].TotalVisitors
		case "09":
			month["09"] = resp[i].TotalVisitors
		case "10":
			month["10"] = resp[i].TotalVisitors
		case "11":
			month["11"] = resp[i].TotalVisitors
		case "12":
			month["12"] = resp[i].TotalVisitors
		}
	}

	return month, nil

}

/****************************************** IOS PANEL *****************************************************/

/********************************** USER PROFILE SETUP ********************************/
func (c *UserService) UserProfileSetup(req *models.UserRequest) (*models.UserLoginResponse, error) {
	var resp models.UserLoginResponse

	var userDetails models.UserDetails

	var users models.Users

	users.FirstName = req.FirstName
	users.LastName = req.LastName
	users.Email = req.Email
	users.Gender = req.Gender
	users.DOB = req.DOB
	users.ReferenceID = req.ReferenceID
	users.ProfileStatus = "Onboarding"

	if req.ReferenceID == "" {
		value, ok := Users.Load(req.PhoneNumber)
		if !ok {
			return nil, errors.New("phoneNumber invalid or missing")
		}

		userDetails, _ = value.(models.UserDetails)

		users.GAuth = userDetails.GAuth
		users.PhoneNo = req.PhoneNumber

	}

	if err := repository.Repo.Insert(&users); err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "users.email") {
			return nil, errors.New("email already exists")
		}
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "users.phone_no") {
			return nil, errors.New("phoneNumber already exists")
		}

		return nil, err
	}

	if req.Gender == "Non-Binary" {

		var groupGender models.GroupGender
		groupGender.UserID = users.ID
		groupGender.GroupCategory = req.GroupCategory
		groupGender.GroupBy = req.GroupBy
		if err := repository.Repo.Insert(&groupGender); err != nil {
			return nil, errors.New("unable to insert user data into gender table")
		}
	}

	var visible models.Visibility
	visible.UserID = users.ID
	visible.GenderVisibility = req.GenderVisibility
	visible.GenderCategoryVisibility = req.GenderCategoryVisibility

	if err := repository.Repo.Insert(&visible); err != nil {
		return nil, errors.New("unable to insert user data into visibility table")
	}

	token, err := GenerateUserToken(&users)
	if err != nil {
		return nil, err
	}

	resp.Token = token
	resp.UserId = users.ID

	return &resp, nil

}

/****************************** EMAIL VALIDATION *************************************/
func (c *UserService) EmailValidation(email string) (*models.ResponseString, error) {
	var users models.Users
	var resp models.ResponseString
	if err := repository.Repo.EmailValidation(&users, email); err != nil {
		resp.Message = "Validation Successful"
		return &resp, nil
	}
	return nil, errors.New("email already exists")
}

/****************************** USER INTERESTS (MALE/FEMALE/NON-BINARY) *************************************/
func (c *UserService) UserInterests(req *models.UserInterest) (*models.ResponseString, error) {
	var interest models.UserInterestedIn

	var resp models.ResponseString

	interest.UserID = req.ID
	interest.Gender = req.Gender

	if err := repository.Repo.FirstOrCreate(&interest, req.ID); err != nil {
		return nil, err
	}

	if err := repository.Repo.UpdateifExist(&models.UserInterestedIn{}, int(interest.UserID), map[string]interface{}{
		"Gender": req.Gender,
	}); err != nil {
		return nil, err
	}

	if err := repository.Repo.UpdatewithUserID(&models.Visibility{}, int(req.ID), map[string]interface{}{
		"InterestVisibility": req.InterestVisibility}); err != nil {
		return nil, err
	}

	resp.Message = "User Interests updated successfully"
	return &resp, nil
}

/**************************************MEETING PURPOSE (DATING/FRIENDSHIP/OTHERS)***********************************/
func (c *UserService) MeetingPurpose(req *models.Purpose) (*models.ResponseString, error) {
	var meetingPurpose models.MeetingPurpose
	var resp models.ResponseString

	meetingPurpose.UserID = req.ID
	meetingPurpose.MeetingPurpose = req.MeetingPurpose

	if err := repository.Repo.FirstOrCreate(&meetingPurpose, req.ID); err != nil {
		return nil, err
	}

	if err := repository.Repo.UpdateifExist(&models.MeetingPurpose{}, int(meetingPurpose.UserID), map[string]interface{}{
		"MeetingPurpose": req.MeetingPurpose,
	}); err != nil {
		return nil, err
	}

	if err := repository.Repo.UpdatewithUserID(&models.Visibility{}, int(req.ID), map[string]interface{}{
		"PurposeVisibility": req.PurposeVisibility}); err != nil {
		return nil, err
	}

	resp.Message = "Meeting purpose updated successfully"
	return &resp, nil
}

/************************************UPDATE USER PROFILE(OCCUPATION/HEIGHT/COUNTRY)*******************************************************/
func (c *UserService) UpdateUserProfile(req *models.UpdateUserProfile) (*models.ResponseString, error) {
	var resp models.ResponseString

	/*****************UPDATING USER PROFILE***********************/
	if err := repository.Repo.UpdateUserProfile(&models.Users{}, int(req.ID), map[string]interface{}{
		"Occupation": req.Occupation,
		"Height":     req.Height,
		"Country":    req.Country,
		"City":       req.City}); err != nil {
		return nil, errors.New("unable to update user profile")
	}

	/******************UPDATING USER HEIGHT VISIBILITY*************/
	if err := repository.Repo.UpdatewithUserID(&models.Visibility{}, int(req.ID), map[string]interface{}{
		"HeightVisibility":     req.HeightVisibility,
		"OccupationVisibility": req.OccupationVisibility}); err != nil {
		return nil, err
	}

	resp.Message = "User Profile updated Successfully"
	return &resp, nil
}

/************************************ UPLOAD PROFILE PICTURE *************************************************/
func (c *UserService) UploadProfilePic(req *models.UploadPhotos, Photo1, Photo2, Photo3, Photo4, Photo5 multipart.File, PhotoHandler1, PhotoHandler2, PhotoHandler3, PhotoHandler4, PhotoHandler5 *multipart.FileHeader) (*models.ResponseString, error) {

	var users models.Users
	var profilephotos models.UserPhotos
	var resp models.ResponseString

	var fileName1, fileName2, fileName3, fileName4, fileName5 string

	url := viper.GetString("s3.url")
	fileName1 = "glampicture/" + PhotoHandler1.Filename
	fileName2 = "glampicture/" + PhotoHandler2.Filename

	errFile1 := make(chan error)
	errFile2 := make(chan error)

	go UploadFileS3(Photo1, fileName1, PhotoHandler1.Size, errFile1)
	go UploadFileS3(Photo2, fileName2, PhotoHandler2.Size, errFile2)

	profilephotos.Photo1 = url + fileName1
	profilephotos.Photo2 = url + fileName2

	err1, err2 := <-errFile1, <-errFile2
	if err1 != nil {
		logging.Logger.WithField("error in Uploading the s3 image1", err1).WithError(err1).Error(constant.INTERNALSERVERERROR, err1)
		return nil, err1
	}
	if err2 != nil {
		logging.Logger.WithField("error in Uploading the s3 image2", err2).WithError(err2).Error(constant.INTERNALSERVERERROR, err2)
		return nil, err2

	}

	if Photo3 != nil {

		fileName3 = "glampicture/" + PhotoHandler3.Filename
		errFile3 := make(chan error)
		go UploadFileS3(Photo3, fileName3, PhotoHandler3.Size, errFile3)
		err3 := <-errFile3
		if err3 != nil {
			logging.Logger.WithField("error in uploading the s3 image3", err3).WithError(err3).Error(constant.INTERNALSERVERERROR, err3)
			return nil, err3

		}
		profilephotos.Photo3 = url + fileName3
	}

	if Photo4 != nil {

		fileName4 = "glampicture/" + PhotoHandler4.Filename
		fmt.Println("Photo4", fileName4)

		errFile4 := make(chan error)
		go UploadFileS3(Photo4, fileName4, PhotoHandler4.Size, errFile4)
		err4 := <-errFile4
		if err4 != nil {
			logging.Logger.WithField("error in uploading the s3 image4", err4).WithError(err4).Error(constant.INTERNALSERVERERROR, err4)
			return nil, err4

		}
		profilephotos.Photo4 = url + fileName4
	}

	if Photo5 != nil {
		fileName5 = "glampicture/" + PhotoHandler5.Filename
		errFile5 := make(chan error)
		go UploadFileS3(Photo5, fileName5, PhotoHandler5.Size, errFile5)
		err5 := <-errFile5
		if err5 != nil {
			logging.Logger.WithField("error in uploading the s3 image5", err5).WithError(err5).Error(constant.INTERNALSERVERERROR, err5)
			return nil, err5

		}
		profilephotos.Photo5 = url + fileName5
	}

	updateprofilePhoto := make(map[string]interface{})

	if req.IsRegistration {
		updateprofilePhoto["ProfileStatus"] = "Prompts"
	}
	updateprofilePhoto["ProfilePhoto"] = url + fileName1

	/***********************UPDATING PROFILEPHOTO IN USERS TABLE****************************/
	if err := repository.Repo.UpdateUserProfile(&users, int(req.ID), updateprofilePhoto); err != nil {
		return nil, errors.New("unable to uploadprofilephoto")
	}

	profilephotos.UserID = req.ID

	if req.IsRegistration {
		if err := repository.Repo.Insert(&profilephotos); err != nil {
			return nil, errors.New("unable to upload profile photos")
		}
	} else {
		if err := repository.Repo.UpdatewithUserID(&models.UserPhotos{}, int(req.ID), map[string]interface{}{
			"Photo1": profilephotos.Photo1,
			"Photo2": profilephotos.Photo2,
			"Photo3": profilephotos.Photo3,
			"Photo4": profilephotos.Photo4,
			"Photo5": profilephotos.Photo5,
		}); err != nil {
			return nil, errors.New("unable to update user profile")
		}
	}

	// if err := repository.Repo.UpdateUserPhotos(&profilephotos, profilephotos.UserID); err != nil {
	// 	return nil, errors.New("unable to insert data in userPhotos table")
	// }

	resp.Message = "Profile photos uploaded successfully"
	return &resp, nil
}

//********************************API TO LIST ALL COUNTRIES*************************************/
func (c *UserService) ListAllCountries() ([]models.AllCountries, error) {

	var resp []models.AllCountries
	if err := repository.Repo.ListAllCountries(&resp); err != nil {
		return nil, err
	}
	return resp, nil
}

/*************************** CITIES *********************************/
func (c *UserService) Cities(request *models.CountryRequest) (*models.CityResponse, error) {

	var response models.CityResponse
	Country := request.Country
	fmt.Println("Country", Country)

	url := "https://countriesnow.space/api/v0.1/countries/cities"
	method := "POST"

	payload := strings.NewReader(`{"country":"` + request.Country + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	json.Unmarshal(body, &response)

	if !response.Error {
		return &response, nil
	}
	return nil, errors.New("no cities found for the selected country")
}

/*****************GET ALL COUNTRIES FROM 3rdPARTY AND INSERT INTO COUNTRIES TABLE********************/
func GetAllCountries() (string, error) {

	var resp models.CountryResponse
	var Countries []models.Countries

	if err := constant.Httpmethod(nil, "https://countriesnow.space/api/v0.1/countries/codes", "GET", &resp); err != nil {
		return "Unable to fetch data from third party", err
	}

	for _, v := range resp.Data {

		var country models.Countries
		country.Name = v.Name
		country.Code = v.Code
		country.DialCode = v.DialCode

		Countries = append(Countries, country)
	}

	for _, v := range Countries {
		if err := repository.Repo.Insert(&v); err != nil {
			return "Unable to insert data into database", err
		}
	}

	return "Data inserted successfully in countries table", nil
}

func (c *UserService) GetUserProfile(request *models.GetUserDetails) (*models.Userdetails, error) {

	var response models.Userdetails
	var interests []models.UserInterests

	if err := repository.Repo.GetUserInterests(&interests, int(request.ID)); err != nil {
		return nil, err
	}

	var interestsArray []string

	for _, v := range interests {
		interestsArray = append(interestsArray, v.Interests)
	}
	response.Interests = interestsArray

	if err := repository.Repo.GetUserProfile(&response, int(request.ID)); err != nil {
		return nil, err
	}
	if err := repository.Repo.FindUserPromptQuestionAnswersById(&response.PromptQuestions, int(request.ID)); err != nil {
		return nil, err
	}
	var status string
	if response.Status == "Requested" {
		status = "Pending"
	} else if response.Status == "Approved" {
		status = "Verified"
	} else if response.Status == "DisApproved" {
		status = "Disapproved"
	} else {
		status = "Not Submitted"
	}

	var profileCompletion int

	if response.ProfileStatus == "NewUser" {
		profileCompletion = 0
	} else if response.ProfileStatus == "Onboarding" {
		profileCompletion = 25
	} else if response.ProfileStatus == "Prompts" {
		profileCompletion = 50
	} else if response.ProfileStatus == "KYC" {
		profileCompletion = 75
	} else if response.ProfileStatus == "Completed" {
		profileCompletion = 100
	} else {
		profileCompletion = 0
	}

	response.Status = status
	response.ProfileCompletion = profileCompletion

	return &response, nil
}

/*********************** UPDATE PREFERENCES *********************************/
func (c *UserService) UpdatePreferences(request *models.Purpose) (*models.ResponseString, error) {

	var response models.ResponseString

	if err := repository.Repo.UpdatewithUserID(&models.MeetingPurpose{}, int(request.ID), map[string]interface{}{
		"MeetingPurpose": request.MeetingPurpose,
	}); err != nil {
		return nil, errors.New("unable to update meeting purpose")
	}

	if err := repository.Repo.UpdatewithUserID(&models.UserInterestedIn{}, int(request.ID), map[string]interface{}{
		"Gender": request.Gender,
	}); err != nil {
		return nil, errors.New("unable to update user interested in")
	}

	if err := repository.Repo.UpdatewithUserID(&models.Visibility{}, int(request.ID), map[string]interface{}{
		"Verified_visibility": request.Verified_visibility,
	}); err != nil {
		return nil, errors.New("unable to update verified toggle")
	}

	response.Message = "Preferences updated successfully"
	return &response, nil
}

/********************* UPDATE USER PROFILE *********************************/
func (c *UserService) UpdateUserDetails(request *models.Userdetails) (*models.ResponseString, error) {

	var users models.Users
	var response models.ResponseString

	if err := repository.Repo.UpdateUserProfile(&users, int(request.ID), map[string]interface{}{
		"FirstName": request.FirstName,
		"LastName":  request.LastName,
		"DOB":       request.DOB,
		"Country":   request.Country,
		"City":      request.City,
		"Height":    request.Height,
		"Status":    request.Status,
	}); err != nil {
		return nil, errors.New("unable to update profile")
	}

	if err := repository.Repo.UpdatewithUserID(&models.Visibility{}, int(request.ID), map[string]interface{}{
		"HideDetailsVisibility":       request.HideDetailsVisibility,
		"AllowprofileslikeVisibility": request.AllowprofileslikeVisibility,
		"EnableVisibility":            request.EnableVisibility,
	}); err != nil {
		return nil, errors.New("unable to update profile")
	}

	response.Message = "Profile updated successfully"
	return &response, nil
}

/**********DEACTIVATE USER ACCOUNT*********************/
func (c *UserService) DeactivateAccount(request *models.DeactivateAccount) (*models.ResponseString, error) {

	var response models.ResponseString

	var user models.Users
	if err := repository.Repo.GetUserByID(&user, int(request.ID)); err != nil {
		return nil, errors.New("unable to get user details")
	}

	Email := user.Email
	PhoneNo := user.PhoneNo
	FirstName := user.FirstName
	LastName := user.LastName
	DOB := user.DOB

	//MORPH THE DETAILS AS RANDOM UUID AND UPDATE THE USER

	Email = uuid.NewString()
	PhoneNo = uuid.NewString()
	FirstName = uuid.NewString()
	LastName = uuid.NewString()
	DOB = uuid.NewString()

	PhoneNo = PhoneNo[0:25]

	if err := repository.Repo.UpdateUserProfile(&user, int(request.ID), map[string]interface{}{
		"Email":            Email,
		"PhoneNo":          PhoneNo,
		"FirstName":        FirstName,
		"LastName":         LastName,
		"DOB":              DOB,
		"ProfileStatus":    "Deactivated",
		"DeactivateReason": request.DeactivateReason,
	}); err != nil {
		return nil, errors.New("unable to deactivate account")
	}

	if err := repository.Repo.SoftDeleteUser(&models.Users{}, int(request.ID)); err != nil {
		return nil, errors.New("unable to deactivate account")
	}

	response.Message = "Account deactivated successfully"
	return &response, nil
}

/**********************UPDATE PHONENUMBER BY SENDING AN EMAIL OTP TO THE USER EMAIL ID*********************/
func (c *UserService) PhoneNumberGenerateOTP(request *models.UserPhoneno) (*models.UserLoginResponse, error) {

	var resp models.UserLoginResponse
	var user models.Users
	var userDetail models.UserDetails

	if err := repository.Repo.GetUserByID(&user, int(request.ID)); err != nil {
		return nil, errors.New("unable to fetch user")
	}

	if request.PhoneNo == user.PhoneNo {
		return nil, errors.New("phone number already exists")
	}

	newUser, err := repository.Repo.FindUserWithPhoneNo(&user, request.PhoneNo)
	if err != nil {
		return nil, err
	}

	fmt.Println("newUser", newUser)
	if !newUser {
		return nil, errors.New("phone number already exists")
	}

	PhoneNo := request.PhoneNo
	gauth := user.GAuth
	fmt.Println("gauth", gauth)

	value, ok := Users.Load(PhoneNo)
	us, _ := value.(models.UserDetails)
	gauth = us.GAuth
	if !ok {
		key, err := totp.Generate(
			totp.GenerateOpts{
				Issuer:      PhoneNo,
				AccountName: PhoneNo,
				Period:      120,
			})
		if err != nil {
			return nil, err
		}
		userDetail.PhoneNumber = PhoneNo
		userDetail.GAuth = key.Secret()
		Users.LoadOrStore(PhoneNo, userDetail)
		gauth = key.Secret()
	}
	fmt.Println("Gauth secret", gauth)
	otp, err := totp.GenerateCodeCustom(gauth, time.Now(), totp.ValidateOpts{
		Digits: 4,
		Period: 120,
	})
	if err != nil {
		return nil, err
	}

	resp.Otp = otp
	resp.Message = "Four-digit code sent successfully"
	go sendSmsNotification(otp)

	// body := `<!DOCTYPE html>
	// <html lang="en">
	// <head>
	// 	<meta charset="UTF-8">
	// 	<title>Verify your OTP</title>
	// 	<style>
	// 		body {
	// 			background: #191818;
	// 			font-family: 'Open Sans', sans-serif;
	// 			font-size: 14px;
	// 			margin: 0;
	// 			padding: 0;
	// 		}
	// 		.container {
	//             background: #222222;
	//             color: #797979;
	//             border-radius: 5px;
	// 			width: 120%;
	// 			max-width: 1200px;
	// 			margin: 0 auto;
	// 			padding: 20px;
	// 		}
	// 		.container h1 {
	// 			font-size: 24px;
	// 			font-weight: 400;
	// 			margin-bottom: 20px;
	// 		}

	// 		.container h3 {
	// 			background: #0b43f9;
	// 			margin: 0 auto;
	// 			width: max-content;
	// 			padding: 0 10px;
	// 			color: #fff;
	// 			border-radius: 4px;
	// 		}

	// 		.container p {
	// 			font-size: 14px;
	// 			line-height: 1.5;
	// 			margin-bottom: 20px;
	// 		}

	// 		.container input {
	// 			width: 25%;
	// 			padding: 10px;
	// 			border: 1px solid #ccc;
	// 			border-radius: 3px;
	// 			margin-bottom: 20px;
	// 		}
	// 		.container button {
	// 			width: 20%;
	// 			padding: 10px;
	// 			background-color: #0b43f9;
	// 			border: none;
	// 			border-radius: 3px;
	// 			color: #fff;
	// 			cursor: pointer;
	// 		}
	// 		.container button:hover {
	// 			background-color: #0b43f9;
	// 		}
	// 		.container button:disabled {
	// 			background-color: #ccc;
	// 			cursor: not-allowed;
	// 		}

	// 		.container .error {
	// 			color: #ff0000;
	// 			font-weight: bold;
	// 			margin-bottom: 20px;
	// 		}
	// 	</style>
	// </head>
	// <body>
	// 	<div class="container">
	// 		<h2>Dear user,</h2>
	// 		<p>We have received a request an OTP request on your mobile to update your phone number to ` + PhoneNo + `. Please use the OTP below to verify your phone number.</p>
	// 		<h3>OTP: ` + otp + `</h3>
	// 		<p>If you did not request to update your phone number, please ignore this email.</p>
	// 		<p>Regards,</p>
	// 		<p>Team</p>
	// 	</div>
	// </body>
	// </html>`
	// sendMail, err := notification.HtmlMail("Alert:Request to change phonenumber", body, user.Email)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(sendMail)
	return &resp, nil
}

/********************  EMAIL GENERATE OTP******************************************/
func (c *UserService) EmailIdGenerateOtp(request *models.UserPhoneno) (*models.UserLoginResponse, error) {

	var resp models.UserLoginResponse
	var user models.Users
	var useremailDetail models.UserEmailDetails

	if err := repository.Repo.GetUserByID(&user, int(request.ID)); err != nil {
		return nil, errors.New("unable to fetch user")
	}

	Name := user.FirstName

	Email := user.Email
	fmt.Println("Email", Email)
	GAuth := user.GAuth
	fmt.Println("GAuth", GAuth)

	value, ok := emailOtp.Load(Email)
	us, _ := value.(models.UserEmailDetails)
	gauth := us.GAuth
	if !ok {
		key, err := totp.Generate(
			totp.GenerateOpts{
				Issuer:      Email,
				AccountName: Email,
				Period:      120,
			})
		if err != nil {
			return nil, err
		}
		useremailDetail.Email = Email
		useremailDetail.GAuth = GAuth
		emailOtp.LoadOrStore(Email, GAuth)
		gauth = key.Secret()
	}
	fmt.Println("Gauth secret", gauth)
	otp, err := totp.GenerateCodeCustom(GAuth, time.Now(), totp.ValidateOpts{
		Digits: 4,
		Period: 120,
	})
	if err != nil {
		return nil, err
	}

	resp.Otp = otp
	go sendSmsNotification(otp)

	body := `<!DOCTYPE html>
   <html lang="en">
   <head>
		
   </head>	
   <body>
		<div class="container">
		
			<h2>Dear ` + Name + `,</h2>
			<p>We have received a request to change your phone number. </p>
			<p>Please use the OTP below to update your phone number.</p>
			
			<h3>OTP: ` + otp + `</h3>

			<p>If you did not request to update your phone number, please ignore this email.</p>
			<p>Regards,</p>
			<p>Team Sondr</p>
		</div>
		   </body>	
		      </html>`

	sendMail, err := notification.HtmlMail("Request to update phone number", body, user.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println(sendMail)

	resp.Message = "Four-digit code has been sent to your email"
	return &resp, nil
}

/**********************VERIFY EMAIL OTP*****************************/

func (us *UserService) VerifyEmailOtp(req *models.RequestLogin) (*models.UserLoginResponse, error) {
	var user models.Users
	var resp models.UserLoginResponse

	if err := repository.Repo.GetUserByID(&user, int(req.ID)); err != nil {
		return nil, errors.New("unable to fetch user")
	}

	Email := user.Email
	GAuth := user.GAuth

	var useremailDetail models.UserEmailDetails
	value, ok := emailOtp.Load(Email)
	if !ok {
		return nil, errors.New("Invalid four-digit code")
	}
	useremailDetail, _ = value.(models.UserEmailDetails)
	fmt.Println("gauth", GAuth, "email", useremailDetail.Email)

	fmt.Println(GAuth, req.OTP)
	valid, err := totp.ValidateCustom(req.OTP, GAuth, time.Now(), totp.ValidateOpts{
		Digits: 4,
		Period: 120,
	})

	fmt.Println(valid)
	if err != nil {
		return nil, errors.New("Invalid four-digit code")
	}

	if valid {
		resp.Message = "Verified sucessfully"
		return &resp, nil
	}

	return nil, errors.New("Invalid four-digit code")
}

/******************UPDATE PHONENUMBER BASED ON ID*********************/
func (c *UserService) UpdatePhoneNumberByID(request *models.UserPhoneno) (*models.UserLoginResponse, error) {

	var resp models.UserLoginResponse
	var users models.Users
	var userDetails models.UserDetails

	value, ok := Users.Load(request.PhoneNo)
	if !ok {
		return nil, errors.New("phoneNumber invalid or missing")
	}
	userDetails, _ = value.(models.UserDetails)

	if err := repository.Repo.UpdateUserProfile(&users, int(request.ID), map[string]interface{}{"phone_no": userDetails.PhoneNumber, "g_auth": userDetails.GAuth}); err != nil {
		return nil, errors.New("unable to update phone number")
	}

	resp.Message = "Phone number updated successfully"
	return &resp, nil
}
