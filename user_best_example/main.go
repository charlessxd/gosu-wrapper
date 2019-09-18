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
)

func init() {
	key = os.Getenv("API_KEY")
	userID = os.Getenv("USER_ID")
}

func main() {
	s := gosu.NewSession(key)

	if ub, e := s.FetchUserBest(gosu.UserBestCall{UserID: userID,}); e != nil {
		fmt.Println(e)
	} else {
		if u, e := s.FetchUser(gosu.UserCall{UserID: userID,}); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(u.Username + "'s  Top 10 Plays")

			for i := 0; i < 10; i++ {
				if b, e := s.FetchBeatmap(gosu.BeatmapCall{BeatmapID: ub.Plays[i].BeatmapID}); e != nil {
					fmt.Println(e)
				} else {
					fmt.Print("\n" + strconv.Itoa(i+1) + ". " + b.Title + " - " + b.Artist)
					fmt.Print(" [" + b.Version + "] +")
					for x := 0; x < len(ub.Plays[i].EnabledMods); x++ {
						fmt.Print(ub.Plays[i].EnabledMods[x])
					}
					fmt.Print(" " + fmt.Sprintf("%.f", ub.Plays[i].PP) + "pp")
				}
			}
		}
	}
}
