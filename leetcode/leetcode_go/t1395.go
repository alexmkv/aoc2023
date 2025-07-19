package main

import (
	"fmt"
)

/*

 */

func main() {
	check([]int{1, 2, 3}, 1)
	check([]int{3, 2, 1}, 1)
	check([]int{3, 2, 4}, 0)
	check([]int{1, 2, 6, 5, 4, 7, 8}, (5+2+2+2+1)+(2+2+2+1)+(1+0+1)+(1)+1)
	// 1,2,3,4,5,6,7,8
	// if we try to do it for 2 elements:
	// straightforward way is to count each element
	// start from 0th element, and check how many elements are greater than
	// for 2 elements it doesn't make any sense it is always possible to form such pairs
	// so for 3 elements we will have 3 inner loups and check if condition is match
	// second element relation to first element will define relation to third element
	// will get complexity O(n^3)

	// first optimization: we can try to select middle element and count number of elements,
	// that is greater than middle one smaller than it. Obviously total_count = smaller+greater
	// we can do it for each side. Than for one specific middle number
	// number of combination will be smaller_left*greater_right + greater_left*smaller_right
	// Complexity: O(n^2)
	// with n = 1000 it can work with such complexity

	// can we do something better?
	// can we maintain number of left/right smaller&greater elements through going through

	// idea - we sort sequence complexity N*logN
	// and using binary search will find position of our element in sorted sequence O(logN)
	// so we immediatily get count of smaller & greater elements
	// as we need to get N searched it is getting us O(N*logN)
	// but it wouldn't get us number of greater and smaller relative to current position
	// it is total number of elements but it isn't clearer are they on left side or on right side
}

func check(rating []int, expected int) {
	r := numTeams(rating)
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

func numTeams(rating []int) int {
	res := 0
	for i := 1; i < len(rating)-1; i++ {
		v := rating[i]
		left_smaller := 0
		left_bigger := 0
		for j := 0; j < i; j++ {
			if v > rating[j] {
				left_smaller++
			} else {
				left_bigger++
			}
		}
		right_bigger := 0
		right_smaller := 0
		for j := i + 1; j < len(rating); j++ {
			if v < rating[j] {
				right_bigger++
			} else {
				right_smaller++
			}
		}
		//fmt.Println(res)
		res += (left_smaller * right_bigger) + (left_bigger * right_smaller)
		//fmt.Println(i, left_smaller, left_bigger, right_smaller, right_bigger, res)
	}
	return res
}
