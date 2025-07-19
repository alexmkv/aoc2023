package main

import (
	"fmt"
)

/*

 */

func main() {
	check([]int{}, []int{}, true)
	check([]int{1}, []int{1}, true)
	check([]int{1, 2}, []int{2, 1}, true)
	check([]int{1, 3, 2}, []int{2, 1, 4}, false)
	check([]int{1, 2, 2}, []int{1, 1, 2}, false)
}

func check(target []int, arr []int, expected bool) {
	r := canBeEqual(target, arr)
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

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	f := make(map[int]int)
	for _, v := range target {
		f[v]++
	}
	for _, v := range arr {
		fc, ok := f[v]
		if !ok || fc <= 0 {
			return false
		}
		if fc == 1 {
			delete(f, v)
		} else {
			f[v] = fc - 1
		}
	}
	return true
}
