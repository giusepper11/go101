package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	// var weight int
	// weight = 10

	weight := 10
	fmt.Println(weight)

	// default values
	var texto string
	var numero int
	var metro float32
	var ehVerdade bool

	fmt.Println(texto)
	fmt.Println(numero)
	fmt.Println(metro)
	fmt.Println(ehVerdade)

	//Conversion
	var num int8 = 127
	var numInt int
	numInt = int(num)
	fmt.Println(numInt)

	b, _ := strconv.ParseBool("true")
	println(b)
	println(reflect.TypeOf(b))

}
