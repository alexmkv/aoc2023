package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	calc("0.txt")
	calc("1.txt")
}

const (
	DIR_N = 1
	DIR_E = 1 << 1
	DIR_W = 1 << 2
	DIR_S = 1 << 3
)

func calc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	scn := bufio.NewScanner(f)
	fld := createField(scn)
	fld.passLight(0, 0, DIR_E)
	fmt.Println(fld.calcEnerg())
}

type Field struct {
	f []string
	m [][]int
}

func createField(sc *bufio.Scanner) *Field {
	fld := Field{}
	for sc.Scan() {
		t := sc.Text()
		fld.f = append(fld.f, t)
		fld.m = append(fld.m, make([]int, len(t)))
	}
	return &fld
}

func (fld *Field) calcEnerg() int {
	sum := 0
	for _, l := range fld.m {
		for _, v := range l {
			if v != 0 {
				sum++
			}
		}
	}
	return sum
}

func (fld *Field) passLight(x, y int, dir int) {
	if x < 0 || y < 0 {
		return
	}
	if x >= len(fld.f[0]) || y >= len(fld.f) {
		return
	}
	if fld.m[y][x]&dir != 0 {
		return
	}
	(*fld).m[y][x] = (*fld).m[y][x] | dir
	switch fld.f[y][x] {
	case '.':
		switch dir {
		case DIR_N:
			fld.passLight(x, y-1, dir)
		case DIR_S:
			fld.passLight(x, y+1, dir)
		case DIR_W:
			fld.passLight(x-1, y, dir)
		case DIR_E:
			fld.passLight(x+1, y, dir)
		}
	case '/':
		switch dir {
		case DIR_N:
			fld.passLight(x+1, y, DIR_E)
		case DIR_S:
			fld.passLight(x-1, y, DIR_W)
		case DIR_W:
			fld.passLight(x, y+1, DIR_S)
		case DIR_E:
			fld.passLight(x, y-1, DIR_N)
		}
	case '\\':
		switch dir {
		case DIR_S:
			fld.passLight(x+1, y, DIR_E)
		case DIR_N:
			fld.passLight(x-1, y, DIR_W)
		case DIR_E:
			fld.passLight(x, y+1, DIR_S)
		case DIR_W:
			fld.passLight(x, y-1, DIR_N)
		}
	case '|':
		switch dir {
		case DIR_N:
			fld.passLight(x, y-1, dir)
		case DIR_S:
			fld.passLight(x, y+1, dir)
		case DIR_E:
			fld.passLight(x, y+1, DIR_S)
			fld.passLight(x, y-1, DIR_N)
		case DIR_W:
			fld.passLight(x, y-1, DIR_N)
			fld.passLight(x, y+1, DIR_S)
		}
	case '-':
		switch dir {
		case DIR_S:
			fld.passLight(x+1, y, DIR_E)
			fld.passLight(x-1, y, DIR_W)
		case DIR_N:
			fld.passLight(x+1, y, DIR_E)
			fld.passLight(x-1, y, DIR_W)
		case DIR_W:
			fld.passLight(x-1, y, dir)
		case DIR_E:
			fld.passLight(x+1, y, dir)
		}
	}
}
