package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "cat.txt"
const BUFFERSIZE = 3000
const NUMOFWORKER = 5

var words = make(chan string, BUFFERSIZE) //job
var total = map[string]int{}
var resultC = make(chan map[string]int)

//one channel likely to be faster
func main() {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	go readText()
	workerPool()
	computeTotal()
	var total = map[string]int{}
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		words <- strings.Trim(word, ".,:;")
		total[word]++
	}
	fmt.Println(total)
}

func computeTotal() {
	for i := 1; i <= NUMOFWORKER; i++ {
		m := <-resultC
		for word, count := range m {
			total[word] += count
		}
	}
}

//call countWord func,
func worker() {
	var tempMap = make(map[string]int)
	for w := range words {
		tempMap[w]++
	}
	resultC <- tempMap

}

func workerPool() {
	for i := 1; i <= NUMOFWORKER; i++ {
		go worker()
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
	close(resultC)
}
