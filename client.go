package osrm

import (
	"context"
	"encoding/json"
	"time"
)

// Client makes a real query to OSRM server
type Client struct {
	Transport

	serverURL string
}

// NewClient creates a client with server url and specific getter
func NewClient(serverURL string, t Transport) Client {
	return Client{Transport: t, serverURL: serverURL}
}

// NewClientWithTimeout creates a client with http.Client
func NewClientWithTimeout(serverURL string, timeout time.Duration) Client {
	return NewClient(serverURL, NewDefaultTransport(timeout))
}

// Serve makes GET request to OSRM server and decodes the given JSON
func (c Client) Serve(ctx context.Context, in *Request, out Response) error {
	url, err := in.URL(c.serverURL)
	if err != nil {
		return err
	}
	bytes, err := c.Get(ctx, url)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}
