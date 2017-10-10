package cli

import (
	"sort"
	"strings"
)

func fmtTopics(label string, topics map[string]string) string {
	text := []string{}

	// Sort the topics keys
	keys := []string{}
	for k, _ := range topics {
		keys = append(keys, k)
	}
	text = append(text, label)

	sort.Strings(keys)
	j := len(label)
	for i, k := range keys {
		if i > 0 {
			text = append(text, ", ")
			j += 2
		}
		text = append(text, k)
		if i == len(keys)-1 {
			text = append(text, ".")
		} else {
			// Wrap the line if needed
			j += len(k)
			if j > 60 {
				text = append(text, "\n")
				j = 0
			}
		}
	}
	text = append(text, "\n\n")
	return strings.Join(text, "")
}
