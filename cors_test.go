package headers

import (
	"testing"
	"time"
)

func TestCorsValue(t *testing.T) {
	var age AccessControlMaxAge
	var method AccessControlRequestMethod

	verify(t, []testcase{
		{&age, "0"},
		{&AccessControlMaxAge{time.Hour}, "3600"},
		{&method, ""},
		{&AccessControlRequestMethod{"POST"}, "POST"},
		{&AccessControlAllowMethods{[]string{"POST", "GET"}}, "POST, GET"},
	})
}
