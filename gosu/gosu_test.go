package gosu

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
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

func ExampleSession_Emit() {
	s := NewSession(os.Getenv("API_KEY"))

	c := UserCall{
		UserID: os.Getenv("USER_ID"),
	}

	u, _ := s.FetchUser(c)

	event := make(chan string)

	s.AddListener(u.UserID, event)

	// Outputs the event
	go func() {
		for {
			fmt.Println(<-event)
		}
	}()

	// Event for when a user's PP changes
	go func() {
		init := u.PPRaw
		for init == u.PPRaw {
			if t, _ := s.FetchUser(c); t.PPRaw != u.PPRaw {
				s.Emit(u.UserID, "PP Changed")
				u, _ = s.FetchUser(c)
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
