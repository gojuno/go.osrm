package osrm

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Transport makes GET request
type Transport interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

// DefaultTransport is default transport implementation based on http.Client
type DefaultTransport struct {
	http.Client
}

// NewDefaultTransport creates new default transport based on http.Client with given timeout
func NewDefaultTransport(timeout time.Duration) *DefaultTransport {
	return &DefaultTransport{http.Client{Timeout: timeout}}
}

// Get implements Transport interface
func (t DefaultTransport) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := t.Client.Do(req.WithContext(ctx))
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
