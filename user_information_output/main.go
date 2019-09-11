package main

import (
	"../gosu"
	"fmt"
	"os"
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
	u := gosu.User{}

	c := gosu.UserCall{
		UserID: userID,
	}

	s := gosu.NewSession(key)

	if user, err := s.FetchUser(c); err != nil {
		fmt.Println(err)
		return
	} else {
		u = user
	}

	fmt.Println(u)
}
