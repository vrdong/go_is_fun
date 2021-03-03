package main

import (
	"fmt"
	"sync"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals, 200)
	go square(squares, naturals)
	printer(squares)
}

func counter(out chan<- int, limit int) {
	for i := 1; i <= limit; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
