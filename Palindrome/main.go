package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	stringX := strconv.Itoa(x)
	start := 0
	end := len(stringX) - 1

	for start < end {
		if stringX[start] != stringX[end] {
			return false
		}
		start++

		end--
	}
	return true
}

func main() {
	var x int
	fmt.Printf("Siema podaj twoja zmienna integer x=")
	fmt.Scan(&x)
	if isPalindrome(x) {
		fmt.Println("Jest palindromem")
	} else {
		fmt.Println("Nie jest polindromem")
	}

}
