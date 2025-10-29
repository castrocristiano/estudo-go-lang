package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.14159
	var raio = 3.2
	area := PI * m.Pow(raio, 2)
	println("Área da circunferência:", area)

	const (
		A = 1
		B = 2
	)
	var ( 
		C = 3
		D = 4
	)
	println("Constantes A e B:", A, B)
	fmt.Printf("Variáveis C e D: %d %d\n", C, D)
	g,h,i :=2, false, "oia"
	fmt.Println(g,h,i)
}