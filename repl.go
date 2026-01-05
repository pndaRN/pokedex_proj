package main

import (
	"regexp"
	"strings"
	"fmt"
	"os"
	"bufio"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		text := cleanInput(scanner.Text())
		firstWord := text[0]

		fmt.Printf("Your command was: %v \n", firstWord)
	}
}

func cleanInput(text string) []string {

	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	cleanText := re.ReplaceAllString(text, " ")

	lowerText := strings.ToLower(cleanText)
	parts := strings.Fields(lowerText)
	return parts
}
