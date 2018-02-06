package osrm

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

// HTTPClient defines minimal interface necessary for making HTTP requests.
// Standard library http.Client{} implements this interface.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// transport makes GET request
type transport interface {
	get(ctx context.Context, url string) ([]byte, error)
}

// defaultTransport is default transport implementation based on http.Client
type defaultTransport struct {
	httpClient HTTPClient
}

// get implements Transport interface
func (t defaultTransport) get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer closeSilently(resp.Body)

	return ioutil.ReadAll(resp.Body)
}

func closeSilently(c io.Closer) {
	// #nosec - make github.com/GoASTScanner/gas linter ignore this
	_ = c.Close() // nothing meaningful to do with this error - so ignore and suppress linter warnings
}
