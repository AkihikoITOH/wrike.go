# wrike.go [![Build Status](https://travis-ci.org/AkihikoITOH/wrike.go.svg?branch=master)](https://travis-ci.org/AkihikoITOH/wrike.go)

Go implementation of [Wrike API](https://developers.wrike.com/documentation/api/overview) client

## Documentation

Find documentation with of examples at [GoDoc](https://godoc.org/github.com/AkihikoITOH/wrike.go)

## Basic usage

```go
import "github.com/AkihikoITOH/wrike.go"

conf := wrike.NewConfig("wrike-access-token", "")
api := wrike.NewAPI(conf)

user := api.QueryUser("KUAAAAHP")
```
