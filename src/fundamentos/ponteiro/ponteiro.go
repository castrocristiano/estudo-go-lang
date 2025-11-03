package main

import "fmt"

func main() {
	i := 1

	// Go não tem operador de ponteiro como C/C++
	// mas tem o operador & para obter o endereço de memória
	// e o operador * para desreferenciar um ponteiro

	p := &i	  // p é um ponteiro para um inteiro
	*p++         // incrementa o valor apontado por p (i)
	fmt.Println(i)   // imprime 2

	var ponteiro *int = nil // nill é igual a null em outras linguagens
	ponteiro = &i

	fmt.Println(ponteiro, *ponteiro, i, &i)
}