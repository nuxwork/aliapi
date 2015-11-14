package ecs

// API Version
const(
	API_V20140526 string = "2014-05-26"
)

const(
	HTTP string = "http"
	HTTPS string = "https"
)

// Format of return value
const(
	FORMAT_XML string = "xml"
	FORMAT_JSON string = "json"
)

// Signature method
const(
	SIGN_HMAC_SHA1 string = "HMAC-SHA1"
)

// Signature Version
const(
	SIGN_V_1_0 string = "1.0"
)

type IProfile interface {
	Set(key string, param string)
	Get(key string) (param string)
	Add(params map[string] string)
	Remove(key string)
	Reset()
	Params()(params map[string] string)
}

type Profile struct {
	params map[string] string
}

func NewProfile() *Profile {
	p := new(Profile)
	initProfile(p, "", "", "")
	return p
}

func NewDefaultProfile(accessKeyId, accessKeySecret, regionId string) *Profile {
	p := new(Profile)
	initProfile(p, accessKeyId, accessKeySecret, regionId)
	return p
}

func initProfile(profile *Profile, accessKeyId, accessKeySecret, regionId string){
	profile.SetAccessKeyId(accessKeyId)
	profile.SetAccessKeySecret(accessKeySecret)
	profile.SetRegionId(regionId)
	profile.SetHttpMethod(HTTPS)
	profile.SetFormat(FORMAT_JSON)
	profile.SetVersion(API_V20140526)
	profile.SetSignatureMethod(SIGN_HMAC_SHA1)
	profile.SetSignatureVersion(SIGN_V_1_0)
}

func (this *Profile) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *Profile) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *Profile) SetHttpMethod(v string){
	this.Set("HttpMethod", v)
}

func (this *Profile) GetHttpMethod() string {
	return this.Get("HttpMethod")
}

func (this *Profile) SetFormat(v string){
	this.Set("Format", v)
}

func (this *Profile) GetFormat() string {
	return this.Get("Format")
}

func (this *Profile) SetVersion(v string){
	this.Set("Version", v)
}

func (this *Profile) GetVersion() string {
	return this.Get("Version")
}

func (this *Profile) SetAccessKeyId(v string){
	this.Set("AccessKeyId", v)
}

func (this *Profile) GetAccessKeyId() string {
	return this.Get("AccessKeyId")
}

func (this *Profile) SetAccessKeySecret(v string){
	this.Set("AccessKeySecret", v)
}

func (this *Profile) GetAccessKeySecret() string {
	return this.Get("AccessKeySecret")
}

func (this *Profile) SetSignatureMethod(v string){
	this.Set("SignatureMethod", v)
}

func (this *Profile) GetSignatureMethod() string {
	return this.Get("SignatureMethod")
}

func (this *Profile) SetSignatureVersion(v string){
	this.Set("SignatureVersion", v)
}

func (this *Profile) GetSignatureVersion() string {
	return this.Get("SignatureVersion")
}

func (this *Profile) SetResourceOwnerAccount(v string){
	this.Set("ResourceOwnerAccount", v)
}

func (this *Profile) GetResourceOwnerAccount() string {
	return this.Get("ResourceOwnerAccount")
}

func (this *Profile) check(){
	if this.params == nil {
		this.params = make(map[string] string, 35)
	}
}

func (this *Profile) Add(params map[string] string){
	this.check()
	if params != nil {
		for k, v := range params {
			this.params[k] = v
		}
	}
}

func (this *Profile) Set(key string, value string){
	this.check()
	this.params[key] = value
}

func (this *Profile) Get(key string) (value string){
	this.check()
	return this.params[key]
}

func (this *Profile) Remove(key string){
	this.check()
	delete(this.params, key)
}

func (this *Profile) Reset(){
	this.check()
	for k, _ := range this.params {
		delete(this.params, k)
	}
}

func (this *Profile) Params()(params map[string] string){
	this.check()
	return this.params
}