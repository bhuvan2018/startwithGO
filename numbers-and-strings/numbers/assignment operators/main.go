package main

import "fmt"

func main() {
	score := 50
	fmt.Println("Initial Score:", score)

	// 1. Add and Assign (score = score + 10)
	score += 10
	fmt.Println("After += 10:", score) // 60

	// 2. Subtract and Assign (score = score - 5)
	score -= 5
	fmt.Println("After -= 5: ", score) // 55

	// 3. Multiply and Assign (score = score * 2)
	score *= 2
	fmt.Println("After *= 2: ", score) // 110

	// 4. Divide and Assign (score = score / 10)
	score /= 10
	fmt.Println("After /= 10:", score) // 11

	// 5. Modulus and Assign (score = score % 3)
	score %= 3
	fmt.Println("After %= 3: ", score) // 2 (Remainder of 11 / 3)
}
