package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Signed Integers (Can hold NEGATIVE and POSITIVE numbers)
	var smallSigned int8 = -128
	fmt.Printf("int8   range: %d to %d (example: %d)\n", math.MinInt8, math.MaxInt8, smallSigned)
	fmt.Printf("int16  range: %d to %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  range: %d to %d\n\n", math.MinInt32, math.MaxInt32)

	// 2. Unsigned Integers ('u' = Unsigned -> POSITIVE numbers & 0 ONLY!)
	var smallUnsigned uint8 = 255
	fmt.Printf("uint8  range: 0 to %d (example: %d)\n", math.MaxUint8, smallUnsigned)
	fmt.Printf("uint16 range: 0 to %d\n\n", math.MaxUint16)

	// 3. Architecture-Dependent Default Types (Most common in daily coding)
	// Automatically uses 64-bit on modern 64-bit computers, or 32-bit on older ones
	defaultInt := 42       // defaults to 'int'
	defaultFloat := 3.14   // defaults to 'float64' (higher precision than float32)
	isReady := true        // defaults to 'bool'
	greeting := "Namaste!" // defaults to 'string'

	fmt.Println("--- Default Inferred Types ---")
	fmt.Printf("defaultInt   -> Type: %T | Value: %v\n", defaultInt, defaultInt)
	fmt.Printf("defaultFloat -> Type: %T | Value: %v\n", defaultFloat, defaultFloat)
	fmt.Printf("isReady      -> Type: %T | Value: %v\n", isReady, isReady)
	fmt.Printf("greeting     -> Type: %T | Value: %v\n", greeting, greeting)
}
