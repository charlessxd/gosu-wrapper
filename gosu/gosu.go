// Package gosu provides a method of accessing osu-api in Go programs.
package gosu

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

// Session holds the API key and rate limiter.
type session struct {
	// Osu API Key
	key string

	// Rate limit
	limiter *RateLimit
}

// NewSession creates a Session using the user's APIKey.
func NewSession(APIKey string) (s session) {
	if APIKey == "" {
		return
	}

	s = session{
		key:     APIKey,
		limiter: NewRateLimit(),
	}

	return s
}

func (s *session) buildCall(endpoint string, v url.Values) string {
	return endpointAPI + endpoint + v.Encode()
}

func (s *session) parseJSON(url string, target interface{}) error {
	if !s.limiter.CanRequest {
		return errors.New("ratelimit exceeded")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}

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
