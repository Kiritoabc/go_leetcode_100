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

var (
	n      = 1010
	parent []int
)

func initialize() {
	for i := 0; i < n; i++ {
		parent[i] = i
	}
}

func find(u int) int {
	if u == parent[u] {
		return u
	}
	parent[u] = find(parent[u])
	return parent[u]
}

func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	parent[v] = u
}

func isSame(u, v int) bool {
	return find(u) == find(v)
}

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
