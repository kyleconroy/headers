# stakmachine/headers
[![GoDoc](https://godoc.org/stackmachine.com/headers?status.svg)](https://godoc.org/stackmachine.com/headers) [![Build Status](https://travis-ci.org/stackmachine/headers.svg?branch=master)](https://travis-ci.org/stackmachine/headers)

stackmachine/headers is a type-safe API for manipulating HTTP headers. Say goodbye to stringly-typed, unsafe code.

## Install

```
dep ensure stackmachine.com/headers
```

## Usage

```go
headers.Set(r, headers.AccessControlMaxAge(time.Minute * 3))
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
