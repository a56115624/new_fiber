package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started jon %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}
func main() {
	numJobs := 5
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Received result: %d\n", result)
	}
}
