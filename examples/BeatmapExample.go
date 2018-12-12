package main

import (
	"../gosu"
	"fmt"
	"os"
)

func main() {
	s := gosu.NewSession(os.Getenv("API_KEY"))

	call := gosu.BeatmapCall{
		BeatmapID: "252002",
	}

	beatmap, _ := s.FetchBeatmap(call)

	fmt.Println(beatmap)
}