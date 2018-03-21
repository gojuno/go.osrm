package osrm

import (
	"context"
	"errors"
	"io"
	"net/http"
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
	httpClientMock := NewHTTPClientMock(t)
	defer httpClientMock.MinimockFinish()

	httpClientMock.DoMock.Return(nil, errors.New("something wrong"))

	c := newClient("/", httpClientMock)
	b, err := c.get(context.Background(), "/")
	require.Nil(t, b)
	require.NotNil(t, err)
}

func Test_doRequestWithBadHTTPCode(t *testing.T) {
	httpClientMock := NewHTTPClientMock(t)
	defer httpClientMock.MinimockFinish()

	httpClientMock.DoMock.Return(&http.Response{
		StatusCode: 500,
		Body:       mockReadCloser{},
	}, nil)

	c := newClient("/", httpClientMock)
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
	}
	err := c.doRequest(context.Background(), &req, nil)
	require.EqualError(t, err, "unexpected http status code 500 with body \"\"")
}

func Test_doRequestWithBodyReadingFailure(t *testing.T) {
	httpClientMock := NewHTTPClientMock(t)
	defer httpClientMock.MinimockFinish()

	httpClientMock.DoMock.Return(&http.Response{
		StatusCode: 200,
		Body:       mockReadCloser{readError: true},
	}, nil)

	c := newClient("/", httpClientMock)
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
	}
	err := c.doRequest(context.Background(), &req, nil)
	require.EqualError(t, err, "failed to read body: unexpected EOF")
}

func Test_doRequestWithBodyUnmarshalFailure(t *testing.T) {
	httpClientMock := NewHTTPClientMock(t)
	defer httpClientMock.MinimockFinish()

	httpClientMock.DoMock.Return(&http.Response{
		StatusCode: 200,
		Body:       mockReadCloser{},
	}, nil)

	c := newClient("/", httpClientMock)
	req := request{
		profile: "something",
		geoPath: geoPath,
		service: "foobar",
	}
	err := c.doRequest(context.Background(), &req, nil)
	require.EqualError(t, err, "failed to unmarshal body \"\": unexpected end of JSON input")
}
