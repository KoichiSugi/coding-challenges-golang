package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	if x < 0 {
		return false
	}
	for i := 0; i < len(str)/2; i++ {
		length := len(str) - 1
		if str[i] != str[length-i] {
			fmt.Println(str[i])
			fmt.Println(str[length-i])
			return false
		}
	}
	return true
}
