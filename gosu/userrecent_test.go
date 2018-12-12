package gosu

import (
	"fmt"
	"os"
)

func ExampleUserRecent() {
	s := NewSession(os.Getenv("API_KEY"))

	c := UserRecentCall{
		UserID: os.Getenv("USER_ID"),
	}

	userrecent, _ := s.FetchUserRecent(c)

	if len(userrecent) > 0 {
		fmt.Println(userrecent[0].UserID)
	}
}
