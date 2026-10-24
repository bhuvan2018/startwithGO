package main

import "fmt"

func main() {
	name := "bhuvan"
	age := 25
	isGo := true
	pi := 3.14

	// 1. Using Println
	fmt.Println("--- Using Println ---")
	fmt.Println("Name:", name, "Age:", age)

	// 2. Using Printf with Escape Sequences (\n for new lines, \t for tabs)
	fmt.Println("\n--- Using Printf ---")
	fmt.Printf("User Details:\n\tName: %s\n\tAge: %d\n", name, age)
	// Note: We used %s for strings and %d for digits/integers, and manual \n at the end!

	// 3. Printing Data Types using %T and generic values with %v
	fmt.Println("\n--- Printing Data Types & Values ---")
	fmt.Printf("Variable name has value %v and its type is: %T\n", name, name)
	fmt.Printf("Variable age has value %v and its type is: %T\n", age, age)
	fmt.Printf("Variable isGo has value %v and its type is: %T\n", isGo, isGo)
	fmt.Printf("Variable pi has value %v and its type is: %T\n", pi, pi)
}
