package main

import (
	"fmt"
	"slices"
)

/*

 */

func main() {
	fmt.Println("1312")
	/*
		checkT(0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
		checkT(123, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 123)
		checkT(123, []int{0, 9, 2, 3, 4, 5, 6, 7, 8, 9}, 923)
		checkT(121, []int{0, 9, 2, 3, 4, 5, 6, 7, 8, 9}, 929)
		checkT(123, []int{0, 1, 2, 8, 4, 5, 6, 7, 8, 9}, 128)
		checkT(123, []int{0, 1, 5, 3, 4, 5, 6, 7, 8, 9}, 153)
	*/
	check([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}, []int{})
	check([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 2}, []int{2, 3})
	check([]int{0, 1, 2, 2, 4, 5, 6, 7, 8, 9}, []int{3, 2}, []int{3, 2})
	check([]int{0, 2, 5, 5, 4, 5, 6, 7, 8, 9}, []int{3, 2, 1}, []int{1, 3, 2})
	check([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
}

func check(mapping []int, nums []int, expected []int) {
	r := sortJumbled(mapping, nums)
	good := true
	if len(r) != len(expected) {
		good = false
	} else {
		good = true
	}
	if good {
		for i := 0; i < len(r); i++ {
			if r[i] != expected[i] {
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
	fmt.Println(mapping, nums, expected, "RESULT: ", r)
}

func checkT(v int, mapping []int, expected int) {
	r := transform(v, mapping)
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
	fmt.Println(v, mapping, expected, "RESULT: ", r)
}

func transform(v int, mapping []int) int {
	if v == 0 {
		return mapping[v]
	}
	r := 0
	m := 1
	for v > 0 {
		v0 := v % 10
		v /= 10
		r += m * mapping[v0]
		m *= 10
		//fmt.Println(m, r, v0)
	}
	return r
}

func sortJumbled(mapping []int, nums []int) []int {
	slices.SortStableFunc(nums, func(a, b int) int {
		return transform(a, mapping) - transform(b, mapping)
	})
	return nums
}
