package main

import "fmt"

func main() {
	firstName, lastName := "Bhuvan", "Go"

	// 1. Using the '+' Operator
	// You must manually add a space " " string between variables!
	fullName := firstName + " " + lastName
	fmt.Println("--- Plus Operator ---")
	fmt.Println("Full Name:", fullName)

	// 2. Plus-Equal Assignment Operator (Appending text)
	phrase := "Learning Go"
	phrase += " is super fun!" // Appends text to the existing variable
	fmt.Println("\n--- Appending Text ---")
	fmt.Println(phrase)

	// 3. Mixing Strings and Numbers (The Go Trap)
	age := 25
	// This line will FAIL if you uncomment it because string + int is illegal:
	// message := "My age is " + age

	// FIX: Use Printf to format different data types together!
	fmt.Println("\n--- Mixing Data Types ---")
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", fullName, age)
}
