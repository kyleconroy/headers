package headers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// The Access-Control-Allow-Credentials response header indicates whether or not
// the response to the request can be exposed to the page. It can be exposed when
// the true value is returned.
//
// Credentials are cookies, authorization headers or TLS client certificates.
//
// When used as part of a response to a preflight request, this indicates whether
// or not the actual request can be made using credentials. Note that simple GET
// requests are not preflighted, and so if a request is made for a resource with
// credentials, if this header is not returned with the resource, the response is
// ignored by the browser and not returned to web content.
//
// The Access-Control-Allow-Credentials header works in conjunction with the
// XMLHttpRequest.withCredentials property or with the credentials option in the
// Request() constructor of the Fetch API. Credentials must be set on both sides
// (the Access-Control-Allow-Credentials header and in the XHR or Fetch request)
// in order for the CORS request with credentials to succeed.
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials
type AccessControlAllowCredentials struct {
}

func (h AccessControlAllowCredentials) Name() string {
	return "Access-Control-Allow-Credentials"
}

func (h AccessControlAllowCredentials) Value() string {
	return "true"
}

func (h *AccessControlAllowCredentials) Parse(hdr string) error {
	if strings.ToLower(hdr) != "true" {
		return fmt.Errorf("The only valid value for Access-Control-Allow-Credentials is true (case-sensitive): got %s", hdr)
	}
	return nil
}

var _ Header = &AccessControlAllowCredentials{}

type AccessControlMaxAge struct {
	Age time.Duration
}

func (h AccessControlMaxAge) Name() string {
	return "Access-Control-Max-Age"
}

func (h AccessControlMaxAge) Value() string {
	return strconv.Itoa(int(time.Duration(h.Age).Seconds()))
}

func (h *AccessControlMaxAge) Parse(hdr string) error {
	age, err := strconv.Atoi(hdr)
	if err != nil {
		return fmt.Errorf("The value for Access-Control-Max-Age must be an integer; got %s", hdr)
	}
	*h = AccessControlMaxAge{time.Duration(age) * time.Second}
	return nil
}

var _ Header = &AccessControlMaxAge{}

type AccessControlRequestMethod struct {
	Method string
}

func (h AccessControlRequestMethod) Name() string {
	return "Access-Control-Request-Method"
}

func (h AccessControlRequestMethod) Value() string {
	return h.Method
}

func (h *AccessControlRequestMethod) Parse(hdr string) error {
	*h = AccessControlRequestMethod{hdr}
	return nil
}

var _ Header = &AccessControlRequestMethod{}

type AccessControlRequestHeaders struct {
	Headers []string
}

func (h AccessControlRequestHeaders) Name() string {
	return "Access-Control-Request-Headers"
}

func (h AccessControlRequestHeaders) Value() string {
	return strings.Join(h.Headers, ", ")
}

func (h *AccessControlRequestHeaders) Parse(hdr string) error {
	*h = AccessControlRequestHeaders{strings.Split(hdr, ", ")}
	return nil
}

var _ Header = &AccessControlRequestHeaders{}

type AccessControlAllowMethods struct {
	Methods []string
}

func (h AccessControlAllowMethods) Name() string {
	return "Access-Control-Allow-Methods"
}

func (h AccessControlAllowMethods) Value() string {
	return strings.Join(h.Methods, ", ")
}

func (h *AccessControlAllowMethods) Parse(hdr string) error {
	h = &AccessControlAllowMethods{strings.Split(hdr, ", ")}
	return nil
}

var _ Header = &AccessControlAllowMethods{}

type AccessControlAllowHeaders struct {
	Headers []string
}

func (h AccessControlAllowHeaders) Name() string {
	return "Access-Control-Allow-Headers"
}

func (h AccessControlAllowHeaders) Value() string {
	return strings.Join(h.Headers, ", ")
}

func (h *AccessControlAllowHeaders) Parse(hdr string) error {
	*h = AccessControlAllowHeaders{strings.Split(hdr, ", ")}
	return nil
}

var _ Header = &AccessControlAllowHeaders{}

type AccessControlExposeHeaders struct {
	Headers []string
}

func (h AccessControlExposeHeaders) Name() string {
	return "Access-Control-Expose-Headers"
}

func (h AccessControlExposeHeaders) Value() string {
	return strings.Join(h.Headers, ", ")
}

func (h *AccessControlExposeHeaders) Parse(hdr string) error {
	*h = AccessControlExposeHeaders{strings.Split(hdr, ", ")}
	return nil
}

type AccessControlAllowOrigin struct {
	Origin string
}

func (h AccessControlAllowOrigin) Name() string {
	return "Access-Control-Allow-Origin"
}

func (h AccessControlAllowOrigin) Value() string {
	return h.Origin
}

func (h *AccessControlAllowOrigin) Parse(hdr string) error {
	*h = AccessControlAllowOrigin{hdr}
	return nil
}

var _ Header = &AccessControlAllowOrigin{}
