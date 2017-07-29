package headers

import "net/http"

type Header interface {
	Name() string
	Value() string
	Parse(hdr string) error
}

func Set(w http.ResponseWriter, h Header) {
	w.Header().Add(h.Name(), h.Value())
}

func Get(r *http.Request, h Header) error {
	return h.Parse(r.Header.Get(h.Name()))
}
