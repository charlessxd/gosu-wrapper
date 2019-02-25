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
type Session struct {
	// Osu API Key
	key string

	// Rate limit
	limiter *RateLimit
}

// NewSession creates a Session using the user's APIKey.
func NewSession(APIKey string) (s Session) {
	if APIKey == "" {
		return
	}

	s = Session{
		key:     APIKey,
		limiter: NewRateLimit(),
	}

	return s
}

// Builds an API Call to osu API v1
func (s *Session) buildCall(endpoint string, v url.Values) string {
	return endpointAPI + endpoint + v.Encode()
}

// ParseJSON parses received JSON from url into a target interface
func (s *Session) parseJSON(url string, target interface{}) error {
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

func (s *Session) Fetch(c, target interface{}) error {
	switch c.(type) {
	case UserCall:
		call := c.(UserCall)
		if _, ok := target.(User); ok {
			if f, e := s.fetchUser(call); e == nil {
				target = &f
			} else {
				return errors.New("user does not exist")
			}
		} else {
			return errors.New("target is not type User when c is type UserCall")
		}
	case BeatmapCall:
		call := c.(BeatmapCall)
		if _, ok := target.(Beatmap); ok {
			if f, e := s.fetchBeatmap(call); e == nil {
				target = f
			} else {
				return errors.New("beatmap does not exist")
			}
		} else {
			return errors.New("target is not type Beatmap when c is type BeatmapCall")
		}
	case BeatmapsCall:
		call := c.(BeatmapsCall)
		if _, ok := target.(Beatmaps); ok {
			if f, e := s.fetchBeatmaps(call); e == nil {
				target = f
			} else {
				return errors.New("beatmaps do not exist")
			}
		} else {
			return errors.New("target is not type Beatmaps when c is type BeatmapsCall")
		}
	case MatchCall:
		call := c.(MatchCall)
		if _, ok := target.(Match); ok {
			if f, e := s.fetchMatch(call); e == nil {
				target = f
			} else {
				return errors.New("match does not exist")
			}
		} else {
			return errors.New("target is not type Match when c is type MatchCall")
		}
	case ScoresCall:
		call := c.(ScoresCall)
		if _, ok := target.(Scores); ok {
			if f, e := s.fetchScores(call); e == nil {
				target = f
			} else {
				return errors.New("beatmap does not exist")
			}
		} else {
			return errors.New("target is not type Scores when c is type ScoresCall")
		}
	case UserBestCall:
		call := c.(UserBestCall)
		if _, ok := target.(UserBest); ok {
			if f, e := s.fetchUserBest(call); e == nil {
				target = f
			} else {
				return errors.New("user does not exist")
			}
		} else {
			return errors.New("target is not type UserBest when c is type UserBestCall")
		}
	case UserRecent:
		call := c.(UserRecentCall)
		if _, ok := target.(UserRecent); ok {
			if f, e := s.fetchUserRecent(call); e == nil {
				target = f
			} else {
				return errors.New("user does not exist")
			}
		} else {
			return errors.New("target is not type UserRecent when c is type UserRecentCall")
		}
	default:
		return errors.New("c is not an api call type; i.e. UserCall, BeatmapCall, etc")
	}
	return nil
}
