package wfclientapi

// Types of params
const (
	Token          = "appid"
	Input          = "input"
	ResponseFormat = "output"
	PodState       = "podstate"
	Include        = "includepodid"
)

// Types of Format
const (
	ResponseFormatJSON = "JSON"
)

// Types of Pods
const (
	PodStateStepByStep = "step-by-step solution"
)

// WolframAPI endpoint
const APIEndpoint = "https://api.wolframalpha.com/v2/query"

// Pod ids
const (
	PodIdResult      = "Result"
	PodIdDecimalAppx = "DecimalApproximation"
)

const (
	PodStateStepByStepTitle = "Possible intermediate steps"
)
