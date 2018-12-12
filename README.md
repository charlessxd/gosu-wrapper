# gosu-wrapper

[![GoDoc](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu?status.svg)](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu) 
[![GoReport](https://goreportcard.com/badge/github.com/charlessxd/gosu-wrapper)](https://goreportcard.com/report/github.com/charlessxd/gosu-wrapper)

This wrapper was initially intended for personal use, but feel free to use it if you feel it's of use. 
The osu-api of which this is based on can be found [here](https://github.com/ppy/osu-api/wiki).


## Installation

If you don't already have a working Go environment, go to [this page](https://golang.org/doc/install) before continuing.

```sh
go get github.com/charlessxd/gosu-wrapper/gosu
```


## Usage

Import the package into your project.

```go
import "github.com/charlessxd/gosu-wrapper/gosu"
```

## Examples
Examples can be found in the [examples directory](https://github.com/charlessxd/gosu-wrapper/tree/master/examples).
##### BeatmapExample.go 
```go
s := gosu.NewSession(os.Getenv("API_KEY"))

call := gosu.BeatmapCall{
	BeatmapID: "252002",
}

beatmap, _ := s.FetchBeatmap(call)

fmt.Println(beatmap)
```

## Documentation
All exported variables and functions have been documented:
 *  [![GoDoc](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu?status.svg)](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu) 
 
## To do
 * Improved error checking
 * Examples