package main

import "fmt"

func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!\n")

	x := 3.14159
	xs := fmt.Sprintf("%.2f", x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", xs)
	fmt.Printf("O valor de x é %.3f", x)

	a := 10
	b := 1.999999999
	c := false
	d := "Palavra"
	
	fmt.Printf("\nA variável a é do tipo %T e tem valor %d", a, a)
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}