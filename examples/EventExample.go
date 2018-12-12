package main

import (
	"../gosu"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	s := gosu.NewSession(os.Getenv("API_KEY"))

	c := gosu.UserCall{
		UserID: os.Getenv("USER_ID"),
	}

	u, _ := s.FetchUser(c)

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
				s.Emit(u.UserID, strconv.FormatFloat(t.PPRaw - u.PPRaw, 'G', -1, 64))
				u, _ = s.FetchUser(c)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}