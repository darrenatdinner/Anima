package server

import "net/http"
import "fmt"

// Instance - Single Running Instance Of An HTTP Server
type Instance struct {
	serverChannel chan *http.Request
}

func (server Instance) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.serverChannel <- request
	fmt.Fprint(writer, "notification sent")
}
