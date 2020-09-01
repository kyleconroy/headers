package headers

import (
	"fmt"
	"time"
)

// All present and future subdomains will be HTTPS for a max-age of 1 year.
// This blocks access to pages or sub domains that can only be served over
// HTTP.
func ExampleStrictTransportSecurity_one_year() {
	h := StrictTransportSecurity{
		MaxAge:            365 * 24 * time.Hour,
		IncludeSubdomains: true,
	}
	fmt.Printf("%s: %s", h.Name(), h.Value())
	// Output: Strict-Transport-Security: max-age=31536000; includeSubDomains
}

// In the following example, max-age is set to 2 years, raised from what was a
// former limit max-age of 1 year. Note that 1 year is acceptable for a domain
// to be included in browsers' HSTS preload lists. 2 years is, however, the
// recommended goal as a website's final HSTS configuration as explained on
// https://hstspreload.org. It also suffixed with preload which is necessary
// for inclusion in most major web browsers' HSTS preload lists, e.g. Chromium,
// Edge, & Firefox.
func ExampleStrictTransportSecurity_two_year() {
	h := StrictTransportSecurity{
		MaxAge:            2 * 365 * 24 * time.Hour,
		IncludeSubdomains: true,
		Preload:           true,
	}
	fmt.Printf("%s: %s", h.Name(), h.Value())
	// Output: Strict-Transport-Security: max-age=63072000; includeSubDomains; preload
}
