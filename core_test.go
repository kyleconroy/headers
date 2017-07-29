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
