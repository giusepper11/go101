package main

import "fmt"

func main() {
	m := map[string]int{
		"sp": 10000000,
		"cg": 900000,
	}
	fmt.Println(m)

	valor, existe := m["rj"]

	if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Nao existe")
	}
}
