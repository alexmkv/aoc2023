package main

import (
	"fmt"
)

/*

 */

func main() {
	check([]int{0}, 1, 1, 1, 0)
	check([]int{1, 2}, 2, 1, 1, 0)
}

func check(nums []int, n int, left int, right int, expected int) {
	r := f()
	good := true
	if r != expected {
		good = false
	} else {
		good = true
	}

	if good {
		fmt.Print("OK ")
	} else {
		fmt.Print("FAIL ")
	}
	fmt.Println(expected, "RESULT: ", r)
}

func rangeSum(nums []int, n int, left int, right int) int {
	// leftmost - is smallest element, rightmost - sum of all elements
	// we can do something with prefix sums
	// prefix sum are obviously increasing
	// rightmost : prefsum[n-1]
	// rightmost - 1, prefsum[n-
}
