package main

import "fmt"

func main() {
	arr := []int32{2, 4}
	arr2 := []int32{16, 32, 96}

	fmt.Println(getTotalX(arr, arr2))
}

func getTotalX(a []int32, b []int32) int32 {
	// en el Array a tendria que agarrar el numero mas pequeño y ver los multiplos que tiene con el numero mayor del Array a
	// siempre y cuando sean menor al numero mas pequeño del Array B y que sean divisibles dando resto 0

	for true {
		if a[0] < a[1] {
			a[0] += a[0]
		}
		if a[0] == a[1] {

		}
	}
	return 0
}
