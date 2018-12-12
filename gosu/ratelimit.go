package gosu

import (
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
func NewRateLimit() RateLimit {
	limiter := RateLimit{
		MaxRequests:     100,
		CurrentRequests: 0,
		CanRequest:      true,
		FirstRequest:    time.Now(),
		TimeInterval:    60.0,
	}

	return limiter
}

// SetRateLimit sets a Session's MaxRequests and TimeInterval to a given amount.
func (s *Session) SetRateLimit(max int, seconds float64) {
	s.limiter = RateLimit{
		MaxRequests:     max,
		CurrentRequests: 0,
		CanRequest:      true,
		FirstRequest:    time.Now(),
		TimeInterval:    seconds,
	}
}

// Update updates the RateLimit's CanRequest and CurrentRequest if
// TimeInterval seconds have passed since the first request.
// First request being the request sent when CurrentRequests is 0.
func (l *RateLimit) update() {
	if time.Since(l.FirstRequest).Seconds() >= l.TimeInterval {
		l.CurrentRequests = 0
		l.CanRequest = true
	}
}

// Iterate tells RateLimit that a request has been made.
// Returns true if successfully iterated, false if not.
func (l *RateLimit) iterate() bool {
	l.update()
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
