package headers

import (
	"fmt"
	"strconv"
)

// These headers are only implemented in Firefox

type LargeAllocation struct {
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
