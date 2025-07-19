package main

import (
	"bufio"
	"fmt"
	"os"
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
	sc := bufio.NewScanner(f)
	a := make([]string, 0)
	sum := int64(0)
	for sc.Scan() {
		s := sc.Text()
		if s == "" {
			sum += proc(a)
			a = make([]string, 0)
			continue
		}
		a = append(a, s)
	}
	sum += proc(a)
	fmt.Println(sum)
}

func proc(a []string) int64 {
	fmt.Println(a)
	for y := 0; y < len(a)-1; y++ {
		if is_hor(a, y) {
			return int64(y+1) * 100
		}
	}
	for x := 0; x < len(a[0])-1; x++ {
		if is_ver(a, x) {
			return int64(x+1) * 1
		}
	}
	fmt.Println(a)
	panic("not found mirror")

	return 0
}

func is_hor(a []string, y int) bool {
	for i := 0; ; i++ {
		y0 := y - i
		y1 := y + 1 + i
		if y0 < 0 {
			return true
		}
		if y1 >= len(a) {
			return true
		}
		for x := 0; x < len(a[0]); x++ {
			if a[y0][x] != a[y1][x] {
				return false
			}
		}
	}
	panic("unexpected is_hor")
}

func is_ver(a []string, x int) bool {
	for i := 0; ; i++ {
		x0 := x - i
		x1 := x + 1 + i
		if x0 < 0 {
			return true
		}
		if x1 >= len(a[0]) {
			return true
		}
		for y := 0; y < len(a); y++ {
			if a[y][x0] != a[y][x1] {
				return false
			}
		}
	}
	panic("unexpected is_hor")
}
