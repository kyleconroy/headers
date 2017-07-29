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
// https://mdn.io/Access-Control-Allow-Credentials
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

// The Access-Control-Max-Age response header indicates how long the results of
// a preflight request (that is the information contained in the
// Access-Control-Allow-Methods and Access-Control-Allow-Headers headers) can
// be cached.
//
// https://mdn.io/Access-Control-Max-Age
type AccessControlMaxAge struct {
	// Maximum number of seconds the results can be cached. Firefox caps this at 24
	// hours (86400 seconds) and Chromium at 10 minutes (600 seconds). Chromium
	// also specifies a default value of 5 seconds. A value of -1 will disable
	// caching, requiring a preflight OPTIONS check for all calls.
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

// The Access-Control-Request-Method request header is used when issuing a
// preflight request to let the server know which HTTP method will be used when
// the actual request is made. This header is necessary as the preflight
// request is always an OPTIONS and doesn't use the same method as the actual
// request.
//
// https://mdn.io/Access-Control-Request-Method
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

// The Access-Control-Request-Headers request header is used when issuing a
// preflight request to let the server know which HTTP headers will be used
// when the actual request is made.
//
// https://mdn.io/Access-Control-Request-Headers
type AccessControlRequestHeaders struct {
	// A list of HTTP headers that are included in the request.
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

// The Access-Control-Allow-Methods response header specifies the method or
// methods allowed when accessing the resource in response to a preflight
// request.
//
// https://mdn.io/Access-Control-Allow-Methods
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

// The Access-Control-Allow-Headers response header is used in response to a
// preflight request to indicate which HTTP headers will be available via
// Access-Control-Expose-Headers when making the actual request.
//
// The simple headers, Accept, Accept-Language, Content-Language, Content-Type
// (but only with a MIME type of its parsed value (ignoring parameters) of
// either application/x-www-form-urlencoded, multipart/form-data, or
// text/plain), are always available and don't need to be listed by this
// header.
//
// This header is required if the request has an Access-Control-Request-Headers
// header.
//
// https://mdn.io/Access-Control-Allow-Headers
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

// The Access-Control-Expose-Headers response header indicates which headers
// can be exposed as part of the response by listing their names.
//
// By default, only the 6 simple response headers are exposed:
// - Cache-Control
// - Content-Language
// - Content-Type
// - Expires
// - Last-Modified
// - Pragma
//
// If you want clients to be able to access other headers, you have to list
// them using the Access-Control-Expose-Headers header.
//
// https://mdn.io/Access-Control-Expose-Headers
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

var _ Header = &AccessControlExposeHeaders{}

// The Access-Control-Allow-Origin response header indicates whether the
// response can be shared with resources with the given origin.
//
// https://mdn.io/Access-Control-Allow-Origin
type AccessControlAllowOrigin struct {
	// Specifies a URI that may access the resource. For requests without
	// credentials, the server may specify "*" as a wildcard, thereby allowing
	// any origin to access the resource.
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
