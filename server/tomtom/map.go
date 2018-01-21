package tomtom

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	baseURL = "https://api.tomtom.com/traffic/services/4/flowSegmentData/absolute/10/json?key=%s&point=%s,%s"
)

type API struct {
	key    string
	secret string
}

func New() *API {
	return &API{
		key:    "TqUq1Sw7zST6k5BKAhOiLKPJGi8SsFkK",
		secret: "ZXhPmCGz89fWxGSd",
	}
}

func (api *API) GetSegmentAtCoordinate(lat string, lon string) (*FlowSegmentData, error) {
	request := fmt.Sprintf(baseURL, api.key, lat, lon)
	resp, err := api.makeRequest("GET", request, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Not okay :( " + strconv.Itoa(resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	segment := &tomTomResp{}
	err = json.Unmarshal(respBody, segment)
	if err != nil {
		return nil, err
	}

	return &segment.FlowSegmentData, nil
}

func (api *API) makeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	query := req.URL.Query()
	req.URL.RawQuery = query.Encode()
	return http.DefaultClient.Do(req)
}
