package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		results <- fmt.Sprintf("worker %d processed job %d", id, job)
	}
}

func main() {
	jobs := make(chan int, 3)
	results := make(chan string, 3)

	go worker(1, jobs, results)
	go worker(2, jobs, results)

	jobs <- 101
	jobs <- 102
	jobs <- 103
	close(jobs)

	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
}
