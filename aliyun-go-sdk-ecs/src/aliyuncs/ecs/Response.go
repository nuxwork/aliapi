package ecs

import(
	"net/http"
)

type IResponse interface {
	OK() bool
	setRaw(*http.Response)
}

type Response struct {
	Raw *http.Response
	RequestId string
	HostId string
	Code string
	Message string
}

func (this *Response) OK() bool {
	return this.Raw != nil && this.Raw.StatusCode >= 200 && this.Raw.StatusCode < 300
}

func (this *Response) setRaw(raw *http.Response) {
	this.Raw = raw
}