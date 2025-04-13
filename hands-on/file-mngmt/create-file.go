package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString("This is some random data.\n")
	_, err = file.WriteString("Random number: " + fmt.Sprintf("%d", rand.Intn(100)) + "\n")
	if err != nil {
		panic(err)
	}
}
