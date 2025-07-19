package main

import (
	"fmt"
)

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
*/

func main() {
	check([]int{}, 0)
	check([]int{1}, 0)
	check([]int{1, 1}, 0)
	check([]int{1, 0, 2}, 1)
	check([]int{1, 0, 2, 0, 1}, 2)
	check([]int{3, 1, 4, 1, 5}, 2+3)
	check([]int{1, 1, 4, 1, 1}, 0)
	check([]int{1, 1, 4, 5, 1}, 0)
	check([]int{1, 2, 3, 4, 4}, 0)
	check([]int{0, 4, 4, 3, 2, 1, 0}, 0)
}

func check(height []int, expected int) {
	r := trap(height)
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

func trap(height []int) int {
	// we will have two pointers: left and right and
	// we'll keep moving the one with lower level, this way we
	// can calculate possible filled value immediately
	if len(height) <= 2 {
		return 0
	}
	li := 0
	ll := height[li]
	ri := len(height) - 1
	rl := height[ri]
	r := 0
	for ri-li > 1 {
		if ll <= rl {
			li++
			if height[li] >= ll {
				ll = height[li]
			} else {
				r += (ll - height[li])
			}
		} else {
			ri--
			if height[ri] >= rl {
				rl = height[ri]
			} else {
				r += (rl - height[ri])
			}
		}
	}
	return r
}
