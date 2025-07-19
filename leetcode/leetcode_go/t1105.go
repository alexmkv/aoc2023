package main

import (
	"fmt"
)

/*

 */

func main() {
	check([][]int{}, 0, 0)
	check([][]int{{1, 1}}, 1, 1)
	check([][]int{{1, 1}, {1, 2}}, 1, 3)
	check([][]int{{5, 4}, {1, 3}, {5, 3}, {1, 2}, {5, 2}, {1, 1}, {1, 1}}, 6, 4+3+2+1)
	check([][]int{{7, 3}, {8, 7}, {2, 7}, {2, 5}}, 10, 15)
}

func Height(a []int) int {
	return a[1]
}

func Width(a []int) int {
	return a[0]
}

func check(books [][]int, shelfWidth int, expected int) {
	r := minHeightShelves(books, shelfWidth)
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

func minHeightShelves(books [][]int, shelfWidth int) int {
	// on each step we must make decision either put book on current shelve or on new shelves
	// when there is not enough space on current shelve or height of book is <= of shelve decision is obvious
	// in other situation we need to decide between to options
	// solutions looks recursive
	// dynamic programming? We calculate for both options and select minimum one
	// space: O(n), time complexity: O(n)
	// can we improve space complexity?
	// in some cases probably we can't make decision up to final book, so we need to calculate all the value, removing
	// this map will increase time complexity to exponential. We can slightly reduce this map size by some edge case,
	// but not sure that it is worth of it, so i stop here
	return minHeightShelvesFrom(books, shelfWidth, 0, 0, make(map[int]int))
}

func minHeightShelvesFrom(books [][]int, shelfWidth int, shelfWidthLeft int, shelfHeight int, cache map[int]int) int {
	if len(books) == 0 {
		return 0
	}
	if th, ok := cache[len(books)]; ok {
		return th
	}
	b := books[0]
	if shelfWidthLeft >= Width(b) && Height(b) <= shelfHeight {
		o1 := minHeightShelvesFrom(books[1:], shelfWidth, shelfWidth-Width(b), shelfHeight, cache)
		fmt.Println(len(books), o1)
		return o1
	}
	if shelfWidthLeft < Width(b) {
		o1 := minHeightShelvesFrom(books[1:], shelfWidth, shelfWidth-Width(b), Height(b), cache) + Height(b)
		fmt.Println(len(books), o1)
		return o1
	}
	o1 := minHeightShelvesFrom(books[1:], shelfWidth, shelfWidth-Width(b), shelfHeight, cache)
	o2 := minHeightShelvesFrom(books[1:], shelfWidth, shelfWidth-Width(b), Height(b), cache) + Height(b)
	r := min(o1, o2)
	fmt.Println(len(books), r)
	cache[len(books)] = r
	return r
}
