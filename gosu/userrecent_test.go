package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_FetchUserRecent() {
	s := NewSession(os.Getenv("API_KEY"))

	c := UserRecentCall{
		UserID: os.Getenv("USER_ID"),
	}

	userrecent, _ := s.FetchUserRecent(c)

	if len(userrecent.Plays) > 0 {
		fmt.Println(userrecent.Plays[0].UserID)
	}
}
