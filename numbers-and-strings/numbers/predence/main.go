package main

import "fmt"

func main() {

	//the format is: PMDMAS (paranthesis, multiplication/division, modulus, addition, subtraction)

	result1 := 10 + 5*2
	fmt.Println(result1)

	result2 := 100 / 5 * 2
	fmt.Println(result2)

	result3 := (10 + 5) * 2
	fmt.Println(result3)
}
