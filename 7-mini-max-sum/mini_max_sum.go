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
	var sumaMinima int32 = 0
	var sumaMaxima int32 = 0
	var minInt32Number int32 = math.MaxInt32
	var maxInt32Number int32 = math.MinInt32
	j := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[i] < minInt32Number {
			minInt32Number = arr[i]
		}
		if arr[j] > maxInt32Number {
			maxInt32Number = arr[j]
		}
		j--
	}
	fmt.Println("El numero maximo es: ", maxInt32Number)
	fmt.Println("El numero minimo es: ", minInt32Number)

	for i := 0; i < len(arr); i++ {
		if arr[i] == maxInt32Number {
			sumaMinima += 0
			sumaMaxima += maxInt32Number
			continue
		}
		if arr[i] == minInt32Number {
			sumaMinima += minInt32Number
			sumaMaxima += 0
			continue
		}
		sumaMinima += arr[i]
		sumaMaxima += arr[i]
	}
	fmt.Println("El numero minimo es: ", sumaMinima)
	fmt.Println("El numero maximo es: ", sumaMaxima)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}

	miniMaxSum(arr)
}
