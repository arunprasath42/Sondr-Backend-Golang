package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sondr-backend/src/models"
	service "sondr-backend/src/service"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/logging"
	"sondr-backend/utils/response"
	"sondr-backend/utils/validator"
	val "sondr-backend/utils/validator"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

/*******************************CREATING SUB-ADMINS**********************************/

func CreateSubadmin(c *gin.Context) {
	role := c.GetString("role")

	if role != "Admin" {
		c.JSON(http.StatusUnauthorized, response.ErrorMessage(constant.UNAUTHORIZED, errors.New("Unauthorized")))
		return
	}

	reqModel := models.Admins{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidatePasswordString(reqModel.Password); !err {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("password must contain atleast 8 characters, 1 uppercase, 1 lowercase, 1 number and 1 special character")))
		return
	}
	var service = service.AdminService{}
	saved, err := service.CreateSubadmin(&reqModel)
	if err != nil {
		log.Error().Msgf("Error inserting data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Inserted Succesfully in the database: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}

/******************************GENARATE PASSWORD MANUALLY***************************/
func GeneratePassword(c *gin.Context) {
	var service = service.AdminService{}
	generatePassword, err := service.GeneratePassword()
	if err != nil {
		log.Error().Msgf("Error in generating password: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Password Generated: %s", generatePassword)
	c.JSON(http.StatusOK, response.SuccessResponse(generatePassword))
}

/**********************************SUBADMIN- SHARETOGMAIL*********************************************/
func SharetoGmail(c *gin.Context) {
	reqModel := &models.Admins{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	readAdmins, err := service.SharetoGmail(reqModel)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Admin details shared sucessfully: %s", readAdmins)
	c.JSON(http.StatusOK, response.SuccessResponse(readAdmins))
}

/*********************************LOGIN*********************************************/
func Login(c *gin.Context) {
	reqModel := models.Request{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	if err := val.ValidateVariable(reqModel.Email, "required", "Email"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.AdminService{}
	saved, err := service.Login(&reqModel)
	if err != nil {
		log.Error().Msgf("Unable to Login: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Login sucessfull: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))

}

/*********************************LISTING SUBADMINS*********************************************/
func ListSubAdmins(c *gin.Context) {

	role := c.GetString("role")

	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	Subadmin, err := service.ListSubAdmins(reqModel.PageNo, reqModel.PageSize, reqModel.SearchFilter, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("SubAdmins listed sucessfully: %s", Subadmin)
	c.JSON(http.StatusOK, response.SuccessResponse(Subadmin))

}

/*********************************READ SUBADMIN DETAILS*********************************************/
func ReadSubAdmin(c *gin.Context) {

	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	readAdmins, err := service.ReadSubAdmin(reqModel.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("SubAdmin details read sucessfully: %s", readAdmins)
	c.JSON(http.StatusOK, response.SuccessResponse(readAdmins))
}

/*********************************UPDATE SUBADMIN DETAILS*********************************************/

func UpdateSubAdmin(c *gin.Context) {
	reqModel := &models.Admins{}
	profilePicture, profilePictureHandler, err := c.Request.FormFile("profilePicture")
	if profilePicture != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
			return
		}
	}

	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var updateSubadmin = service.AdminService{}
	resp, err := updateSubadmin.UpdateSubAdmin(reqModel, profilePicture, profilePictureHandler)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusAccepted, response.SuccessResponse(resp))
}

/***************************BLOCK SUBADMIN*************************************/
func BlockSubAdmin(c *gin.Context) {

	role := c.GetString("role")

	if role != "Admin" {
		c.JSON(http.StatusUnauthorized, response.ErrorMessage(constant.UNAUTHORIZED, errors.New("Unauthorized")))
		return
	}
	reqModel := &models.AdminBlock{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	resp, err := service.BlockSubAdmin(reqModel)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusAccepted, response.SuccessResponse(resp))
}

/******************************VERIFY PASSWORD*********************************************/
func VerifyPassword(c *gin.Context) {
	reqModel := models.Admins{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	saved, err := service.VerifyPassword(&reqModel)
	if err != nil {
		log.Error().Msgf("Unable to Verify: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Password verified sucessfully: %s", saved)
	c.JSON(http.StatusOK, response.SuccessResponse(saved))
}

/*********************************CHANGE PASSWORD*********************************************/

func ChangePassword(c *gin.Context) {
	reqModel := models.Admin{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	changePassword, err := service.ChangePassword(&reqModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Password changed sucessfully: %s", changePassword)
	c.JSON(http.StatusOK, response.SuccessResponse(changePassword))
}

/*********************************FORGET PASSWORD*********************************************/
func ForgetPassword(c *gin.Context) {
	reqModel := models.AdminForgotPassword{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	if err := val.Validate(reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.AdminService{}
	forgetPassword, err := service.ForgetPassword(&reqModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Password changed sucessfully: %s", forgetPassword)
	c.JSON(http.StatusOK, response.SuccessResponse(forgetPassword))
}

/*********************************DELETE SUBADMIN*********************************************/
func DeleteSubAdmin(c *gin.Context) {

	role := c.GetString("role")
	if role != "Admin" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("permission denied")))
		return
	}

	reqModel := models.Admins{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.AdminService{}
	deleted, err := service.DeleteSubAdmin(&reqModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("SubAdmin deleted sucessfully: %s", deleted)
	c.JSON(http.StatusOK, response.SuccessResponse(deleted))
}

/**************************************DASHBOARD****************************************************/
func DashboardCount(c *gin.Context) {
	reqmodel := &models.Countrequest{}
	if err := c.ShouldBind(&reqmodel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.AdminService{}
	dashboardCount, err := service.DashboardCount(reqmodel.FromTime, reqmodel.ToTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Dashboard Count: %s", dashboardCount)
	c.JSON(http.StatusOK, response.SuccessResponse(dashboardCount))
}

/************************LAST 10 Registered USERS*************************************/
func LastTenUsers(c *gin.Context) {
	var service = service.UserService{}

	result, err := service.ListLastTenUsers()
	if err != nil {
		log.Error().Msgf("Error fetching data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
	}
	logging.Logger.Info("Retrieved all users: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

//TopTenCheckedInLocations
func TopTenCheckedInLocations(c *gin.Context) {
	var service = service.UserService{}

	result, err := service.TopTenLocationUsers()
	if err != nil {
		log.Error().Msgf("Error fetching data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
	}
	logging.Logger.Info("Retrieved all users: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

//KYC PIE CHART
func KycVerifiedUsers(c *gin.Context) {
	reqModel := &models.DailyMatchesReportreq{}

	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}

	result, err := service.KYCVerifiedUnverified(reqModel.FromTime, reqModel.ToTime)
	if err != nil {
		log.Error().Msgf("Error fetching data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
	}
	logging.Logger.Info("Retrieved all users: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/****DAILY MATCHES REPORT FROM MATCHES TABLE****/
func DailyMatchesReport(c *gin.Context) {
	var service = service.UserService{}

	result, err := service.DailyMatchesReport()
	if err != nil {
		log.Error().Msgf("Error fetching data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
	}
	logging.Logger.Info("Retrieved all users: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

//SITE VISITOR ANALYTICS
func SiteVisitorAnalytics(c *gin.Context) {
	reqModel := &models.DailyMatchesReportreq{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}

	result, err := service.SiteVisitorAnalytics(reqModel.FromTime, reqModel.ToTime)
	if err != nil {
		log.Error().Msgf("Error fetching data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
	}
	logging.Logger.Info("Retrieved all users: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/************************************** IOS-PANEL ******************************************************/
/******************** PROFILE SETUP ****************** ***********************/
func UserProfileSetup(c *gin.Context) {

	reqModel := models.UserRequest{}

	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}
	result, err := service.UserProfileSetup(&reqModel)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*************************************** USER INTERESTS **********************/
func UserInterests(c *gin.Context) {

	reqModel := models.UserInterest{}

	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}
	result, err := service.UserInterests(&reqModel)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*************************************** MEETING PURPOSE **********************/
func MeetingPurpose(c *gin.Context) {

	reqModel := models.Purpose{}

	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}
	result, err := service.MeetingPurpose(&reqModel)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*******************************UPDATE USER PROFILE*************************************/
func UpdateUserProfile(c *gin.Context) {

	reqModel := models.UpdateUserProfile{}

	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}
	result, err := service.UpdateUserProfile(&reqModel)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*********************************** UPLOAD PROFILE PIC *************************************/
func UploadProfilePic(c *gin.Context) {

	reqModel := &models.UploadPhotos{}

	/***************************** UPLOAD FILE - 1 *************************************/
	Photo1, PhotoHandler1, err := c.Request.FormFile("photo1")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	/***************************** UPLOAD FILE - 2 *************************************/
	Photo2, PhotoHandler2, err := c.Request.FormFile("photo2")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	/***************************** UPLOAD FILE - 3 *************************************/
	Photo3, PhotoHandler3, err := c.Request.FormFile("photo3")
	if Photo3 != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
			return
		}
	}
	/***************************** UPLOAD FILE - 4 *************************************/
	Photo4, PhotoHandler4, err := c.Request.FormFile("photo4")
	if Photo4 != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
			return
		}
	}
	/***************************** UPLOAD FILE - 5 *************************************/
	Photo5, PhotoHandler5, err := c.Request.FormFile("photo5")
	if Photo5 != nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
			return
		}
	}

	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	result, err := service.UploadProfilePic(reqModel, Photo1, Photo2, Photo3, Photo4, Photo5, PhotoHandler1, PhotoHandler2, PhotoHandler3, PhotoHandler4, PhotoHandler5)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))

}

/************************LIST ALL CITIES BASED ON COUNTRY *************************************/
func Cities(c *gin.Context) {
	reqModel := models.CountryRequest{}
	reqModel.Country = c.Request.FormValue("country")
	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}

	result, err := service.Cities(&reqModel)
	if err != nil {
		log.Error().Msgf("Error while fetching data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/******************************LIST ALL COUNTRIES*************************************/
func ListAllCountries(c *gin.Context) {
	var service = service.UserService{}

	result, err := service.ListAllCountries()
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
	}
	logging.Logger.Info("Retrieved all countries: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*****************************GET USER PROFILE *************************************/
func GetUserProfile(c *gin.Context) {
	reqModel := models.GetUserDetails{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}

	result, err := service.GetUserProfile(&reqModel)
	if err != nil {
		log.Error().Msgf("Error fetching data from the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Retrieved all data's: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/************************EMAIL VALIDATION**********************************/
func EmailValidation(c *gin.Context) {
	email := c.Request.FormValue("Email")
	if err := val.ValidateVariable(email, "required", "email"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("invalid email")))
		return
	}

	var service = service.UserService{}
	result, err := service.EmailValidation(email)
	if err != nil {
		log.Error().Msgf("Error fetching data from the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Retrieved all data's: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*******************************UPDATE MEETING PURPOSE*************************************/
func UpdatePreferences(c *gin.Context) {

	reqModel := models.Purpose{}

	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}
	result, err := service.UpdatePreferences(&reqModel)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("Preferences Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*******************************UPDATE USER DETAILS*************************************/
func UpdateUserDetails(c *gin.Context) {

	reqModel := models.Userdetails{}

	if err := c.ShouldBind(&reqModel); err != nil {
		log.Error().Msgf("Unable to bind the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.Validate(reqModel); err != nil {
		log.Error().Msgf("Unable to validate the reqModel: %s", err.Error())
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	var service = service.UserService{}
	result, err := service.UpdateUserDetails(&reqModel)
	if err != nil {
		log.Error().Msgf("Error updating data into the database: %s", err.Error())
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Updated Succesfully in the database: %s", result)
	c.JSON(http.StatusOK, response.SuccessResponse(result))
}

/*********************************DELETE SUBADMIN*********************************************/
func DeleteUser(c *gin.Context) {

	reqModel := models.DeactivateAccount{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := val.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	var service = service.UserService{}
	deleted, err := service.DeactivateAccount(&reqModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	logging.Logger.Info("User Deactivated sucessfully: %s", deleted)
	c.JSON(http.StatusOK, response.SuccessResponse(deleted))
}

/******************* UPDATE USER PHONE NUMBER*************************************/
func PhoneNumberGenerateOTP(c *gin.Context) {
	var req models.UserPhoneno
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := val.Validate(req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	var service = service.UserService{}
	resp, err := service.PhoneNumberGenerateOTP(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

/******************** EmailIdGenerateOtp *************************************/
func EmailIdGenerateOtp(c *gin.Context) {
	var req models.UserPhoneno
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := validator.ValidateVariable(req.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	var service = service.UserService{}
	resp, err := service.EmailIdGenerateOtp(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

/********************* VerifyEmailOtp *************************************/
func VerifyEmailOtp(c *gin.Context) {
	var req models.RequestLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := val.Validate(req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	fmt.Println(req)
	if err := validator.ValidateVariable(req.ID, "required", "ID"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := validator.ValidateVariable(req.OTP, "required", "otp"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	var service = service.UserService{}
	resp, err := service.VerifyEmailOtp(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

/***********************UPDATE USER PHONENUMBER*************************************/
func UpdatePhoneNumberByID(c *gin.Context) {
	var req models.UserPhoneno

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := val.Validate(req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	var service = service.UserService{}
	resp, err := service.UpdatePhoneNumberByID(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}
