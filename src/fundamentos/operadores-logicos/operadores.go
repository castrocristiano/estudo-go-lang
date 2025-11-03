package main

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarSorvete := trabalho1 || trabalho2
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2
	ficarSaudavel := !comprarSorvete
	return comprarSorvete, comprarTv50, comprarTv32 && ficarSaudavel
}

func main() {
	sorvete, tv50, saudavel := compras(true, false)
	println("Sorvete:", sorvete)
	println("TV 50\":", tv50)
	println("Ficar saudável:", saudavel)

	sorvete, tv50, saudavel = compras(true, true)
	println("Sorvete:", sorvete)
	println("TV 50\":", tv50)
	println("Ficar saudável:", saudavel)

	sorvete, tv50, saudavel = compras(false, true)
	println("Sorvete:", sorvete)
	println("TV 50\":", tv50)
	println("Ficar saudável:", saudavel)

	sorvete, tv50, saudavel = compras(false, false)
	println("Sorvete:", sorvete)
	println("TV 50\":", tv50)
	println("Ficar saudável:", saudavel)
}