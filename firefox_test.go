package headers

import "testing"

func TestFirefox(t *testing.T) {
	verify(t, []testcase{
		{&LargeAllocation{}, "0"},
		{&LargeAllocation{500}, "500"},
	})
}
