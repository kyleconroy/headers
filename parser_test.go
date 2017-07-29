package headers

import "testing"

func TestValid(t *testing.T) {
	for _, test := range []string{
		"foo",
		"foo;",
		"foo; bar;",
		"foo; bar;",
		"foo; bar; baz=3",
		"foo; bar; baz=\"3\"",
		"foo; bar; baz=\"3\\\"\"",
		"foo=",
	} {
		t.Run(test, func(t *testing.T) {
			p := newParser(test)
			if _, err := p.parse(); err != nil {
				t.Errorf("%#v", p)
				t.Errorf("%s", err)
			}
		})
	}
}

func TestInvalid(t *testing.T) {
	for _, test := range []string{
		"fo{o",
		"foo; bar=\"3; bat",
		"foo; bar; baz=\"3\\",
	} {
		t.Run(test, func(t *testing.T) {
			p := newParser(test)
			if _, err := p.parse(); err == nil {
				t.Errorf("expected err %#v", p)
			}
		})
	}
}
