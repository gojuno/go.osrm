package osrm

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

// Transport makes GET request
type Transport interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type DefaultTransport struct {
	http.Client
}

func NewDefaultTransport(timeout time.Duration) *DefaultTransport {
	return &DefaultTransport{http.Client{Timeout: timeout}}
}

func (t DefaultTransport) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := t.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
