package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	do("0.txt")
	do("1.txt")
}

func do(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	a := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		a = append(a, sc.Text())
	}
	fld := makeField(a)
	fld2 := makeField(a)
	i := int64(0)
	max := int64(1000000000)
	first := int64(0)
	second := int64(0)
	for ; ; i++ {

		rotate(fld)
		rotate(fld2)
		rotate(fld2)
		if compareFields(fld, fld2) {
			if first == 0 {
				first = i
			} else if second == 0 {
				second = i
				cl := second - first
				max = ((max - first) % cl) + first + 2*cl - 1
			}
		}
		if i == max {
			break
		}
	}
	//fmt.Println(i)
	//fmt.Println(compareFields(fld, fld2))
	fmt.Println(fld.calcLoad())
}

func rotate(fld *Field) {
	fld.moveRocksN()
	fld.moveRocksW()
	fld.moveRocksS()
	fld.moveRocksE()
}

func compareFields(f *Field, f2 *Field) bool {
	for i, l := range f.f {
		for j, r := range l {
			if r != f2.f[i][j] {
				return false
			}
		}
	}
	return true
}

type Field struct {
	f [][]rune
}

func (f *Field) calcLoad() int64 {
	cnt := len(f.f)
	sum := int64(0)
	for i, l := range f.f {
		sl := int64(0)
		for _, r := range l {
			if r == 'O' {
				sl += 1
			}
		}
		sum += sl * int64(cnt-i)
	}
	return sum
}

func (f *Field) moveRocksN() {
	W := len(f.f[0])
	H := len(f.f)
	for x := 0; x < W; x++ {
		lastDot := 0
		for y := 0; y < H; y++ {
			r := f.f[y][x]
			if r == 'O' {
				if lastDot != y {
					f.f[lastDot][x] = 'O'
					f.f[y][x] = '.'
					lastDot = lastDot + 1
				} else {
					lastDot = y + 1
				}
			} else if r == '#' {
				lastDot = y + 1
			}
		}
	}
}

func (f *Field) moveRocksS() {
	W := len(f.f[0])
	H := len(f.f)
	for x := 0; x < W; x++ {
		lastDot := H - 1
		for y := H - 1; y >= 0; y-- {
			r := f.f[y][x]
			if r == 'O' {
				if lastDot != y {
					f.f[lastDot][x] = 'O'
					f.f[y][x] = '.'
					lastDot = lastDot - 1
				} else {
					lastDot = y - 1
				}
			} else if r == '#' {
				lastDot = y - 1
			}
		}
	}
}

func (f *Field) moveRocksW() {
	W := len(f.f[0])
	H := len(f.f)
	for y := 0; y < H; y++ {
		lastDot := 0
		for x := 0; x < W; x++ {
			r := f.f[y][x]
			if r == 'O' {
				if lastDot != x {
					f.f[y][lastDot] = 'O'
					f.f[y][x] = '.'
					lastDot = lastDot + 1
				} else {
					lastDot = x + 1
				}
			} else if r == '#' {
				lastDot = x + 1
			}
		}
	}
}

func (f *Field) moveRocksE() {
	W := len(f.f[0])
	H := len(f.f)
	for y := 0; y < H; y++ {
		lastDot := W - 1
		for x := W - 1; x >= 0; x-- {
			r := f.f[y][x]
			if r == 'O' {
				if lastDot != x {
					f.f[y][lastDot] = 'O'
					f.f[y][x] = '.'
					lastDot = lastDot - 1
				} else {
					lastDot = x - 1
				}
			} else if r == '#' {
				lastDot = x - 1
			}
		}
	}
}

func (f *Field) print() {
	for _, l := range f.f {
		for _, r := range l {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func makeField(a []string) *Field {
	f := Field{}
	for _, s := range a {
		in := make([]rune, 0)
		for _, r := range s {
			in = append(in, r)
		}
		f.f = append(f.f, in)
	}
	return &f
}
