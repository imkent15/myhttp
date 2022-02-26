package interfaces

import "net/http"

type IHTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}
