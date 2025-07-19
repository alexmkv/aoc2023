package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	calc("0.txt")
	calc("1.txt")
}

type El struct {
	dir string
	n   int64
	clr string
}

type Pt struct {
	x, y int64
}

func calc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	scn := bufio.NewScanner(f)
	els := []El{}
	for scn.Scan() {
		//fmt.Println(scn.Text())
		res := regexp.MustCompile(`(.) ([0-9]+) \(#(.*)\)`).FindStringSubmatch(scn.Text())
		e := El{dir: res[1], clr: res[3]}
		switch e.clr[len(e.clr)-1] {
		case '0':
			e.dir = "R"
		case '1':
			e.dir = "D"
		case '2':
			e.dir = "L"
		case '3':
			e.dir = "U"
		}
		nums := e.clr[:len(e.clr)-1]
		e.n, _ = strconv.ParseInt(nums, 16, 64)
		fmt.Println(e.clr, e.dir, e.n)
		//e.n, _ = strconv.ParseInt(res[2], 10, 64)
		els = append(els, e)
	}
	//fmt.Println(els)
	p := Pt{0, 0}
	ps := []Pt{p}
	l := int64(0)
	var minx, maxx, miny, maxy int64
	for _, e := range els {
		switch e.dir {
		case "R":
			p.x += e.n
			l += e.n
		case "L":
			p.x -= e.n
			l += e.n
		case "U":
			p.y -= e.n
			l += e.n
		case "D":
			p.y += e.n
			l += e.n
		}
		//fmt.Println(e)
		//fmt.Println(p)
		ps = append(ps, p)
		minx = min(minx, p.x)
		miny = min(miny, p.y)
		maxy = max(maxy, p.y)
		maxx = max(maxx, p.x)
	}
	s := int64(0)
	for i := 1; i < len(ps); i++ {
		//s += (ps[i-1].y + ps[i].y) * (ps[i-1].x - ps[i].x)
		s += (ps[i-1].x * ps[i].y) - (ps[i].x * ps[i-1].y)
	}
	fmt.Println(l / 2)
	fmt.Println(s/2 + l/2 + 1)
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
