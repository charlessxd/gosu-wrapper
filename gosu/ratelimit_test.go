package gosu

import (
	"testing"
	"time"
)

func TestRateLimiter_Iterate(t *testing.T) {
	l := NewRateLimit()

	for l.CanRequest {
		l.iterate()
		if l.CurrentRequests > l.MaxRequests {
			t.Fatalf("Expected (Total Requests <= Max Requests) but got (Total Requests > Max Requests).")
		}
	}
	if l.CurrentRequests != l.MaxRequests {
		t.Fatal("Expected true but got false.")
	}
}

func TestRateLimit_Update(t *testing.T) {
	l := RateLimit{
		MaxRequests:     100,
		CurrentRequests: 99,
		CanRequest:      true,
		FirstRequest:    time.Now(),
		TimeInterval:    1.0,
	}

	l.iterate()
	l.update()
	if l.CanRequest {
		t.Fatal("Expected false but got true.")
	}
	if l.CurrentRequests != 100 {
		t.Fatalf("Expected %d but got %d", 100, l.CurrentRequests)
	}

	time.Sleep(1.0 * time.Second)
	l.update()
	if !l.CanRequest {
		t.Fatal("Expected true but got false.")
	}
	if l.CurrentRequests != 0 {
		t.Fatalf("Expected %d but got %d", 0, l.CurrentRequests)
	}
}
