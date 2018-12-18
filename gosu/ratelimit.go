package gosu

import (
	"fmt"
	"time"
)

// RateLimit limits the amount of requests per TimeInterval seconds.
type RateLimit struct {
	MaxRequests     int
	CurrentRequests int
	CanRequest      bool
	FirstRequest    time.Time
	TimeInterval    float64
}

// NewRateLimit returns an instantiated RateLimit.
func NewRateLimit() *RateLimit {
	limiter := &RateLimit{
		MaxRequests:     100,
		CurrentRequests: 0,
		CanRequest:      true,
		FirstRequest:    time.Now(),
		TimeInterval:    60.0,
	}

	// Updates limiter every TimeInterval seconds.
	go func(s *RateLimit) {
		for {
			d, _ := time.ParseDuration(fmt.Sprintf("%fs", limiter.TimeInterval))
			time.Sleep(time.Second * d)

			limiter.CanRequest = true
		}
	}(limiter)

	return limiter
}

// SetRateLimit sets a Session's MaxRequests and TimeInterval to a given amount.
func (s *Session) SetRateLimit(max int, seconds float64) {
	s.limiter = &RateLimit{
		MaxRequests:     max,
		CurrentRequests: 0,
		CanRequest:      true,
		FirstRequest:    time.Now(),
		TimeInterval:    seconds,
	}
}

// Iterate tells RateLimit that a request has been made.
// Returns true if successfully iterated, false if not.
func (l *RateLimit) iterate() bool {
	if l.CanRequest {
		if l.CurrentRequests == 0 {
			l.FirstRequest = time.Now()
		}

		l.CurrentRequests++

		if l.CurrentRequests >= l.MaxRequests {
			l.CanRequest = false
		}

		return true
	}

	return false
}
