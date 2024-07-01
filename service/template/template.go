package template

func Shorter(s string) string {
	if len(s) > 20 {
		return s[:5] + "..." + s[len(s)-5:]
	}
	return s
}
