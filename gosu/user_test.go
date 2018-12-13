package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_FetchUser() {
	s := NewSession(os.Getenv("API_KEY"))

	c := UserCall{
		UserID: os.Getenv("USER_ID"),
	}

	user, _ := s.FetchUser(c)

	fmt.Println(user.UserID)
}
