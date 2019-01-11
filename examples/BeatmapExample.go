package main

import (
	"../gosu"
	"fmt"
	"os"
)

func main() {
	s := gosu.NewSession(os.Getenv("API_KEY"))

	/*call := gosu.BeatmapCall{
		BeatmapID: os.Getenv("BEATMAP_ID"),
	}*/

	beatmap, err := s.FetchBeatmap(gosu.BeatmapCall{
		BeatmapID: os.Getenv("BEATMAP_ID"),
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(beatmap)
}
