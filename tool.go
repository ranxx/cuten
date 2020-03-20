package cuten

import (
	"fmt"
	"runtime"
)

func fileinfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return fmt.Sprintf("%s:%d", "<<<?>>>", 0)
	}
	return fmt.Sprintf("%s:%d", file, line)
}
