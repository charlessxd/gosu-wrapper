package rank_change_notifier

import (
	"OsuAPI/gosu-wrapper/gosu"
	"flag"
	"fmt"
	"strconv"
	"time"
)

var (
	APIKey string
	UserID string
)

func init() {
	flag.StringVar(&APIKey, "k", "", "API Key")
	flag.StringVar(&UserID, "u", "", "User's ID")
}

func main() {
	s := gosu.NewSession(APIKey)

	c := gosu.UserCall{
		UserID: UserID,
	}

	u, err := s.FetchUser(c)
	if err != nil {
		fmt.Println(err)
	}

	event := make(chan string)

	go func(e chan string) {
		for {
			t := u
			u.Update()

			if t.PPRank != u.PPRank {
				change := ""
				if t.PPRank < u.PPRank {
					change = fmt.Sprintf("+%d", strconv.FormatInt(u.PPRank-t.PPRank, 10))
				} else {
					change = fmt.Sprintf("%d", strconv.FormatInt(t.PPRank-u.PPRank, 10))
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
