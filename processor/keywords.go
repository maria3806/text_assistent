package processor

import (
	"regexp"
	"sort"
	"strings"
)

func GetKeywords(text string, top int) []string {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[a-zA-Zа-яА-ЯґҐєЄіІїЇ0-9]+`)
	words := re.FindAllString(text, -1)

	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	type kv struct {
		Key   string
		Value int
	}

	var sorted []kv
	for k, v := range freq {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	var result []string
	for i, kv := range sorted {
		if i >= top {
			break
		}
		result = append(result, kv.Key)
	}
	return result
}
