package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5}
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista[0]: ", lista[0])
	fmt.Println("Lista[1]: ", lista[1])
	fmt.Println("Tamanho da lista: ", len(lista))

	lista = append(lista, 8)
	fmt.Println(lista)

	listavazia := make([]int, 1)
	listavazia = append(listavazia, 1)
	listavazia = append(listavazia, 2)
	listavazia = append(listavazia, 4)
	fmt.Printf("lista: %v\n", listavazia)

	somaTotal := 0
	for i := 0; i < len(listavazia); i++ {
		somaTotal += listavazia[i]
	}
	fmt.Println("Media: ", somaTotal/len(listavazia))

	// slice
	listaToda := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(listaToda[0:1])
}
