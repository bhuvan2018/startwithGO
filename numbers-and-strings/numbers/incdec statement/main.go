package main

import "fmt"

func main() {
	counter := 10

	//1. Increment
	counter++
	fmt.Println(counter)

	counter--
	fmt.Println(counter)

	// Cannot use it as an expression / assignment
	// result := counter++

	// Pre-increment (++counter) does NOT exist in Go
	// ++counter
}
