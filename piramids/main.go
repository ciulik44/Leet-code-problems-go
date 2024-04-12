package main

import "fmt"

func main() {
	var rows int
	fmt.Printf("Podaj ilość wierszy: ")
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		for k := 0; k < rows-i; k++ {
			fmt.Print(" ")
		}

		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
