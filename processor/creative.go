package processor

import "fmt"

func RewriteInStyle(text, style string) string {
	prompt := fmt.Sprintf("Перепиши цей текст у стилі: %s\n\nТекст:\n%s", style, text)
	return chatGPTRequest(prompt)
}
