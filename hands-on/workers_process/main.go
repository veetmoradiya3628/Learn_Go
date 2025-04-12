package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %v\n", id, job)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numWorkers := 5
	numTasks := 10

	jobs := make(chan []int, numTasks)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= numTasks; i++ {
		length := rand.Intn(5) + 1 // random length 1 to 5
		jobData := make([]int, length)
		for j := range jobData {
			jobData[j] = rand.Intn(100)
		}
		jobs <- jobData
		fmt.Printf("Sent job %d: %v\n", i, jobData)
	}

	close(jobs)
	wg.Wait()

	fmt.Println("All jobs processed.")
}
