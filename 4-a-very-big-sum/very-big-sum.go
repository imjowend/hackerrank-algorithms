package main

import "fmt"

func main() {
	var ar []int64
	ar = append(ar, 10, 11, 12, 13, 14)
	fmt.Println(veryBigSum(ar))
}

func veryBigSum(ar []int64) int64 {
	var result int64

	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}

	return result
}
