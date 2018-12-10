package gosu

import (
	"time"
)

type RateLimit struct {
	MaxRequests     int
	CurrentRequests int
	CanRequest      bool
	FirstRequest    time.Time
	TimeInterval    float64
}

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

func (s *Session) SetRateLimit(max int, seconds float64) {
	s.Limiter = RateLimit{
		MaxRequests:     max,
		CurrentRequests: 0,
		CanRequest:      true,
		FirstRequest:    time.Now(),
		TimeInterval:    seconds,
	}
}

func (l *RateLimit) Update() {
	if time.Since(l.FirstRequest).Seconds() >= l.TimeInterval {
		l.CurrentRequests = 0
		l.CanRequest = true
	}
}

func (l *RateLimit) Iterate() bool {
	l.Update()
	if l.CanRequest {
		if l.CurrentRequests == 0 {
			l.FirstRequest = time.Now()
		}

		l.CurrentRequests ++

		if l.CurrentRequests >= l.MaxRequests {
			l.CanRequest = false
		}

		return true
	}

	return false
}
