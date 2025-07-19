package main

import (
	"fmt"
)

/*

 */

func main() {
	check(0)
}

func check(expected int) {
	r := f()
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

func f() int {
	return 0
}
