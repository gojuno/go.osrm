package osrm

import "errors"

type ErrCode string

const (
	OK                ErrCode = "Ok"
	ErrInvalidURL     ErrCode = "InvalidUrl"
	ErrInvalidService ErrCode = "InvalidService"
	ErrInvalidVersion ErrCode = "InvalidVersion"
	ErrInvalidOptions ErrCode = "InvalidOptions"
	ErrInvalidQuery   ErrCode = "InvalidQuery"
	ErrInvalidValue   ErrCode = "InvalidValue"
	ErrNoSegment      ErrCode = "NoSegment"
	ErrTooBig         ErrCode = "TooBig"
	ErrNoRoute        ErrCode = "NoRoute"
	ErrNoTable        ErrCode = "NoTable"
	ErrNoMatch        ErrCode = "NoMatch"

	// ErrInternal for errors which don't related to OSRM
	ErrInternal ErrCode = "Internal"
)

var (
	ErrEmptyServiceName = errors.New("osrm5: the request should contain a service name")
	ErrEmptyProfileName = errors.New("osrm5: the request should contain a profile name")
	ErrNoCoordinates    = errors.New("osrm5: the request should contain coordinates")
)

type Error interface {
	ErrCode() ErrCode
	Error() string
}

type ResponseError struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"message"`
}

// WrapError creates an error with internal code
func WrapError(err error) Error {
	return ResponseError{ErrInternal, err.Error()}
}

func (err ResponseError) ErrCode() ErrCode {
	return err.Code
}

func (err ResponseError) Error() string {
	return err.Msg
}
