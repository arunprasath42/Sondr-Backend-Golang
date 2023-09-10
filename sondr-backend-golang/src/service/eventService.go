package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sondr-backend/src/models"
	"sondr-backend/src/repository"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/distance"
	"sondr-backend/utils/logging"
	mail "sondr-backend/utils/notification"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var Users sync.Map

var StartTimmings []models.StartTimeQueue
var EndTimmings []models.EndTimeQueue

type StartTimeSlice []models.StartTimeQueue
type EndTimeSlice []models.EndTimeQueue

var ChannelStart chan bool
var ChannelEnd chan bool

type EventService struct{}

func (es *EventService) ListAllEventService(pageNo, pageSize int, searchfilter string, from string, to string) (int, *models.EventResponse, error) {
	var listEvent []*models.ListEvent
	var resp models.EventResponse
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllEvents(&listEvent, pageNo, pageSize, from, to, searchfilter)
	if err != nil {
		logging.Logger.WithField("error in listing Event", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return constant.INTERNALSERVERERROR, nil, err
	}
	resp.Events = listEvent
	resp.Count = count
	return constant.SUCESS, &resp, nil
}

func (es *EventService) GetEventService(id uint) (int, *models.EventResponse, error) {
	var resp models.EventResponse
	var event models.Event

	err := repository.Repo.FindEventById(&event, id)
	if err != nil {
		logging.Logger.WithField("error in Get event by id", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return constant.INTERNALSERVERERROR, nil, err

	}
	fmt.Println("created :", event.CreatedAt)
	resp.Event = &event
	return constant.SUCESS, &resp, nil
}

func (es *EventService) InvitedUserService(id uint) (int, *models.EventResponse, error) {
	var resp models.EventResponse
	var invitedGuest []*models.InvitedGuest

	invitedGuestCount, err := repository.Repo.FetchInvitedUsers(&invitedGuest, id)
	if err != nil {
		logging.Logger.WithField("error in  InvitedUser while Fetching", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return constant.INTERNALSERVERERROR, nil, err

	}
	resp.InvitedGuest = invitedGuest
	resp.InvitedGuestCount = invitedGuestCount

	return constant.SUCESS, &resp, nil

}

func (es *EventService) GetAttendieEventService(id uint) (int, *models.EventResponse, error) {
	var resp models.EventResponse

	var AttendedGuest []*models.AttendedGuest

	attendGuestCount, err := repository.Repo.FetchAttendiesUsers(&AttendedGuest, id)
	if err != nil {
		logging.Logger.WithField("error in GetAttendieUser fetching", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		fmt.Println("error", err)
		return constant.INTERNALSERVERERROR, nil, err

	}
	resp.AttendedGuest = AttendedGuest
	resp.AttendedGuestCount = attendGuestCount
	return constant.SUCESS, &resp, nil

}

func (es *EventService) CancelEventService(req *models.Events) (int, *models.EventResponse, error) {
	var resp models.EventResponse
	var event models.Events

	if err := repository.Repo.CancelEvent(&event, req.Reason, req.ID); err != nil {
		logging.Logger.WithField("error in updating the cancleling event ", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return constant.INTERNALSERVERERROR, nil, err

	}
	go sendNotification(req.ID, event.EventName, event.Reason)

	resp.Message = "Event Cancelled Successfully"
	return constant.SUCESS, &resp, nil

}

func sendNotification(eventid uint, eventName string, reason string) {
	var mailId []models.EmailId

	if err := repository.Repo.GetUserEmailId(&mailId, eventid); err != nil {
		fmt.Println("error in fetching mailId", err)
	}
	fmt.Println("mailID :", mailId)
	var email []string
	for _, val := range mailId {
		email = append(email, val.Email)
	}
	body := "Your " + eventName + " event has cancelled"
	mail.SendMail("Event cancelled", body, email...)

}

func (es *EventService) CreateEvent(req *models.EventRequest) (*models.EventResponse, error) {
	var resp models.EventResponse
	var event models.Events
	var startQueue models.StartTimeQueue
	var endQueue models.EndTimeQueue

	startTime, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)

	event.HostUserID = req.HostUserID
	event.EventName = req.EventName
	event.Location = req.Location
	event.EventMode = req.EventMode
	if req.EventMode == "Private" {
		event.Password = req.Password
	}
	event.Date = req.Date
	event.StartTime = startTime
	event.EndTime = endTime
	event.Coordinates = req.Coordinates
	event.Status = "Planned"
	if err := repository.Repo.Insert(&event); err != nil {
		logging.Logger.WithField("error in Inserting the  event table", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return nil, err
	}
	fmt.Println("event Id", event.ID)
	for _, userId := range req.InvitedUserId {
		var eventMetadata models.EventMetadatas
		eventMetadata.EventId = event.ID
		eventMetadata.InvitedUserId = userId
		eventMetadata.IsAttended = false
		eventMetadata.NoOfCheckIn = 0
		if err := repository.Repo.Insert(&eventMetadata); err != nil {
			logging.Logger.WithField("error in Inserting the eventmetadata ", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
			return nil, err
		}
	}
	t1 := time.Now()
	todayDate := t1.Format("2006-01-02")
	if req.Date == todayDate {
		startQueue.EventID = event.ID
		startQueue.StartedAt = event.StartTime
		StartTimmings = append(StartTimmings, startQueue)
		endQueue.EventID = event.ID
		endQueue.EndAt = event.EndTime
		EndTimmings = append(EndTimmings, endQueue)
		sort.Sort(StartTimeSlice(StartTimmings))
		sort.Sort(EndTimeSlice(EndTimmings))
		ChannelStart <- true
		ChannelEnd <- true
	}
	resp.EventID = event.ID
	resp.Message = "event Added SuccessFully"
	return &resp, nil
}
func (es *EventService) FetchEventById(id uint) (*models.EventResponse, error) {
	var resp models.EventResponse
	var event models.Event
	var userId []*models.InviUsersID
	if err := repository.Repo.FindEventById(&event, id); err != nil {
		return nil, err
	}
	count, err := repository.Repo.ListInvitedUserId(&userId, id)
	if err != nil {
		return nil, err
	}

	for _, usrId := range userId {
		resp.InvitedUserId = append(resp.InvitedUserId, &usrId.InvitedUserId)
	}

	resp.InvitedGuestCount = count
	//resp.InvitedUserId = userId
	resp.Event = &event
	//resp.InvitedUserId = userId

	return &resp, nil
}

func (es *EventService) UpdateEvent(req *models.EventRequest) (*models.EventResponse, error) {
	var resp models.EventResponse
	var updateEvent models.Events
	var startQueue models.StartTimeQueue
	var endQueue models.EndTimeQueue
	var userId []*models.InviUsersID

	startTime, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)

	updateEvent.HostUserID = req.HostUserID
	updateEvent.EventName = req.EventName
	updateEvent.Location = req.Location
	updateEvent.EventMode = req.EventMode
	if req.EventMode == "Private" {
		updateEvent.Password = req.Password
	}
	updateEvent.Date = req.Date
	updateEvent.StartTime = startTime
	updateEvent.EndTime = endTime
	updateEvent.Coordinates = req.Coordinates
	updateEvent.Status = req.Status
	if err := repository.Repo.UpdateEvent(&models.Events{}, req.EventId, updateEvent); err != nil {
		logging.Logger.WithField("error in Update the event ", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return nil, err
	}

	_, err := repository.Repo.ListInvitedUserId(&userId, req.EventId)
	if err != nil {
		return nil, err
	}

	for _, id := range req.InvitedUserId {
		contain := IsContains(userId, id)
		if !contain {
			var eventMetadata models.EventMetadatas
			eventMetadata.EventId = req.EventId
			eventMetadata.InvitedUserId = id
			eventMetadata.IsAttended = false
			eventMetadata.NoOfCheckIn = 0
			if err := repository.Repo.Insert(&eventMetadata); err != nil {
				logging.Logger.WithField("error in Inserting the eventmetadata ", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
				return nil, err
			}
		}
	}
	t1 := time.Now().UTC()
	todayDate := t1.Format("2006-01-02")
	timestring := t1.Format("2006-01-02 15:04:05")
	tie, err := time.Parse("2006-01-02 15:04:05", timestring)
	if err != nil {
		fmt.Println("error in parsing time", err)
	}

	if req.Date == todayDate {
		fmt.Println(tie, startTime)
		b := tie.After(startTime)
		if !b {
			fmt.Println("startTime ", b)
			startQueue.EventID = req.EventId
			startQueue.StartedAt = updateEvent.StartTime
			StartTimmings = append(StartTimmings, startQueue)
			sort.Sort(StartTimeSlice(StartTimmings))
			ChannelStart <- true
		}
		e := tie.After(endTime)
		if !e {
			fmt.Println("endTime ", e)

			endQueue.EventID = req.EventId
			endQueue.EndAt = updateEvent.EndTime
			EndTimmings = append(EndTimmings, endQueue)
			sort.Sort(EndTimeSlice(EndTimmings))
			ChannelEnd <- true
		}

	}

	resp.Message = "updated Successfully"
	return &resp, nil

}

func (es *EventService) InvitedEvents(id uint) (*models.EventResponse, error) {
	var resp models.EventResponse
	var invitedEvents []*models.ListInvitedEvents

	count, err := repository.Repo.InvitedEvents(&invitedEvents, id)
	if err != nil {
		logging.Logger.WithField("error in  Getting the InvitedEvents ", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)

		return nil, err
	}

	for _, value := range invitedEvents {
		var invitedGuest []*models.InvitedGuest
		invitedGuestCount, err := repository.Repo.FetchInvitedUsers(&invitedGuest, value.ID)
		if err != nil {
			return nil, err
		}
		value.InvitedUserCount = invitedGuestCount

	}

	resp.InivitedEventCount = count
	resp.InivitedEvents = invitedEvents

	return &resp, nil
}

func (es *EventService) HostedEvents(id uint) (*models.EventResponse, error) {
	var resp models.EventResponse
	var hostedEvents []*models.ListHostedEvents
	count, err := repository.Repo.HostedEvents(&hostedEvents, id)
	if err != nil {
		logging.Logger.WithField("error in  Getting the HostedEvent ", err).WithError(err).Error(constant.INTERNALSERVERERROR, err)
		return nil, err
	}
	for _, value := range hostedEvents {
		var invitedGuest []*models.InvitedGuest
		invitedGuestCount, err := repository.Repo.FetchInvitedUsers(&invitedGuest, value.ID)
		if err != nil {
			return nil, err
		}
		value.InvitedUserCount = invitedGuestCount
	}
	resp.HostedEvents = hostedEvents
	resp.HostedEventCount = count

	return &resp, nil
}

func (es *EventService) EventCheckIn(req *models.EventCheckInRequest) (*models.EventResponse, error) {
	var resp models.EventResponse
	var event models.Events

	if err := repository.Repo.FindById(&event, int(req.EventId)); err != nil {
		return nil, err
	}
	if event.EventMode == "Private" {
		if event.Password != req.Password {
			return nil, errors.New("invalid Event Password")
		}
	}
	coordinates := strings.Split(event.Coordinates, ",")
	lat1, err := strconv.ParseFloat(req.Latitude, 64)
	if err != nil {
		return nil, err
	}
	long1, err := strconv.ParseFloat(req.Longitude, 64)
	if err != nil {
		return nil, err
	}
	lat2, err := strconv.ParseFloat(coordinates[0], 64)
	if err != nil {
		return nil, err
	}
	long2, err := strconv.ParseFloat(coordinates[1], 64)
	if err != nil {
		return nil, err
	}
	if distance.Distance(lat1, long1, lat2, long2) > 1 {
		return nil, errors.New("you can't check in to the location which is more than 1 mile away")
	}
	if err := repository.Repo.VerifyEventCheckin(req.EventId, req.UserID); err != nil {
		return nil, err
	}
	var eventmeta models.EventMetadatas
	up, err := repository.Repo.GetEventMetadata(&eventmeta, req.EventId, req.UserID)
	if err != nil {
		return nil, err
	}
	if up {
		update := make(map[string]interface{})
		if eventmeta.IsAttended {
			update["no_of_check_in"] = eventmeta.NoOfCheckIn + 1
			update["check_out"] = sql.NullTime{}
		} else {
			update["check_in"] = time.Now()
			update["is_attended"] = true
			update["no_of_check_in"] = 1

		}

		if err := repository.Repo.UpdateEventMetadata(&models.EventMetadatas{}, req.EventId, req.UserID, update); err != nil {
			return nil, err
		}
		resp.Message = "Event Checked in successfully"
		return &resp, nil
	}

	var eventMetadata models.EventMetadatas
	eventMetadata.EventId = req.EventId
	eventMetadata.InvitedUserId = req.UserID
	eventMetadata.IsAttended = true
	eventMetadata.CheckIn = time.Now()
	eventMetadata.NoOfCheckIn = 1
	if err := repository.Repo.Insert(&eventMetadata); err != nil {
		return nil, err
	}
	resp.Message = "Event Checked in successfully"

	return &resp, nil
}

func (es *EventService) EventCheckOut(req *models.EventCheckInRequest) (*models.EventResponse, error) {
	var resp models.EventResponse
	var event models.Events
	//var eventmetadata models.EventMetadatas
	if err := repository.Repo.FindById(&event, int(req.EventId)); err != nil {
		return nil, err
	}
	//eventmetadata.CheckOut = time.Now()
	if err := repository.Repo.EventCheckout(&models.EventMetadatas{}, req.EventId, req.UserID); err != nil {
		return nil, err
	}
	resp.Message = "Event Checked out successfully"

	return &resp, nil
}

func (es *EventService) ListProfilesEventCheckIn(userId, eventId uint) (*models.EventResponse, error) {
	var resp models.EventResponse
	var users []*models.MatchedProfileResponse
	var count int64

	count, err := repository.Repo.ListProfilesEventCheckIn(&users, eventId, userId)
	if err != nil {
		return nil, err
	}
	resp.Count = count
	resp.ListProfilesEventCheckIn = users
	return &resp, nil
}

/*********************Cron ************************************************/

func EventServiceCronJob() {
	fmt.Println("**********************Cron service for Event  Started *******************************")
	var events []*models.Events
	t1 := time.Now().UTC()
	fmt.Println("time", t1)
	todayDate := t1.Format("2006-01-02 15:04:05")

	fmt.Println(todayDate)
	if err := repository.Repo.FindEventByStartTime(&events, todayDate); err != nil {
		fmt.Println("error in fetching user with date", err)
	}
	for _, event := range events {
		var startQueue models.StartTimeQueue
		startQueue.EventID = event.ID
		startQueue.StartedAt = event.StartTime
		StartTimmings = append(StartTimmings, startQueue)
	}
	fmt.Println("Today's total event start time :", len(events))

	if len(events) > 0 {
		sort.Sort(StartTimeSlice(StartTimmings))
		ChannelStart <- true
	}
	if err := repository.Repo.FindEventByEndTime(&events, todayDate); err != nil {
		fmt.Println("error in fetching user with date", err)
	}
	for _, event := range events {
		var endQueue models.EndTimeQueue
		endQueue.EventID = event.ID
		endQueue.EndAt = event.EndTime
		//fmt.Println("time in end", event.EndTime.Local())
		EndTimmings = append(EndTimmings, endQueue)
	}
	fmt.Println("Today's total event end time :", len(events))
	if len(events) > 0 {
		sort.Sort(EndTimeSlice(EndTimmings))
		ChannelEnd <- true

	}
	fmt.Println("**********************Cron service for Event Completed *******************************")

}

func LoadEvents() {
	fmt.Println("**********************Load  Event  Service Started *******************************")
	var events []*models.Events
	t1 := time.Now().UTC()
	fmt.Println("time", t1)

	todayDate := t1.Format("2006-01-02 15:04:05")
	fmt.Println(todayDate)
	if err := repository.Repo.FindEventByStartTime(&events, todayDate); err != nil {
		fmt.Println("error in fetching user with date", err)
	}
	for _, event := range events {
		var startQueue models.StartTimeQueue
		startQueue.EventID = event.ID
		startQueue.StartedAt = event.StartTime
		StartTimmings = append(StartTimmings, startQueue)
	}
	if len(events) > 0 {
		fmt.Println("called", len(events))
		sort.Sort(StartTimeSlice(StartTimmings))
		ChannelStart <- true
	}
	if err := repository.Repo.FindEventByEndTime(&events, todayDate); err != nil {
		fmt.Println("error in fetching user with date", err)
	}
	//fmt.Println("startTime", StartTimmings[0].EventID)
	for _, event := range events {
		var endQueue models.EndTimeQueue
		endQueue.EventID = event.ID
		endQueue.EndAt = event.EndTime
		fmt.Println("time in end", event.EndTime.Local())
		EndTimmings = append(EndTimmings, endQueue)
	}
	fmt.Println("length", len(events))
	if len(events) > 0 {
		fmt.Println("called")
		sort.Sort(EndTimeSlice(EndTimmings))

		ChannelEnd <- true

	}
	//	fmt.Println("endTime", EndTimmings[0].EventID)

	fmt.Println("**********************Load Event service Completed *******************************")

}

/************** IOS CREATEOTP ***************************/

func (us *UserService) GenerateOTP(phoneNo string) (*models.UserLoginResponse, error) {
	var user models.Users
	var resp models.UserLoginResponse
	var userDetail models.UserDetails
	newUser, err := repository.Repo.FindUserWithPhoneNo(&user, phoneNo)
	if err != nil {
		return nil, err
	}
	gauth := user.GAuth
	if newUser {

		value, ok := Users.Load(phoneNo)
		us, _ := value.(models.UserDetails)
		gauth = us.GAuth
		if !ok {
			key, err := totp.Generate(
				totp.GenerateOpts{
					Issuer:      phoneNo,
					AccountName: phoneNo,
					Period:      120,
				})
			if err != nil {
				return nil, err
			}
			userDetail.PhoneNumber = phoneNo
			userDetail.GAuth = key.Secret()
			Users.LoadOrStore(phoneNo, userDetail)
			gauth = key.Secret()
		}

	}
	//resp.Gauth = user.GAuth
	value, err := totp.GenerateCodeCustom(gauth, time.Now(), totp.ValidateOpts{
		Digits: 4,
		Period: 120,
	})
	if err != nil {
		return nil, err
	}
	resp.Otp = value
	go sendSmsNotification(value)

	var account_sid, auth_token, from_number string

	/*if viper.GetString("twilio.env.staging.account_sid") != "" {
		account_sid = viper.GetString("twilio.env.staging.account_sid")
		auth_token = viper.GetString("twilio.env.staging.auth_token")
		from_number = viper.GetString("twilio.env.staging.from_number")
	}*/

	if viper.GetString("twilio.env.production.account_sid") != "" {
		account_sid = viper.GetString("twilio.env.production.account_sid")
		auth_token = viper.GetString("twilio.env.production.auth_token")
		from_number = viper.GetString("twilio.env.production.from_number")
	}

	/*if viper.GetString("twilio.env.development.account_sid") != "" {
		account_sid = viper.GetString("twilio.env.development.account_sid")
		auth_token = viper.GetString("twilio.env.development.auth_token")
		from_number = viper.GetString("twilio.env.development.from_number")
	}*/

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: account_sid,
		Password: auth_token,
	})

	params := &openapi.CreateMessageParams{}
	params.SetFrom(from_number)
	params.SetTo(phoneNo)

	params.SetBody("<#> Sondr Login OTP is " + value)
	fmt.Println("OTP", value)

	resptwilio, err := client.Api.CreateMessage(params)
	fmt.Println("OTP SENT SUCCESSFULLY", resp)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resptwilio)
		fmt.Println("Response: " + string(response))
	}

	resp.Message = "Four-digit code sent successfully"

	return &resp, nil
}

func sendSmsNotification(value string) {
	log.Println("otp", value)

}

/***************** Verify OTP IOS panel *********************/

func (us *UserService) VerifyOtp(phoneNo string, otp string) (*models.UserLoginResponse, error) {
	var user models.Users
	var resp models.UserLoginResponse
	var gauth string
	fmt.Println("Phonenumber", phoneNo)
	newUser, err := repository.Repo.FindUserWithPhoneNo(&user, phoneNo)
	if err != nil {
		return nil, err
	}
	// if !newUser && user.IsLoggedIn {
	// 	return nil, errors.New("Please logout from the other device")
	// }
	// if !newUser && user.Blocked {
	// 	return nil, errors.New("Your Profile is Blocked")
	// }
	gauth = user.GAuth
	resp.ProfileStatus = user.ProfileStatus
	if newUser {
		fmt.Println("user", newUser)
		var userDetail models.UserDetails
		value, ok := Users.Load(phoneNo)
		if !ok {
			return nil, errors.New("Invalid four-digit code")
		}
		userDetail, _ = value.(models.UserDetails)
		fmt.Println("gauth", userDetail.GAuth, "phonenumber", userDetail.PhoneNumber)
		gauth = userDetail.GAuth
		resp.ProfileStatus = "NewUser"

	}
	fmt.Println(gauth, otp)
	valid, err := totp.ValidateCustom(otp, gauth, time.Now(), totp.ValidateOpts{
		Digits: 4,
		Period: 120,
	})

	fmt.Println(valid)
	if err != nil {
		return nil, errors.New("Invalid four-digit code")
	}
	if valid {
		resp.Message = "Verified sucessfully"
		if user.ProfileStatus != "" && user.ProfileStatus != "NewUser" {
			resp.UserId = user.ID
			token, err := GenerateUserToken(&user)
			if err != nil {
				return nil, err
			}
			resp.Token = token
			// if err := repository.Repo.UpdateUserProfile(&models.Users{}, int(user.ID), map[string]interface{}{
			// 	"is_logged_in": true,
			// }); err != nil {
			// 	return nil, errors.New("unable to update user profile")
			// }
		}
		return &resp, nil
	}

	return nil, errors.New("Invalid four-digit code")
}

func (us *UserService) SocialLogin(details map[string]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	var user models.Users
	new, err := repository.Repo.FindUserWithReferenceID(&user, details["id"].(string))
	if err != nil {
		return nil, err

	}
	if new {
		return details, nil
	}

	res["profileStatus"] = user.ProfileStatus
	res["userID"] = user.ID
	token, err := GenerateUserToken(&user)
	if err != nil {
		return nil, err
	}
	res["token"] = token
	return res, nil
}

func GenerateUserToken(user *models.Users) (string, error) {
	var mySigningKey = []byte(viper.GetString("secret.UserKey"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user.Email

	claims["id"] = strconv.FormatUint(uint64(user.ID), 10)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func EventDurationStarting() {
	//prevTime := time.Now()
	var timer *time.Timer
	for {

		value := <-ChannelStart
		fmt.Println("New Start time Found", value)
		if len(StartTimmings) > 1 && timer != nil {
			timer.Stop()
		}
		prevTime := time.Now()
		log.Println("length of starttimings ", len(StartTimmings))
		log.Println("starting  event id", StartTimmings[0].EventID)
		log.Println("Start Time duration ", time.Duration(StartTimmings[0].StartedAt.Sub(prevTime)))
		timer = time.NewTimer(time.Duration(StartTimmings[0].StartedAt.Sub(prevTime)))
		go func() {
			select {
			case <-timer.C:
				log.Println("Timer start channel got in case", StartTimmings[0].EventID)
				if err := repository.Repo.UpdateEventStartingStatus(&models.Events{}, StartTimmings[0].EventID); err != nil {
					fmt.Println("error in update ", err)
				}
				fmt.Println("updated Successfully", StartTimmings[0].EventID)
				StartTimmings = StartTimmings[1:]

				if len(StartTimmings) > 0 {
					//fmt.Println("after removed updated element ", StartTimmings[0].EventID)
					ChannelStart <- true
				}
			}
		}()
	}
}

func EventDurationEnding() {
	//prevTime := time.Now()
	//fmt.Println("today time", prevTime)
	var endtimer *time.Timer
	for {
		value := <-ChannelEnd
		fmt.Println("New Value founf in EndTimming array", value)
		if len(EndTimmings) > 1 && endtimer != nil {
			endtimer.Stop()
		}
		prevTime := time.Now()
		log.Println("length of Endtimings ", len(EndTimmings))
		log.Println("ending  event id", EndTimmings[0].EventID)
		log.Println("end Time duration ", time.Duration(EndTimmings[0].EndAt.Sub(prevTime)))
		endtimer = time.NewTimer(time.Duration(EndTimmings[0].EndAt.Sub(prevTime)))
		//	prevTime = EndTimmings[0].EndAt
		go func() {
			select {
			case <-endtimer.C:
				log.Println("Timer End channel got in case", EndTimmings[0].EventID)

				if err := repository.Repo.UpdateEventExpiryStatus(&models.Events{}, EndTimmings[0].EventID); err != nil {
					fmt.Println("error in update ", err)
				}
				EndTimmings = EndTimmings[1:]
				if len(EndTimmings) > 0 {
					ChannelEnd <- true
				}
			}
		}()

	}
}

func (p StartTimeSlice) Len() int {

	return len(p)
}

// Define compare

func (p StartTimeSlice) Less(i, j int) bool {

	return p[i].StartedAt.Before(p[j].StartedAt)
}

// Define swap over an array

func (p StartTimeSlice) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]
}

func (p EndTimeSlice) Len() int {

	return len(p)
}

// Define compare

func (p EndTimeSlice) Less(i, j int) bool {

	return p[i].EndAt.Before(p[j].EndAt)
}

// Define swap over an array

func (p EndTimeSlice) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]
}

func IsContains(myconfig []*models.InviUsersID, id uint) bool {
	var found bool
	for _, v := range myconfig {
		if v.InvitedUserId == id {
			// Found!
			found = true
		}

	}
	return found
}
