package main

import "fmt"

func main() {
	txt := "palavra"
	fmt.Println("Quantidade: ", len(txt))

	for i := 0; i < len(txt); i++ {
		fmt.Println(string(txt[i]))
		if string(txt[i]) == "r" {
			break
		}
	}

	
}


