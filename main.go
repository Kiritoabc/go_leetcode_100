package main

import (
	"fmt"
	"sync"
	"time"
)

func Loop(ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("send %d\n", v)
			wg.Done()
		}
	}
}

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go Loop(ch)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- i
		}(i)
		wg.Wait()
	}

	time.Sleep(3 * time.Second)
}
