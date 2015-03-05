package between

import (
	"strings"
)

func Between(str, start, end string) string {
	_start := strings.Index(str, start)
	_end := strings.Index(str, end)

	if _start == -1 || _end == -1 || _end <= _start {
		return ""
	}

	_start += len(start)

	return str[_start:_end]
}
