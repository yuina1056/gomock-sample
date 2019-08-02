package httpaccess

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TestAPI interface {
	GET(urlString string, authorization string, responseJSON interface{}) (statuscode int, err error)
}

type testapi struct {
	BaseURL string
}

func NewTestAPI(baseURL string) TestAPI {
	return &testapi{baseURL}
}

// GET return error
func (api testapi) GET(urlString string, authorization string, responseJSON interface{}) (statuscode int, err error) {
	url := api.BaseURL + urlString
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", authorization)
	request.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, err
	}
	json.Unmarshal(body, &responseJSON)
	return response.StatusCode, nil
}
