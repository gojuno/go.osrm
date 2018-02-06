package osrm

import (
	"context"
	"encoding/json"
)

// client makes a real query to OSRM server
type client struct {
	transport

	serverURL string
}

// newClient creates a client with server url and specific getter
func newClient(serverURL string, t transport) client {
	return client{transport: t, serverURL: serverURL}
}

// doRequest makes GET request to OSRM server and decodes the given JSON
func (c client) doRequest(ctx context.Context, in *request, out interface{}) error {
	url, err := in.URL(c.serverURL)
	if err != nil {
		return err
	}
	bytes, err := c.get(ctx, url)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}
