package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2023/day/1
// 54431
func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	for i := 0; i < 10; i++ {
		m[strconv.Itoa(i)] = i
	}
	re := "("
	idx := len(m)
	for k, _ := range m {
		idx--
		if idx == 0 {
			re += k
		} else {
			re += k + "|"
		}
	}
	re += ")"
	fmt.Println(re)
	file, err := os.Open("puzzles/01_21")
	if err != nil {
		panic(err)
	}

	res := regexp.MustCompile("^.*?" + re)
	ree := regexp.MustCompile("^.*" + re)
	scanner := bufio.NewScanner(file)
	var sum int = 0
	for scanner.Scan() {
		s := scanner.Text()
		//matched, err := regexp.Match("^.*?("+re+")", []byte(s))
		//fmt.Println(s)
		ress := res.FindStringSubmatch(s)[1]
		rese := ree.FindStringSubmatch(s)[1]
		//fmt.Println(ress)
		f := m[ress]
		l := m[rese]
		//fmt.Println(ress)
		if err != nil {
			panic(err)
		}

		sum += f*10 + l
		//fmt.Println(f)
		//fmt.Println(l)
		//fmt.Println(f - '0')
		//fmt.Println(l - '0')
		//fmt.Println(s)
	}
	fmt.Println(sum)

	//one, two, three, four, five, six, seven, eight, and nine
}
