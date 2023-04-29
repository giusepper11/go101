package main

import (
	"fmt"
	"os"
)

func ShowText() {
	fmt.Println("Finished writing file")
}

func ReadFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}()
	_, err := os.Open("./test1.txt")
	if err != nil {
		panic(err)
	}
}

func main() {
	// panic("error")
	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}

	// defer will be executed after last instruction
	// defer file.Close()
	// defer ShowText()

	file.Write([]byte("test"))
	file.Close()

	ReadFile()

	fmt.Println("End")
}
