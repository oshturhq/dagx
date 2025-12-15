package dagx

import (
	"errors"
)

var (
	ErrCycleDetected = errors.New("dagx.cycle_detected")
	ErrEmptyGraph    = errors.New("dagx.empty_graph")
)
