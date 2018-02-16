package osrm

import "errors"

// Error codes that could be returned from OSRM
const (
	ErrorCodeInvalidURL     = "InvalidUrl"
	ErrorCodeInvalidService = "InvalidService"
	ErrorCodeInvalidVersion = "InvalidVersion"
	ErrorCodeInvalidOptions = "InvalidOptions"
	ErrorCodeInvalidQuery   = "InvalidQuery"
	ErrorCodeInvalidValue   = "InvalidValue"
	ErrorCodeNoSegment      = "NoSegment"
	ErrorCodeTooBig         = "TooBig"
	ErrorCodeNoRoute        = "NoRoute"
	ErrorCodeNoTable        = "NoTable"
	ErrorCodeNoMatch        = "NoMatch"
	errorCodeOK             = "Ok" // "Ok" error code never returned to library client, thus not exported
)

// Invalid request errors
var (
	ErrEmptyProfileName = errors.New("osrm5: the request should contain a profile name")
	ErrNoCoordinates    = errors.New("osrm5: the request should contain coordinates")
	ErrEmptyServiceName = errors.New("osrm5: the request should contain a service name")
)

// ResponseError represent OSRM API error
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrCode returns error code from OSRM response
func (err ResponseError) ErrCode() string {
	return err.Code
}

func (err ResponseError) Error() string {
	return err.Code + " - " + err.Message
}

type responseStatus struct {
	ResponseError
}

func (r responseStatus) apiError() error {
	if r.Code != errorCodeOK {
		return r.ResponseError
	}
	return nil
}
