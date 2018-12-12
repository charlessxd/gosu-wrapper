package gosu

import (
	"fmt"
	"net/url"
	"os"
	"testing"
)

func TestSession_BuildCall(t *testing.T) {
	s := NewSession("12345")

	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)
	v.Add(EndpointParamBeatmapID, "696969")

	result := s.buildCall(EndpointBeatmaps, v)
	expected := EndpointAPI + EndpointBeatmaps + "b=696969&k=12345"

	if result != expected {
		t.Fatal("Expected \"" + expected + "\" but got \"" + result + "\".")
	}
}

func ExampleNewSession() {
	s := NewSession(os.Getenv("API_KEY"))

	fmt.Println(s)
}
