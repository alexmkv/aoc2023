package main

import (
	"fmt"
)

/*

 */

func main() {
	check([]int{}, []int{})
	check([]int{1}, []int{1})
	check([]int{1, 2}, []int{1, 2})
	check([]int{2, 1}, []int{1, 2})
	check([]int{3, 2, 1}, []int{1, 2, 3})
	check([]int{2, 3, 1}, []int{1, 2, 3})
	check([]int{1, 2, 3, 4, 5, 11, 10, 8, 9, 7, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	check([]int{1, 2, 3, 5, 4, 11, 10, 8, 9, 7, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	check([]int{1, 2, 3, 5, 4, 11, 10, 8, 9, 7, 6, 12}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})

}

func check(nums []int, expected []int) {
	r := sortArray(nums)
	good := true
	if len(r) == len(expected) {
		for i, v := range r {
			if v != expected[i] {
				good = false
				break
			}
		}
	} else {
		good = false
	}

	if good {
		fmt.Print("OK ")
	} else {
		fmt.Print("FAIL ")
	}
	fmt.Println(expected, "RESULT: ", r)
}

func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	heapify(nums)
	//fmt.Println(nums)
	for s := len(nums) - 1; s > 0; s-- {
		nums[s], nums[0] = nums[0], nums[s]
		siftDown(nums[:s], 0)
	}
}

func heapify(nums []int) {
	for st := parent(len(nums) - 1); st >= 0; st-- {
		siftDown(nums, st)
	}
}

func siftDown(nums []int, a int) {
	ch1 := child(a, 0)
	ch2 := ch1 + 1
	if ch1 >= len(nums) {
		return
	}
	if ch2 >= len(nums) {
		if nums[a] < nums[ch1] {
			nums[a], nums[ch1] = nums[ch1], nums[a]
			siftDown(nums, ch1)
		}
	} else {
		if nums[a] < nums[ch1] && (nums[ch1] >= nums[ch2] || nums[a] >= nums[ch2]) {
			nums[a], nums[ch1] = nums[ch1], nums[a]
			siftDown(nums, ch1)
		} else if nums[a] < nums[ch2] {
			nums[a], nums[ch2] = nums[ch2], nums[a]
			siftDown(nums, ch2)
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func child(i int, p int) int {
	return i*2 + 1 + p
}

func sortArray2(nums []int) []int {
	r := make([]int, len(nums))
	for i, v := range nums {
		r[i] = v
	}
	mergeSort(nums, 0, len(nums), r)
	return nums
}

func mergeSort(nums []int, i, j int, r []int) {
	if j-i <= 1 {
		return
	}
	m := (j + i) / 2
	mergeSort(r, i, m, nums)
	mergeSort(r, m, j, nums)
	a := i
	b := m
	for k := i; k < j; k++ {
		if a >= m || (b < j && r[b] < r[a]) {
			nums[k] = r[b]
			b++
		} else {
			nums[k] = r[a]
			a++
		}
	}
}

/**
1 4 5 2 3 6
k= 0, a =0 , b = 3
1 4 5 2 3 6
k= 1, a=2 , b = 3, a0=3, a_len = 0
1 2 5 4 3 6
k= 2, a =2 , b = 4, , a0=3, a_len = 1
1 2 3 4 5 6
k= 3, a =3 , b = 4, , a0=3, a_len = 0


2(a,k) 3 0(bs, b, a_s, a_e) 1 .
0 3(a,k) 2(bs,a_s) 1(b,a_e) .
0 1 2(a,k, bs,a_s) 3(b) .(a_e)

2(a,k) 3 0(bs, b, a_s, a_e) 5 .
0 3(a,k) 2(bs,a_s) 5(b,a_e) .
0 3(a,k) 2(bs,a_s) 5(b,a_e) .

*/
