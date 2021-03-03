package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a int64 = 10
	ch := make(chan int64, 3)
	ch <- a
	//ch <- b
	//ch <- c
	sum := doSth(ch)
	fmt.Println(sum)
}

func doSth(ch <-chan int64) int64 {
	sizes := make(chan int64, 3)
	var wg sync.WaitGroup
	for f := range ch {
		wg.Add(1)
		param := func(f int64) {
			defer wg.Done()
			sizes <- f
		}
		go param(f)
	}
	time.Sleep(time.Second)
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
