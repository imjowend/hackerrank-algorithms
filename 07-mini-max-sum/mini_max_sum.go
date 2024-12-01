package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	var total int64 = 0
	var sumaMinima int64 = 0
	var sumaMaxima int64 = 0
	var minInt64Number int64 = math.MaxInt64
	var maxInt64Number int64 = math.MinInt64
	j := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if int64(arr[i]) <= minInt64Number {
			minInt64Number = int64(arr[i])
		}
		if int64(arr[j]) >= maxInt64Number {
			maxInt64Number = int64(arr[j])
		}
		j--
		total += int64(arr[i])
	}

	sumaMinima = total - maxInt64Number
	sumaMaxima = total - minInt64Number
	fmt.Println(sumaMinima, sumaMaxima)
}

func main() {
	arr := []int32{1, 1, 3, 5, 5}
	arr2 := []int32{1, 1, 1, 1, 1}
	arr3 := []int32{4, 4, 4, 4, 1}

	miniMaxSum(arr)
	miniMaxSum(arr2)
	miniMaxSum(arr3)
}
