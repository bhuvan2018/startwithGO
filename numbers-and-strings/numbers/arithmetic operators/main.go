package main

import "fmt"

func main() {

	//standard whole numbers (integers)
	a := 10
	b := 3
	fmt.Println("---Integer Arithmetic---")
	fmt.Println("Addition of", a, "and", b, "is:", a+b)
	fmt.Println("Subtraction of", a, "and", b, "is:", a-b)
	fmt.Println("Multiplication of", a, "and", b, "is:", a*b)

	// Integer division cuts off the decimal part, bcz a & b is chosen as integer by GO
	fmt.Println("Division of", a, "and", b, "is:", a/b)
	fmt.Println("Modulus of", a, "and", b, "is:", a%b)

	//Mixing type
	price := 100
	tax := 1.5

	// this will give error as price is int and tax is float64
	// total:=price+tax

	// we can convert price to float64 to avoid error
	total := float64(price) + tax
	fmt.Println("\n--- Mixed Type Arithmetic (Converted) ---")
	fmt.Println("Total Price with Tax:", total)
}
