package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	base int
	fib  int64
}

var resultC = make(chan Result, NUM)
var numbers = make(chan int, NUM)

const NUM int = 100
const NUMOFWORKER = 5

func main() {
	startTime := time.Now()
	go splitNumbers()
	go print()
	workerPool()
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Println("Time taken for reading the book", timeTaken)
}
func splitNumbers() {
	for i := 1; i <= NUM; i++ {
		numbers <- i
	}
	close(numbers)
}
func workerPool() {
	var wg sync.WaitGroup
	for i := 1; i <= NUMOFWORKER; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println("all goroutines finished")
	close(resultC)
}
func worker(wg *sync.WaitGroup) {
	for number := range numbers {
		resultC <- Result{number, fib(number)}
	}
	wg.Done()
}

func print() {
	for e := range resultC {
		fmt.Printf("base num : %d, fib num: %d \n", e.base, e.fib)
	}
}

//Dynamic Programming— O(N) Time, O(N) Space
//Create a cache previously computed values so that we can reference them for later use.
func fib(num int) int64 {
	baseCases := map[int]int64{
		1: 0,
		2: 1,
	}
	return getCache(num, baseCases)
}

//helper fucntion takes n's fib value along with map storing a pair
func getCache(n int, cmap map[int]int64) int64 {
	if v, found := cmap[n]; found {
		return v
	}
	cmap[n] = getCache(n-1, cmap) + getCache(n-2, cmap) //recursive calls to calculate a fib number
	return cmap[n]
}
