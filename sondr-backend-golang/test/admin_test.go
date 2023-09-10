package test

import (
	"fmt"
	"mime/multipart"
	"os"
	"sondr-backend/src/models"
	"sondr-backend/src/service"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func TestAdminService_CreateSubadminValid(t *testing.T) {
	type args struct {
		insert *models.Admins
	}

	test := args{
		insert: &models.Admins{
			Name:     "test",
			Email:    "test3@gmail.com",
			Password: "test@123",
		},
	}
	c := service.AdminService{}
	_, err := c.CreateSubadmin(test.insert)
	fmt.Println(err)
	//database.DB.Unscoped().Where("name = ?", "test").Delete(&models.Admins{})
	if err != nil {
		t.Error(err.Error())
	}
}

func Test_CreateSubadminInvalid(t *testing.T) {
	type args struct {
		insert *models.Admins
	}

	test := args{
		insert: &models.Admins{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "test@123",
		},
	}
	c := service.AdminService{}
	_, err := c.CreateSubadmin(test.insert)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestAdminService_GeneratePassword(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := service.AdminService{}
			_, err := c.GeneratePassword()
			if (err != nil) != tt.wantErr {
				t.Errorf("AdminService.GeneratePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAdminService_LoginValid(t *testing.T) {
	type args struct {
		req *models.Request
	}

	test := args{
		req: &models.Request{
			Email:    "arun.ps@accubits.com",
			Password: "P@ssword123",
		},
	}
	c := service.AdminService{}
	_, err := c.Login(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestAdminService_LoginInvalid(t *testing.T) {
	type args struct {
		req *models.Request
	}

	test := args{
		req: &models.Request{
			Email:    "abc@gmail.com",
			Password: "abc",
		},
	}
	c := service.AdminService{}
	_, err := c.Login(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestGenerateTokenValid(t *testing.T) {
	type args struct {
		admin *models.Admins
	}

	test := args{
		admin: &models.Admins{
			Name:     "test",
			Email:    "arun.ps@gmail.com",
			Password: "P@ssword123",
		},
	}
	_, err := service.GenerateToken(test.admin)

	if err != nil {
		t.Error(err.Error())
	}

}

func Test_ForgetPasswordValid(t *testing.T) {
	type args struct {
		admin *models.AdminForgotPassword
	}

	test := args{
		admin: &models.AdminForgotPassword{
			Email: "arun.ps@accubits.com",
		},
	}
	c := service.AdminService{}
	_, err := c.ForgetPassword(test.admin)

	if err != nil {
		t.Error(err.Error())
	}

}

func Test_ForgetPasswordInValid(t *testing.T) {
	type args struct {
		admin *models.AdminForgotPassword
	}

	test := args{
		admin: &models.AdminForgotPassword{
			Email: "arun.p@accubits.com",
		},
	}
	c := service.AdminService{}
	_, err := c.ForgetPassword(test.admin)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestAdminService_ListSubAdminsValid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchFilter string
		role         string
	}

	test := args{
		pageNo:       1,
		pageSize:     10,
		searchFilter: "",
		role:         "",
	}
	c := service.AdminService{}
	_, err := c.ListSubAdmins(test.pageNo, test.pageSize, test.searchFilter, test.role)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestAdminService_ListSubAdminsInvalid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchFilter string
		role         string
	}

	test := args{
		pageNo:       1,
		pageSize:     10,
		searchFilter: "0",
		role:         "",
	}
	c := service.AdminService{}
	val, _ := c.ListSubAdmins(test.pageNo, test.pageSize, test.searchFilter, test.role)

	if val.Count != 0 {
		t.Error("error msg in find all users invalid")
	}

}

func TestAdminService_ReadSubAdminValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		id: 1,
	}
	c := service.AdminService{}
	_, err := c.ReadSubAdmin(test.id)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestAdminService_ReadSubAdminInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		id: 200,
	}
	c := service.AdminService{}
	_, err := c.ReadSubAdmin(test.id)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestAdminService_UpdateSubAdminInvalid(t *testing.T) {
	type args struct {
		req                   *models.Admins
		profilePicture        multipart.File
		profilePictureHandler *multipart.FileHeader
	}
	var err error
	var fileheader multipart.FileHeader
	file, err := os.Open("../../demo.jpg")

	if err != nil {
		t.Error("error in file length", err)
	}
	ff, err := file.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader.Size = ff.Size()
	fileheader.Filename = ff.Name()

	test := args{
		//pass the parameter here
		req: &models.Admins{
			Model: gorm.Model{
				ID: 0,
			},
		},
		profilePicture:        file,
		profilePictureHandler: &fileheader,
	}
	c := service.AdminService{}
	_, err = c.UpdateSubAdmin(test.req, test.profilePicture, test.profilePictureHandler)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestAdminService_UpdateSubAdminValid(t *testing.T) {
	type args struct {
		req                   *models.Admins
		profilePicture        multipart.File
		profilePictureHandler *multipart.FileHeader
	}
	var err error
	var fileheader multipart.FileHeader
	file, err := os.Open("../../demo.jpg")

	if err != nil {
		t.Error("error in file length", err)
	}
	ff, err := file.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader.Size = ff.Size()
	fileheader.Filename = ff.Name()

	test := args{
		//pass the parameter here
		req: &models.Admins{
			Model: gorm.Model{
				ID: 1,
			},
		},
		profilePicture:        file,
		profilePictureHandler: &fileheader,
	}
	c := service.AdminService{}
	_, err = c.UpdateSubAdmin(test.req, test.profilePicture, test.profilePictureHandler)

	if err != nil {
		t.Error(err.Error())
	}

}

/*
func TestAdminService_UpdateSubAdminValid(t *testing.T) {
	type args struct {
		update *models.Admins
	}

	test := args{
		update: &models.Admins{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "Arunprasath",
			Email:    "arun.ps@accubits.com",
			Password: "P@ssword123",
		},
	}
	c := service.AdminService{}
	_, err := c.UpdateSubAdmin(test.update)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestAdminService_UpdateSubAdminInvalid(t *testing.T) {
	type args struct {
		update *models.Admins
	}

	test := args{
		update: &models.Admins{
			Model: gorm.Model{
				ID: 200,
			},
			Name:     "test",
			Email:    "abc@gmail.com",
			Password: "abc",
		},
	}
	c := service.AdminService{}
	_, err := c.UpdateSubAdmin(test.update)

	if err == nil {
		t.Error(err.Error())
	}

}*/

func TestAdminService_VerifyPasswordValid(t *testing.T) {
	type args struct {
		req *models.Admins
	}

	test := args{
		req: &models.Admins{
			Model: gorm.Model{
				ID: 1,
			},
			Email:    "arun.ps@accubits.com",
			Password: "P@ssword123",
		},
	}
	c := service.AdminService{}
	_, err := c.VerifyPassword(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestAdminService_VerifyPasswordInvalid(t *testing.T) {
	type args struct {
		req *models.Admins
	}

	test := args{
		//pass the parameter here
		req: &models.Admins{
			Model: gorm.Model{
				ID: 200,
			},
			Email:    "test@gmail.com",
			Password: "P@ssword123",
		},
	}
	c := service.AdminService{}
	_, err := c.VerifyPassword(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestAdminService_ChangePasswordValid(t *testing.T) {
	type args struct {
		req *models.Admin
	}

	test := args{
		//pass the parameter here
		req: &models.Admin{
			ID:          1,
			NewPassword: "P@ssword1234",
			Password:    "P@ssword1234",
		},
	}
	c := service.AdminService{}
	_, err := c.ChangePassword(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestAdminService_ChangePasswordInvalid(t *testing.T) {
	type args struct {
		req *models.Admin
	}

	test := args{
		//pass the parameter here
		req: &models.Admin{
			ID:          2,
			NewPassword: "P@ssword123",
			Password:    "P@ssword123",
		},
	}
	c := service.AdminService{}
	_, err := c.ChangePassword(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

/***************************SHARETOGMAIL - SUBADMIN DETAILS*****************************************************/
func TestAdminService_SharetoGmailValid(t *testing.T) {
	type args struct {
		req *models.Admins
	}

	test := args{
		//pass the parameter here
		req: &models.Admins{
			Model: gorm.Model{
				ID: 1,
			},
			Email:    "arun.ps@accubits.com",
			Password: "P@ssword123",
		},
	}
	c := service.AdminService{}
	_, err := c.SharetoGmail(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestAdminService_SharetoGmailInvalid(t *testing.T) {
	type args struct {
		req *models.Admins
	}

	test := args{
		req: &models.Admins{
			Model: gorm.Model{
				ID: 200,
			},
		},
	}
	c := service.AdminService{}
	_, err := c.SharetoGmail(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestAdminService_DeleteSubAdminValid(t *testing.T) {
	type args struct {
		del *models.Admins
	}

	test := args{
		//pass the parameter here
		del: &models.Admins{
			Model: gorm.Model{
				ID: 3,
			},
		},
	}
	c := service.AdminService{}
	_, err := c.DeleteSubAdmin(test.del)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestAdminService_DeleteSubAdminInvalid(t *testing.T) {
	type args struct {
		del *models.Admins
	}

	test := args{
		//pass the parameter here
		del: &models.Admins{
			Model: gorm.Model{
				ID: 200,
			},
		},
	}
	c := service.AdminService{}
	_, err := c.DeleteSubAdmin(test.del)

	if err == nil {
		t.Error(err.Error())
	}

}

/*****BLOCK SUBADMIN VALID*****/
func TestAdminService_BlockSubAdminValid(t *testing.T) {
	type args struct {
		req *models.AdminBlock
	}

	test := args{
		req: &models.AdminBlock{
			ID:      1,
			Blocked: true,
		},
	}
	c := service.AdminService{}
	_, err := c.BlockSubAdmin(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

/*****BLOCK SUBADMIN INVALID*****/
func TestAdminService_BlockSubAdminInvalid(t *testing.T) {
	type args struct {
		req *models.AdminBlock
	}

	test := args{
		req: &models.AdminBlock{
			ID:      200,
			Blocked: true,
		},
	}
	c := service.AdminService{}
	_, err := c.BlockSubAdmin(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

/****************************************ADMIN DASHBOARD TESTS*********************************/

//DASHBOARD COUNT VALID
func TestAdminService_DashboardCountValid(t *testing.T) {
	type args struct {
		fromtime string
		totime   string
	}

	test := args{
		fromtime: "2020-01-01",
		totime:   "2022-07-01",
	}
	c := service.AdminService{}
	_, err := c.DashboardCount(test.fromtime, test.totime)

	if err != nil {
		t.Error(err.Error())
	}

}

//DASHBOARD COUNT INVALID
//NOT WORKING
// func TestAdminService_DashboardCountInvalid(t *testing.T) {
// 	type args struct {
// 		fromtime string
// 		totime   string
// 	}

// 	test := args{

// 		fromtime: "2090-03-19",
// 		totime:   "2090-07-01",
// 	}
// 	c := service.AdminService{}
// 	_, err := c.DashboardCount(test.fromtime, test.totime)
// 	fmt.Println(err)
// 	if err == nil {
// 		t.Error(err.Error())
// 	}

// }

//LAST TEN USERS VALID
func TestUserService_ListLastTenUsersValid(t *testing.T) {
	c := service.UserService{}
	_, err := c.ListLastTenUsers()

	if err != nil {
		t.Error(err.Error())
	}

}

//TOP TEN USERS VALID
func TestUserService_TopTenLocationUsersValid(t *testing.T) {

	c := service.UserService{}
	_, err := c.TopTenLocationUsers()

	if err != nil {
		t.Error(err.Error())
	}

}

//KYC VERIFED/UNVERIFIED PERCENTAGE VALID
func TestUserService_KYCVerifiedUnverifiedValid(t *testing.T) {

	type args struct {
		fromtime string
		totime   string
	}

	test := args{
		fromtime: "2020-01-01",
		totime:   "2022-07-01",
	}

	c := service.UserService{}
	_, err := c.KYCVerifiedUnverified(test.fromtime, test.totime)

	if err != nil {
		t.Error(err.Error())
	}

}

//DAILY MATCHES REPORT VALID
func TestUserService_DailyMatchesReportValid(t *testing.T) {

	c := service.UserService{}
	_, err := c.DailyMatchesReport()

	if err != nil {
		t.Error(err.Error())
	}

}

//SITEVISITOR ANALYTICS VALID
func TestUserService_SiteVisitorAnalyticsValid(t *testing.T) {
	type args struct {
		fromtime string
		totime   string
	}

	test := args{
		fromtime: "2020-01-01",
		totime:   "2020-01-01",
	}
	c := service.UserService{}
	_, err := c.SiteVisitorAnalytics(test.fromtime, test.totime)

	if err != nil {
		t.Error(err.Error())
	}

}

/*****************************IOS USER PROFILE TESTS*********************************/
func TestUserService_UserProfileSetupInvalid(t *testing.T) {
	type args struct {
		req *models.UserRequest
	}

	test := args{
		//pass the parameter here
		req: &models.UserRequest{
			ID:          200,
			FirstName:   "Test",
			LastName:    "User",
			Email:       "testgmail.com",
			PhoneNumber: "+911122334455",
			DOB:         "2020-01-01",
		},
	}
	c := service.UserService{}
	_, err := c.UserProfileSetup(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_UserProfileSetupValid(t *testing.T) {

	type args struct {
		req *models.UserRequest
	}
	test := args{
		req: &models.UserRequest{
			ID:          1,
			FirstName:   "Test",
			LastName:    "User",
			Email:       "test@gmail.com",
			PhoneNumber: "+919994810592",
			DOB:         "2020-01-01",
		},
	}
	c := service.UserService{}

	_, err := c.UserProfileSetup(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

//EMAIL VALIDATION INVALID
// func TestUserService_EmailValidationInvalid(t *testing.T) {
// 	type args struct {
// 		email string
// 	}

// 	test := args{
// 		//pass the parameter here
// 		email: "arun.ps@accubits.com",
// 	}
// 	c := service.UserService{}

// 	_, err := c.EmailValidation(test.email)
// 	fmt.Println(err)
// 	if err == nil {
// 		t.Error(err.Error())
// 	}

// }

//EMAIL VALIDATION VALID
func TestUserService_EmailValidationValid(t *testing.T) {
	type args struct {
		email string
	}

	test := args{
		//pass the parameter here
		email: "test@gmail.com",
	}
	c := service.UserService{}

	_, err := c.EmailValidation(test.email)

	if err != nil {
		t.Error(err.Error())
	}

}

//USER INTERESTS INVALID
func TestUserService_UserInterestsInvalid(t *testing.T) {
	type args struct {
		req *models.UserInterest
	}

	test := args{
		//pass the parameter here
		req: &models.UserInterest{
			ID:                 100,
			Gender:             "Male",
			InterestVisibility: true,
		},
	}
	c := service.UserService{}

	_, err := c.UserInterests(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

//USER INTERESTS VALID
func TestUserService_UserInterestsValid(t *testing.T) {
	type args struct {
		req *models.UserInterest
	}

	test := args{
		//pass the parameter here
		req: &models.UserInterest{
			ID:                 1,
			Gender:             "Male",
			InterestVisibility: true,
		},
	}
	c := service.UserService{}

	_, err := c.UserInterests(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

//MEETING PURPOSE INVALID
func TestUserService_MeetingPurposeInvalid(t *testing.T) {
	type args struct {
		req *models.Purpose
	}

	test := args{
		//pass the parameter here
		req: &models.Purpose{
			ID:             100,
			MeetingPurpose: "Meeting",
		},
	}
	c := service.UserService{}

	_, err := c.MeetingPurpose(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_MeetingPurposeValid(t *testing.T) {
	type args struct {
		req *models.Purpose
	}

	test := args{
		req: &models.Purpose{
			ID:             1,
			MeetingPurpose: "Meeting",
		},
	}
	c := service.UserService{}

	_, err := c.MeetingPurpose(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_UpdateUserProfileInvalid(t *testing.T) {
	type args struct {
		req *models.UpdateUserProfile
	}

	test := args{
		//pass the parameter here
		req: &models.UpdateUserProfile{
			ID: 100,
		},
	}
	c := service.UserService{}
	_, err := c.UpdateUserProfile(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_UpdateUserProfileValid(t *testing.T) {
	type args struct {
		req *models.UpdateUserProfile
	}

	test := args{
		//pass the parameter here
		req: &models.UpdateUserProfile{
			ID: 1,
		},
	}
	c := service.UserService{}

	_, err := c.UpdateUserProfile(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_UploadProfilePicInvalid(t *testing.T) {
	type args struct {
		req           *models.UploadPhotos
		Photo1        multipart.File
		Photo2        multipart.File
		Photo3        multipart.File
		Photo4        multipart.File
		Photo5        multipart.File
		PhotoHandler1 *multipart.FileHeader
		PhotoHandler2 *multipart.FileHeader
		PhotoHandler3 *multipart.FileHeader
		PhotoHandler4 *multipart.FileHeader
		PhotoHandler5 *multipart.FileHeader
	}

	var err error
	var fileheader1 multipart.FileHeader
	var fileheader2 multipart.FileHeader

	file1, err := os.Open("../../demo.jpg")
	if err != nil {
		t.Error("error in file length", err)
	}
	ff1, err := file1.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader1.Size = ff1.Size()
	fileheader1.Filename = ff1.Name()

	file2, err := os.Open("../../demo.jpg")
	if err != nil {
		t.Error("error in file2 length", err)
	}
	ff2, err := file2.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader2.Size = ff2.Size()
	fileheader2.Filename = ff2.Name()

	test := args{
		req: &models.UploadPhotos{
			ID: 100,
		},
		Photo1:        file1,
		Photo2:        file2,
		PhotoHandler1: &fileheader1,
		PhotoHandler2: &fileheader2,
	}
	c := service.UserService{}

	_, err = c.UploadProfilePic(test.req, test.Photo1, test.Photo2, test.Photo3, test.Photo4, test.Photo5, test.PhotoHandler1, test.PhotoHandler2, test.PhotoHandler3, test.PhotoHandler4, test.PhotoHandler5)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_UploadProfilePicValid(t *testing.T) {

	type args struct {
		req           *models.UploadPhotos
		Photo1        multipart.File
		Photo2        multipart.File
		Photo3        multipart.File
		Photo4        multipart.File
		Photo5        multipart.File
		PhotoHandler1 *multipart.FileHeader
		PhotoHandler2 *multipart.FileHeader
		PhotoHandler3 *multipart.FileHeader
		PhotoHandler4 *multipart.FileHeader
		PhotoHandler5 *multipart.FileHeader
	}

	var err error
	var fileheader1 multipart.FileHeader
	var fileheader2 multipart.FileHeader

	file1, err := os.Open("../../download.jpeg")
	if err != nil {
		t.Error("error in file length", err)
	}
	ff1, err := file1.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader1.Size = ff1.Size()
	fileheader1.Filename = ff1.Name()

	file2, err := os.Open("../../download.jpeg")
	if err != nil {
		t.Error("error in file2 length", err)
	}
	ff2, err := file2.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader2.Size = ff2.Size()
	fileheader2.Filename = ff2.Name()

	test := args{
		//pass the parameter here
		req: &models.UploadPhotos{
			ID: 1,
		},

		Photo1:        file1,
		Photo2:        file2,
		PhotoHandler1: &fileheader1,
		PhotoHandler2: &fileheader2,
	}
	c := service.UserService{}
	_, err = c.UploadProfilePic(test.req, test.Photo1, test.Photo2, test.Photo3, test.Photo4, test.Photo5, test.PhotoHandler1, test.PhotoHandler2, test.PhotoHandler3, test.PhotoHandler4, test.PhotoHandler5)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_ListAllCountriesValid(t *testing.T) {
	c := service.UserService{}

	_, err := c.ListAllCountries()

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_CitiesInvalid(t *testing.T) {
	type args struct {
		request *models.CountryRequest
	}

	test := args{
		request: &models.CountryRequest{
			Country: "UK",
		},
	}
	c := service.UserService{}

	_, err := c.Cities(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_CitiesValid(t *testing.T) {
	type args struct {
		request *models.CountryRequest
	}

	test := args{
		//pass the parameter here
		request: &models.CountryRequest{
			Country: "India",
		},
	}
	c := service.UserService{}

	_, err := c.Cities(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_GetAllCountriesValid(t *testing.T) {

	_, err := service.GetAllCountries()

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_GetUserProfileInvalid(t *testing.T) {
	type args struct {
		request *models.GetUserDetails
	}

	test := args{
		//pass the parameter here
		request: &models.GetUserDetails{
			ID: 100,
		},
	}
	c := service.UserService{}

	_, err := c.GetUserProfile(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_GetUserProfileValid(t *testing.T) {
	type args struct {
		request *models.GetUserDetails
	}

	test := args{
		//pass the parameter here
		request: &models.GetUserDetails{
			ID: 1,
		},
	}
	c := service.UserService{}

	_, err := c.GetUserProfile(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_UpdatePreferencesInvalid(t *testing.T) {
	type args struct {
		request *models.Purpose
	}

	test := args{
		//pass the parameter here
		request: &models.Purpose{
			ID: 100,
		},
	}
	c := service.UserService{}

	_, err := c.UpdatePreferences(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_UpdatePreferencesValid(t *testing.T) {
	type args struct {
		request *models.Purpose
	}

	test := args{
		//pass the parameter here
		request: &models.Purpose{
			ID: 1,
		},
	}
	c := service.UserService{}

	_, err := c.UpdatePreferences(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_UpdateUserDetailsInvalid(t *testing.T) {
	type args struct {
		request *models.Userdetails
	}

	test := args{
		//pass the parameter here
		request: &models.Userdetails{
			ID: 100,
		},
	}
	c := service.UserService{}

	_, err := c.UpdateUserDetails(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_UpdateUserDetailsValid(t *testing.T) {
	type args struct {
		request *models.Userdetails
	}

	test := args{
		//pass the parameter here
		request: &models.Userdetails{
			ID: 1,
		},
	}
	c := service.UserService{}

	_, err := c.UpdateUserDetails(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_DeactivateAccountInvalid(t *testing.T) {
	type args struct {
		request *models.DeactivateAccount
	}

	test := args{
		//pass the parameter here
		request: &models.DeactivateAccount{
			ID: 100,
		},
	}
	c := service.UserService{}

	_, err := c.DeactivateAccount(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_DeactivateAccountValid(t *testing.T) {
	type args struct {
		request *models.DeactivateAccount
	}

	test := args{
		//pass the parameter here
		request: &models.DeactivateAccount{
			ID: 1,
		},
	}
	c := service.UserService{}

	_, err := c.DeactivateAccount(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_PhoneNumberGenerateOTPInvalid(t *testing.T) {
	type args struct {
		request *models.UserPhoneno
	}

	test := args{
		//pass the parameter here
		request: &models.UserPhoneno{
			ID:      100,
			PhoneNo: "+919999999999",
		},
	}
	c := service.UserService{}

	_, err := c.PhoneNumberGenerateOTP(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_PhoneNumberGenerateOTPValid(t *testing.T) {
	type args struct {
		request *models.UserPhoneno
	}

	test := args{
		//pass the parameter here
		request: &models.UserPhoneno{
			ID:      1,
			PhoneNo: "+919999999999",
		},
	}
	c := service.UserService{}

	_, err := c.PhoneNumberGenerateOTP(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUserService_EmailIdGenerateOTPValid(t *testing.T) {
	type args struct {
		request *models.UserPhoneno
	}

	test := args{
		//pass the parameter here
		request: &models.UserPhoneno{
			ID: 1,
		},
	}
	c := service.UserService{}

	_, err := c.EmailIdGenerateOtp(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}

//VerifyEmailOtp
func TestUserService_VerifyEmailOtpInvalid(t *testing.T) {
	type args struct {
		request *models.RequestLogin
	}

	test := args{
		//pass the parameter here
		request: &models.RequestLogin{
			ID:  1,
			OTP: "3444",
		},
	}
	c := service.UserService{}

	_, err := c.VerifyEmailOtp(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestUserService_UpdatePhoneNumberByIDInvalid(t *testing.T) {
	type args struct {
		request *models.UserPhoneno
	}

	test := args{
		//pass the parameter here
		request: &models.UserPhoneno{
			ID:      100,
			PhoneNo: "+919999999999",
		},
	}
	c := service.UserService{}

	_, err := c.UpdatePhoneNumberByID(test.request)

	if err == nil {
		t.Error(err.Error())
	}

}
func TestUserService_UpdatePhoneNumberByIDValid(t *testing.T) {
	type args struct {
		request *models.UserPhoneno
	}

	test := args{
		//pass the parameter here
		request: &models.UserPhoneno{
			ID:      1,
			PhoneNo: "+919999999999",
		},
	}
	c := service.UserService{}

	_, err := c.UpdatePhoneNumberByID(test.request)

	if err != nil {
		t.Error(err.Error())
	}

}
