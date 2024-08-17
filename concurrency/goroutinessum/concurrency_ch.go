package main

import (
	"fmt"
	"sync"
)

func initialize(ch chan int) {
	ch <- 0
}

func threadSum(ch chan int, wg *sync.WaitGroup, id int) {
	fmt.Println("#", id, " Started")

	defer wg.Done()

	value := <-ch

	fmt.Println("#", id, " Value: ", value)
	
	go func() {
		ch <- value + 1
		fmt.Println("#", id, " Finished")
	}()

	// wg.Add(1)
}

func result(ch chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

func main() {
	var wg sync.WaitGroup

	threads_count := 100

	ch := make(chan int)

	fmt.Println("inicializando channel")
	
	go initialize(ch)
	
	fmt.Println("channel inicializado")

	wg.Add(threads_count)

	for i := 0; i < threads_count; i++ {
		go threadSum(ch, &wg, i)
	}

	wg.Wait()

	result := <-ch

	fmt.Println("result: ", result)
}
