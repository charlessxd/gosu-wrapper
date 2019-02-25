package rank_change_notifier

import (
	"OsuAPI/gosu-wrapper/gosu"
	"fmt"
	"os"
	"strconv"
	"time"
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
	s := gosu.NewSession(Key)

	c := gosu.UserCall{
		UserID: UserID,
	}

	u := gosu.User{}

	if e := s.Fetch(&c, &u); e != nil {
		fmt.Println(e)
		return
	}

	event := make(chan string)

	go func(e chan string) {
		for {
			t := u
			u.Update()

			if t.PPRank != u.PPRank {
				change := ""
				if t.PPRank < u.PPRank {
					change = fmt.Sprintf("+%s", strconv.FormatInt(u.PPRank-t.PPRank, 10))
				} else {
					change = fmt.Sprintf("%s", strconv.FormatInt(t.PPRank-u.PPRank, 10))
				}

				e <- fmt.Sprintf("%s rank change: %s", u.Username, change)
			}

			time.Sleep(time.Second * 5)
		}
	}(event)

	for {
		select {
		case msg := <-event:
			fmt.Println(msg)
		}
	}
}
