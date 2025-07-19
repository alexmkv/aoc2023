package main

import (
	"fmt"
)

/*

 */

func main() {
	fmt.Println("1312")
	/*check("", 0)
	check("a", 0)
	check("aa", 0)
	check("aba", 0)
	check("abca", 1)
	check("ab", 1)
	check("abcc", 2)
	check("abzcdeffedcba", 1)*/
	check("abcdeffedczba", 1)
	check("abcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczba"+
		"abcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczbaabcdeffedczba", 22)
}

func check(s string, expected int) {
	r := minInsertions(s)
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

func minInsertions(s string) int {
	// check that we already have palindrome
	// if it is so -> 0
	// we maintain two pointer on start and end of the string
	// up to the moment when difference between them is <= 1
	// on each step if chars on pointer are same we reduce both of
	// them. if not we reduce each of them separatly and return mininum
	// of calculation between them
	cache := make(map[string]int)
	return minInsertionC(s, 0, len(s)-1, cache)
}

func minInsertionC(s string, i, j int, cache map[string]int) int {
	if j <= i {
		return 0
	}
	//fmt.Println(i, j, s[i:j+1])
	v, ok := cache[s[i:j+1]]
	if ok {
		return v
	}
	if s[i] == s[j] {
		return minInsertionC(s, i+1, j-1, cache)
	}

	v1 := minInsertionC(s, i+1, j, cache)
	v2 := minInsertionC(s, i, j-1, cache)
	r := min(v1+1, v2+1)
	cache[s[i:j+1]] = r
	return r
}
