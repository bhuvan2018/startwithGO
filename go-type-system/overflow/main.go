package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. COMPILE-TIME PROTECTION
	// Uncommenting the line below causes a compiler error: "cannot use 128 (untyped int constant) as int8 value in variable declaration (overflows)"
	// var tooBig int8 = 128

	// 2. RUNTIME OVERFLOW (Signed Integer: int8)
	// MaxInt8 is 127. Adding 1 wraps it all the way around to the minimum (-128)!
	var signedMax int8 = math.MaxInt8 // 127
	fmt.Println("--- Signed Overflow (int8) ---")
	fmt.Println("Before (+0):", signedMax)
	signedMax++                            // Adds 1 at runtime!
	fmt.Println("After  (+1):", signedMax) // Prints -128!

	// 3. RUNTIME OVERFLOW (Unsigned Integer: uint8)
	// MaxUint8 is 255. Adding 1 wraps it back to 0!
	var unsignedMax uint8 = math.MaxUint8 // 255
	fmt.Println("\n--- Unsigned Overflow (uint8) ---")
	fmt.Println("Before (+0):", unsignedMax)
	unsignedMax++                            // Adds 1 at runtime!
	fmt.Println("After  (+1):", unsignedMax) // Prints 0!

	// 4. UNDERFLOW (Going below the minimum!)
	// MinUint8 is 0. Subtracting 1 wraps it all the way up to 255!
	var unsignedMin uint8 = 0
	fmt.Println("\n--- Unsigned Underflow (uint8) ---")
	fmt.Println("Before (-0):", unsignedMin)
	unsignedMin--                            // Subtracts 1 at runtime!
	fmt.Println("After  (-1):", unsignedMin) // Prints 255!
}
