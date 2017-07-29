package headers

import "fmt"

type DNSPrefetchControl struct {
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
