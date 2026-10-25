package main

import "fmt"

func main() {
	// A simple math problem: 0.1 + 0.2 should equal 0.3, right?
	result := 0.1 + 0.2

	fmt.Println("--- The Float Surprise ---")
	// 1. Let's see what Go prints out natively:
	fmt.Println("Result using Println:", result)

	// 2. Let's look closely at the deep decimal values using Printf:
	fmt.Printf("Deep float precision: %.20f\n", result)

	fmt.Println("\n--- The Comparison Trap ---")
	// 3. Trying to compare it strictly to 0.3 will fail!
	if result == 0.3 {
		fmt.Println("Math works! It matches 0.3")
	} else {
		fmt.Println("Warning: It does NOT strictly equal 0.3!")
	}
}
