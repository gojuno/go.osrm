package osrm

import (
	"context"
	"time"
)

const (
	defaultTimeout   = time.Second
	defaultServerURL = "http://127.0.0.1:5000"

	version = "v1"
)

// OSRM implements the common OSRM API v5.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md for details.
// TODO: implement (nearest, trip, tile) methods
type OSRM struct {
	client Client
}

// New creates a client with default server url
func New() *OSRM {
	return NewFromURL(defaultServerURL)
}

// NewFromURL creates a client with custom server url
func NewFromURL(serverURL string) *OSRM {
	return NewFromURLWithTimeout(serverURL, defaultTimeout)
}

// NewFromURLWithTimeout creates a client with custom timeout connection
func NewFromURLWithTimeout(serverURL string, timeout time.Duration) *OSRM {
	return NewWithClient(NewClientWithTimeout(serverURL, timeout))
}

// NewWithClient creates a client with custom transport layer
func NewWithClient(client Client) *OSRM {
	return &OSRM{client}
}

func (o OSRM) query(ctx context.Context, in *Request, out Response) Error {
	if err := o.client.Serve(ctx, in, out); err != nil {
		return WrapError(err)
	}
	if out.ErrCode() != OK {
		return ResponseError{out.ErrCode(), out.Error()}
	}
	return nil
}

// Route searches the shortest path between given coordinates.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#service-route for details.
func (o OSRM) Route(ctx context.Context, r RouteRequest) (*RouteResponse, Error) {
	resp := RouteResponse{}
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Table computes duration tables for the given locations.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#service-table for details.
func (o OSRM) Table(ctx context.Context, r TableRequest) (*TableResponse, Error) {
	resp := TableResponse{}
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Match matches given GPS points to the road network in the most plausible way.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#service-match for details.
func (o OSRM) Match(ctx context.Context, r MatchRequest) (*MatchResponse, Error) {
	resp := MatchResponse{}
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
