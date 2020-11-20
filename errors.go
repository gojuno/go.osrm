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
	ErrorCodeNoTrips        = "NoTrips"
	errorCodeOK             = "Ok" // "Ok" error code never returned to library client, thus not exported
)

// Invalid request errors
var (
	ErrorNotImplemented = errors.New("osrm5: the request is not implemented")
	ErrEmptyProfileName = errors.New("osrm5: the request should contain a profile name")
	ErrNoCoordinates    = errors.New("osrm5: the request should contain coordinates")
	ErrEmptyServiceName = errors.New("osrm5: the request should contain a service name")
)
