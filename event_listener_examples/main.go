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

type topPlayEvent struct {
	user        gosu.User
	play        gosu.UserBestPlay
	playRanking int
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
	rankEvent := make(chan userRankEvent)

	// Go routine to check for changes in a user's rank in terms of Performance Points.
	// Checks for rank changes every 5 seconds.
	go func(e chan userRankEvent) {
		for {
			t := u
			if err := u.Update(); err != nil {
				fmt.Println(err)
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
	}(rankEvent)

	tpEvent := make(chan topPlayEvent)

	ub, _ := s.FetchUserBest(gosu.UserBestCall{UserID: u.UserID})

	go func(e chan topPlayEvent) {
		for {
			t := ub
			if err := u.Update(); err != nil {
				fmt.Println(err)
			}

			for i, play := range ub.Plays {
				if ub.Plays[i].Date != t.Plays[i].Date {
					if i <= 5 {
						e <- topPlayEvent{u, play, i + 1}
					}
				}
			}
		}
	}(tpEvent)

	for {
		select {
		case e := <-rankEvent: // When rank change event has occurred.
			fmt.Println(e.changeF)
			fmt.Println(fmt.Sprintf("%s is now rank: %d", e.user.Username, e.user.PPRank))

		case e := <-tpEvent:
			fmt.Println(u.Username + " has a new top 5 play!")
			if b, err := s.FetchBeatmap(gosu.BeatmapCall{BeatmapID: e.play.BeatmapID}); err != nil {
				fmt.Println(err)
			} else {
				fmt.Print("\n" + strconv.Itoa(e.playRanking) + ". " + b.Title + " - " + b.Artist)
				fmt.Print(" [" + b.Version + "] +")
				for x := 0; x < len(e.play.EnabledMods); x++ {
					fmt.Print(e.play.EnabledMods[x])
				}
				fmt.Print(" " + fmt.Sprintf("%.f", e.play.PP) + "pp")
				fmt.Print(" " + strconv.Itoa(e.play.MaxCombo) + "/" + strconv.Itoa(b.MaxCombo))
				fmt.Print(" " + fmt.Sprintf("%.2f", e.play.Accuracy) + "%")
			}
		}
	}
}
