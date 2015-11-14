package ecs

import(
	"net/http"
	"net/url"
	"sort"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"time"
	"strings"

	"github.com/pborman/uuid"
)

type IClient interface {
	Do(IRequest) IResponse
}

type Client struct {
	profile Profile
}

func NewClient(profile Profile) *Client{
	client := &Client{profile:profile}
	return client
}

func NewDefaultClient(accessKeyId, accessKeySecret, regionId string) *Client{
	profile := NewDefaultProfile(accessKeyId, accessKeySecret, regionId)
	client  := NewClient(*profile)
	return client
}

func (this *Client) Do(request IRequest, response IResponse) error {
	params := make(map[string] string, 35)

	makeParams(this.profile, request, params)

	makeSign(this.profile, params)

	url := makeUrl(this.profile, params)

	fmt.Printf("\n", url)

	httpResp, err := http.Get(url)
	if err != nil {
		parseHttpResponse(httpResp, response, params["Format"])
	}

	return err
}

func parseHttpResponse(httpResp *http.Response, ecsResp IResponse, format string) {
	ecsResp.setRaw(httpResp)
	contents, _ := ioutil.ReadAll(httpResp.Body)

	if strings.EqualFold(FORMAT_JSON, format) {
		json.Unmarshal(contents, ecsResp)
	} else if strings.EqualFold(FORMAT_XML, format) {
		xml.Unmarshal(contents, ecsResp)
	}

	//fmt.Printf("\nresp = %s\n", string(contents))
}

	// 1. make public params
	// 2. delete signture
func makeParams(profile Profile, request IRequest, params map[string] string){
	r_params := request.Params()

	params["Format"] = profile.GetFormat()
	params["Version"] = profile.GetVersion()
	params["AccessKeyId"] = profile.GetAccessKeyId()
	params["SignatureMethod"] = profile.GetSignatureMethod()
	params["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	params["SignatureVersion"] = profile.GetSignatureVersion()
	params["SignatureNonce"] = uuid.New()
	//params["ResourceOwnerAccount"] = profile.GetResourceOwnerAccount()

	for k, v := range r_params {
		params[k] = v
	}
}

	// 3. get sign
	// 3.1 sort params
	// 3.2 ...
	// 4. [signture] = sign
func makeSign(profile Profile, params map[string] string){
	delete(params, "Signature")

	leng := len(params)
	keys := make([]string, 0, leng)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer

	for i, k := range keys {
		buf.WriteString(percentEncode(k))
		buf.WriteByte('=')
		buf.WriteString(percentEncode(params[k]))
		if i < leng-1 {
			buf.WriteByte('&')
		}
	}

	HTTPMethod := "GET"
	secret := profile.GetAccessKeySecret()

	var signStr bytes.Buffer
	signStr.WriteString(HTTPMethod)
	signStr.WriteByte('&')
	signStr.WriteString(percentEncode("/"))
	signStr.WriteByte('&')
	signStr.WriteString(percentEncode(buf.String()))

	sign := signature(signStr.String(), secret+"&")
	params["Signature"] = sign
}

func percentEncode(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}

	if hexCount == 0 {
		return s
	}

	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case shouldEscape(c):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

func shouldEscape(c byte) bool {
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	switch c {
	case '-', '_', '.', '~': 
		return false
	}
	
	return true
}

func signature(source string, secret string) (sign string){
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(source))
	sign = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return sign
}

	// 5. make url
func makeUrl(profile Profile, params map[string] string) string{
	if params == nil {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(profile.GetHttpMethod())
	buf.WriteString("://ecs-cn-hangzhou.aliyuncs.com/?")

	i := 0
	for k, v := range params {
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
		if i < len(params)-1 {
			buf.WriteByte('&')
		}
		i++
	}

	return buf.String()
}
