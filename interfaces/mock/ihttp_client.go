package mock

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type MockIHTTPClient struct {
}

var (
	noErrorBodyWithCloser = ioutil.NopCloser(bytes.NewBuffer([]byte("body response from noerror.com")))
	noErrorResponse       = &http.Response{
		StatusCode: http.StatusOK,
		Body:       noErrorBodyWithCloser,
	}

	errorResponse = &http.Response{
		StatusCode: http.StatusBadRequest,
	}
)

func (m *MockIHTTPClient) Do(request *http.Request) (*http.Response, error) {
	switch request.URL.String() {
	case "http://noerror.com":
		return noErrorResponse, nil
	case "http://error.com":
		return errorResponse, nil
	default:
		return nil, errors.New("address not supported by mock client")
	}
}
