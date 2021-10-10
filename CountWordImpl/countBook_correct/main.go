package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const FILENAME = "cat.txt"
const BUFFERSIZE = 3000
const NUMOFWORKER = 5

var words = make(chan string, BUFFERSIZE) //job
var resultC = make(chan Result, BUFFERSIZE)

var total = map[string]int{}

type Result struct {
	word  string
	count int
}

func main() {
	computeTotalDone := make(chan struct{})
	startTime := time.Now()
	go readText()
	go computeTotal(computeTotalDone)
	workerPool()       //blocking
	<-computeTotalDone // stop race condition between computeTotal and fmt.println(total)
	fmt.Println(total)
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Println("total words: ", len(total))
	fmt.Println("Time taken for reading the book", timeTaken)
}

func computeTotal(done chan struct{}) {
	defer close(done)
	i := 0
	for e := range resultC {
		total[e.word] += e.count
		i += 1
		fmt.Println(i)
	}
}

//ensure close words at the right timing
func readText() {

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		words <- strings.Trim(word, ".,:;")

	}
	//time.Sleep(1 * time.Second)
	close(words)
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

//call countWord func,
func worker(wg *sync.WaitGroup) {
	var tempMap = make(map[string]int)
	for w := range words {
		resultC <- countWord(w, tempMap) //retuns Result value
	}
	wg.Done()

}

//creates a map each word
func countWord(word string, tempMap map[string]int) Result {
	_, ok := tempMap[word]
	if ok {
		tempMap[word]++
		return Result{word, tempMap[word] + 1}

	}
	return Result{word, 1}

}
