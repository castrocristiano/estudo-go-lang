package main

import (
	"fmt"
	"strconv")

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste de conversões")

	var i int = 100
	var b byte = byte(i)
	fmt.Println(i, b)

	var l int64 = int64(i)
	fmt.Println(i, l)

	var f float64 = float64(i)
	fmt.Println(i, f)

	var u uint = uint(i)
	fmt.Println(i, u)

	fmt.Println("Teste" + strconv.Itoa(i)) // int para string

	num, _ := strconv.Atoi("123") // string para int
	fmt.Println(num - 122)

	valor, _ := strconv.ParseBool("true") // string para bool
	valor2, _ := strconv.ParseBool("0") // string para bool
	if valor {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
	if valor2 {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}