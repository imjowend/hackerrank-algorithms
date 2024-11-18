package main

import "fmt"

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var arraySum int32
	for _, value := range ar {
		arraySum += value
	}
	return arraySum
}

func main() {
	// Example array
	ar := []int32{1, 2, 3, 4, 5}

	result := simpleArraySum(ar)

	fmt.Println(result)
}
