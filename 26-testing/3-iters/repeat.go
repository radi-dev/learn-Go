package iters

func Repeat(s string, n int) string {
	var result string
	for range n {
		result += s
	}
	return result
}
