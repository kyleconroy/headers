package headers

import (
	"testing"
	"time"
)

func TestCorsValue(t *testing.T) {
	var age AccessControlMaxAge
	var creds AccessControlAllowCredentials
	var method AccessControlRequestMethod

	verify(t, []testcase{
		{&age, "0"},
		{&creds, "true"},
		{&method, ""},
		{&AccessControlMaxAge{time.Hour}, "3600"},
		{&AccessControlRequestMethod{"POST"}, "POST"},
		{&AccessControlRequestHeaders{[]string{"Content-Length", "Host"}}, "Content-Length, Host"},
		{&AccessControlAllowMethods{[]string{"POST", "GET"}}, "POST, GET"},
		{&AccessControlAllowHeaders{[]string{"Content-Length", "Host"}}, "Content-Length, Host"},
	})
}
