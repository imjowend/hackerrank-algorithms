package main

import "fmt"

func main() {
	grade := []int32{67, 75, 37, 39}
	fmt.Println(grades(grade))
}

func grades(grades []int32) []int32 {
	var roundedGrade []int32
	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			roundedGrade = append(roundedGrade, grades[i])
			continue
		}
		if !((grades[i] % 5) < 3) {
			newGrade := grades[i] + grades[i]%5
			roundedGrade = append(roundedGrade, newGrade)
			continue
		}
		roundedGrade = append(roundedGrade, grades[i])
	}

	return roundedGrade
}
