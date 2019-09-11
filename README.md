# gosu-wrapper

[![GoDoc](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu?status.svg)](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu) 
[![GoReport](https://goreportcard.com/badge/github.com/charlessxd/gosu-wrapper)](https://goreportcard.com/report/github.com/charlessxd/gosu-wrapper)

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





##### Example of FetchUser() to get user metadata.
```go
package main

import (
	"github.com/charlessxd/gosu-wrapper/gosu"
	"fmt"
	"os"
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
	u := gosu.User{}

	c := gosu.UserCall{
		UserID: UserID,
	}

	s := gosu.NewSession(Key)

	if user, err := s.FetchUser(c); err != nil {
		fmt.Println(err)
		return
	} else {
		u = user
	}

	fmt.Println(u.Username)
}
```

## Documentation
All exported variables and functions have been documented: 

[![GoDoc](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu?status.svg)](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu) 

## To do
* Mod support

