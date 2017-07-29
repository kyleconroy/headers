package headers

import (
	"net/url"
	"testing"
	"time"
)

func TestStrictTransportSecurity(t *testing.T) {
	var sts StrictTransportSecurity
	verify(t, []testcase{
		{&sts,
			"max-age=0"},
		{&StrictTransportSecurity{},
			"max-age=0"},
		{&StrictTransportSecurity{MaxAge: time.Hour},
			"max-age=3600"},
		{&StrictTransportSecurity{MaxAge: time.Hour, IncludeSubdomains: true},
			"max-age=3600; includeSubDomains"},
		{&StrictTransportSecurity{MaxAge: time.Hour, IncludeSubdomains: true, Preload: true},
			"max-age=3600; includeSubDomains; preload"},
	})
}

func TestFrameOptions(t *testing.T) {
	uri, _ := url.Parse("http://example.com")
	verify(t, []testcase{
		{&FrameOptions{}, "DENY"},
		{&FrameOptions{Directive: FrameDirectiveSameOrigin}, "SAMEORIGIN"},
		{FrameOptionsAllow(uri), "ALLOW-FROM http://example.com"},
	})
}

func TestXSSProtection(t *testing.T) {
	verify(t, []testcase{
		{&XSSProtection{Disabled: true}, "0"},
		{&XSSProtection{Disabled: true, Report: "http://example.com"}, "0"},
		{&XSSProtection{Disabled: true, Block: true}, "0"},
		{&XSSProtection{}, "1"},
		{&XSSProtection{Block: true}, "1; mode=block"},
		{&XSSProtection{Report: "http://example.com"}, "1; report=\"http://example.com\""},
	})
}

func TestContentTypeOptions(t *testing.T) {
	var nosniff ContentTypeOptions
	verify(t, []testcase{
		{&nosniff, "nosniff"},
	})
}
