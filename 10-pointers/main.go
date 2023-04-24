package main

import "fmt"

func main() {
	x := 5
	y := &x
	*y = 10

	fmt.Println(x, *y)
	fmt.Println(&x, y)

	printVal(&x, y)

	fmt.Println(x, *y)
}

func printVal(x *int, y *int) {
	fmt.Println(*x, *y)
	fmt.Println(x, y)

	*x = 20
	fmt.Println(*x, *y)
}
