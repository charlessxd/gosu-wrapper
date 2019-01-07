package gosu

import (
	"net/url"
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
