package headers

// These headers are only implemented in Firefox

import (
	"fmt"
	"strconv"
)

// The non-standard Large-Allocation response header tells the browser that the
// page being loaded is going to want to perform a large allocation. It is
// currently only implemented in Firefox, but is harmless to send to every
// browser.
//
// WebAssembly or asm.js applications can use large contiguous blocks of
// allocated memory. For complex games, for example, these allocations can be
// quite large, sometimes as large as 1GB. The Large-Allocation tells the
// browser that the web content in the to-be-loaded page is going to want to
// perform a large contiguous memory allocation and the browser can react to
// this header by starting a dedicated process for the to-be-loaded document,
// for example.
//
// https://mdn.io/Large-Allocation
type LargeAllocation struct {
	// The expected size of the allocation to be performed, in megabytes. 0 is a
	// special value which represents uncertainty as to what the size of the
	// allocation is.
	Megabytes int
}

func (h LargeAllocation) Name() string {
	return "Large-Allocation"
}

func (h LargeAllocation) Value() string {
	return strconv.Itoa(h.Megabytes)
}

func (h *LargeAllocation) Parse(hdr string) error {
	size, err := strconv.Atoi(hdr)
	if err != nil {
		return fmt.Errorf("The value for Large-Allocation must be an integer; got %s", hdr)
	}
	h.Megabytes = size
	return nil
}

var _ Header = &LargeAllocation{}
