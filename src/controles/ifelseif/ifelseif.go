package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 {
		return "A"
	} else if nota >= 7 {
		return "B"
	} else if nota >= 5 {
		return "C"
	} else if nota >= 3 {
		return "D"
	}
	return "E"
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(6.7))
	fmt.Println(notaParaConceito(4.3))
	fmt.Println(notaParaConceito(2.1))
}