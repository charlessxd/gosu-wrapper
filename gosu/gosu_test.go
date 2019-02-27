package gosu

import (
	"fmt"
	"net/url"
	"os"
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

func TestSession_Fetch(t *testing.T) {
	s := NewSession("12345")

	c := UserCall{}

	b := Beatmap{}

	if e := s.Fetch(c, b); e == nil {
		t.Fatal("Expected target mismatch error")
	}
}

func ExampleSession_Fetch() {
	s := NewSession(os.Getenv("API_KEY"))

	call := UserCall{UserID: os.Getenv("USER_ID")}

	user := User{}

	s.Fetch(call, user)

	fmt.Println(user.Username)
}
