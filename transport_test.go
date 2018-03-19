package osrm

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockReadCloser struct{}

func (m mockReadCloser) Close() error               { return nil }
func (m mockReadCloser) Read(_ []byte) (int, error) { return 0, nil }

func Test_getWithBadHTTPCode(t *testing.T) {
	httpClientMock := NewHTTPClientMock(t)
	defer httpClientMock.MinimockFinish()

	httpClientMock.DoMock.Return(&http.Response{
		StatusCode: 500,
		Body:       mockReadCloser{},
	}, nil)

	defaultTransport := defaultTransport{httpClient: httpClientMock}
	b, err := defaultTransport.get(context.Background(), "/")
	require.Nil(t, b)
	require.EqualError(t, err, "http code 500 is not OK")
}
