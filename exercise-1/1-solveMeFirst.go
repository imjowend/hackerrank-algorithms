/*
Function Description

Complete the solveMeFirst function in the editor below.

solveMeFirst has the following parameters:

int a: the first value
int b: the second value
Returns
- int: the sum of a and b

constraints
1<= a,b <= 1000
*/
package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	if a >= 1 && a <= 1000 {
		if b >= 1 && b <= 1000 {
			return (a + b)
		}
	}
	return uint32(0)
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	if res == 0 {
		fmt.Println("Error with constraints")
	}
	fmt.Println(res)
}
