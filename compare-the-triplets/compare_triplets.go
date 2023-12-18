package main

import "fmt"

func main() {

	var a []int32
	a = append(a, 1, 2, 3)

	var b []int32
	b = append(b, 3, 2, 1)

	fmt.Println(compareTriplets(a, b))

}

func compareTriplets(a []int32, b []int32) []int32 {
	var answer []int32
	var aPoint int32
	var bPoint int32
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if a[i] > b[i] {
				aPoint++
			} else {
				bPoint++
			}
		}
	}
	answer = append(answer, aPoint, bPoint)
	return answer
}
