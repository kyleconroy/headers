package headers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// The HTTP Strict-Transport-Security response header (often abbreviated as
// HSTS) is a security feature that lets a web site tell browsers that it should
// only be communicated with using HTTPS, instead of using HTTP.
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security#Preloading_Strict_Transport_Security
type StrictTransportSecurity struct {
	// The duration that the browser should remember that this site is only to be accessed using HTTPS.
	MaxAge time.Duration
	// If true, this rule applies to all of the site's subdomains as well.
	IncludeSubdomains bool
	// Google maintains an HSTS preload service. By following the guidelines and
	// successfully submitting your domain, browsers will never connect to your
	// domain using an insecure connection. While the service is hosted by Google,
	// all browsers have stated an intent to use (or actually started using) the
	// preload list.

	// Information regarding the HSTS preload list in Chrome : https://www.chromium.org/hsts
	// Consultation of the Firefox HSTS preload list : https://dxr.mozilla.org/comm-central/source/mozilla/security/manager/ssl/nsSTSPreloadList.inc
	Preload bool
}

func (h StrictTransportSecurity) Name() string {
	return "Strict-Transport-Security"
}

func (h StrictTransportSecurity) Value() string {
	v := fmt.Sprintf("max-age=%d", int(h.MaxAge.Seconds()))
	if h.IncludeSubdomains {
		v += "; includeSubDomains"
	}
	if h.Preload {
		v += "; preload"
	}
	return v
}

func (h *StrictTransportSecurity) Parse(hdr string) error {
	directives, err := ParseDirectives(hdr)
	if err != nil {
		return err
	}
	val := StrictTransportSecurity{}
	for name, value := range directives {
		name = strings.TrimSpace(strings.ToLower(name))
		if name == "preload" {
			val.Preload = true
		} else if name == "includesubdomains" {
			val.IncludeSubdomains = true
		} else if name == "max-age" {
			age, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			val.MaxAge = time.Duration(age) * time.Second
		}
	}
	*h = val
	return nil
}

// If you specify FrameOptionsDeny, not only will attempts to load the page in a
// frame fail when loaded from other sites, attempts to do so will fail when
// loaded from the same site. On the other hand, if you specify
// FrameOptionsSameOrigin, you can still use the page in a frame as long as the
// site including it in a frame is the same as the one serving the page.
type FrameDirective int8

const (
	// The page cannot be displayed in a frame, regardless of the site attempting to do so.
	FrameDirectiveDeny FrameDirective = iota
	// The page can only be displayed in a frame on the same origin as the page itself.
	FrameDirectiveSameOrigin
	// The page can only be displayed in a frame on the specified origin.
	FrameDirectiveAllowFrom
)

// The X-Frame-Options HTTP response header can be used to indicate whether or not
// a browser should be allowed to render a page in a <frame>, <iframe> or <object>
// . Sites can use this to avoid clickjacking attacks, by ensuring that their
// content is not embedded into other sites.
//
// The added security is only provided if the user accessing the document is using
// a browser supporting X-Frame-Options.
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
type FrameOptions struct {
	Directive FrameDirective
	URL       *url.URL
}

func (h FrameOptions) Name() string {
	return "X-Frame-Options"
}

func (h FrameOptions) Value() string {
	switch h.Directive {
	case FrameDirectiveAllowFrom:
		return "ALLOW-FROM " + h.URL.String()
	case FrameDirectiveSameOrigin:
		return "SAMEORIGIN"
	default:
		return "DENY"
	}
}

func (h *FrameOptions) Parse(hdr string) error {
	val := FrameOptions{}
	switch {
	case hdr == "DENY":
		val.Directive = FrameDirectiveDeny
	case hdr == "SAMEORIGIN":
		val.Directive = FrameDirectiveSameOrigin
	case strings.HasPrefix(hdr, "ALLOW-FROM "):
		uri, err := url.Parse(hdr[11:])
		if err != nil {
			return err
		}
		val.Directive = FrameDirectiveAllowFrom
		val.URL = uri
	default:
		return fmt.Errorf("Unknown X-Frame-Options directive: %s", hdr)
	}
	*h = val
	return nil
}

// The page can only be displayed in a frame on the specified origin.
func FrameOptionsAllow(uri *url.URL) Header {
	return &FrameOptions{FrameDirectiveAllowFrom, uri}
}

type XSSProtection struct {
	Disabled bool
	Block    bool
	Report   string
}

func (h XSSProtection) Name() string {
	return "X-XSS-Protection"
}

func (h XSSProtection) Value() string {
	if h.Disabled {
		return "0"
	}
	v := "1"
	if h.Block {
		v += "; mode=block"
	}
	if h.Report != "" {
		v += "; report=\"" + h.Report + "\""
	}
	return v
}

func (h *XSSProtection) Parse(hdr string) error {
	directives, err := ParseDirectives(hdr)
	if err != nil {
		return err
	}
	val := XSSProtection{}
	if _, disabled := directives["0"]; disabled {
		val.Disabled = true
	} else {
		for name, value := range directives {
			name = strings.TrimSpace(strings.ToLower(name))
			if name == "mode" {
				val.Block = true
			} else if name == "report" {
				val.Report = value
			}
		}
	}
	*h = val
	return nil
}

type ContentTypeOptions struct {
}

func (h ContentTypeOptions) Name() string {
	return "X-Content-Type-Options"
}

func (h ContentTypeOptions) Value() string {
	return "nosniff"
}

func (h *ContentTypeOptions) Parse(hdr string) error {
	return nil
}

var _ Header = &ContentTypeOptions{}
