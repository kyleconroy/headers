package headers

import (
	"fmt"
	"testing"
)

type testcase struct {
	Header   Header
	Expected string
}

func marshal(t *testing.T, cases []testcase) {
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d/marshal/%s", i, c.Header.Name()), func(t *testing.T) {
			if c.Header.Value() != c.Expected {
				t.Errorf("Expected '%s', not '%s'\n", c.Expected, c.Header.Value())
			}
		})
	}
}

func unmarshal(t *testing.T, cases []testcase) {
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d/unmarshal/%s", i, c.Header.Name()), func(t *testing.T) {
			if err := c.Header.Parse(c.Expected); err != nil {
				t.Fatalf("Error parseing '%s': %s", c.Expected, err)
			}
			if c.Header.Value() != c.Expected {
				t.Errorf("Expected '%s', not '%s'\n", c.Expected, c.Header.Value())
			}
		})
	}
}

func verify(t *testing.T, cases []testcase) {
	marshal(t, cases)
	unmarshal(t, cases)
}
