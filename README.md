# stakmachine/headers
[![PkgGoDev](https://pkg.go.dev/badge/github.com/stackmachine/headers)](https://pkg.go.dev/github.com/stackmachine/headers) ![GithubActions](https://github.com/stackmachine/headers/workflows/ci/badge.svg?branch=master)

stackmachine/headers is a type-safe API for manipulating HTTP headers. Say goodbye to stringly-typed, unsafe code.

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

	"github.com/stackmachine/headers"
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
