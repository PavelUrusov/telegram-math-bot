package wfclientapi

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

type WfClient interface {
	MakeRequest(params url.Values) (*ApiResponse, error)
	MakeArithmeticRequest(input string) (*Solution, error)
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
		logrus.Infof("values %v\n", params)
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
		logrus.Infof("ApiResponse: %v\n", apiResponse)
	}
	return apiResponse, nil
}

func (wfc *wfClient) MakeArithmeticRequest(input string) (*Solution, error) {
	if wfc.debug {
		logrus.Infof("input %s\n", input)
	}
	params := wfc.arithmeticParams(input)
	resp, err := wfc.MakeRequest(params)
	if err != nil {
		return nil, err
	}
	sln, err := wfc.toArithmeticSolution(resp)
	if err != nil {
		return nil, err
	}
	if wfc.debug {
		logrus.Infof("sln %v\n", sln)
	}
	return sln, nil
}

func (wfc *wfClient) buildParams(params url.Values) string {
	params.Add(Token, wfc.token)
	params.Add(ResponseFormat, ResponseFormatJSON)
	return params.Encode()
}

func (wfc *wfClient) arithmeticParams(input string) url.Values {
	params := url.Values{}
	params.Add(Input, input)
	params.Add(Include, PodIdResult)
	params.Add(Include, PodIdDecimalAppx)
	params.Add(PodState, PodStateStepByStep)
	return params
}

func (wfc *wfClient) toArithmeticSolution(r *ApiResponse) (*Solution, error) {
	qr := r.QueryResult
	sln := &Solution{}
	pod, ok := qr.Pods[PodIdDecimalAppx]
	if ok {
		sln.Answers = append(sln.Answers, pod.Subpods[0].Plaintext)
		return sln, nil
	}
	pod, ok = qr.Pods[PodIdResult]
	if !ok {
		return nil, errors.New("can't find a solution")
	}
	for _, s := range pod.Subpods {
		if s.Title == PodStateStepByStepTitle {
			sln.ImageStepsURL = s.Img.Src
			continue
		}
		sln.Answers = append(sln.Answers, s.Plaintext)
	}

	return sln, nil
}
