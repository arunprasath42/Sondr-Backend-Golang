package test

import (
	"fmt"
	"sondr-backend/src/models"
	"sondr-backend/src/service"
	"testing"
)

func TestFindAllUserServiceValid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchfilter string
		from         string
		to           string
	}

	test := args{
		//pass the parameter here
		pageNo:   0,
		pageSize: 10,
	}
	// Add External package name in ____
	es := &service.UserService{}
	_, err := es.FindAllUsers(test.searchfilter, test.pageNo, test.pageSize, test.from, test.to)
	if err != nil {
		t.Error("error msg in find all user valid", err.Error())
	}

}

func TestFindAllUserServiceInvalid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchfilter string
		from         string
		to           string
	}

	test := args{
		//pass the parameter here
		pageNo:       0,
		pageSize:     10,
		searchfilter: "0",
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.FindAllUsers(test.searchfilter, test.pageNo, test.pageSize, test.from, test.to)
	if resp.Count == 0 {
		t.Error("error msg in find all users invalid")
	}
}

func TestFindAllReportedUserServiceValid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchfilter string
		from         string
		to           string
	}

	test := args{
		//pass the parameter here
		pageNo:   0,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.FindAllReportedUsers(test.searchfilter, test.pageNo, test.pageSize, test.from, test.to)

	if err != nil {
		t.Error("error msg in find all reported users valid", err.Error())
	}

}

func TestFindAllReportedUserServiceInValid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchfilter string
		from         string
		to           string
	}

	test := args{
		//pass the parameter here
		pageNo:       0,
		pageSize:     0,
		searchfilter: "z",
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.FindAllReportedUsers(test.searchfilter, test.pageNo, test.pageSize, test.from, test.to)

	if resp.Count != 0 {
		t.Error("error msg in find all reported users invalid")
	}

}

func TestFindAllBlockedUserServiceValid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchfilter string
		from         string
		to           string
	}

	test := args{
		//pass the parameter here
		pageNo:   0,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.FindAllReportedUsers(test.searchfilter, test.pageNo, test.pageSize, test.from, test.to)

	if err != nil {
		t.Error("error msg in find all blocked users valid", err.Error())
	}

}

func TestFindAllBlockedUserServiceInValid(t *testing.T) {
	type args struct {
		pageNo       int
		pageSize     int
		searchfilter string
		from         string
		to           string
	}

	test := args{
		//pass the parameter here
		pageNo:       1,
		pageSize:     10,
		searchfilter: "0",
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.FindAllReportedUsers(test.searchfilter, test.pageNo, test.pageSize, test.from, test.to)

	if resp.Count != 0 {
		t.Error("error msg in find all blocked users invalid")
	}

}

func TestFindAllHostedEventServiceValid(t *testing.T) {
	type args struct {
		userId   uint
		pageNo   int
		pageSize int
		from     string
		to       string
	}

	test := args{
		//pass the parameter here
		userId:   1,
		pageNo:   1,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.FindAllHostedEvents(test.userId, test.pageNo, test.pageSize, test.from, test.to)

	if err != nil {
		t.Error("error msg in find all hosted events of user valid", err.Error())
	}

}

func TestFindAllHostedEventServiceInvalid(t *testing.T) {
	type args struct {
		userId   uint
		pageNo   int
		pageSize int
		from     string
		to       string
	}

	test := args{
		//pass the parameter here
		userId: 0,
	} // Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.FindAllHostedEvents(test.userId, test.pageNo, test.pageSize, test.from, test.to)
	fmt.Println(resp.Count)
	if resp.Count != 0 {
		t.Error("error msg in find all hosted events of user invalid")
	}

}

func TestFindAllReportServiceValid(t *testing.T) {
	type args struct {
		userId   uint
		pageNo   int
		pageSize int
		from     string
		to       string
	}

	test := args{
		//pass the parameter here
		userId:   1,
		pageNo:   1,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.FindAllReports(test.userId, test.pageNo, test.pageSize, test.from, test.to)
	if err != nil {
		t.Error("error msg in find all reports of user valid", err.Error())
	}

}

func TestFindAllReportsServiceInvalid(t *testing.T) {
	type args struct {
		userId   uint
		pageNo   int
		pageSize int
		from     string
		to       string
	}

	test := args{
		//pass the parameter here
		userId:   0,
		pageNo:   1,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.FindAllReports(test.userId, test.pageNo, test.pageSize, test.from, test.to)

	if resp.Count != 0 {
		t.Error("error msg in find all reports of user invalid")
	}

}

func TestFindUserInfoServiceValid(t *testing.T) {
	type args struct {
		userId uint
	}

	test := args{
		//pass the parameter here
		userId: 1,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.AboutUser(test.userId)
	if err != nil {
		t.Error("error msg in find user info valid", err.Error())
	}

}
func TestFindUserInfoServiceInvalid(t *testing.T) {
	type args struct {
		userId uint
	}

	test := args{
		//pass the parameter here
		userId: 0,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.AboutUser(test.userId)
	if err == nil {
		t.Error("error msg in find user info invalid")
	}

}

func TestGetUploadedPhotosOfUserValid(t *testing.T) {
	type args struct {
		userId uint
	}

	test := args{
		//pass the parameter here
		userId: 1,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetUploadedPhotos(test.userId)
	if err != nil {
		t.Error("error msg in get uploaded photos of user valid", err.Error())
	}

}
func TestGetUploadedphotoServiceInvalid(t *testing.T) {
	type args struct {
		userId uint
	}

	test := args{
		//pass the parameter here
		userId: 0,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetUploadedPhotos(test.userId)
	if err == nil {
		t.Error("error msg in get uploaded photos of user invalid")
	}

}

func TestGetUserMetadataServiceValid(t *testing.T) {
	type args struct {
		userId uint
	}

	test := args{
		//pass the parameter here
		userId: 1,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetUsersMetadata(test.userId)
	if err != nil {
		t.Error("error msg in get user metadata valid", err.Error())
	}

}

func TestGetUserMetadataServiceInValid(t *testing.T) {
	type args struct {
		userId uint
	}
	test := args{
		//pass the parameter here
		userId: 0,
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.GetUsersMetadata(test.userId)
	if resp.TotalCheckIns != 0 {
		t.Error("error msg in get user metadata invalid")
	}

}
func TestBlockAndUnblockUserServiceValid(t *testing.T) {
	args := models.Request{
		UserId: 1,
		ID:     1,
	}

	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.BlockAndUnblockUser(args.ID, args.UserId)
	if err != nil {
		t.Error("error msg in block and unblock user valid")
	}

}

func TestBlockAndUnblockUserServiceInvalid(t *testing.T) {
	args := models.Request{
		UserId: 0,
		ID:     1,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.BlockAndUnblockUser(args.ID, args.UserId)
	if err == nil {
		t.Error("error msg in block and unblock user invalid")
	}

}

func TestGetPromptQuestionsServiceValid(t *testing.T) {
	// Add External package name in ____
	us := &service.UserService{}
	resp, err := us.GetPromptQuestion()
	if resp.Count == 0 {
		t.Error("error msg in get prompt questions valid", err.Error())
	}

}

func TestGetPromptQuestionsServiceInValid(t *testing.T) {
	// Add External package name in ____
	us := &service.UserService{}
	resp, err := us.GetPromptQuestion()
	fmt.Println(resp, err)
	if resp.Count != 0 {
		t.Error("error msg in get prompt questions invalid")
	}

}

func TestAddPromptAnswerServiceValid(t *testing.T) {
	type args struct {
		userId uint
		test   []*models.AnswerRequest
	}
	test :=
		args{
			userId: 1,
			test: []*models.AnswerRequest{
				//pass the parameter here
				{
					QuestionId: 1,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 2,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 3,
					Answer:     "Hi i am fine.",
				}},
		}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.AddAnswerOfPromptQuestion(test.userId, test.test)
	if err != nil {
		t.Error("error msg in add prompt answer valid", err.Error())
	}

}

func TestAddPromptAnswerServiceInValid(t *testing.T) {

	type args struct {
		userId uint
		test   []*models.AnswerRequest
	}
	test :=
		args{
			userId: 1,
			test: []*models.AnswerRequest{
				//pass the parameter here
				{
					QuestionId: 1,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 2,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 3,
					Answer:     "Hi i am fine.",
				}},
		}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.AddAnswerOfPromptQuestion(test.userId, test.test)
	if err == nil {
		t.Error("error msg in add prompt answer invalid")
	}

}

func TestGetAllPromptsOfUserServiceValid(t *testing.T) {
	type args struct {
		userId uint
	}

	test := args{
		//pass the parameter here
		userId: 1,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetAllPromptOfUser(test.userId)
	if err != nil {
		t.Error("error msg in get all prompts of user valid", err.Error())
	}

}

func TestGetAllPromptsOfUserServiceInValid(t *testing.T) {
	type args struct {
		userId uint
	}
	test := args{
		//pass the parameter here
		userId: 0,
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.GetAllPromptOfUser(test.userId)
	if resp.Count != 0 {
		t.Error("error msg in get all prompts of user invalid")
	}

}

func TestEditPromptAnswerServiceValid(t *testing.T) {
	type args struct {
		userId int64
		test   []*models.AnswerRequest
	}
	test :=
		args{
			userId: 1,
			test: []*models.AnswerRequest{
				//pass the parameter here
				{
					QuestionId: 1,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 2,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 3,
					Answer:     "Hi i am fine.",
				}},
		}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.EditPromptAnswers(test.userId, test.test)
	if err != nil {
		t.Error("error msg in edit prompt answer valid", err.Error())
	}

}

func TestEditPromptAnswerServiceInValid(t *testing.T) {
	type args struct {
		userId int64
		test   []*models.AnswerRequest
	}
	test :=
		args{
			userId: 1,
			test: []*models.AnswerRequest{
				//pass the parameter here
				{
					QuestionId: 1,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 2,
					Answer:     "Hi i am fine.",
				},
				{
					QuestionId: 3,
					Answer:     "Hi i am fine.",
				}},
		}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.EditPromptAnswers(test.userId, test.test)
	if err == nil {
		t.Error("error msg in edit prompt answer invalid")
	}

}

func TestMatchedprofilesServiceValid(t *testing.T) {
	type args struct {
		userId       uint
		pageNo       int
		pageSize     int
		searchFilter string
	}
	test := args{
		//pass the parameter here
		userId:   1,
		pageNo:   1,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetMatchedProfiles(test.userId, test.pageNo, test.pageSize, test.searchFilter)
	if err != nil {
		t.Error("error msg in edit prompt answer valid", err.Error())
	}

}

func TestMatchedprofilesServiceInValid(t *testing.T) {
	type args struct {
		userId       uint
		pageNo       int
		pageSize     int
		searchFilter string
	}
	test := args{
		//pass the parameter here
		userId:   0,
		pageNo:   0,
		pageSize: 10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	resp, _ := us.GetMatchedProfiles(test.userId, test.pageNo, test.pageSize, test.searchFilter)
	if resp.Count != 0 {
		t.Error("error msg in edit prompt answer invalid")
	}

}

func TestGetUserProfilesByIdServiceValid(t *testing.T) {
	type args struct {
		id     int64
		userId int64
	}
	test := args{
		//pass the parameter here
		id:     1,
		userId: 1,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetUserProfileById(test.id, test.userId)
	if err != nil {
		t.Error("error msg in edit prompt answer valid", err.Error())
	}

}

func TestGetuserProfilesbyIdServiceInValid(t *testing.T) {
	type args struct {
		id     int64
		userId int64
	}
	test := args{
		//pass the parameter here
		id:     1,
		userId: 0,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.GetUserProfileById(test.id, test.userId)
	if err == nil {
		t.Error("error msg in edit prompt answer invalid")
	}

}

func TestLocationCheckInServiceValid(t *testing.T) {
	type args struct {
		userId               uint
		currentLatitude      string
		currentLongitude     string
		destinationLatitude  string
		destinationLongitude string
		locationName         string
	}
	test := args{
		//pass the parameter here
		userId:               1,
		currentLatitude:      "30.9697",
		currentLongitude:     "-92.80322",
		destinationLatitude:  "30.9697",
		destinationLongitude: "-92.80322",
		locationName:         "Trivandrum",
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.LocationCheckIn(test.userId, test.currentLatitude, test.currentLongitude, test.destinationLatitude, test.destinationLongitude, test.locationName)
	if err != nil {
		t.Error("error msg in edit prompt answer valid", err.Error())
	}

}

func TestLocationCheckInServiceInValid(t *testing.T) {
	type args struct {
		userId               uint
		currentLatitude      string
		currentLongitude     string
		destinationLatitude  string
		destinationLongitude string
		locationName         string
	}
	test := args{
		//pass the parameter here
		userId:               0,
		currentLatitude:      "30.9697",
		currentLongitude:     "-92.80322",
		destinationLatitude:  "90.9697",
		destinationLongitude: "-82.80322",
		locationName:         "Trivandrum",
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.LocationCheckIn(test.userId, test.currentLatitude, test.currentLongitude, test.destinationLatitude, test.destinationLongitude, test.locationName)
	if err == nil {
		t.Error("error msg in edit prompt answer invalid")
	}

}

func TestLocationCheckInProfilesServiceValid(t *testing.T) {
	type args struct {
		userId    uint
		latitude  string
		longitude string
		pageNo    int
		pageSize  int
	}
	test := args{
		//pass the parameter here
		userId:    1,
		latitude:  "30.9697",
		longitude: "-92.80322",
		pageNo:    0,
		pageSize:  0,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.ListOfProfilesLocationCheckIn(test.userId, test.latitude, test.longitude, test.pageNo, test.pageSize)
	if err != nil {
		t.Error("error msg in edit prompt answer valid", err.Error())
	}

}

func TestLocationCheckInProfilesServiceInValid(t *testing.T) {
	type args struct {
		userId    uint
		latitude  string
		longitude string
		pageNo    int
		pageSize  int
	}
	test := args{
		//pass the parameter here
		userId:    0,
		latitude:  "",
		longitude: "",
		pageNo:    0,
		pageSize:  10,
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.ListOfProfilesLocationCheckIn(test.userId, test.latitude, test.longitude, test.pageNo, test.pageSize)
	if err == nil {
		t.Error("error msg in edit prompt answer invalid")
	}

}

func TestListOfEventsOnNearByLocationsServiceValid(t *testing.T) {
	type args struct {
		userId    int64
		latitude  string
		longitude string
		pageNo    int
		pageSize  int
	}
	test := args{
		//pass the parameter here
		userId:    1,
		latitude:  "30.9697",
		longitude: "-92.80322",
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.ListOfEventsAndUserLocations(test.userId, test.latitude, test.longitude)
	if err != nil {
		t.Error("error msg in edit prompt answer valid", err.Error())
	}

}

func TestListOfEventsOnNearByLocationsServiceInValid(t *testing.T) {
	type args struct {
		userId    int64
		latitude  string
		longitude string
	}
	test := args{
		//pass the parameter here
		userId:    0,
		latitude:  "",
		longitude: "",
	}
	// Add External package name in ____
	us := &service.UserService{}
	_, err := us.ListOfEventsAndUserLocations(test.userId, test.latitude, test.longitude)
	if err == nil {
		t.Error("error msg in edit prompt answer invalid")
	}

}

func TestLocationCheckoutValid(t *testing.T) {
	type args struct {
		id int64
	}
	test := args{
		id: 1,
	}
	us := service.UserService{}
	_, err := us.LocationCheckOut(test.id)
	if err != nil {
		t.Error("error msg in location checkout valid.", err.Error())
	}
}

func TestLocationCheckoutInValid(t *testing.T) {
	type args struct {
		id int64
	}
	test := args{
		id: 0,
	}
	us := service.UserService{}
	_, err := us.LocationCheckOut(test.id)
	if err == nil {
		t.Error("error msg in location checkout invalid.", err.Error())
	}
}

func TestLikeUserProfileValid(t *testing.T) {
	type args struct {
		senderUserId   uint
		receiverUserId uint
	}
	test := args{
		senderUserId:   1,
		receiverUserId: 2,
	}
	us := service.UserService{}
	_, err := us.LikeAndUnlikeUserProfile(test.senderUserId, test.receiverUserId)
	if err != nil {
		t.Error("error msg in like user profile valid", err.Error())
	}
}

func TestLikeUserProfileInValid(t *testing.T) {
	type args struct {
		senderUserId   uint
		receiverUserId uint
	}
	test := args{
		senderUserId:   0,
		receiverUserId: 0,
	}
	us := service.UserService{}
	_, err := us.LikeAndUnlikeUserProfile(test.senderUserId, test.receiverUserId)
	if err == nil {
		t.Error("error msg in like user profile invalid", err.Error())
	}
}

func TestGetAllNotificationValid(t *testing.T) {
	type args struct {
		userId   int64
		pageNo   int64
		pageSize int64
	}
	test := args{
		userId:   1,
		pageNo:   1,
		pageSize: 10,
	}
	us := service.UserService{}
	_, err := us.GetAllNotifications(test.userId, test.pageNo, test.pageSize)
	if err != nil {
		t.Error("error msg in get all notifications of user valid", err.Error())
	}
}

func TestGetAllNotificationInvalid(t *testing.T) {
	type args struct {
		userId   int64
		pageNo   int64
		pageSize int64
	}
	test := args{
		userId:   11,
		pageNo:   1,
		pageSize: 10,
	}
	us := service.UserService{}
	resp, _ := us.GetAllNotifications(test.userId, test.pageNo, test.pageSize)
	if resp.Count != 0 {
		t.Error("error msg in get all notifications of user invalid")
	}
}

func TestReadNotificationValid(t *testing.T) {
	type args struct {
		id     uint
		userId uint
	}
	test := args{
		id:     9,
		userId: 0,
	}
	us := service.UserService{}
	_, err := us.ReadNotification(test.id, test.userId)
	if err != nil {
		t.Error("error msg in read all notifications of user valid", err.Error())
	}
}

func TestReadNotificationInValid(t *testing.T) {
	type args struct {
		id     uint
		userId uint
	}
	test := args{
		id:     1,
		userId: 0,
	}
	us := service.UserService{}
	_, err := us.ReadNotification(test.id, test.userId)
	if err == nil {
		t.Error("error msg in read all notifications of user invalid", err.Error())
	}
}

func TestReportProfileValid(t *testing.T) {
	type args struct {
		reporterUserId uint
		reporteeUserId uint
		reason         string
		comment        string
	}
	test := args{
		reporterUserId: 1,
		reporteeUserId: 1,
		reason:         "Abusive",
		comment:        "Bad behaviour",
	}
	us := service.UserService{}
	_, err := us.ReportProfile(test.reporterUserId, test.reporteeUserId, test.reason, test.comment)
	if err != nil {
		t.Error("error msg in read all notifications of user valid", err.Error())
	}
}

func TestReportProfileInValid(t *testing.T) {
	type args struct {
		reporterUserId uint
		reporteeUserId uint
		reason         string
		comment        string
	}
	test := args{
		reporterUserId: 0,
		reporteeUserId: 0,
		reason:         "",
		comment:        "",
	}
	us := service.UserService{}
	_, err := us.ReportProfile(test.reporterUserId, test.reporteeUserId, test.reason, test.comment)
	if err == nil {
		t.Error("error msg in read all notifications of user invalid", err.Error())
	}
}
func TestInsertUserInterestsValid(t *testing.T) {
	type args struct {
		userId    uint
		interests []*models.InterestRequest
	}
	test := args{
		userId: 1,
		interests: []*models.InterestRequest{
			{GenreType: "Creativity", Interests: []string{"Art", "Design"}},
			{GenreType: "Music", Interests: []string{"Rap", "Romantic"}},
		},
	}
	us := service.UserService{}
	_, err := us.InsertUserInterests(test.userId, test.interests)
	if err != nil {
		t.Error("error msg in inserting interests of user valid", err.Error())
	}
}

func TestInsertUserInterestsInValid(t *testing.T) {
	type args struct {
		userId    uint
		interests []*models.InterestRequest
	}
	test := args{
		userId: 0,
		interests: []*models.InterestRequest{
			{GenreType: "", Interests: []string{"", ""}},
			{GenreType: "", Interests: []string{"", ""}},
		},
	}

	us := service.UserService{}
	_, err := us.InsertUserInterests(test.userId, test.interests)
	if err == nil {
		t.Error("error msg in inserting interests of user invalid", err.Error())
	}
}

func TestLogoutValid(t *testing.T) {
	type args struct {
		userId uint
	}
	test := args{
		userId: 1,
	}
	us := service.UserService{}
	_, err := us.Logout(test.userId)
	if err != nil {
		t.Error("error msg in logout user valid", err.Error())
	}
}

func TestLogoutInValid(t *testing.T) {
	type args struct {
		userId uint
	}
	test := args{
		userId: 0,
	}
	us := service.UserService{}
	_, err := us.Logout(test.userId)
	if err == nil {
		t.Error("error msg in logout user invalid")
	}
}

func TestRejectProfileValid(t *testing.T) {
	type args struct {
		id     uint
		userId uint
	}
	test := args{
		id:     12,
		userId: 1,
	}
	us := service.UserService{}
	_, err := us.RejectProfile(test.id, test.userId)
	if err != nil {
		t.Error("error msg in rejecting user profile valid", err.Error())
	}
}

func TestRejectProfileInValid(t *testing.T) {
	type args struct {
		id     uint
		userId uint
	}
	test := args{
		id:     1,
		userId: 0,
	}

	us := service.UserService{}
	_, err := us.RejectProfile(test.id, test.userId)
	if err == nil {
		t.Error("error msg in rejecting user profile invalid", err.Error())
	}
}

func TestUpdateLastVisitedStatusValid(t *testing.T) {
	type args struct {
		userId uint
	}
	test := args{
		userId: 1,
	}
	us := service.UserService{}
	_, err := us.UpdateLastVisitedOfUser(test.userId)
	if err != nil {
		t.Error("error msg in updating last visited of user valid", err.Error())
	}
}

func TestUpdateLastVisitedStatusInValid(t *testing.T) {
	type args struct {
		userId uint
	}
	test := args{
		userId: 0,
	}

	us := service.UserService{}
	_, err := us.UpdateLastVisitedOfUser(test.userId)
	if err == nil {
		t.Error("error msg in updating last visited of user invalid", err.Error())
	}
}
