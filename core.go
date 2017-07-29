package headers

import "net/url"

//The SourceMap HTTP response header links generated code to a source map,
//enabling the browser to reconstruct the original source and present the
//reconstructed original in the debugger.
//
// https://mdn.io/SourceMap
type SourceMap struct {
	// A relative (to the request URL) or absolute URL pointing to a source map file.
	URL *url.URL
}

func (h SourceMap) Name() string {
	return "SourceMap"
}

func (h SourceMap) Value() string {
	return h.URL.String()
}

func (h *SourceMap) Parse(hdr string) error {
	smap, err := url.Parse(hdr)
	if err != nil {
		return err
	}
	h.URL = smap
	return nil
}

var _ Header = &SourceMap{}
