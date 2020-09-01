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
	fmt.Printf("%s: %s", h.Name(), h.Value())
	// Output: Access-Control-Allow-Headers: X-Custom-Header
}

// This example shows Access-Control-Allow-Headers when it specifies support
// for multiple headers.
func ExampleReverseAccessControlAllowHeaders_multiple_headers() {
	h := AccessControlAllowHeaders{
		Headers: []string{"X-Custom-Header", "Upgrade-Insecure-Requests"},
	}
	fmt.Printf("%s: %s", h.Name(), h.Value())
	// Output: Access-Control-Allow-Headers: X-Custom-Header, Upgrade-Insecure-Requests
}

// Although CORS-safelisted request headers are always allowed and don't
// usually need to be listed in Access-Control-Allow-Headers, listing them
// anyway will circumvent the additional restrictions that apply.
func ExampleReverseAccessControlAllowHeaders_bypassing_additional_restrictions() {
	h := AccessControlAllowHeaders{
		Headers: []string{"Accept"},
	}
	fmt.Printf("%s: %s", h.Name(), h.Value())
	// Output: Access-Control-Allow-Headers: Accept
}
