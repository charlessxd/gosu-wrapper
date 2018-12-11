package gosu

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"testing"
)

func TestSession_BuildCall(t *testing.T) {
	s := NewSession("12345")

	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)
	v.Add(EndpointBeatmapsBeatmapID, "696969")

	result := s.buildCall(EndpointBeatmaps, v)
	expected := EndpointAPI + EndpointBeatmaps + "b=696969&k=12345"

	if result != expected {
		t.Fatal("Expected \"" + expected + "\" but got \"" + result + "\".")
	}
}

func ExampleNewSession() {
	s := NewSession(os.Getenv("API-KEY"))

	fmt.Println(json.MarshalIndent(s, "", "\t"))
	// Output:
	// {
	//		Key = "API-KEY",
	//		Limiter = {
	// 			MaxRequests:     100,
	//			CurrentRequests: 0,
	//			CanRequest:      true,
	//			FirstRequest:    time.Now(),
	//			TimeInterval:    60.0,
	//		}
	//	}
}
