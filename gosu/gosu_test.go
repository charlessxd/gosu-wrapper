package gosu

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestSession_buildCall(t *testing.T) {
	s := NewSession("12345")

	v := url.Values{}
	v.Add(endpointAPIKey, s.key)
	v.Add(endpointParamBeatmapID, "696969")

	result := s.buildCall(endpointBeatmaps, v)
	expected := endpointAPI + endpointBeatmaps + "b=696969&k=12345"

	if result != expected {
		t.Fatal("Expected \"" + expected + "\" but got \"" + result + "\".")
	}
}

func Test_getMods(t *testing.T) {
	if strings.Join(getMods(781), "") != "[NoFail TouchDevice Hidden HalfTime]" {
		t.Fatal("Expected \"[NoFail TouchDevice Hidden HalfTime]\" but got \"" + strings.Join(getMods(781), "") + "\".")
	}
}

func ExampleSession_Fetch() {
	s := NewSession(os.Getenv("API_KEY"))

	call := UserCall{UserID: os.Getenv("USER_ID")}

	if u, e := s.FetchUser(call); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(u.Username)
	}
}
