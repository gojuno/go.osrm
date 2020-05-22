package osrm

import (
	"bytes"
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
		usePOST    bool
	}
)

// newClient creates a client with server url and specific getter
func newClient(httpClient HTTPClient, serverURL string, usePOST bool) client {
	return client{httpClient: httpClient, serverURL: serverURL, usePOST: usePOST}
}

// doRequest makes GET request to OSRM server and decodes the given JSON
func (c client) doRequest(ctx context.Context, in *request, out interface{}) error {
	path, err := in.URLPath()
	if err != nil {
		return err
	}

	resp, err := c.httpRequest(ctx, path)
	if err != nil {
		return err
	}
	defer closeSilently(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	// OSRM returns both codes 200 and 400 in a case with a body.
	// In other cases, it returns an unexpected error without a body.
	// http://project-osrm.org/docs/v5.5.1/api/#responses
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("unexpected http status code %d with body %q", resp.StatusCode, body)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("failed to unmarshal body %q: %v", body, err)
	}

	return nil
}

func (c client) httpRequest(ctx context.Context, path string) (*http.Response, error) {
	if c.usePOST {
		return c.post(ctx, path)
	}
	return c.get(ctx, path)
}

func (c client) get(ctx context.Context, path string) (*http.Response, error) {
	url := c.serverURL + "/" + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(req.WithContext(ctx))
}

func (c client) post(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.serverURL, bytes.NewReader([]byte(path)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-uri")
	return c.httpClient.Do(req.WithContext(ctx))
}

func closeSilently(c io.Closer) {
	// #nosec - make github.com/GoASTScanner/gas linter ignore this
	_ = c.Close() // nothing meaningful to do with this error - so ignore and suppress linter warnings
}
