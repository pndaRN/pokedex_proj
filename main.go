package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
