package main

import (
	"regexp"
	"strconv"
	"strings"
)

func transformText(text string) string {
	words := strings.Fields(text)
	words = splitAttachedCommands(words)
	var result []string
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
		case "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
		case "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = capitalizeFirstLetter(result[len(result)-1])
			}
		case "(up,":
			if i+1 < len(words) {
				count, err := parseCount(words[i+1])
				if err == nil {
					applyToLastN(result, count, "up")
				}
			}
			i++ // Skip the count word
		case "(low,":
			if i+1 < len(words) {
				count, err := parseCount(words[i+1])
				if err == nil {
					applyToLastN(result, count, "low")
				}
			}
			i++ // Skip the count word
		case "(cap,":
			if i+1 < len(words) {
				count, err := parseCount(words[i+1])
				if err == nil {
					applyToLastN(result, count, "cap")
				}
			}
			i++ // Skip the count word
		case "(hex)":
			if len(result) > 0 {
				hex := result[len(result)-1]
				decimal, err := hexToDecimal(hex)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(decimal, 10)
				}
			}
		case "(bin)":
			if len(result) > 0 {
				bin := result[len(result)-1]
				decimal, err := binToDecimal(bin)
				if err == nil {
					result[len(result)-1] = strconv.FormatInt(decimal, 10)
				}
			}
		default:
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}

func fixPunctuation(text string) string {
	// merge separated punctuation like "! ?" → "!?"
	re := regexp.MustCompile(`([!?.,;:])\s+([!?.,;:])`)
	text = re.ReplaceAllString(text, "$1$2")

	// remove spaces BEFORE punctuation
	re = regexp.MustCompile(`\s+([.,!?;:])`)
	text = re.ReplaceAllString(text, "$1")

	// remove spaces before grouped punctuation (... !! !?)
	re = regexp.MustCompile(`\s+([.,!?]{2,})`)
	text = re.ReplaceAllString(text, "$1")

	// ensure space AFTER punctuation
	re = regexp.MustCompile(`([.,!?;:])([^\s.,!?;:])`)
	text = re.ReplaceAllString(text, "$1 $2")

	return text
}

func fixQuotes(text string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		next := strings.ToLower(words[i+1])

		if word == "a" || word == "A" {
			if len(next) > 0 {
				first := next[0]
				if strings.ContainsRune("aeiouh", rune(first)) {
					if word == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}
