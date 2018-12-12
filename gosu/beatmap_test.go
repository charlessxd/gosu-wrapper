package gosu

import (
	"fmt"
	"os"
)

func ExampleBeatmap() {
	session := NewSession(os.Getenv("API_KEY"))

	c := BeatmapCall{
		BeatmapID: os.Getenv("BEATMAP_ID"),
	}

	bm, _ := session.FetchBeatmap(c)

	fmt.Println(bm.BeatmapID)
}
