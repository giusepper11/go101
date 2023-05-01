package main

import "fmt"

func main() {

	channel := make(chan int, 100)
	go setList(channel)
	PrintMSG(channel)
}

func setList(c chan<- int) {
	defer close(c)
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Println("sending: ", i)
	}
}

func PrintMSG(c <-chan int) {
	for i := range c {
		fmt.Println("receiving: ", i)
	}

}
