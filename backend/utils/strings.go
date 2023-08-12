package utils

import "strings"

func JoinWithPrefix(prefix, separator string, input ...string) string {
	if len(input) == 0 || len(separator) == 0 {
		return Empty
	}

	return prefix + strings.Join(input, separator)
}
