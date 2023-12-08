package wfclientapi

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

type WfClient interface {
	MakeRequest(params url.Values) (*ApiResponse, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type wfClient struct {
	client      HTTPClient
	debug       bool
	token       string
	apiEndpoint string
}

func NewWfAPIClient(token string, debug bool) (WfClient, error) {
	client := &wfClient{
		client:      &http.Client{},
		token:       token,
		apiEndpoint: APIEndpoint,
		debug:       debug,
	}

	return client, nil
}

func (wfc *wfClient) MakeRequest(params url.Values) (*ApiResponse, error) {
	if wfc.debug {
		logrus.Debugf("values %v\n", params)
	}
	req, err := http.NewRequest("GET", APIEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = wfc.buildParams(params)
	resp, err := wfc.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	aresp, err := wfc.ToApiResponse(resp)
	if err != nil {
		return nil, err
	}
	return aresp, nil
}

func (wfc *wfClient) buildParams(params url.Values) string {
	params.Add(Token, wfc.token)
	params.Add(ResponseFormat, ResponseFormatJSON)
	return params.Encode()
}

func (wfc *wfClient) ToApiResponse(response *http.Response) (*ApiResponse, error) {
	apiResponse := &ApiResponse{StatusCode: response.StatusCode}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}
	if wfc.debug {
		logrus.Debugf("ApiResponse: %v\n", apiResponse)
	}
	return apiResponse, nil
}
