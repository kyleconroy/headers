package headers

import "testing"

func TestDNSPrefetchControl(t *testing.T) {
	verify(t, []testcase{
		{&DNSPrefetchControl{}, "on"},
		{&DNSPrefetchControl{Disabled: true}, "off"},
	})
}
