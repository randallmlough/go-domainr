# Domainr Go Library

## A simple golang library that interacts with the domainr.com API

Refer to the [domainr docs](https://domainr.com/docs/api/v2) on how to get your API key.

**This library is still a work in progress** it is meant as a personal learning experience. It may not be 100% idiomatic Go, but I'm working towards it. Open to any contributions and/or suggestions.

The basics
```go
package main

import "github.com/randallmlough/go-domainr"

func main() {
	
    // Or by initializing and passing in the config struct
    cfg := domainr.Config{
    	APIEndpoint: "https://domainr.p.mashape.com",
        AuthKey: "XXX",
    }
    drCfg := domainr.NewClient().SetCfg(cfg)
    drCfg.Search.Search("google.com", nil)
}
``` 
