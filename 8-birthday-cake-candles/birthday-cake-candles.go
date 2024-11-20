package main

import (
	"fmt"
	"math"
)

func main() {
	ar := []int32{1, 2, 3, 4, 4}

	result := birthdayCakeCandles(ar)

	fmt.Println(result)
}

func birthdayCakeCandles(candles []int32) int32 {
	var tallestCandle int32 = math.MinInt32
	counterTallestCandle := 0
	for i := 0; i < len(candles); i++ {
		if candles[i] >= tallestCandle {
			tallestCandle = candles[i]
		}
	}
	for i := 0; i < len(candles); i++ {
		if candles[i] == tallestCandle {
			counterTallestCandle++
		}
	}
	return int32(counterTallestCandle)
}
