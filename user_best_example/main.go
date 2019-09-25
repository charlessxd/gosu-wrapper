package main

import (
	"../gosu"
	"fmt"
	"os"
	"strconv"
)

var (
	key    string
	userID string
	limit  string
)

func init() {
	key = os.Getenv("API_KEY")
	userID = os.Getenv("USER_ID")
	limit = os.Getenv("SCORE_LIMIT")
}

func main() {
	s := gosu.NewSession(key)
	s.SetRateLimit(105, 60)

	if ub, e := s.FetchUserBest(gosu.UserBestCall{UserID: userID, Limit: limit}); e != nil {
		fmt.Println(e)
	} else {
		if u, e := s.FetchUser(gosu.UserCall{UserID: userID,}); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(u.Username + "'s  Top 10 Plays")

			for i, play := range ub.Plays {
				if b, e := s.FetchBeatmap(gosu.BeatmapCall{BeatmapID: play.BeatmapID}); e != nil {
					fmt.Println(e)
				} else {
					fmt.Print("\n" + strconv.Itoa(i+1) + ". " + b.Title + " - " + b.Artist)
					fmt.Print(" [" + b.Version + "] +")
					for x := 0; x < len(play.EnabledMods); x++ {
						fmt.Print(play.EnabledMods[x])
					}
					fmt.Print(" " + fmt.Sprintf("%.f", play.PP) + "pp")
					fmt.Print(" " + strconv.Itoa(play.MaxCombo) + "/" + strconv.Itoa(b.MaxCombo))
					fmt.Print(" " + fmt.Sprintf("%.2f", play.Accuracy) + "%")
				}
			}
		}
	}
}
