package main

import "fmt"

type House struct {
	startPoint int32
	endPoint   int32
}

type AppleTree struct {
	position int32
	apples   []int32
}

type OrangeTree struct {
	position int32
	oranges  []int32
}

func main() {
	house := House{7, 11}
	appleTree := AppleTree{5, []int32{-2, 2, 1}}
	orangeTree := OrangeTree{15, []int32{5, -6}}

	countApplesAndOranges(house.startPoint, house.endPoint, appleTree.position, orangeTree.position, appleTree.apples, orangeTree.oranges)
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	appleCounter := 0
	orangeCounter := 0
	for i := 0; i < len(apples); i++ {
		if (a+apples[i]) >= s && (a+apples[i]) <= t {
			appleCounter++
		}
	}

	for i := 0; i < len(oranges); i++ {
		if (b+oranges[i]) <= t && (b+oranges[i]) >= s {
			orangeCounter++
		}
	}

	fmt.Println(appleCounter)
	fmt.Println(orangeCounter)
}
