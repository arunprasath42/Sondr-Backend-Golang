package constant

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	SUCESS               = 200
	BADREQUEST           = 400
	UNAUTHORIZED         = 401
	FORBIDDEN            = 403
	NOTFOUND             = 404
	METHODNOTALLOWED     = 405
	INTERNALSERVERERROR  = 500
	NOTIMPLEMENTED       = 501
	BADGATEWAY           = 502
	SERVICEUNAVAILABLE   = 503
	GATEWAYTIMEOUT       = 504
	UNSUPPORTEDMEDIATYPE = 415
	UNPROCESSABLEENTITY  = 422
	Subject              = "Password for Sondr Login"
	SubjectforSubadmin   = "Subadmin credentials"
	S3URL                = "https://sondr-app-dev.s3.sa-east-1.amazonaws.com/"

	ADMINBLOCKED         = "ADMIN BLOCKED"
	INVALIDUSER          = "INVALID USER"
	MATCHSTATUSREQUESTED = "Requested"
	MATCHSTATUSMATCHED   = "Matched"
	MATCHSTATUSREJECTED  = "Rejected"
	NOTIFICATIONTYPELIKE = "LIKE"
)

/*********************************GENERIC FUNCTION FOR JSON FORMATTING**************************************/
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func Httpmethod(req interface{}, url string, method string, resp interface{}) error {

	client := &http.Client{}
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	//request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &resp)
	return nil
}
