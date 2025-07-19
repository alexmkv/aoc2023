package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	proc("1.txt")
	// 9556712
	proc("11.txt")
}

func proc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	sc := bufio.NewScanner(f)
	fd := make([]string, 0)
	extraY := make(map[int]bool)
	extraX := make(map[int]bool)
	for sc.Scan() {
		s := sc.Text()
		//fmt.Println(s)
		fd = append(fd, s)

		//extraY = append(extraY, len(fd)-1)
		if strings.Index(s, "#") == -1 {
			//fd = append(fd, strings.Repeat(".", len(s)))
			extraY[len(fd)-1] = true
		}
	}
	for i := 0; i < len(fd[0]); i++ {
		fnd := false
		for _, s := range fd {
			if s[i] == '#' {
				fnd = true
				break
			}
		}
		if !fnd {
			/*for j := 0; j < len(fd); j++ {
				fd[j] = fd[j][:i] + "." + fd[j][i:]
			}*/
			//extraX = append(extraX, i)
			extraX[i] = true
		}
	}
	sum := int64(0)
	type Pt struct {
		x, y int
	}
	srt := func(v0, v1 int) (int, int) {
		if v0 > v1 {
			return v1, v0
		}
		return v0, v1
	}
	PLUS := int64(1000000 - 1)
	// 678626878094 (high)
	dist := func(p0, p1 Pt) int64 {
		x0, x1 := srt(p0.x, p1.x)
		y0, y1 := srt(p0.y, p1.y)
		f := func(i0, i1 int, extraI map[int]bool) int64 {
			sum := int64(0)
			for i := i0; i < i1; i++ {
				sum++
				if extraI[i] {
					sum += PLUS
				}
			}
			return sum
		}
		return f(x0, x1, extraX) + f(y0, y1, extraY)
	}
	fmt.Println(dist(Pt{3, 0}, Pt{7, 8}))
	if false {
		return
	}
	pts := make([]Pt, 0)
	for x := 0; x < len(fd[0]); x++ {
		for y := 0; y < len(fd); y++ {
			if fd[y][x] == '#' {
				pc := Pt{x, y}
				for _, p0 := range pts {
					sum += dist(pc, p0)
				}
				pts = append(pts, pc)
				//sum += checkAround(fd, x, y)
				//fmt.Println(sum)
			}
		}
	}
	fmt.Println(sum)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func checkAround(fd []string, x, y int) int {
	cp := func(x0, y0 int) bool {
		if x0 < 0 || y0 < 0 || x0 >= len(fd[0]) || y0 >= len(fd) {
			return false
		}
		if fd[y0][x0] == '#' {
			//fd[y0] = fd[y0][:x0] + "." + fd[y0][x0+1:]
			return true
		}
		return false
	}
	for i := 1; ; i++ {
		for dx := 0; dx <= i; dx++ {
			dy := i - dx
			if cp(x-dx, y-dy) || cp(x+dx, y-dy) || cp(x-dx, y+dy) || cp(x+dx, y+dy) {
				return i
			}
		}
		if i == 100 {
			break
		}
	}
	return -100000
}
