package ecs

import(
	"net/http"
)

type IResponse interface {
	OK() bool
	GetHttpResponse() http.Response
	GetRequestId() string
	GetHostId() string
	GetCode() string
	GetMessage() string

	// Set(key string, param string)
	Get(key string) (param string)
	// Add(params map[string] string)
	// Remove(key string)
	// Reset()
	Params()(params map[string] string)
}

type Response struct {
	Raw *http.Response
	params map[string] string
}

func (this *Response) OK() bool {
	return this.Raw.StatusCode >= 200 && this.Raw.StatusCode < 300
}

func (this *Response) GetRequestId() string {
	return this.Get("RequestId")
}

func (this *Response) GetHostId() string{
	return this.Get("HostId")
}

func (this *Response) GetCode() string {
	return this.Get("Code")
}

func (this *Response) GetMessage() string{
	return this.Get("Message")
}

func (this *Response) GetHttpResponse() *http.Response{
	return this.Raw
}

func (this *Response) check(){
	if this.params == nil {
		this.params = make(map[string] string, 35)
	}
}
/*
func (this *Response) Add(params map[string] string){
	this.check()
	if params != nil {
		for k, v := range params {
			this.params[k] = v
		}
	}
}

func (this *Response) Set(key string, value string){
	this.check()
	this.params[key] = value
}
*/
func (this *Response) Get(key string) (value string){
	this.check()
	return this.params[key]
}
/*
func (this *Response) Remove(key string){
	this.check()
	delete(this.params, key)
}

func (this *Response) Reset(){
	this.check()
	for k, _ := range this.params {
		delete(this.params, k)
	}
}*/

func (this *Response) Params()(params map[string] string){
	this.check()
	return this.params
}