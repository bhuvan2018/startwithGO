package main

import "fmt"

func hello() {
	const ok = "Hello"
	var h1 = "Bye"
}

func main() {
	fmt.Println(ok, h1)
}
