package main

import "fmt"

func init() {
	fmt.Println("Start init1")
}

func init() {
	fmt.Println("Start init2")
}

func main() {
	printMSG("msg XYZ")

	fmt.Println("Resultado soma(3,4): ", soma(3, 4))

	soma, sub, mult, div := all_calcs(3, 4)
	fmt.Println("Resultado all_calcs(3,4): ", soma, sub, mult, div)
}

func printMSG(msg string) {
	fmt.Println(msg)
}

func soma(n1 int, n2 int) int {
	return n1 + n2
}

func all_calcs(n1 int, n2 int) (int, int, int, int) {
	soma := n1 + n2
	sub := n1 - n2
	mult := n1 * n2
	div := n1 / n2

	return soma, sub, mult, div
}

// func all_calcs_named(n1 int, n2 int) (soma int, sub int, mult int, div int) {
// 	soma = n1 + n2
// 	sub = n1 - n2
// 	mult = n1 * n2
// 	div = n1 / n2
// 	return
// }
