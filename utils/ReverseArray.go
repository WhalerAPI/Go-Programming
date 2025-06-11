package utils

func ReverseArrayCopy(s []string) []string {
	n := len(s)
	out := make([]string, n)

	for i := 0; i < n; i++ {
		out[n-1-i] = s[i]
	}
	return out
}
