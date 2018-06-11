package api

import "net/http"

// IController - Interface For An API Controller
type IController interface {
	HandleRequest(http.ResponseWriter, *http.Request) error
}
