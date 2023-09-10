package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sondr-backend/src/models"
	"sondr-backend/src/service"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/response"
	"sondr-backend/utils/validator"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/instagram"
)

var oauthStateStringFb string = ""
var oauthConfFb = &oauth2.Config{}

var oauthConfGl = &oauth2.Config{}
var oauthStateStringGl = ""

var oauthConfIn = &oauth2.Config{}

func LoadSocialMediaCreditional() {
	oauthConfFb = &oauth2.Config{
		ClientID:     viper.GetString("fb.AppID"),
		ClientSecret: viper.GetString("fb.AppSecret"),
		RedirectURL:  viper.GetString("fb.ReDirectURL"),
		Scopes:       []string{"public_profile", "email", "user_gender", "user_birthday"},
		Endpoint:     facebook.Endpoint,
	}
	oauthConfGl = &oauth2.Config{
		ClientID:     viper.GetString("gl.ClientID"),
		ClientSecret: viper.GetString("gl.ClientSecret"),
		RedirectURL:  viper.GetString("gl.ReDirectURL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	oauthConfIn = &oauth2.Config{
		ClientID:     viper.GetString("insta.AppID"),
		ClientSecret: viper.GetString("insta.AppSecret"),
		RedirectURL:  viper.GetString("insta.ReDirectURL"),
		Scopes:       []string{"user_profile,user_media"},
		Endpoint:     instagram.Endpoint,
	}

}

func ListAllEvents(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	statusCode, resp, err := es.ListAllEventService(reqModel.PageNo, reqModel.PageSize, reqModel.SearchFilter, reqModel.From, reqModel.To)
	if err != nil {
		c.JSON(statusCode, response.ErrorMessage(statusCode, err))
		return
	}
	c.JSON(statusCode, response.SuccessResponse(resp))
}

func GetEventById(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	statusCode, resp, err := es.GetEventService(reqModel.ID)
	if err != nil {
		c.JSON(statusCode, response.ErrorMessage(statusCode, err))
		return
	}
	c.JSON(statusCode, response.SuccessResponse(resp))

}
func GetInvitedUsers(c *gin.Context) {

	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	statusCode, resp, err := es.InvitedUserService(reqModel.ID)
	if err != nil {
		c.JSON(statusCode, response.ErrorMessage(statusCode, err))
		return
	}
	c.JSON(statusCode, response.SuccessResponse(resp))
}

func GetAttendedUsers(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	statuscode, resp, err := es.GetAttendieEventService(reqModel.ID)
	if err != nil {
		c.JSON(statuscode, response.ErrorMessage(statuscode, err))
		return
	}
	c.JSON(statuscode, response.SuccessResponse(resp))
}

func CancelEvent(c *gin.Context) {
	var req models.Events
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	if err := validator.ValidateVariable(req.ID, "required", "id"); err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	statusCode, resp, err := es.CancelEventService(&req)
	if err != nil {
		c.JSON(statusCode, response.ErrorMessage(statusCode, err))
		return
	}
	c.JSON(statusCode, response.SuccessResponse(resp))

}

/***************************************************************IOS Module *******************************************************************/
func CreateEvent(c *gin.Context) {
	var req models.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	es := service.EventService{}
	resp, err := es.CreateEvent(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}
func FetchEventById(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.ValidateVariable(reqModel.ID, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	resp, err := es.FetchEventById(reqModel.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}
func UpdateEvent(c *gin.Context) {
	var req models.EventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	resp, err := es.UpdateEvent(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

func InvitedEvents(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if err := validator.ValidateVariable(req.UserId, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	resp, err := es.InvitedEvents(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}

func HostedEvents(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	if err := validator.ValidateVariable(req.UserId, "required", "id"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	es := service.EventService{}
	resp, err := es.HostedEvents(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}
func EventCheckIn(c *gin.Context) {
	var req models.EventCheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	if err := validator.Validate(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	es := service.EventService{}
	resp, err := es.EventCheckIn(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

func EventCheckOut(c *gin.Context) {
	var req models.EventCheckInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	if err := validator.ValidateVariable(req.EventId, "required", "eventID"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	if err := validator.ValidateVariable(req.UserID, "required", "userID"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	es := service.EventService{}
	resp, err := es.EventCheckOut(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

func ListProfilesEventCheckIn(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	if err := validator.ValidateVariable(req.UserId, "required", "UserId"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	if err := validator.ValidateVariable(req.EventId, "required", "eventId"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(constant.BADGATEWAY, err))
		return
	}
	es := service.EventService{}
	resp, err := es.ListProfilesEventCheckIn(req.UserId, req.EventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

/**************IOS create OTP **************************/

func CreateOTP(c *gin.Context) {

	var req models.Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := validator.ValidateVariable(req.PhoneNumber, "required", "phoneNumber"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	us := service.UserService{}
	resp, err := us.GenerateOTP(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}
func VerifyOtp(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := validator.ValidateVariable(req.PhoneNumber, "required", "PhoneNumber"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	if err := validator.ValidateVariable(req.OTP, "required", "otp"); err != nil {
		c.JSON(http.StatusBadGateway, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	us := service.UserService{}
	resp, err := us.VerifyOtp(req.PhoneNumber, req.OTP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}

func FaceBookLogin(c *gin.Context) {
	fmt.Println("redirectURl", oauthConfFb.RedirectURL)
	URL, err := url.Parse(oauthConfFb.Endpoint.AuthURL)
	if err != nil {
		fmt.Println("error in login", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConfFb.ClientID)
	parameters.Add("scope", strings.Join(oauthConfFb.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfFb.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateStringFb)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Println(url)
	c.Redirect(http.StatusTemporaryRedirect, url)

}
func FacebookCallBack(c *gin.Context) {
	fmt.Println("Callback-fb..")

	state := c.Request.FormValue("state")
	fmt.Println(state)
	if state != oauthStateStringFb {
		fmt.Println("invalid oauth state, expected " + oauthStateStringFb + ", got " + state + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Request.FormValue("code")
	fmt.Println("code", code)

	if code == "" {
		fmt.Println("Code not found..")
		c.JSON(http.StatusOK, response.CustomErrorMessage(http.StatusBadGateway, "Code Not Found to provide AccessToken..\n"))

		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			c.JSON(http.StatusOK, response.CustomErrorMessage(http.StatusBadGateway, "User has denied Permission.."))
			return
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := oauthConfFb.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		fmt.Println("TOKEN>> AccessToken>> " + token.AccessToken)
		fmt.Println("TOKEN>> Expiration Time>> " + token.Expiry.String())
		fmt.Println("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://graph.facebook.com/me?access_token=" + url.QueryEscape(token.AccessToken) + "&fields=email,gender,birthday,first_name,last_name")
		if err != nil {
			fmt.Println("Get: " + err.Error() + "\n")

			return
		}
		defer resp.Body.Close()

		responsebyte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ReadAll: " + err.Error() + "\n")
			return
		}

		fmt.Println("parseResponseBody: " + string(responsebyte) + "\n")
		var mapresp map[string]interface{}
		if err := json.Unmarshal(responsebyte, &mapresp); err != nil {
			fmt.Println("error in json unmarsahal", err)
			return
		}
		if mapresp["id"] == nil {
			c.JSON(http.StatusInternalServerError, response.CustomErrorMessage(http.StatusInternalServerError, string(responsebyte)))
		}
		fmt.Println("response MAp", mapresp["id"])
		us := service.UserService{}
		result, err := us.SocialLogin(mapresp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, response.SuccessResponse(result))
		return
	}
}

func GoogleLogin(c *gin.Context) {
	URL, err := url.Parse(oauthConfGl.Endpoint.AuthURL)
	if err != nil {
		fmt.Println("error in login", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConfGl.ClientID)
	parameters.Add("scope", strings.Join(oauthConfGl.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfGl.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateStringGl)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Println("url", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallBackGoogle(c *gin.Context) {
	fmt.Println("Callback-gl..")

	state := c.Request.FormValue("state")
	fmt.Println(state)
	if state != oauthStateStringGl {
		fmt.Println("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Request.FormValue("code")
	fmt.Println("code", code)

	if code == "" {
		fmt.Println("Code not found..")
		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			c.JSON(http.StatusBadRequest, response.CustomErrorMessage(http.StatusBadRequest, "User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		fmt.Println("TOKEN>> AccessToken>> " + token.AccessToken)
		fmt.Println("TOKEN>> Expiration Time>> " + token.Expiry.String())
		fmt.Println("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			fmt.Println("Get: " + err.Error() + "\n")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		defer resp.Body.Close()

		responsebytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ReadAll: " + err.Error() + "\n")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		fmt.Println("parseResponseBody: " + string(responsebytes) + "\n")
		var mapresp map[string]interface{}
		if err := json.Unmarshal(responsebytes, &mapresp); err != nil {
			fmt.Println("error in json unmarsahal", err)
		}

		fmt.Println("response MAp", mapresp["id"])
		us := service.UserService{}
		result, err := us.SocialLogin(mapresp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
			return
		}

		c.JSON(http.StatusOK, response.SuccessResponse(result))
		return
	}

}

func InstagramLogin(c *gin.Context) {
	URL, err := url.Parse(oauthConfIn.Endpoint.AuthURL)
	if err != nil {
		fmt.Println("error in login", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConfIn.ClientID)
	parameters.Add("scope", strings.Join(oauthConfIn.Scopes, " "))
	parameters.Add("redirect_uri", oauthConfIn.RedirectURL)
	parameters.Add("response_type", "code")
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Println("url", url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallBackInstagram(c *gin.Context) {
	fmt.Println("Callback-Insta..")

	code := c.Request.FormValue("code")
	fmt.Println("code", code)

	if code == "" {
		fmt.Println("Code not found..")
		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			c.JSON(http.StatusBadRequest, response.CustomErrorMessage(http.StatusBadRequest, "User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := oauthConfIn.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		fmt.Println("TOKEN>> AccessToken>> " + token.AccessToken)
		fmt.Println("TOKEN>> Expiration Time>> " + token.Expiry.String())
		fmt.Println("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://graph.instagram.com/me?fields=id,username&access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			fmt.Println("Get: " + err.Error() + "\n")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		defer resp.Body.Close()

		responsebytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ReadAll: " + err.Error() + "\n")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		fmt.Println("parseResponseBody: " + string(responsebytes) + "\n")
		var mapresp map[string]interface{}
		if err := json.Unmarshal(responsebytes, &mapresp); err != nil {
			fmt.Println("error in json unmarsahal", err)
		}

		fmt.Println("response MAp", mapresp["id"])
		c.JSON(http.StatusOK, response.SuccessResponse(mapresp))
		return
	}

}
