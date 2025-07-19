package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	calc("0.txt")
	calc("1.txt")
	// 1108 - low
	calc("2.txt")
}

type LR struct {
	l string
	r string
}

func calc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	path := scanner.Text()
	scanner.Scan()
	mm := make(map[string]LR, 0)
	pp := make[[]string, 0]
	for scanner.Scan() {
		s := scanner.Text()
		//AAA = (BBB, CCC)
		sm := regexp.MustCompile(`(\w+) = \((\w+?), (\w+)\)`).FindStringSubmatch(s)
		//fmt.Println(s)
		mm[sm[1]] = LR{sm[2], sm[3]}
		if sm[1][2] == 'A' {
			pp = append()
		}
	}
		//fmt.Println(mm[sm[1]])
	}
	fmt.Println(path)
	steps := 0
	p := "AAA"

	for {

		if steps > 500000 {
			break
		}
		for _, s := range path {
			steps++
			//fmt.Println(s)
			el := mm[p]
			if s == 'L' {
				p = el.l
			} else {
				p = el.r
			}
			if p == "ZZZ" {
				break
			}
		}
		if p == "ZZZ" {
			break
		}

	}
	fmt.Println(steps)
}
