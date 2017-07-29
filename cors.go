package headers

import (
	"strconv"
	"strings"
	"time"
)

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
		return err
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
