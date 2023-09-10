package test

import (
	"fmt"
	"sondr-backend/src/models"
	"sondr-backend/src/service"
	"time"

	"testing"

	"gorm.io/gorm"
)

func TestEventService_CreateEventValid(t *testing.T) {
	type args struct {
		req *models.EventRequest
	}
	t1 := time.Now().Add(time.Minute).UTC()
	//startTime, _ := time.Parse("2006-01-02 15:04:05", t1.String())
	endTime := time.Now().Add(time.Hour).UTC()

	test := args{
		//pass the parameter here
		req: &models.EventRequest{
			HostUserID:    1,
			EventName:     "Testing Purpose",
			Location:      "Accubits",
			EventMode:     "Open",
			Date:          t1.Format("2006-01-02"),
			Coordinates:   "13.058370,80.273682",
			StartTime:     t1.Format("2006-01-02 15:04:05"),
			EndTime:       endTime.Format("2006-01-02 15:04:05"),
			InvitedUserId: []uint{3, 2},
		},
	}
	es := &service.EventService{}

	_, err := es.CreateEvent(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEventService_CreateEventInvalid(t *testing.T) {
	type args struct {
		req *models.EventRequest
	}
	t1 := time.Now()
	//startTime, _ := time.Parse("2006-01-02 15:04:05", t1.String())
	endTime := time.Now().Add(time.Hour)

	test := args{
		//pass the parameter here
		req: &models.EventRequest{
			HostUserID:    0,
			EventName:     "Testing Purpose",
			Location:      "Accubits",
			EventMode:     "Open",
			Date:          t1.Format("2006-01-02"),
			Coordinates:   "13.058370,80.273682",
			StartTime:     t1.Format("2006-01-02 15:04:05"),
			EndTime:       endTime.Format("2006-01-02 15:04:05"),
			InvitedUserId: []uint{3, 2},
		},
	}
	es := &service.EventService{}

	_, err := es.CreateEvent(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_FetchEventByIdValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 1,
	}
	es := &service.EventService{}

	_, err := es.FetchEventById(test.id)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEventService_FetchEventByIdInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 0,
	}
	es := &service.EventService{}

	_, err := es.FetchEventById(test.id)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_UpdateEventValid(t *testing.T) {
	type args struct {
		req *models.EventRequest
	}
	t1 := time.Now()
	//startTime, _ := time.Parse("2006-01-02 15:04:05", t1.String())
	endTime := time.Now().Add(time.Hour)

	test := args{
		//pass the parameter here]
		req: &models.EventRequest{
			HostUserID:    1,
			EventName:     "Testing Purpose",
			Location:      "Accubits",
			EventMode:     "Open",
			Date:          t1.Format("2006-01-02"),
			Coordinates:   "13.058370,80.273682",
			StartTime:     t1.Format("2006-01-02 15:04:05"),
			EndTime:       endTime.Format("2006-01-02 15:04:05"),
			InvitedUserId: []uint{3, 2},
		},
	}
	es := &service.EventService{}

	_, err := es.UpdateEvent(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestEventService_UpdateEventInvalid(t *testing.T) {
	type args struct {
		req *models.EventRequest
	}
	t1 := time.Now()
	//startTime, _ := time.Parse("2006-01-02 15:04:05", t1.String())
	endTime := time.Now().Add(time.Hour)

	test := args{
		//pass the parameter here
		req: &models.EventRequest{
			HostUserID:    0,
			EventName:     "Testing Purpose",
			Location:      "Accubits",
			EventMode:     "Open",
			Date:          t1.Format("2006-01-02"),
			Coordinates:   "13.058370,80.273682",
			StartTime:     t1.Format("2006-01-02 15:04:05"),
			EndTime:       endTime.Format("2006-01-02 15:04:05"),
			InvitedUserId: []uint{3, 2},
		},
	}
	es := &service.EventService{}

	_, err := es.UpdateEvent(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_InvitedEventsValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 0,
	}
	es := &service.EventService{}

	_, err := es.InvitedEvents(test.id)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEventService_InvitedEventsInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 1,
	}
	es := &service.EventService{}

	_, err := es.InvitedEvents(test.id)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_HostedEventsValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 0,
	}
	es := &service.EventService{}

	_, err := es.HostedEvents(test.id)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEventService_HostedEventsInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 1,
	}
	es := &service.EventService{}

	_, err := es.HostedEvents(test.id)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_EventCheckInValid(t *testing.T) {
	type args struct {
		req *models.EventCheckInRequest
	}

	test := args{
		//pass the parameter here
		req: &models.EventCheckInRequest{
			UserID:    1,
			EventId:   1,
			Latitude:  "13.058370",
			Longitude: "80.273682",
		},
	}
	es := &service.EventService{}

	_, err := es.EventCheckIn(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}
func TestEventService_EventCheckInInvalid(t *testing.T) {
	type args struct {
		req *models.EventCheckInRequest
	}

	test := args{
		//pass the parameter here
		req: &models.EventCheckInRequest{
			UserID:    1,
			EventId:   1,
			Latitude:  "13.058370",
			Longitude: "34.273682",
		},
	}
	es := &service.EventService{}

	_, err := es.EventCheckIn(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_ListProfilesEventCheckInValid(t *testing.T) {
	type args struct {
		userId  uint
		eventId uint
	}

	test := args{
		//pass the parameter here
		userId:  1,
		eventId: 1,
	}
	es := &service.EventService{}

	_, err := es.ListProfilesEventCheckIn(test.userId, test.eventId)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEventService_ListProfilesEventCheckInInvalid(t *testing.T) {
	type args struct {
		userId  uint
		eventId uint
	}

	test := args{
		//pass the parameter here
		userId:  0,
		eventId: 1,
	}
	es := &service.EventService{}

	_, err := es.ListProfilesEventCheckIn(test.userId, test.eventId)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestEventService_EventCheckOutValid(t *testing.T) {
	type args struct {
		req *models.EventCheckInRequest
	}

	test := args{
		//pass the parameter here
		req: &models.EventCheckInRequest{
			UserID:  1,
			EventId: 1,
		},
	}
	es := &service.EventService{}

	_, err := es.EventCheckOut(test.req)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestEventService_EventCheckOutInvalid(t *testing.T) {
	type args struct {
		req *models.EventCheckInRequest
	}

	test := args{
		//pass the parameter here
		req: &models.EventCheckInRequest{
			UserID:  1,
			EventId: 0,
		},
	}
	es := &service.EventService{}

	_, err := es.EventCheckOut(test.req)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestListAllEventServiceValid(t *testing.T) {
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
	es := &service.EventService{}
	_, _, err := es.ListAllEventService(test.pageNo, test.pageSize, test.searchfilter, test.from, test.to)

	if err != nil {
		t.Error("error msg in list event valid", err.Error())
	}

}

func TestListAllEventServiceInvalid(t *testing.T) {
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
	es := &service.EventService{}
	_, resp, _ := es.ListAllEventService(test.pageNo, test.pageSize, test.searchfilter, test.from, test.to)
	fmt.Println(resp)
	if resp.Count != 0 {
		t.Error("data not found")
	}
}

func TestGetEventServiceValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 1,
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.GetEventService(test.id)

	if err != nil {
		t.Error("error msg in geteventbyID valid", err.Error())
	}

}

func TestGetEventServiceInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 0,
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.GetEventService(test.id)

	if err == nil {
		t.Error("error msg in geteventbyID Invalid", err.Error())
	}

}

func TestInvitedUserServiceValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 1,
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.InvitedUserService(test.id)

	if err != nil {
		t.Error("error msg in InvitedUser valid", err.Error())
	}

}

func TestInvitedUserServiceInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 0,
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, resp, _ := es.InvitedUserService(test.id)

	if resp.InvitedGuestCount != 0 {
		fmt.Println("error called")
		t.Error("data missing")
	}

}

func TestGetAttendieEventServiceValid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 5,
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.GetAttendieEventService(test.id)

	if err != nil {
		t.Error("error msg in attendedusers valid", err.Error())
	}

}

func TestGetAttendieEventServiceInvalid(t *testing.T) {
	type args struct {
		id uint
	}

	test := args{
		//pass the parameter here
		id: 0,
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.GetAttendieEventService(test.id)

	if err == nil {
		t.Error("error msg in attendedusers invalid", err.Error())
	}

}

func TestCancelEventServiceValid(t *testing.T) {
	type args struct {
		req *models.Events
	}

	test := args{
		//pass the parameter here
		req: &models.Events{
			Model: gorm.Model{
				ID: 1,
			},
			Reason: "ilegal event",
		},
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.CancelEventService(test.req)

	if err != nil {
		t.Error("error msg in cancelEvent valid", err.Error())
	}

}
func TestCancelEventServiceInvalid(t *testing.T) {
	type args struct {
		req *models.Events
	}

	test := args{
		//pass the parameter here
		req: &models.Events{
			Model: gorm.Model{
				ID: 1,
			},
			Reason: " ",
		},
	}
	// Add External package name in ____
	es := &service.EventService{}
	_, _, err := es.CancelEventService(test.req)

	if err == nil {
		t.Error("error msg in CancelEvent Invalid", err.Error())
	}

}
