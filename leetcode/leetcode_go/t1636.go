package main

import (
	"fmt"
	"slices"
)

/*
Given an array of integers nums, sort the array in increasing order based on the frequency of the values. If multiple values have the same frequency, sort them in decreasing order.

Return the sorted array.
*/

func main() {
	fmt.Println("1636")
	check([]int{}, []int{})
	check([]int{1, 2, 3}, []int{3, 2, 1})
	check([]int{1, 2, 2, 3}, []int{3, 1, 2, 2})
	check([]int{1, 2, 2, 3, 3}, []int{1, 3, 3, 2, 2})
}

func check(nums []int, expected []int) {
	r := frequencySort(nums)
	good := true
	if len(r) != len(expected) {
		good = false
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] != expected[i] {
				good = false
				break
			}
		}
	}

	if good {
		fmt.Print("OK ")
	} else {
		fmt.Print("FAIL ")
	}
	fmt.Println(nums, expected, "RESULT: ", r)
}

func frequencySort(nums []int) []int {
	fr := make(map[int]int)
	for _, v := range nums {
		fr[v]++
	}
	slices.SortFunc(nums, func(a int, b int) int {
		fr1 := fr[a]
		fr2 := fr[b]
		if fr1 != fr2 {
			return fr1 - fr2
		}
		return b - a
	})
	return nums
}
