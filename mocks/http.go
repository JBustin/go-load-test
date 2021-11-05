package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HTTPClientMock struct {
	// DoFunc will be executed whenever Do function is executed
	// so we'll be able to create a custom response
	DoFunc func(*http.Request) (*http.Response, error)
}

func (H HTTPClientMock) Do(r *http.Request) (*http.Response, error) {
	return H.DoFunc(r)
}

func HTTPresponse(statusCode int, response string, err error) HTTPClientMock {
	body := ioutil.NopCloser(bytes.NewReader([]byte(response)))

	return HTTPClientMock{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{Body: body, StatusCode: statusCode}, err
		},
	}
}
