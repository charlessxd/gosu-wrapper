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
		Mode:   "3",
	}

	u, err := s.FetchUser(c)
	fmt.Println(u.PPRaw)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

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
		for {
			// EVENT
			time.Sleep(1 * time.Second)
			if t, _ := s.FetchUser(c); t.PPRaw != u.PPRaw {
				s.Emit(u.UserID, strconv.FormatFloat(t.PPRaw-u.PPRaw, 'G', -1, 64))
				u = t
			}
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}

	s.RemoveListener(u.UserID, event)
}
