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
	n   int
	clr string
}

type Pt struct {
	x, y int
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
		e.n, _ = strconv.Atoi(res[2])
		els = append(els, e)
	}
	//fmt.Println(els)
	p := Pt{0, 0}
	ps := []Pt{p}
	var minx, maxx, miny, maxy int
	for _, e := range els {
		switch e.dir {
		case "R":
			p.x += e.n
		case "L":
			p.x -= e.n
		case "U":
			p.y -= e.n
		case "D":
			p.y += e.n
		}
		//fmt.Println(e)
		//fmt.Println(p)
		ps = append(ps, p)
		minx = min(minx, p.x)
		miny = min(miny, p.y)
		maxy = max(maxy, p.y)
		maxx = max(maxx, p.x)
	}
	fmt.Println(minx, miny, maxx, maxy)
	H := maxy - miny + 3
	W := maxx - minx + 3
	fmt.Println(ps[len(ps)-1])
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = make([]int, W)
	}
	p = Pt{0 - minx + 1, 0 - miny + 1}
	fmt.Println(p)
	a[p.y][p.x] = 2
	for _, e := range els {
		switch e.dir {
		case "R":
			for i := 1; i <= e.n; i++ {
				a[p.y][p.x+i] = 2
			}
			p.x += e.n

		case "L":
			for i := 1; i <= e.n; i++ {
				a[p.y][p.x-i] = 2
			}
			p.x -= e.n
		case "U":
			for i := 1; i <= e.n; i++ {
				a[p.y-i][p.x] = 2
			}
			p.y -= e.n
		case "D":
			for i := 1; i <= e.n; i++ {
				a[p.y+i][p.x] = 2
			}
			p.y += e.n
		}
		//fmt.Println(e)
		//fmt.Println(p)
		//ps = append(ps, p)
	}
	chk := []Pt{Pt{0, 0}}
	for len(chk) > 0 {
		p = chk[len(chk)-1]
		chk = chk[:len(chk)-1]
		if p.x < 0 || p.y < 0 || p.x >= W || p.y >= H {
			continue
		}
		if a[p.y][p.x] != 0 {
			continue
		}
		a[p.y][p.x] = 1
		chk = append(chk, Pt{p.x - 1, p.y})
		chk = append(chk, Pt{p.x + 1, p.y})
		chk = append(chk, Pt{p.x, p.y - 1})
		chk = append(chk, Pt{p.x, p.y + 1})
	}
	sm := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			fmt.Print(a[y][x])
			if a[y][x] != 1 {
				sm += 1
			}
		}
		fmt.Println()
	}
	fmt.Println(sm)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
