package between

import (
	"strings"
)

func Between(str, start, end string) string {
	indexStart := strings.Index(str, start)

	if indexStart == -1 {
		return ""
	}

	indexStart += len(start)
	strNew := str[indexStart:]

	indexEnd := strings.Index(strNew, end)

	if indexEnd == -1 {
		return ""
	}

	return strNew[:indexEnd]
}
