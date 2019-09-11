package main

import (
	"../gosu"
	"fmt"
	"os"
)

var (
	Key    string
	UserID string
)

func init() {
	Key = os.Getenv("API_KEY")
	UserID = os.Getenv("USER_ID")
}

func main() {
	u := gosu.User{}

	c := gosu.UserCall{
		UserID: UserID,
	}

	s := gosu.NewSession(Key)

	if user, err := s.FetchUser(c); err != nil {
		fmt.Println(err)
		return
	} else {
		u = user
	}

	fmt.Println(u.Username)
}
