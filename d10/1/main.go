package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//calc("0.txt")
	//calc("1.txt")
	//calc("3.txt")
	calc("4.txt")
}

type Pt struct {
	x, y int
}

func calc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	mz := make([]string, 0)
	anPt := Pt{}
	for sc.Scan() {
		t := sc.Text()
		mz = append(mz, t)
		idx := strings.Index(t, "S")
		if idx != -1 {
			anPt.x = idx
			anPt.y = len(mz) - 1
		}
		//fmt.Println(t)
	}

	get := func(pt Pt) uint8 {
		return mz[pt.y][pt.x]
	}
	moveLoop := func(pt Pt, dx, dy int) (Pt, Pt) {
		res := Pt{pt.x + dx, pt.y + dy}
		if res.x < 0 || res.y < 0 || res.y >= len(mz) || res.x >= len(mz[res.y]) {
			return Pt{-1, -1}, Pt{0, 0}
		}
		//fmt.Println(res)
		return res, Pt{dx, dy}
	}
	moveByCh := func(pt Pt, dirPt Pt, ch uint8) (Pt, Pt) {
		if ch == '.' {
			return Pt{-1, -1}, Pt{0, 0}
		}
		if ch == 'J' {
			if dirPt.x == 1 {
				return moveLoop(pt, 0, -1)
			}
			return moveLoop(pt, -1, 0)
		}
		if ch == 'F' {
			if dirPt.x == -1 {
				return moveLoop(pt, 0, 1)
			}
			return moveLoop(pt, 1, 0)
		}
		if ch == '7' {
			if dirPt.x == 1 {
				return moveLoop(pt, 0, 1)
			}
			return moveLoop(pt, -1, 0)
		}
		if ch == 'L' {
			// ch == 'L'
			if dirPt.x == -1 {
				return moveLoop(pt, 0, -1)
			}
			return moveLoop(pt, 1, 0)
		}
		if ch == '|' {
			return moveLoop(pt, 0, dirPt.y)
		}
		if ch == '-' {
			return moveLoop(pt, dirPt.x, 0)
		}
		panic(fmt.Sprintf("Unknown char '%v'", string(ch)))
	}
	find := func() int {
		for i := 0; i < 3; i++ {
			pt := anPt
			var dirPt Pt
			if i == 0 {
				pt, dirPt = moveLoop(pt, -1, 0)
			} else if i == 1 {
				pt, dirPt = moveLoop(pt, 1, 0)
			} else if i == 2 {
				pt, dirPt = moveLoop(pt, 0, 1)
			}
			moves := 1
			for pt.x != -1 {
				s := get(pt)
				if s == 'S' {
					break
				}
				pt, dirPt = moveByCh(pt, dirPt, s)
				moves++
				if moves == 100000 {
					break
				}
			}
			if pt.x == -1 || moves == 100000 {
				//fmt.Println("not found", moves)
				continue
			}

			//fmt.Println("found", moves)
			if moves%2 == 0 {
				return moves / 2
			}
			return moves/2 + 1
		}
		panic("not found")
	}
	fmt.Println(find())

}
