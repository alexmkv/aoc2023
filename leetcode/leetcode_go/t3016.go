package main

import (
	"fmt"
	"slices"
)

/*
 8 keys
lowercase english letters
*/

func main() {
	check("", 0)
	check("a", 1)
	check("aaaaaaaaaaaaaaaa", 16)
	check("abcdefhj", 8)
	check("abcdefhji", 10)
	check("abcdefhjii", 11)
	check("abcdefhjiklmnoprstuvwxyz", 8+16+24)
	check("abcdefhjiklmnoprstuvwxyzabcdefhjiklmnoprstuvwxyzabcdefhjiklmnoprstuvwxyz", (8+16+24)*3)
	check("abcdefhjiklmnoprstuvwxyzabcdefhjiklmnoprstuvwxzabcdefhjiklmnoprstuvwxyz", (8+16+24)*3-3)
	check("abcdefffhjiklmmmnoprstuvwxyz", 8+16+24+4)
}

func check(word string, expected int) {
	r := minimumPushes(word)
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

func minimumPushes(word string) int {
	// get frequency of each letter
	// first 8 letters with top priorities -> we don't even need to map them we add 1 press for push
	// next 8 letters -> 2 pushes and so on
	fr := make([]int, 26)
	for _, ch := range word {
		fr[int(ch)-int('a')]++
	}
	//fmt.Println(fr)
	slices.SortFunc(fr, func(a, b int) int {
		return b - a
	})
	//fmt.Println(fr)
	res := 0
	for i, v := range fr {
		if v == 0 {
			break
		}
		res += v * (i/8 + 1)
		//fmt.Println(res, v, i, (i%8 + 1))
	}
	return res
}
