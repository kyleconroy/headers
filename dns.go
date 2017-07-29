package headers

import "fmt"

// The X-DNS-Prefetch-Control HTTP response header controls DNS prefetching, a
// feature by which browsers proactively perform domain name resolution on both
// links that the user may choose to follow as well as URLs for items
// referenced by the document, including images, CSS, JavaScript, and so forth.
//
// This prefetching is performed in the background, so that the DNS is likely
// to have been resolved by the time the referenced items are needed. This
// reduces latency when the user clicks a link.
//
// By default, this will DNS prefetching. This is what browsers do, if they
// support the feature, when this header is not present
//
// https://mdn.io/X-DNS-Prefetch-Control
type DNSPrefetchControl struct {
	// Disables DNS prefetching. This is useful if you don't control the link
	// on the pages, or knows that you don't want to lead information to these
	// domains.
	Disabled bool
}

func (h DNSPrefetchControl) Name() string {
	return "X-DNS-Prefetch-Control"
}

func (h DNSPrefetchControl) Value() string {
	if h.Disabled {
		return "off"
	}
	return "on"
}

func (h *DNSPrefetchControl) Parse(hdr string) error {
	switch hdr {
	case "on":
		h.Disabled = false
	case "off":
		h.Disabled = true
	default:
		return fmt.Errorf("X-DNS-Prefetch-Control must be either 'on' or 'off', not '%s'", hdr)
	}
	return nil
}

var _ Header = &DNSPrefetchControl{}
