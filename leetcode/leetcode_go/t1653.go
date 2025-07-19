package main

import (
	"fmt"
)

/*
You are given a string s consisting only of characters 'a' and 'b'.
You can delete any number of characters in s to make s balanced. s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.
Return the minimum number of deletions needed to make s balanced.
*/

func main() {
	check("", 0)
	check("a", 0)
	check("b", 0)
	check("ab", 0)
	check("ba", 1)
	check("bba", 1)
	check("bbbbbabb", 1)
	check("aaaabaaa", 1)
	check("baaaabbaaabbbaa", 5)
	// what options do we have:
	// we can quickly check two options: delete all a and all b (just count of them) -> that will balance string
	// for other cases we need to find middle point
}

func check(s string, expected int) {
	r := minimumDeletions(s)
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
	fmt.Println(s, expected, "RESULT: ", r)
}

func minimumDeletions0(s string) int {
	a_to_del := 0
	b_to_del := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a_to_del++
		}
	}
	if a_to_del == 0 || a_to_del == len(s) {
		return 0
	}
	min_del := a_to_del
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a_to_del--
		} else {
			b_to_del++
		}
		if min_del > a_to_del+b_to_del {
			min_del = a_to_del + b_to_del
		}
	}
	return min_del
}

func minimumDeletions(s string) int {
	res := 0
	b_count := 0
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]), s[:i+1], res, b_count)
		if s[i] == 'a' {
			res = min(res+1, b_count)
		} else {
			b_count++
		}
		fmt.Println(s[:i+1], res, b_count)
	}
	return res
}
