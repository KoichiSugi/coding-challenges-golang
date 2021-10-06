package main

import (
	"fmt"
	"sync"
)

type Result struct {
	base int
	fib  int64
}

var resultC = make(chan Result, NUM)

const NUM int = 100

//Dynamic Programming— O(N) Time, O(N) Space
//Create a cache previously computed values so that we can reference them for later use.
func main() {
	done := make(chan bool)
	go print(done)
	createWorkerPool()
	<-done
}

func print(c chan bool) {
	for e := range resultC {
		fmt.Printf("base num : %d, fib num: %d \n", e.base, e.fib)
	}
	c <- true
}

func createWorkerPool() {
	var wg sync.WaitGroup
	for i := 1; i <= NUM; i++ {
		wg.Add(1)
		go calculateFibonacci(i, &wg)
	}
	wg.Wait()
	close(resultC)
}
func calculateFibonacci(num int, wg *sync.WaitGroup) {
	resultC <- Result{num, fib(num)}
	wg.Done()
}

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
