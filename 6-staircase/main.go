package main

import "fmt"

func staircase(n int32) {
	// Write your code here

	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n); j++ {
			if (int(n) - i - 1) <= j {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func main() {
	number := 6
	numberInt32 := int32(number)
	staircase(numberInt32)
}
