package gosu

import (
	"net/url"
	"testing"
)

func TestSession_BuildCall(t *testing.T) {
	s := NewSession("12345")

	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)
	v.Add(EndpointBeatmapsBeatmapID, "696969")

	result := s.BuildCall(EndpointBeatmaps, v)
	expected := EndpointAPI + EndpointBeatmaps + "b=696969&k=12345"

	if result != expected {
		t.Fatal("Expected \"" + expected + "\" but got \"" + result + "\".")
	}
}
