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
	Key string

	// Rate limit
	Limiter RateLimit
}

// NewSession creates a Session using the user's APIKey.
func NewSession(APIKey string) (session Session) {
	if APIKey == "" {
		return
	}

	session = Session{
		Key:     APIKey,
		Limiter: NewRateLimit(),
	}

	return session
}

// Builds an API Call to osu API v1
func (s *Session) buildCall(endpoint string, v url.Values) string {
	return EndpointAPI + endpoint + v.Encode()
}

// ParseJSON parses received JSON from url into a target interface
func (s *Session) parseJSON(url string, target interface{}) error {
	if !s.Limiter.CanRequest {
		return errors.New("ratelimit exceded (Limit: " + strconv.Itoa(s.Limiter.MaxRequests) + " requests.)")
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
	s.Limiter.iterate()

	return err
}
