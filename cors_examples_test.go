package headers

import "fmt"

// Here's an example of what an Access-Control-Allow-Headers header might look
// like. It indicates that in addition to the CORS-safelisted request headers,
// a custom header named X-Custom-Header is supported by CORS requests to the
// server.
func ExampleReverseAccessControlAllowHeaders_a_custom_header() {
	h := AccessControlAllowHeaders{
		Headers: []string{"X-Custom-Header"},
	}
	fmt.Printf("%s: %s\n", h.Name(), h.Value())
	// Output: Access-Control-Allow-Headers: X-Custom-Header
}
