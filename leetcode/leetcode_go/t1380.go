package t1380

import (
	"fmt"
	"slices"
)

/*
Given an m x n matrix of distinct numbers, return all lucky numbers in the matrix in any order.

A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.
*/

func main() {
	fmt.Println("1380")
	check(nil, []int{})
	check([][]int{{1}}, []int{1})
	check([][]int{{20, 10}, {1, 2}}, []int{10})
}

func check(matix [][]int, expected []int) {
	r := luckyNumbers(matix)
	slices.Sort(expected)
	slices.Sort(r)
	good := true
	if len(expected) != len(r) {
		good = false
	} else {
		for i := 0; i < len(r); i++ {
			if expected[i] != r[i] {
				fmt.Println("fail n ", i, expected[i], r[i])
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
	fmt.Println(matix, expected, "RESULT: ", r)
}

func luckyNumbers(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	r := make([]int, 0)
	mxc := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		mxc[i] = matrix[0][i]
		for j := 1; j < len(matrix); j++ {
			mxc[i] = max(mxc[i], matrix[j][i])
		}
	}
	for i := 0; i < len(matrix); i++ {
		mn := matrix[i][0]
		p := 0
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < mn {
				p = j
				mn = matrix[i][j]
			}
		}
		if matrix[i][p] == mxc[p] {
			r = append(r, mxc[p])
		}
	}
	return r
}
