package main

import (
	"../gosu"
	"fmt"
	"os"
	"time"
)

func main() {
	s := gosu.NewSession(os.Getenv("API_KEY"))

	c := gosu.UserCall{
		UserID: os.Getenv("USER_ID"),
	}

	u, err := s.FetchUser(c)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	event := make(chan string)

	//s.AddListener(u.UserID, event)
	u.AddListener(gosu.PPGained, event)
	u.AddListener(gosu.RankChange, event)

	// Outputs the event
	go func() {
		for {
			fmt.Println(<-event)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}

	//s.RemoveListener(u.UserID, event)
}
