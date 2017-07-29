package headers

import (
	"net/url"
	"testing"
	"time"
)

func TestSourceMap(t *testing.T) {
	foo, _ := url.Parse("/foo")
	bar, _ := url.Parse("https://example.com/bar")

	verify(t, []testcase{
		{&SourceMap{foo}, "/foo"},
		{&SourceMap{bar}, "https://example.com/bar"},
	})
}

func TestAge(t *testing.T) {
	verify(t, []testcase{
		{&Age{}, "0"},
		{&Age{1 * time.Hour}, "3600"},
	})
}

func TestDate(t *testing.T) {
	now := time.Now()
	verify(t, []testcase{
		{&Date{}, "Mon, 01 Jan 0001 00:00:00 UTC"},
		{&Date{now}, now.Format(time.RFC1123)},
	})
}

func TestDNT(t *testing.T) {
	verify(t, []testcase{
		{&DoNotTrack{}, "1"},
		{&DoNotTrack{AllowTracking: false}, "1"},
		{&DoNotTrack{AllowTracking: true}, "0"},
	})
}

func TestRetryAfter(t *testing.T) {
	now := time.Now()
	verify(t, []testcase{
		{&RetryAfter{}, "0"},
		{&RetryAfter{Delay: 5 * time.Second}, "5"},
		{&RetryAfter{Date: &now, Delay: 5 * time.Second}, now.Format(time.RFC1123)},
	})
}
