package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2 

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Módulo:", a%b)

	// Bitwise
	fmt.Println("E (AND):", a&b)
	fmt.Println("Ou (OR):", a|b)
	fmt.Println("Xor (XOR):", a^b)
	fmt.Println("And Not (AND NOT):", a&^b)
	fmt.Println("Deslocamento para esquerda:", a<<1)
	fmt.Println("Deslocamento para direita:", a>>1)

	c := 9.0
	d := 2.0

	fmt.Println("Maior:", math.Max(c, d))
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Expoenciação:", math.Pow(c, d))
	fmt.Println("Raiz quadrada de c:", math.Sqrt(c))
}