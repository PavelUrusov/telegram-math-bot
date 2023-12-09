package wfclientapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

type WfClient interface {
	MakeRequest(params url.Values) (*ApiResponse, error)
	MakeElementaryMathRequest(input string) (*Solution, error)
	MakeEquationRequest(input string) (*Solution, error)
	MakePlotRequest(input string) (*Solution, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type wfClient struct {
	client      HTTPClient
	Debug       bool
	token       string
	apiEndpoint string
}

func NewWfAPIClient(token string) (WfClient, error) {
	wfc := &wfClient{
		client:      &http.Client{},
		token:       token,
		apiEndpoint: APIEndpoint,
		Debug:       false,
	}
	if err := wfc.validateAPIKey(); err != nil {
		return nil, err
	}
	return wfc, nil
}

func (wfc *wfClient) MakeRequest(params url.Values) (*ApiResponse, error) {
	if wfc.Debug {
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
	aresp, err := wfc.toApiResponse(resp)
	if err != nil {
		return nil, err
	}
	return aresp, nil
}

func (wfc *wfClient) MakeElementaryMathRequest(input string) (*Solution, error) {
	if wfc.Debug {
		logrus.Infof("input %s\n", input)
	}
	params := wfc.arithmeticParams(input)
	resp, err := wfc.MakeRequest(params)
	if err != nil {
		return nil, err
	}
	sln, err := wfc.toSolution(resp)
	if err != nil {
		return nil, err
	}
	if wfc.Debug {
		logrus.Infof("sln %v\n", sln)
	}
	return sln, nil
}
func (wfc *wfClient) MakeEquationRequest(input string) (*Solution, error) {
	if wfc.Debug {
		logrus.Infof("input string %s\n", input)
	}
	params := wfc.equationParams(input)
	resp, err := wfc.MakeRequest(params)
	if err != nil {
		return nil, err
	}
	sln, err := wfc.toSolution(resp)
	if err != nil {
		return nil, err
	}
	if wfc.Debug {
		logrus.Infof("solution %v\n", sln)
	}
	return sln, nil
}
func (wfc *wfClient) MakePlotRequest(input string) (*Solution, error) {
	if wfc.Debug {
		logrus.Infof("input string %s\n", input)
	}
	params := wfc.plotParams(input)
	resp, err := wfc.MakeRequest(params)
	if err != nil {
		return nil, err
	}
	sln, err := wfc.toSolution(resp)
	if err != nil {
		return nil, err
	}
	if wfc.Debug {
		logrus.Infof("solution %v\n", sln)
	}
	return sln, nil
}

func (wfc *wfClient) arithmeticParams(input string) url.Values {
	params := url.Values{}
	params.Add(Input, input)
	params.Add(Include, PodIdResult)
	params.Add(Include, PodIdDecimalAppx)
	params.Add(PodState, PodStateStepByStep)
	return params
}
func (wfc *wfClient) equationParams(input string) url.Values {
	params := url.Values{}
	params.Add(Input, input)
	params.Add(PodState, PodStateStepByStep)
	params.Add(Include, PodIdSolution)
	params.Add(Include, PodIdRealSolution)
	params.Add(Include, PodIdComplexSolution)

	return params
}
func (wfc *wfClient) plotParams(input string) url.Values {
	params := url.Values{}
	params.Add(Input, input)
	params.Add(Include, PodIdPlot)

	return params
}

func (wfc *wfClient) toSolution(r *ApiResponse) (*Solution, error) {
	qr := r.QueryResult
	sln := &Solution{}
	if len(qr.Pods) == 0 {
		return nil, errors.New("can't find a solution")
	}
	for _, p := range qr.Pods {
		if p.ID == PodIdPlot {
			sln.ImageStepsURL = append(sln.ImageStepsURL, p.Subpods[0].Img.Src)
			continue
		}
		for _, sp := range p.Subpods {
			if sp.Title == PodStateStepByStepTitle {
				sln.ImageStepsURL = append(sln.ImageStepsURL, sp.Img.Src)
				continue
			}
			sln.Answers = append(sln.Answers, sp.Plaintext)
		}
	}

	return sln, nil
}
func (wfc *wfClient) buildParams(params url.Values) string {
	params.Add(Token, wfc.token)
	params.Add(ResponseFormat, ResponseFormatJSON)
	return params.Encode()
}
func (wfc *wfClient) toApiResponse(response *http.Response) (*ApiResponse, error) {
	apiResponse := &ApiResponse{StatusCode: response.StatusCode}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}
	if wfc.Debug {
		logrus.Infof("ApiResponse: %v\n", apiResponse)
	}
	return apiResponse, nil
}
func (wfc *wfClient) validateAPIKey() error {
	_, err := wfc.MakeElementaryMathRequest("2+2")
	if err != nil {
		return fmt.Errorf("API key validation failed: %v", err)
	}
	if wfc.Debug {
		logrus.Infof("API key validated successfully")
	}
	return nil
}
