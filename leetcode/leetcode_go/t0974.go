package main

import (
	"fmt"
)

/*
Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.

A subarray is a contiguous part of an array.
*/

func main() {
	fmt.Println("1312")
	check([]int{}, 5, 0)
	check([]int{1}, 5, 0)
	check([]int{1, 2, 1, 1}, 5, 0)
	check([]int{5, 1, 4}, 5, 3)
	check([]int{1, 5, 1, 4, 4}, 5, 4)
	check([]int{1, 1}, 2, 1)
	check([]int{-1, -1}, 2, 1)
	check([]int{-1, -1}, 2, 1)
	check([]int{-10}, 5, 1)
	check([]int{-1, 6, 4}, 5, 2)
	check([]int{-10, 5}, 5, 3)
	check([]int{-15, 5, 1, 4, 15}, 5, 10)
	check([]int{-3, 2, 1}, 5, 1)
	big := make([]int, 10000)
	for i, _ := range big {
		big[i] = i
	}
	//check(big, 5, 17999000)
	/*check("", 0)
	check("a", 0)
	check("aa", 0)
	check("aba", 0)
	check("abca", 1)
	check("ab", 1)
	check("abcc", 2)
	check("abzcdeffedcba", 1)*/
}

func check(nums []int, k int, expected int) {
	r := subarraysDivByK(nums, k)
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
	fmt.Println(nums, k, expected, "RESULT: ", r)
}

func norm(v, k int) int {
	nn := v % k
	if nn < 0 {
		nn = k + nn
	}
	return nn
}

func subarraysDivByK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	pr := make([]int, len(nums))
	cn := make(map[int]int)
	pr[0] = nums[0]
	cn[norm(pr[0], k)] = 1
	for i := 1; i < len(nums); i++ {
		pr[i] = pr[i-1] + nums[i]
		cn[norm(pr[i], k)] += 1
	}
	//fmt.Println(cn)
	res := 0
	for i, v := range cn {
		// v = 1 -> 0
		// v = 2 -> 1
		// v = 3 -> 2 + 1 = 3
		// v = 4 -> 3 + 2 + 1 = 6
		if i == 0 {
			v += 1
		}
		if v > 1 {
			res += ((v - 1) + 1) * (v - 1) / 2
		}
	}
	return res
}
