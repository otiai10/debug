package debug

import (
	"fmt"
	"runtime"
)

var (
	// DefaultSkip ...
	DefaultSkip = 2
)

func newCaller(skip int) *caller {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}
	return &caller{
		File: file,
		Line: line,
	}
}

type caller struct {
	File string
	Line int
}

// String ...
func (c *caller) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("%v:%v", c.File, c.Line)
}
