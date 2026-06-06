package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func worker(workerId int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker is", workerId, "Job is", job)
		results <- job * 2
	}
}

func main () {
	fmt.Println("UUID: ", uuid.New().String())

	// Worker pool pattern
	numJobs := 20
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Create go routines
	for i:=1; i<=numWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, jobs, results)
		}(i)
	}

	// start sending data into the channels
	for i:=1; i <= numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	go func () {
		wg.Wait()

		close(results)
	}()

	for result := range results {
		fmt.Println("Range is :", result)
	}

	// Pipeline pattern

	data := []int{1, 2, 3, 4, 5}
	inputCh := make(chan int, len(data))


	for i:=0; i<len(data); i++ {
		inputCh <- data[i]
	}
	close(inputCh)

	doubleOpCh := make(chan int, len(data))
	go func () {
		defer close(doubleOpCh)
		for value := range inputCh {
			doubleOpCh <- value * 2
		}
	}()

	squareOpCh := make(chan int, len(data))
	go func() {
		defer close(squareOpCh)
		for value := range(doubleOpCh) {
			squareOpCh <- value * value
		}
	}()

	for value := range squareOpCh {
		fmt.Println("Final value is: ", value)
	}
}






























