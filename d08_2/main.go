package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	//calc("0.txt")
	//calc("1.txt")
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
	pp := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		//AAA = (BBB, CCC)
		sm := regexp.MustCompile(`(\w+) = \((\w+?), (\w+)\)`).FindStringSubmatch(s)
		//fmt.Println(s)
		mm[sm[1]] = LR{sm[2], sm[3]}
		if sm[1][2] == 'A' {
			pp = append(pp, sm[1])
		}
		//fmt.Println(mm[sm[1]])
	}
	fmt.Println(path)

	var nextP = func(p string, s int64) string {
		i := s % int64(len(path))
		el := mm[p]
		if path[i] == 'L' {
			p = el.l
		} else {
			p = el.r
		}
		return p
	}

	var findCycle = func(p string) (int64, int64) {
		i1 := int64(0)
		i2 := int64(0)
		p1 := p
		p2 := p
		// LRRRLRRLLRRLRRLRRLRRLRLLRLRLLRRLRLRRRLRRLRRLLRLRLRLRRRLRRRLLRLRRRLLRRRLRLLRRRLLRRLRLRLRRRLLRRLRRRLLRRLRLRRRLLRRRLRRLRLRRRLLRRLRRRLRRLLRRLRRLRRRLRRRLRRRLRRLRRRLLRLRLRLRRRLRRLRRRLRRLRLRRLRLRRRLRRRLRRLRRRLLRRRLLRRLRLRRRLRLRLRRRLRLRLRLRRLRLRRLRRLLRRRLRLLRRLRRRLRRRLLRRLRLLLLRRLRRRR
		//fmt.Println(len(path))
		for {
			p1 = nextP(p1, i1)
			i1++
			p2 = nextP(p2, i2)
			i2++
			p2 = nextP(p2, i2)
			i2++
			if i1 > 100000 {
				fmt.Println(`Too many cycles`)
				break
			}
			//if p1 == p2 && i1 >= int64(len(path)) && p2[2] == 'Z' {
			if p1 == p2 && i1%int64(len(path)) == i2%int64(len(path)) {
				//if p2[2] == 'Z' {

				fmt.Println(`Cycle found`, i1)
				break
			}
			//}
		}
		return i1, i2
	}

	res := int64(0)
	for _, p := range pp {
		//fmt.Println(findCycle(p))
		r, _ := findCycle(p)
		if res == 0 {
			res = r
		} else {
			res = LCM(r, res)
		}
		fmt.Println(res)
		//break
	}
	fmt.Println(res)

	//}
	//fmt.Println(steps)
}

func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
