package osrm

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockReadCloser struct{ readError bool }

func (m mockReadCloser) Close() error { return nil }
func (m mockReadCloser) Read(p []byte) (int, error) {
	if m.readError {
		return 0, io.ErrUnexpectedEOF
	}
	return 0, io.EOF
}

func Test_getWithError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprint(w, "something wrong")
	}))
	defer ts.Close()

	c := newClient("/", ts.Client())
	b, err := c.get(context.Background(), "/")
	require.Nil(t, b)
	require.NotNil(t, err)
}

func Test_doRequestWithBadHTTPCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprint(w, "<html><head>")
	}))
	defer ts.Close()

	c := newClient(ts.URL, ts.Client())
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
	}
	err := c.doRequest(context.Background(), &req, nil)
	require.EqualError(t, err, "unexpected http status code 500 with body \"<html><head>\"")
}

func Test_doRequestWithBodyUnmarshalFailure(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	c := newClient(ts.URL, ts.Client())
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
	}
	err := c.doRequest(context.Background(), &req, nil)
	require.EqualError(t, err, "failed to unmarshal body \"\": unexpected end of JSON input")
}
