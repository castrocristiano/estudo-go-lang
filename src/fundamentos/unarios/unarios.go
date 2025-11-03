package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++
	y--

	fmt.Printf("X = %v, Y= %v\n", x, y)
}