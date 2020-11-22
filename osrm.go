package osrm

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultTimeout   = time.Second
	defaultServerURL = "http://127.0.0.1:5000"

	version = "v1"
)

// OSRM implements the common OSRM API v5.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md for details.
// TODO: implement (trip, tile) methods
type OSRM struct {
	client
}

// Config represents OSRM client configuration options
type Config struct {
	// ServerURL is OSRM server URL to be used for queries.
	// Local http://127.0.0.1:5000 URL will be used as default if not set.
	ServerURL string
	// Client is custom pre-configured http client to be used for queries.
	// New http.Client instance with default settings and one second timeout will be used if not set.
	Client HTTPClient
}

// ResponseStatus represent OSRM API response
type ResponseStatus struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	DataVersion string `json:"data_version"`
}

// ErrCode returns error code from OSRM response
func (r ResponseStatus) ErrCode() string {
	return r.Code
}

func (r ResponseStatus) Error() string {
	return r.Code + " - " + r.Message
}

func (r ResponseStatus) apiError() error {
	if r.Code != errorCodeOK {
		return r
	}
	return nil
}

type response interface {
	apiError() error
}

// New creates a client with default server url and default timeout
func New() *OSRM {
	return NewWithConfig(Config{})
}

// NewFromURL creates a client with custom server url and default timeout
func NewFromURL(serverURL string) *OSRM {
	return NewWithConfig(Config{ServerURL: serverURL})
}

// NewFromURLWithTimeout creates a client with custom timeout connection
func NewFromURLWithTimeout(serverURL string, timeout time.Duration) *OSRM {
	return NewWithConfig(Config{
		ServerURL: serverURL,
		Client:    &http.Client{Timeout: timeout},
	})
}

// NewWithConfig creates a client with given config
func NewWithConfig(cfg Config) *OSRM {
	if cfg.ServerURL == "" {
		cfg.ServerURL = defaultServerURL
	}
	if cfg.Client == nil {
		cfg.Client = &http.Client{Timeout: defaultTimeout}
	}

	return &OSRM{client: newClient(cfg.ServerURL, cfg.Client)}
}

func (o OSRM) query(ctx context.Context, in *request, out response) error {
	if err := o.client.doRequest(ctx, in, out); err != nil {
		return err
	}
	return out.apiError()
}

// Route searches the shortest path between given coordinates.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#route-service for details.
func (o OSRM) Route(ctx context.Context, r RouteRequest) (*RouteResponse, error) {
	var resp RouteResponse
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Table computes duration tables for the given locations.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#table-service for details.
func (o OSRM) Table(ctx context.Context, r TableRequest) (*TableResponse, error) {
	var resp TableResponse
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Match matches given GPS points to the road network in the most plausible way.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#match-service for details.
func (o OSRM) Match(ctx context.Context, r MatchRequest) (*MatchResponse, error) {
	var resp MatchResponse
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Nearest matches given GPS point to the nearest road network.
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#nearest-service for details.
func (o OSRM) Nearest(ctx context.Context, r NearestRequest) (*NearestResponse, error) {
	var resp NearestResponse
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// The trip plugin solves the Traveling Salesman Problem using a greedy heuristic (farthest-insertion algorithm).
// See https://github.com/Project-OSRM/osrm-backend/blob/master/docs/http.md#trip-service for details
func (o OSRM) Trip(ctx context.Context, r TripRequest) (*TripResponse, error) {
	var resp TripResponse
	if !r.IsSupported() {
		return nil, ErrorNotImplemented
	}
	if err := o.query(ctx, r.request(), &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
