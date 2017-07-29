# stakmachine/headers
[![GoDoc](https://godoc.org/stackmachine.com/headers?status.svg)](https://godoc.org/stackmachine.com/headers) [![Build Status](https://travis-ci.org/stackmachine/headers.svg?branch=master)](https://travis-ci.org/stackmachine/headers)

Structs for headers

I want to be able to do three things:

Parse the header into the struct
WRite it into the header

```go
headers.Set(r, AccessControlMaxAge(time.Minute * 3))
```

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"stackmachine.com/headers"
)

func middleware(next http.Handler) http.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		headers.Set(w, headers.StrictTransportSecurity{
			MaxAge:            time.Hour * 24,
			IncludeSubDomains: true,
			Preload:           true,
		})
		next.ServeHTTP(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome :D")
	})
	http.ListenAndServe(":8080", middleware(mux))
}
```

## Install

```
dep ensure stackmachine.com/headers
```
