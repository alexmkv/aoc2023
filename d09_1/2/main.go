package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//calc("0.txt")
	calc("1.txt")
}

func calc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sum := 0
	i := 0
	for sc.Scan() {
		i++
		if i == -1 {
			break
		}
		s := sc.Text()
		rr := regexp.MustCompile(`[\-0-9]+`).FindAllString(s, -1)
		vals := make([]int, 0)
		for _, rrs := range rr {
			v, _ := strconv.Atoi(rrs)
			vals = append(vals, v)
		}
		nxt := next(vals)
		//fmt.Println(nxt)
		sum += nxt
	}
	fmt.Println(sum)
}

func next(vals []int) int {
	if len(vals) == 0 {
		panic(`len is 0`)
	}
	all_zero := true
	diffs := make([]int, 0)
	prev := 0
	for i, v := range vals {
		if v != 0 {
			all_zero = false
		}
		if i == 0 {
			prev = v
		} else {
			diffs = append(diffs, v-prev)
			prev = v
		}
	}
	if all_zero {
		return 0
	}

	//fmt.Println(diffs)
	nxt := next(diffs)
	//fmt.Println(diffs)
	//fmt.Println(nxt)
	//return vals[len(vals)-1] + nxt
	return vals[0] - nxt
}
