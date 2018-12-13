package main

import (
	"../gosu"
	"fmt"
	"os"
)

func main() {
	s := gosu.NewSession(os.Getenv("API_KEY"))

	call := gosu.BeatmapCall{
		BeatmapID: os.Getenv("BEATMAP_ID"),
	}

	beatmap, err := s.FetchBeatmap(call)


	fmt.Println(beatmap)
}
