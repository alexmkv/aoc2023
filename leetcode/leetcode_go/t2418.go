package main

import (
	"fmt"
	"slices"
)

/*
You are given an array of strings names, and an array heights that consists of distinct positive integers. Both arrays are of length n.

For each index i, names[i] and heights[i] denote the name and height of the ith person.

Return names sorted in descending order by the people's heights.
*/

func main() {
	fmt.Println("t2418")
	check([]string{}, []int{}, []string{})
	check([]string{"a"}, []int{1}, []string{"a"})
	check([]string{"b", "a"}, []int{30, 20}, []string{"b", "a"})
	check([]string{"b", "a"}, []int{20, 30}, []string{"a", "b"})
	check([]string{"b", "a", "c"}, []int{20, 30, 100}, []string{"c", "a", "b"})
}

func check(names []string, heights []int, expected []string) {
	r := sortPeople(names, heights)
	good := true
	if len(expected) != len(r) {
		good = false
	} else {
		for i := 0; i < len(r); i++ {
			if expected[i] != r[i] {
				good = false
				break
			}
		}
	}
	if good {
		fmt.Print("OK")
	} else {
		fmt.Print("FAIL")
	}
	fmt.Println(names, heights, expected, "RESULT: ", r)
}

func sortPeople(names []string, heights []int) []string {
	r := make([]int, len(names))
	for i, _ := range r {
		r[i] = i
	}
	slices.SortFunc(r, func(a, b int) int {
		return heights[b] - heights[a]
	})
	res := make([]string, len(r))
	for i, _ := range r {
		res[i] = names[r[i]]
	}
	return res
}
