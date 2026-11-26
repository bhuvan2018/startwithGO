package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	var roomTemp Celsius = 28.5
	var boilingWater Fahrenheit = 212.0

	fmt.Printf("Room Temp:     %v (Type: %T)\n", roomTemp, roomTemp)
	fmt.Printf("Boiling Water: %v (Type: %T)\n\n", boilingWater, boilingWater)

	// 1. Same defined types can do math together:
	warmerRoom := roomTemp + 2.0 // Untyped constant 2.0 adapts automatically!
	fmt.Println("Warmer Room:", warmerRoom)

	// 2. THE SAFETY TRAP (Different types CANNOT mix!)
	// Uncommenting the line below causes a compiler error:
	// "invalid operation: roomTemp + boilingWater (mismatched types Celsius and Fahrenheit)"
	// total := roomTemp + boilingWater

	// 3. FIX: Explicitly convert one type to match the other
	// Formula: (212°F - 32) * 5/9 = 100°C
	boilingInCelsius := Celsius((boilingWater - 32) * 5 / 9)
	fmt.Printf("Boiling Water in Celsius: %v (Type: %T)\n", boilingInCelsius, boilingInCelsius)
}
