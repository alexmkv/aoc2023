package main

import (
	"fmt"
)

/*

 */

func main() {
	fmt.Println("1380")
	check([]int{}, []int{})
	check([]int{1}, []int{1})
	check2([]int{1, 1, 2, 3, 4, 5, 6}, []int{22})
	check2([]int{2, 4, 6}, []int{12})
	check2([]int{2, 4, 6}, []int{6, 6})
	check2([]int{3, 8}, []int{4, 7})
	//check([][]int{{20, 10}, {1, 2}}, []int{10})
}

func check(k int, rowConditions [][]int, colConditions [][]int) {
	r := buildMatrix(rowSum, colSum)
	good := true
	if len(rowSum) != len(r) {
		good = false
	} else {
		for i := 0; i < len(r); i++ {
			s := 0
			for j := 0; j < len(r[i]); j++ {
				s += r[i][j]
			}
			if rowSum[i] != s {
				good = false
				break
			}
		}
	}
	if len(r) == 0 {
		if len(colSum) != 0 {
			good = false
		}
	} else if len(colSum) != len(r[0]) {
		good = false
	} else {
		for i := 0; i < len(r[0]); i++ {
			s := 0
			for j := 0; j < len(r); j++ {
				s += r[j][i]
			}
			if colSum[i] != s {
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
	fmt.Println(rowSum, colSum, "RESULT: ", r)
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {

}
