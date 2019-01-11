package gosu

import (
	"fmt"
	"time"
)

// RateLimit limits the amount of requests per TimeInterval seconds.
type RateLimit struct {
	MaxRequests     int
	CurrentRequests int
	FirstRequest    time.Time
	TimeInterval    float64
	CanRequest      bool
}

// NewRateLimit returns an instantiated RateLimit.
func NewRateLimit() *RateLimit {
	limiter := &RateLimit{
		MaxRequests:     100,
		CurrentRequests: 0,
		FirstRequest:    time.Now(),
		TimeInterval:    60.0,
		CanRequest:      true,
	}

	// Updates limiter every TimeInterval seconds.
	go func(l *RateLimit) {
		for {
			d, _ := time.ParseDuration(fmt.Sprintf("%fs", l.TimeInterval))

			if l.CurrentRequests >= l.MaxRequests {
				l.CanRequest = false
			}
			if time.Since(l.FirstRequest) >= d && !l.CanRequest {
				l.CanRequest = true
				l.CurrentRequests = 0
			}
		}
	}(limiter)

	return limiter
}

// SetRateLimit sets a Session's MaxRequests and TimeInterval to a given amount.
func (s *Session) SetRateLimit(max int, seconds float64) {
	s.limiter = &RateLimit{
		MaxRequests:     max,
		CurrentRequests: 0,
		FirstRequest:    time.Now(),
		TimeInterval:    seconds,
		CanRequest:      true,
	}
}

// Iterate tells RateLimit that a request has been made.
// Returns true if successfully iterated, false if not.
func (l *RateLimit) iterate() bool {
	if l.CurrentRequests < l.MaxRequests {
		if l.CurrentRequests == 0 {
			l.FirstRequest = time.Now()
		}
		l.CurrentRequests++
		return true
	}
	return false
}
