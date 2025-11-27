package processor

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"text_assistent/config"
)

func GetMainGoal(text string) string {
	return chatGPTRequest("Визнач головну мету цього тексту:\n" + text)
}

func ShortSummary(text string) string {
	return chatGPTRequest("Зроби короткий виклад цього тексту:\n" + text)
}

func chatGPTRequest(prompt string) string {
	apiKey := config.GetAPIKey()
	if apiKey == "" {
		return "API ключ не знайдено"
	}

	requestBody := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	body, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Помилка при зверненні до API"
	}
	defer resp.Body.Close()

	res, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(res, &data)

	choicesRaw, ok := data["choices"].([]interface{})
	if !ok || len(choicesRaw) == 0 {
		return "Пустий відповідь від AI"
	}

	messageMap, ok := choicesRaw[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		return "Некоректний формат відповіді AI"
	}

	content, ok := messageMap["content"].(string)
	if !ok {
		return "Немає тексту у відповіді AI"
	}

	return content
}
