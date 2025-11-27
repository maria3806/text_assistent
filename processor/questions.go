package processor

import "fmt"

func GenerateQuestions(text string, count int) string {
	prompt := fmt.Sprintf("Згенеруй %d запитань до цього тексту:\n%s", count, text)
	return chatGPTRequest(prompt)
}
