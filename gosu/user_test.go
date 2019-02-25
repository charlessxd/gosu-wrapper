package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_fetchUser() {
	s := NewSession(os.Getenv("API_KEY"))
	c := UserCall{
		UserID: os.Getenv("USER_ID"),
	}
	u := User{}

	s.Fetch(c, &u)

	fmt.Println(u.UserID)
}
