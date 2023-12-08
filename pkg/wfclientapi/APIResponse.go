package wfclientapi

import "encoding/json"

type ApiResponse struct {
	QueryResult *QueryResult
	StatusCode  int
}
type QueryResult struct {
	Success       bool    `json:"success"`
	Error         bool    `json:"error"`
	NumPods       int     `json:"numpods"`
	DataTypes     string  `json:"datatypes"`
	TimedOut      string  `json:"timedout"`
	TimedOutPods  string  `json:"timedoutpods"`
	Timing        float64 `json:"timing"`
	ParseTiming   float64 `json:"parsetiming"`
	ParseTimedOut bool    `json:"parsetimedout"`
	Recalculate   string  `json:"recalculate"`
	ID            string  `json:"id"`
	Host          string  `json:"host"`
	Server        string  `json:"server"`
	Related       string  `json:"related"`
	Version       string  `json:"version"`
	InputString   string  `json:"inputstring"`
	Pods          PodsMap `json:"pods"`
}

type Pod struct {
	Title           string      `json:"title"`
	Scanner         string      `json:"scanner"`
	ID              string      `json:"id"`
	Position        int         `json:"position"`
	Error           bool        `json:"error"`
	NumSubpods      int         `json:"numsubpods"`
	Subpods         []Subpod    `json:"subpods"`
	ExpressionTypes interface{} `json:"expressiontypes"` // Can be Default or an array of Default
	States          interface{} `json:"states"`          // Can be Default or an array of Default
}

type PodsMap map[string]Pod

func (pm *PodsMap) UnmarshalJSON(data []byte) error {
	var pods []Pod
	if err := json.Unmarshal(data, &pods); err != nil {
		return err
	}

	podsMap := make(map[string]Pod, len(pods))
	for _, p := range pods {
		podsMap[p.ID] = p
	}

	*pm = podsMap
	return nil
}

type Subpod struct {
	Title     string `json:"title"`
	Img       Image  `json:"img"`
	Plaintext string `json:"plaintext"`
}

type Image struct {
	Src             string `json:"src"`
	Alt             string `json:"alt"`
	Title           string `json:"title"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Type            string `json:"type"`
	Themes          string `json:"themes"`
	ColorInvertable bool   `json:"colorinvertable"`
	ContentType     string `json:"contenttype"`
}
