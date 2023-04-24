package main

import "fmt"

func main() {
	array := [2]int{4, 7}
	soma := 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}

	fmt.Println("Resultado EX1", soma)

	array2 := []int{2, 8, 3, 10, 5, 4, 7, 9, 1, 7}
	somaAte5 := 0
	soma6Ate10 := 0

	for v := range array2 {
		if v <= 5 {
			somaAte5 += v
		} else {
			soma6Ate10 += v
		}
	}

	fmt.Println("SomaAte5", somaAte5)
	fmt.Println("soma6Ate10", soma6Ate10)

}
