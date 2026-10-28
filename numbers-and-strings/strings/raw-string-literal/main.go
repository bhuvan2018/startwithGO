package main

import "fmt"

func main() {
	// 1. Normal String Literal (Using double quotes "")
	// You MUST use escape sequences like \n or \t for layout styling.
	normalStr := "Hello Bro!\n\tWelcome.\n"
	fmt.Println("--- Normal String ---")
	fmt.Println(normalStr)

	// 2. Raw String Literal (Using backticks ``)
	// Whatever you type inside backticks is printed EXACTLY as written.
	// Line breaks, tabs, and special characters are preserved natively!
	rawStr := `Hello Bro!
	Welcome.
Looking good! \n \t (Escape codes are ignored here!)`

	fmt.Println("--- Raw String ---")
	fmt.Println(rawStr)
}
