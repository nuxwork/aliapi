package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
	"fmt"
	"github.com/pborman/uuid"
)

type IClient interface {
	Do(IRequest) IResponse
}

type Client struct {
	profile *Profile
	autoRetry bool
	maxRetryNumber int
}

func NewClient(profile *Profile) *Client {
	client := &Client{profile: profile, autoRetry: true, maxRetryNumber: 2}
	return client
}

func NewDefaultClient(accessKeyId, accessKeySecret, regionId string) *Client {
	profile := NewProfile(accessKeyId, accessKeySecret, regionId)
	client := NewClient(profile)
	return client
}

func (this *Client) SetAutoRetry(autoRetry bool) {
	this.autoRetry = autoRetry
}

func (this *Client) IsAutoRetry() bool {
	return this.autoRetry
}

func (this *Client) SetMaxRetryNumber(maxRetryNumber int) {
	this.maxRetryNumber = maxRetryNumber
}

func (this *Client) GetMaxRetryNumber() int {
	return this.maxRetryNumber
}

func (this *Client) Do(request IRequest, response IResponse) error {
	params := make(map[string]string, 35)

	makeParams(this.profile, request, params)

	makeSign(this.profile, params)

	url := makeUrl(this.profile, params)

	fmt.Printf("\n%s\n", url)

	httpResp, err := http.Get(url)
	if err == nil {
		parseHttpResponse(httpResp, response, params["Format"])
	}else if this.autoRetry {
		for i:=0; i<this.maxRetryNumber; i++ {
			httpResp, err = http.Get(url)
			if err == nil {
				parseHttpResponse(httpResp, response, params["Format"])
				break
			}
		}
	}
	// var err error

	return err
}

func parseHttpResponse(httpResp *http.Response, resp IResponse, format string) {
	resp.setHttpResponse(httpResp)
	contents, _ := ioutil.ReadAll(httpResp.Body)

	if strings.EqualFold(FORMAT_JSON, format) {
		json.Unmarshal(contents, resp)
	} else if strings.EqualFold(FORMAT_XML, format) {
		xml.Unmarshal(contents, resp)
	}
}

func makeParams(profile *Profile, request IRequest, params map[string]string) {
	r_params := request.Params()

	params["Format"] = profile.GetFormat()
	params["AccessKeyId"] = profile.GetAccessKeyId()
	params["SignatureMethod"] = profile.GetSignatureMethod()
	params["SignatureVersion"] = profile.GetSignatureVersion()
	params["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	params["SignatureNonce"] = uuid.New()
	//params["ResourceOwnerAccount"] = profile.GetResourceOwnerAccount()

	for k, v := range r_params {
		params[k] = v
	}
}

func makeSign(profile *Profile, params map[string]string) {
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

func signature(source string, secret string) (sign string) {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(source))
	sign = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return sign
}

// 5. make url
func makeUrl(profile *Profile, params map[string]string) string {
	if params == nil {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(profile.GetProtocol())
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
