package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go callDatabase(&wg)
	go callAPI(&wg)
	go internalProcess(&wg)

	wg.Wait()

	var m sync.Mutex
	i := 0
	for x := 0; x < 10000; x++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			i++
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(i)
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("callDatabase finished")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("callAPI finished")
	wg.Done()
}

func internalProcess(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("internalProcess finished")
	wg.Done()
}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber

}
