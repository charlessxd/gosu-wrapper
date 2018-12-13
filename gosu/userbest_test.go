package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_FetchUserBest() {
	s := NewSession(os.Getenv("API_KEY"))

	c := UserBestCall{
		UserID: os.Getenv("USER_ID"),
	}

	userbest, _ := s.FetchUserBest(c)

	if len(userbest) > 0 {
		fmt.Println(userbest[0].UserID)
	}
}
