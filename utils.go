package main

import (
	"regexp"
	"strconv"
	"strings"
)

func capitalizeFirstLetter(text string) string {
	if len(text) == 0 {
		return text
	}
	return strings.ToUpper(text[:1]) + strings.ToLower(text[1:])
}

func applyToLastN(words []string, n int, mode string) {
	if n > len(words) {
		n = len(words)
	}
	for i := len(words) - n; i < len(words); i++ {
		switch mode {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = capitalizeFirstLetter(words[i])
		}
	}
}

func parseCount(text string) (int, error) {
	text = strings.TrimSuffix(text, ")")
	return strconv.Atoi(text)
}

func hexToDecimal(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)
}

func binToDecimal(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func splitAttachedCommands(words []string) []string {
	re := regexp.MustCompile(`^(.*?)(\((?:up|low|cap)\))$`)
	var result []string

	for _, word := range words {
		matches := re.FindStringSubmatch(word)
		if len(matches) == 3 {
			if matches[1] != "" {
				result = append(result, matches[1])
			}
			result = append(result, matches[2])
		} else {
			result = append(result, word)
		}
	}

	return result
}
