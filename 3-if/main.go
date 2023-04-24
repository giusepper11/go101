package main

import "fmt"

func main() {
	salario := 1999.00
	var salarioPlusBonus float64

	salarioPlusBonus = salario

	if salario < 1000 {
		salarioPlusBonus += 100
	} else {
		salarioPlusBonus += 50
	}

	fmt.Println("Salario: ", salarioPlusBonus)


	var ehCarro bool = true
	var valorDoAutomovel = 1000.00

	if ehCarro {
		valorDoAutomovel += 55.50
	}

	fmt.Println("Valor final: ", valorDoAutomovel)
}
