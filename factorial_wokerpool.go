package main

import (
	"fmt"
	"sync"
)

type Factorial struct {
	original int
	fNum     int
}

func factorial(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum = sum * i
	}
	return sum
}

func worker(num int, wg *sync.WaitGroup) {
	results <- Factorial{num, factorial(num)}
	wg.Done()
}

func createWorkerPool(noOfJobs int) {
	var wg sync.WaitGroup
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go worker(nums[i], &wg)
	}
	wg.Wait()
	close(results) //close the channel all the jobs are done
}

func result(c chan bool) {
	for e := range results {
		fmt.Printf("Original number : %d, factorial: %d \n", e.original, e.fNum)
	}
	c <- true
}

var results = make(chan Factorial, 10)

var nums = []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

func main() {
	done := make(chan bool)
	go result(done)
	createWorkerPool(len(nums))
	<-done
}
