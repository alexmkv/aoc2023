package main

import (
	"fmt"
)

/*

 */

func main() {
	fmt.Println("2341")
}

func check(nums []int, expected []int) {
	r := numberOfPairs(nums)
	good := true
	if expected[0] == r[0] && expected[1] == r[1] {
		good = true
	} else {
		good = false
	}
	if good {
		fmt.Print("OK ")
	} else {
		fmt.Print("FAIL ")
	}
	fmt.Println(nums, expected, "RESULT: ", r)
}

func numberOfPairs(nums []int) []int {

}
