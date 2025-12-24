package main

import (
	"regexp"
	"strings"
)

func cleanInput(text string) []string {

	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	cleanText := re.ReplaceAllString(text, " ")

	lowerText := strings.ToLower(cleanText)
	parts := strings.Fields(lowerText)
	return parts
}
