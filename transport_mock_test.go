package osrm

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "HTTPClient" can be found in github.com/gojuno/go.osrm/http_client
*/
import (
	http "net/http"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//HTTPClientMock implements github.com/gojuno/go.osrm/http_client.HTTPClient
type HTTPClientMock struct {
	t minimock.Tester

	DoFunc       func(p *http.Request) (r *http.Response, r1 error)
	DoCounter    uint64
	DoPreCounter uint64
	DoMock       mHTTPClientMockDo
}

//NewHTTPClientMock returns a mock for github.com/gojuno/go.osrm/http_client.HTTPClient
func NewHTTPClientMock(t minimock.Tester) *HTTPClientMock {
	m := &HTTPClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DoMock = mHTTPClientMockDo{mock: m}

	return m
}

type mHTTPClientMockDo struct {
	mock             *HTTPClientMock
	mockExpectations *HTTPClientMockDoParams
}

//HTTPClientMockDoParams represents input parameters of the HTTPClient.Do
type HTTPClientMockDoParams struct {
	p *http.Request
}

//Expect sets up expected params for the HTTPClient.Do
func (m *mHTTPClientMockDo) Expect(p *http.Request) *mHTTPClientMockDo {
	m.mockExpectations = &HTTPClientMockDoParams{p}
	return m
}

//Return sets up a mock for HTTPClient.Do to return Return's arguments
func (m *mHTTPClientMockDo) Return(r *http.Response, r1 error) *HTTPClientMock {
	m.mock.DoFunc = func(p *http.Request) (*http.Response, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of HTTPClient.Do method
func (m *mHTTPClientMockDo) Set(f func(p *http.Request) (r *http.Response, r1 error)) *HTTPClientMock {
	m.mock.DoFunc = f
	return m.mock
}

//Do implements github.com/gojuno/go.osrm/http_client.HTTPClient interface
func (m *HTTPClientMock) Do(p *http.Request) (r *http.Response, r1 error) {
	atomic.AddUint64(&m.DoPreCounter, 1)
	defer atomic.AddUint64(&m.DoCounter, 1)

	if m.DoMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.DoMock.mockExpectations, HTTPClientMockDoParams{p},
			"HTTPClient.Do got unexpected parameters")

		if m.DoFunc == nil {

			m.t.Fatal("No results are set for the HTTPClientMock.Do")

			return
		}
	}

	if m.DoFunc == nil {
		m.t.Fatal("Unexpected call to HTTPClientMock.Do")
		return
	}

	return m.DoFunc(p)
}

//DoMinimockCounter returns a count of HTTPClientMock.DoFunc invocations
func (m *HTTPClientMock) DoMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DoCounter)
}

//DoMinimockPreCounter returns the value of HTTPClientMock.Do invocations
func (m *HTTPClientMock) DoMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.DoPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *HTTPClientMock) ValidateCallCounters() {

	if m.DoFunc != nil && atomic.LoadUint64(&m.DoCounter) == 0 {
		m.t.Fatal("Expected call to HTTPClientMock.Do")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *HTTPClientMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *HTTPClientMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *HTTPClientMock) MinimockFinish() {

	if m.DoFunc != nil && atomic.LoadUint64(&m.DoCounter) == 0 {
		m.t.Fatal("Expected call to HTTPClientMock.Do")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *HTTPClientMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *HTTPClientMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.DoFunc == nil || atomic.LoadUint64(&m.DoCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.DoFunc != nil && atomic.LoadUint64(&m.DoCounter) == 0 {
				m.t.Error("Expected call to HTTPClientMock.Do")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *HTTPClientMock) AllMocksCalled() bool {

	if m.DoFunc != nil && atomic.LoadUint64(&m.DoCounter) == 0 {
		return false
	}

	return true
}
