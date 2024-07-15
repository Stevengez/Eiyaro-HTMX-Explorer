package template

import (
	"time"
)

func Shorter(s string) string {
	if len(s) > 20 {
		return s[:5] + "..." + s[len(s)-5:]
	}
	return s
}

func UnixToDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)

	// Formatear la fecha como dd/mm/yyyy
	return t.UTC().Format("2006-01-02 15:04:05 UTC")
}
