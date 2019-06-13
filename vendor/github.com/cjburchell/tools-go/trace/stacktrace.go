package trace

import (
	"bytes"
	"fmt"
	"runtime"
)

const maxStackTrace = 40

// GetStack gets the stack trace
func GetStack(min int) string {
	var buffer bytes.Buffer
	_, err := buffer.WriteString(fmt.Sprintf("Stacktrace:\n"))
	if err != nil {
		return ""
	}

	for i := min; i < maxStackTrace; i++ {
		if function1, file1, line1, ok := runtime.Caller(i); ok {
			_, err = buffer.WriteString(fmt.Sprintf("      at %s (%s:%d)\n", runtime.FuncForPC(function1).Name(), file1, line1))
			if err != nil {
				return ""
			}
		} else {
			break
		}
	}

	return buffer.String()
}
