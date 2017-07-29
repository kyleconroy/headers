package headers

import (
	"net/url"
	"testing"
)

func TestSourceMap(t *testing.T) {
	foo, _ := url.Parse("/foo")
	bar, _ := url.Parse("https://example.com/bar")

	verify(t, []testcase{
		{&SourceMap{foo}, "/foo"},
		{&SourceMap{bar}, "https://example.com/bar"},
	})
}
