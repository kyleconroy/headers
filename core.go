package headers

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

//The SourceMap HTTP response header links generated code to a source map,
//enabling the browser to reconstruct the original source and present the
//reconstructed original in the debugger.
//
// https://mdn.io/HTTP/SourceMap
type SourceMap struct {
	// A relative (to the request URL) or absolute URL pointing to a source map file.
	URL *url.URL
}

func (h SourceMap) Name() string {
	return "SourceMap"
}

func (h SourceMap) Value() string {
	return h.URL.String()
}

func (h *SourceMap) Parse(hdr string) error {
	smap, err := url.Parse(hdr)
	if err != nil {
		return err
	}
	h.URL = smap
	return nil
}

var _ Header = &SourceMap{}

// The Age header contains the time in seconds the object has been in a proxy
// cache.
//
// The Age header is usually close to zero. If it is Age: 0, it was probably
// just fetched from the origin server; otherwise It is usually calculated as a
// difference between the proxy's current date and the Date general header
// included in the HTTP response.
//
// https://mdn.io/Age
type Age struct {
	// A non-negative integer, representing time in seconds the object has been
	// in a proxy cache.
	Cached time.Duration
}

func (h Age) Name() string {
	return "Age"
}

func (h Age) Value() string {
	return strconv.Itoa(int(time.Duration(h.Cached).Seconds()))
}

func (h *Age) Parse(hdr string) error {
	age, err := strconv.Atoi(hdr)
	if err != nil {
		return fmt.Errorf("The value for Age must be an integer; got %s", hdr)
	}
	*h = Age{time.Duration(age) * time.Second}
	return nil
}
