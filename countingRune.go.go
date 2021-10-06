package main

import "fmt"

func main() {
	arr := []rune{'a', 'b', 'a', 'a', 'b', 'c', 'd', 'e', 'f', 'c', 'a', 'b', 'a', 'd'}
	m := make(map[rune]int)
	for _, e := range arr {
		_, ok := m[e] //search through map with a rune
		if ok {
			//if a rune value exists
			//increment key's value by 1
			m[rune(e)]++
		} else {
			//if does not exist, create a new pair of map with value of 1
			m[rune(e)] = 1
		}
	}
	//iterate map
	for k, v := range m {
		fmt.Printf("%c found %d times\n", k, v)
	}

}
