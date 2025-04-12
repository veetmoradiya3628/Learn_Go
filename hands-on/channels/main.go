package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go Lang")

	myCh := make(chan int, 10)
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println("Channel is open, value is: ", val)
		} else {
			fmt.Println("Channel is closed")
		}

		// for i := 0; i < 5; i++ {
		// 	fmt.Println(<-ch)
		// }
		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// for i := 0; i < 5; i++ {
		// 	ch <- i
		// }
		ch <- 1
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
