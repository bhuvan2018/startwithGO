package main

import "fmt"

func main() {
	//method 1: explicitly define the name, type, and the value
	var name string = "Bhuvan"
	var age int = 24

	//method 2: use short variable declaration operator (GO guess the type based on the value)
	isGo := true
	score := 98

	// You can change values later using regular '='
	age = 26
	score = 100

	fmt.Println("User:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is it Go?:", isGo)
	fmt.Println("Final Score:", score)
}
