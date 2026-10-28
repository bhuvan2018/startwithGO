package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Go Programming is Awesome!"

	fmt.Println("--- Original Text ---")
	fmt.Println("Text:", message)

	// 1. Get the length of a string (Built-in len function)
	// Counts total characters including spaces!
	fmt.Println("\n1. String Length:", len(message))

	// 2. Change Text Case (Upper and Lower)
	fmt.Println("2. Uppercase:", strings.ToUpper(message))
	fmt.Println("3. Lowercase:", strings.ToLower(message))

	// 3. Trim Spaces (Removes leading and trailing whitespaces)
	cleanMessage := strings.TrimSpace(message)
	fmt.Println("4. Trimmed:", cleanMessage)

	// 4. Search / Check Content (Contains)
	// Checks if a specific word exists inside the text (Returns true/false)
	hasAwesome := strings.Contains(cleanMessage, "Awesome")
	fmt.Println("5. Contains 'Awesome'?:", hasAwesome)

	// 5. Replace Text
	// Main text, Target word, replacement word, and how many times (-1 means replace ALL instances)
	fixedMessage := strings.Replace(cleanMessage, "Awesome", "Fantastic", -1)
	fmt.Println("6. Replaced:  ", fixedMessage)
}
