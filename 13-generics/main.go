package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Println(reverse(s1))

	s2 := []string{"a", "b", "c", "d"}
	fmt.Println(reverse(s2))
}

type constraintCustom interface {
	int | string
}

func reverse[T constraintCustom](slice []T) []T {

	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}
