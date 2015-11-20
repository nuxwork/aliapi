package core

import (
	"net/http"
)

type IResponse interface {
	OK() bool
	setHttpResponse(*http.Response)
}

type Response struct {
	HttpResponse       *http.Response
	RequestId string
	HostId    string
	Code      string
	Message   string
}

func (this *Response) OK() bool {
	return this.HttpResponse != nil && this.HttpResponse.StatusCode >= 200 && this.HttpResponse.StatusCode < 300
}

func (this *Response) setHttpResponse(r *http.Response) {
	this.HttpResponse = r
}
