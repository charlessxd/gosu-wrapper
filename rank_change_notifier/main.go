package main

import (
	"../gosu"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	key    string
	userID string
)

func init() {
	key = os.Getenv("API_KEY")
	userID = os.Getenv("USER_ID")
}

type userRankEvent struct {
	user      gosu.User
	changeRaw int64
	changeF   string
}

/*
	This is an example of a rank tracker. This example shows how you can go
	about tracking a user's rank, and outputting the change when it occurs.

	Showcases how to use the wrapper.
*/
func main() {
	// Create a session to access the osu-api.
	s := gosu.NewSession(key)

	// Create a UserCall to get user metadata.
	c := gosu.UserCall{
		UserID: userID,
	}

	// Create a User to hold the user metadata.
	u := gosu.User{}

	if user, err := s.FetchUser(c); err != nil {
		fmt.Println(err)
		return
	} else {
		u = user
		fmt.Println(u.Username)
	}

	// userRankEvent channel containing information about the rank change.
	event := make(chan userRankEvent)

	// Go routine to check for changes in a user's rank in terms of Performance Points.
	// Checks for rank changes every 5 seconds.
	go func(e chan userRankEvent) {
		for {
			t := u
			if e := u.Update(); e != nil {
				fmt.Println(e)
			}

			if t.PPRank != u.PPRank {
				change := ""
				if t.PPRank < u.PPRank {
					change = fmt.Sprintf("+%s", strconv.FormatInt(u.PPRank-t.PPRank, 10))
				} else {
					change = fmt.Sprintf("%s", strconv.FormatInt(t.PPRank-u.PPRank, 10))
				}

				// Puts rank change information into the Channel String.
				e <- userRankEvent{u,
					t.PPRank - u.PPRank,
					fmt.Sprintf("%s rank change: %s", u.Username, change)}
			}

			time.Sleep(time.Second * 5)
		}
	}(event)

	for {
		select {
		case e := <-event: // When rank change event has occurred.
			fmt.Println(e.changeF)
			fmt.Println(fmt.Sprintf("%s is now rank: %d", e.user.Username, e.user.PPRank))
		}
	}
}
