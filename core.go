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
	return strconv.Itoa(int(h.Cached.Seconds()))
}

func (h *Age) Parse(hdr string) error {
	age, err := strconv.Atoi(hdr)
	if err != nil {
		return fmt.Errorf("The value for Age must be an integer; got %s", hdr)
	}
	*h = Age{time.Duration(age) * time.Second}
	return nil
}

// The Date general HTTP header contains the date and time at which the message
// was originated.
//
// https://mdn.io/HTTP/Date
type Date struct {
	Time time.Time
}

func (h Date) Name() string {
	return "Date"
}

func (h Date) Value() string {
	return h.Time.Format(time.RFC1123)
}

func (h *Date) Parse(hdr string) error {
	t, err := time.Parse(time.RFC1123, hdr)
	if err != nil {
		return fmt.Errorf("Invalid Date value; got %s", hdr)
	}
	*h = Date{t}
	return nil
}

// The DNT (Do Not Track) request header indicates the user's tracking
// preference. It lets users indicate whether they would prefer privacy rather
// than personalized content.
//
// https://mdn.io/DNT
type DoNotTrack struct {
	// By default, the user prefers not to be tracked on the target site.
	AllowTracking bool
}

func (h DoNotTrack) Name() string {
	return "DNT"
}

func (h DoNotTrack) Value() string {
	if h.AllowTracking {
		return "0"
	}
	return "1"
}

func (h *DoNotTrack) Parse(hdr string) error {
	switch hdr {
	case "0":
		h.AllowTracking = true
	case "1":
		h.AllowTracking = false
	default:
		return fmt.Errorf("DNT must be either '0' or '1', got '%s'", hdr)
	}
	return nil
}

// The Retry-After response HTTP header indicates how long the user agent
// should wait before making a follow-up request. There are two main cases this
// header is used:
//
// - When sent with a 503 (Service Unavailable) response, it indicates how long
// the service is expected to be unavailable.
//
// - When sent with a redirect response, such as 301 (Moved Permanently), it
// indicates the minimum time that the user agent is asked to wait before
// issuing the redirected request.
//
// https://mdn.io/Retry-After
type RetryAfter struct {
	// A date after which to retry.
	Date *time.Time
	// Delay duration after the response is received.
	Delay time.Duration
}

func (h RetryAfter) Name() string {
	return "Retry-After"
}

func (h RetryAfter) Value() string {
	if h.Date != nil {
		return h.Date.Format(time.RFC1123)
	}
	return strconv.Itoa(int(h.Delay.Seconds()))
}

func (h *RetryAfter) Parse(hdr string) error {
	if age, err := strconv.Atoi(hdr); err == nil {
		*h = RetryAfter{Delay: time.Duration(age) * time.Second}
		return nil
	}
	if t, err := time.Parse(time.RFC1123, hdr); err == nil {
		*h = RetryAfter{Date: &t}
		return nil
	}
	return fmt.Errorf("The value for Retry-After must be an integer or date; got %s", hdr)
}
