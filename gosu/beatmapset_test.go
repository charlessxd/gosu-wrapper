package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_FetchBeatmaps() {
	session := NewSession(os.Getenv("API_KEY"))

	c := BeatmapsCall{
		BeatmapSetID: os.Getenv("BEATMAPSET_ID"),
	}

	bs, _ := session.FetchBeatmaps(c)

	if len(bs.Beatmaps) > 0 {
		fmt.Println(bs.Beatmaps[0])
	}
}