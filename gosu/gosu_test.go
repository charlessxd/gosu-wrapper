package gosu

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
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

func ExampleSession_Emit() {
	s := NewSession(os.Getenv("API_KEY"))

	c := UserCall{
		UserID: os.Getenv("USER_ID"),
	}

	u, err := s.FetchUser(c)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

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
				s.Emit(u.UserID, strconv.FormatFloat(t.PPRaw-u.PPRaw, 'G', -1, 64))
				u, _ = s.FetchUser(c)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
