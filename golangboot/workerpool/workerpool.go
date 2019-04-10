package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job descreve uma serviço
type Job struct {
	id       int
	randomno int
}

// Result descreve um resultado
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
	startTime := time.Now()

	noOfJobs := 100

	go allocate(noOfJobs)

	done := make(chan bool)

	go result(done)

	noOfWorkers := 10

	createWorkerPool(noOfWorkers)

	<-done

	endTime := time.Now()

	diff := endTime.Sub(startTime)

	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
