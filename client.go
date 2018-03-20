package osrm

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	// HTTPClient defines minimal interface necessary for making HTTP requests.
	// Standard library http.Client{} implements this interface.
	// A non-2xx status code doesn't cause an error.
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}

	// client makes a real query to OSRM server
	client struct {
		httpClient HTTPClient
		serverURL  string
	}
)

// newClient creates a client with server url and specific getter
func newClient(serverURL string, c HTTPClient) client {
	return client{c, serverURL}
}

// doRequest makes GET request to OSRM server and decodes the given JSON
func (c client) doRequest(ctx context.Context, in *request, out interface{}) error {
	url, err := in.URL(c.serverURL)
	if err != nil {
		return err
	}

	resp, err := c.get(ctx, url)
	if err != nil {
		return err
	}
	defer closeSilently(resp.Body)

	// OSRM returns HTTP status code 200 in case of success or HTTP status code 400 in case of error
	// In other cases it has to return error
	// http://project-osrm.org/docs/v5.5.1/api/#requests
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("unexpected http status code %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("failed to unmarshal body %q: %v", bytes, err)
	}

	return nil
}

// get implements Transport interface
func (c client) get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func closeSilently(c io.Closer) {
	// #nosec - make github.com/GoASTScanner/gas linter ignore this
	_ = c.Close() // nothing meaningful to do with this error - so ignore and suppress linter warnings
}