package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sayWord(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}

func sumRange(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func fibbonaci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(2+rand.Intn(2)) * time.Second)
	fmt.Printf("Task %d completed\n", id)
}
func main() {
	// go routine

	// go sayWord("Veet")
	// sayWord("Golang")

	// normal channel

	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// c := make(chan int)
	// go sumRange(s[:len(s)/2], c)
	// go sumRange(s[len(s)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println("Sum:", x, y, x+y)

	// buffered channel

	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// range and close

	// c := make(chan int, 10)
	// go fib(cap(c), c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	// select - multiple channel operations
	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibbonaci(c, quit)

	// WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go heavyTask(i, &wg)
	}
	wg.Wait()

}
