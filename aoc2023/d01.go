package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2023/day/1
// 54431
func main() {
	if 0 == 1 {
		main1()
	}
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
		fmt.Println(s)
		ress := res.FindStringSubmatch(s)[1]
		rese := ree.FindStringSubmatch(s)[1]
		//fmt.Println(ress)
		f := m[ress]
		l := m[rese]
		//fmt.Println(ress)
		if err != nil {
			panic(err)
		}

		/*fmt.Println(matched, err)

		  f -= '0'
		  l -= '0'*/
		sum += f*10 + l
		fmt.Println(f)
		fmt.Println(l)
		//fmt.Println(f - '0')
		//fmt.Println(l - '0')
		//fmt.Println(s)
	}
	fmt.Println(sum)

	//one, two, three, four, five, six, seven, eight, and nine
}

func main1() {
	file, err := os.Open("puzzles/01_1")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var sum int32 = 0
	for scanner.Scan() {
		s := scanner.Text()
		f := 'a'
		l := 'a'
		for _, ch := range s {
			if ch >= '0' && ch <= '9' {
				if f == 'a' {
					f = ch
				}
				l = ch
			}
		}
		if l == 'a' {
			panic(" l == a")
		}
		f -= '0'
		l -= '0'
		sum += f*10 + l
		//fmt.Println(f - '0')
		//fmt.Println(l - '0')
		//fmt.Println(s)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
