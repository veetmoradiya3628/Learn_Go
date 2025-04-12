package main

import (
	"fmt"
	"time"
)

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go greeter("Hello")
	greeter("World")
}
