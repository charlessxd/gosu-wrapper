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



## Examples
Examples can be found in the [examples directory](https://github.com/charlessxd/gosu-wrapper/tree/master/examples).

##### Example of FetchUser()
```go
package main

import (
	"fmt"
	"github.com/charlessxd/gosu-wrapper/gosu"
)

func main() {
    s := gosu.NewSession(os.Getenv("API_KEY"))
	
    c := gosu.UserCall{
        UserID: "1",
    }
    
    u := gosu.User{}
    
    if err = s.Fetch(c, u); err != nil {
	    fmt.Println(err.Error())    
    } else {
        fmt.Println(u.Username) 
    }
}
```

## Documentation
All exported variables and functions have been documented: 

[![GoDoc](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu?status.svg)](https://godoc.org/github.com/charlessxd/gosu-wrapper/gosu) 

## To do
* Mod support

