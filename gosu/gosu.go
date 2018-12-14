// Package gosu provides a method of accessing osu-api in Go programs.
package gosu

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Session holds the API key and rate limiter.
type Session struct {
	// Osu API Key
	key string

	// Rate limit
	limiter RateLimit

	// Listeners
	//listeners map[string][]chan string
	listeners map[*User][]chan string
}

// NewSession creates a Session using the user's APIKey.
func NewSession(APIKey string) (s Session) {
	if APIKey == "" {
		return
	}

	s = Session{
		key:       APIKey,
		limiter:   NewRateLimit(),
		listeners: nil,
	}

	// Listens for user events
	go func() {
		for {
			for u := range s.listeners {
				for _, check := range u.listeners {
					updated, _ := s.FetchUser(UserCall{UserID:u.UserID})
					check(&s, u, updated) // Checks if an event has occured
				}
			}
		}
	}()

	return s
}

// AddListener adds an event listener to the Session struct instance
/*func (s *Session) AddListener(e string, ch chan string) {
	if s.listeners == nil {
		s.listeners = make(map[string][]chan string)
	}

	if _, ok := s.listeners[e]; ok {
		s.listeners[e] = append(s.listeners[e], ch)
	} else {
		s.listeners[e] = []chan string{ch}
	}
}*/

// RemoveListener removes an event listener from the Session struct instance
/*func (s *Session) RemoveListener(e string, ch chan string) {
	if _, ok := s.listeners[e]; ok {
		for i := range s.listeners[e] {
			if s.listeners[e][i] == ch {
				s.listeners[e] = append(s.listeners[e][:i], s.listeners[e][i+1:]...)
				break
			}
		}
	}
}*/

// Emit emits an event on the Session struct instance
func (s *Session) Emit(e User, response string) {
	if _, ok := s.listeners[&e]; ok {
		for _, handler := range s.listeners[&e] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

// Builds an API Call to osu API v1
func (s *Session) buildCall(endpoint string, v url.Values) string {
	return endpointAPI + endpoint + v.Encode()
}

// ParseJSON parses received JSON from url into a target interface
func (s *Session) parseJSON(url string, target interface{}) error {
	if !s.limiter.CanRequest {
		return errors.New("ratelimit exceded (Limit: " + strconv.Itoa(s.limiter.MaxRequests) + " requests.)")
	}

	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	json.Unmarshal([]byte(buf.String()), &target)
	s.limiter.iterate()

	return err

}

func remove(s []int, i int) []int {
	if len(s) >= 1 {
		s[i] = s[len(s)-1]
		return s[:len(s)-1]
	} else {
		return []int{}
	}
}