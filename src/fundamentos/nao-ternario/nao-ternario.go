package	main

import "fmt"

func obterResultado(nota float64) string {
	// Não há operador ternário em Go
	// return nota >= 7 ? "Aprovado" : "Reprovado"
	if nota >= 7 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(8.5))
	fmt.Println(obterResultado(6.3))
}