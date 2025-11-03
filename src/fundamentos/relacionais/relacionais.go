package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Igualdade:", "Banana" == "Banana")
	fmt.Println("Diferença:", 3 != 2)
	fmt.Println("Maior que:", 3 > 2)
	fmt.Println("Menor que:", 3 < 2)
	fmt.Println("Maior ou igual a:", 3 >= 2)
	fmt.Println("Menor ou igual a:", 3 <= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas iguais:", d1 == d2)
	fmt.Println("Datas iguais:", d1.Equal(d2))

	type Pessoa struct {
		nome string
		idade int
	}

	p1 := Pessoa{"João", 30}
	p2 := Pessoa{"Jão", 30}

	fmt.Println("Pessoas iguais:", p1 == p2)

}
