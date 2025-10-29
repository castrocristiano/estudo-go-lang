package main

import (
	"fmt"
	"reflect"
	"math"
)

func main() {
	fmt.Println("Tipos em Go")

	var a int = 10
	var b float64 = 20.5
	var c string = "Hello, Go!"
	var d bool = true

	fmt.Printf("Tipo de a: %s, Valor de a: %d\n", reflect.TypeOf(a), a)
	fmt.Printf("Tipo de b: %s, Valor de b: %.2f\n", reflect.TypeOf(b), b)
	fmt.Printf("Tipo de c: %s, Valor de c: %s\n", reflect.TypeOf(c), c)
	fmt.Printf("Tipo de d: %s, Valor de d: %t\n", reflect.TypeOf(d), d)

	var b2 byte = 255
	fmt.Printf("Tipo de b: %s, Valor de b: %d, Caractere de b: %c\n", reflect.TypeOf(b2), b2, b2)

	i1 := math.MaxInt64
	fmt.Printf("Tipo de i1: %s, Valor de i1: %d\n", reflect.TypeOf(i1), i1)

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))

	var x float32 = 10.99
	fmt.Printf("Tipo de x: %s, Valor de x: %.2f\n", reflect.TypeOf(x), x)
	fmt.Println("O tipo do literal de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal de 49.99 é", reflect.TypeOf(49.99))

	bo := true
	fmt.Printf("Tipo de bo: %s, Valor de bo: %t\n", reflect.TypeOf(bo), bo)
	fmt.Println(!bo)

	s1 := "Olá meu nome é Cristiano"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Texto com
	várias
	linhas`
	fmt.Println("s2 =", s2)
	fmt.Println("O tamanho da string s2 é", len(s2))

	char := 'a'
	fmt.Printf("Tipo de char: %s, Valor de char: %c, Código Unicode: %d\n", reflect.TypeOf(char), char, char)
}