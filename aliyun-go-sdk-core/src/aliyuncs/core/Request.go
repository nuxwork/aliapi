package core

import (
	"encoding/json"
	"strconv"
)

type IRequest interface {
	Set(key string, param string)
	Get(key string) (param string)
	Add(params map[string]string)
	Remove(key string)
	Reset()
	Params() (params map[string]string)
	Product() string
	Version() string
	Action() string
}

type Request struct {
	params map[string]string
	product string
}


func (this *Request) Init(product, version, action string) {
	this.product = product
	this.Set("Version", version)
	this.Set("Action", action)
}

func (this *Request) Product() string {
	return this.product
}

func (this *Request) Version() string {
	return this.Get("Version")
}

func (this *Request) Action() string {
	return this.Get("Action")
}

func (this *Request) check() {
	if this.params == nil {
		this.params = make(map[string]string, 35)
	}
}

func (this *Request) Add(params map[string]string) {
	this.check()
	if params != nil {
		for k, v := range params {
			this.params[k] = v
		}
	}
}

func (this *Request) Set(key string, value string) {
	this.check()
	this.params[key] = value
}

func (this *Request) SetInt(key string, value int) {
	this.Set(key, strconv.Itoa(value))
}

func (this *Request) SetBool(key string, value bool) {
	this.Set(key, strconv.FormatBool(value))
}

func (this *Request) SetStringArray(key string, value []string) {
	s, _ := json.Marshal(value)
	this.Set(key, string(s))
}

func (this *Request) Get(key string) string {
	this.check()
	return this.params[key]
}

func (this *Request) GetInt(key string) int {
	s := this.Get(key)
	v, _ := strconv.Atoi(s)
	return v
}

func (this *Request) GetBool(key string) bool {
	s := this.Get(key)
	v, _ := strconv.ParseBool(s)
	return v
}

func (this *Request) GetStringArray(key string) []string {
	s := this.Get(key)
	var v []string
	json.Unmarshal([]byte(s), &v)
	return v
}

func (this *Request) Remove(key string) {
	this.check()
	delete(this.params, key)
}

func (this *Request) Reset() {
	this.check()
	for k, _ := range this.params {
		delete(this.params, k)
	}
}

func (this *Request) Params() (params map[string]string) {
	this.check()
	return this.params
}
