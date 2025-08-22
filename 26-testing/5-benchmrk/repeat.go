package benchmrk

import "strings"

func Repeat(s string, n int) string {
	var result string
	for range n {
		result += s
	}
	return result
}

func EfficientRepeat(s string, n int) string {
	var result strings.Builder
	for range n {
		result.WriteString(s)
	}
	return result.String()
}
