package ecs

import (
	"strconv"
	"encoding/json"
)

type IRequest interface {
	Set(key string, param string)
	Get(key string) (param string)
	Add(params map[string] string)
	Remove(key string)
	Reset()
	Params()(params map[string] string)
	Version() string
}

type Request struct {
	params map[string] string
}

func (this *Request) GetAction() string {
	return this.Get("Action")
}

/*
func (this *Request) SetFormat(v string){
	this.Set("Format", v)
}

func (this *Request) GetFormat() string {
	return this.Get("Format")
}

func (this *Request) SetVersion(v string){
	this.Set("Version", v)
}

func (this *Request) GetVersion() string {
	return this.Get("Version")
}

func (this *Request) SetAccessKeyId(v string){
	this.Set("AccessKeyId", v)
}

func (this *Request) GetAccessKeyId() string {
	return this.Get("AccessKeyId")
}

func (this *Request) SetSignatureMethod(v string){
	this.Set("SignatureMethod", v)
}

func (this *Request) GetSignatureMethod() string {
	return this.Get("SignatureMethod")
}

func (this *Request) SetTimestamp(v string){
	this.Set("Timestamp", v)
}

func (this *Request) GetTimestamp() string {
	return this.Get("Timestamp")
}

func (this *Request) SetSignatureVersion(v string){
	this.Set("SignatureVersion", v)
}

func (this *Request) GetSignatureVersion() string {
	return this.Get("SignatureVersion")
}

func (this *Request) SetSignatureNonce(v string){
	this.Set("SignatureNonce", v)
}

func (this *Request) GetSignatureNonce() string {
	return this.Get("SignatureNonce")
}

func (this *Request) SetResourceOwnerAccount(v string){
	this.Set("ResourceOwnerAccount", v)
}

func (this *Request) GetResourceOwnerAccount() string {
	return this.Get("ResourceOwnerAccount")
}
*/
func (this *Request) check(){
	if this.params == nil {
		this.params = make(map[string] string, 35)
	}
}

func (this *Request) Add(params map[string] string){
	this.check()
	if params != nil {
		for k, v := range params {
			this.params[k] = v
		}
	}
}

func (this *Request) Set(key string, value string){
	this.check()
	this.params[key] = value
}

func (this *Request) SetInt(key string, value int){
	this.Set(key, strconv.Itoa(value))
}

func (this *Request) SetBool(key string, value bool){
	this.Set(key, strconv.FormatBool(value))
}

func (this *Request) SetStringArray(key string, value []string){
	s, _ := json.Marshal(value)
	this.Set(key, string(s))
}

func (this *Request) Get(key string) string{
	this.check()
	return this.params[key]
}

func (this *Request) GetInt(key string) int{
	s := this.Get(key)
	v, _ := strconv.Atoi(s)
	return v
}

func (this *Request) GetBool(key string) bool{
	s := this.Get(key)
	v, _ := strconv.ParseBool(s)
	return v
}

func (this *Request) GetStringArray(key string) []string{
	s := this.Get(key)
	var v []string
	json.Unmarshal([]byte(s), &v)
	return v
}

func (this *Request) Remove(key string){
	this.check()
	delete(this.params, key)
}

func (this *Request) Reset(){
	this.check()
	for k, _ := range this.params {
		delete(this.params, k)
	}
}

func (this *Request) Params()(params map[string] string){
	this.check()
	return this.params
}

func (this *Request) Version() string {
	return API_V20140526
}