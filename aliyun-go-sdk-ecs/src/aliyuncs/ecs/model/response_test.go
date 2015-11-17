package model

import (
	"encoding/json"
	"encoding/xml"
	"bytes"
	"testing"
	"aliyuncs/ecs"
)

func JsonTest(data string, resp ecs.IResponse, t *testing.T) {
	err1 := json.Unmarshal([]byte(data), resp)
	if err1 != nil {
		t.Error("error: ", err1)
	}

	resp_json, err2 := json.Marshal(resp)
	if err2 != nil {
		t.Error("error: ", err2)
	}

	ds := escape(data)
	rs := escape(string(resp_json))

	result := compare(rs, ds)

	if !result {
		t.Errorf("error: %s\n%s\n%s\n", "not equal", rs, ds)
	}
}

func XmlTest(data string, resp ecs.IResponse, t *testing.T) {
	err1 := xml.Unmarshal([]byte(data), resp)
	if err1 != nil {
		t.Error("error: ", err1)
	}

	resp_json, err2 := xml.Marshal(resp)
	if err2 != nil {
		t.Error("error: ", err2)
	}

	ds := escape(data)
	rs := escape(string(resp_json))

	result := compare(rs, ds)

	if !result {
		t.Errorf("error: %s\n%s\n%s\n", "not equal", rs, ds)
	}
}

func escape(s string) string {
	var buf bytes.Buffer

	for _, r := range s {
		if !shouldEscape(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func shouldEscape(r rune) bool {
	if r >=0 && r <= 32 {
		return true
	}

	return false
}

func compare(a, b string) bool {
	ca := 0
	for c := range a {
		ca += c
	}

	cb := 0
	for c := range b {
		cb += c
	}

	return ca == cb
}