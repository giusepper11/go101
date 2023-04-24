package main

import "strings"

func main() {
	m := map[string]int{
		"sp": 10000000,
		"cg": 900000,
		"rj": 6000000,
	}

	delete(m, "rj")

	for k, v := range m {
		println("Cidade: ", strings.ToUpper(k), ", Valor", v)
	}
}
