package processor

import "strings"

func CompressText(text string, words int) string {
	split := strings.Fields(text)
	if len(split) <= words {
		return text
	}
	return strings.Join(split[:words], " ") + "..."
}
