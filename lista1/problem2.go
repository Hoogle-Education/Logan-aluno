package main

import (
	"fmt"
	"os"
)

func main() {

	var input string
	fmt.Scan(&input)

	// tratamente de exceção
	for i := 0; i < len(input); i++ {
		value := input[i] - 48
		if value >= 0 && value <= 9 {
			os.Exit(3)
		}
	}
	fmt.Println("Output:")
	result := Sottostringhe(input)
	counting := 0
	for allowedSize := 2; true; allowedSize++ {
		counting += allowedPrint(result, allowedSize)
		if counting == len(result) {
			break
		}
	}
}

func allowedPrint(result map[string]int, allowedSize int) int {

	count := 0
	for key, element := range result {

		if len(key) == allowedSize {
			fmt.Println(key, " ", element)
			count++
		}
	}

	return count
}

func Sottostringhe(input string) map[string]int {
	subsCount := make(map[string]int, 0)
	// marker
	for i := 0; i < len(input)-1; i++ {
		// push
		substringSize := 1
		for j := i + 1; j < len(input)+1; j++ {

			if input[i] < input[j] {
				substringSize++
			}

			if (input[i] >= input[j]) || j == len(input)-1 {
				if substringSize > 1 {
					// fmt.Println(i, " ", j-1)
					if !(j == len(input)-1) {
						subsCount[input[i:j]]++
					} else {
						subsCount[input[i:j+1]]++
					}
				}
				break
			}

		}
	}

	return subsCount
}
