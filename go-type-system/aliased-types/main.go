package main

import "fmt"

// Create an Aliased Type using the '=' sign!
// 'Score' is NOT a new type—it is literally an exact nickname for 'int'
type Score = int

func main() {
	var player1Score Score = 100
	var bonusPoints int = 25

	// 1. Check the actual type printed by %T
	// Even though we declared it as 'Score', Go sees it strictly as 'int'!
	fmt.Printf("player1Score -> Value: %d | Actual Type: %T\n", player1Score, player1Score)
	fmt.Printf("bonusPoints  -> Value: %d | Actual Type: %T\n\n", bonusPoints, bonusPoints)

	// 2. NO CONVERSION NEEDED!
	// Because Score and int are identical, you can mix and match them freely:
	total := player1Score + bonusPoints
	fmt.Println("Total Score (no conversion needed):", total)

	// 3. Built-in Go Aliases you already use!
	// 'byte' is a built-in alias for 'uint8'
	// 'rune' is a built-in alias for 'int32'
	var letter byte = 'A'
	var emoji rune = '🚀'
	fmt.Printf("\nBuilt-in aliases -> byte type: %T | rune type: %T\n", letter, emoji)
}
