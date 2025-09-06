package main

import "sync"

var (
	counter int
	mu      sync.Mutex
)

func incr(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incr(&wg)
	}
	wg.Wait()
	println("Final Counter:", counter)
}
