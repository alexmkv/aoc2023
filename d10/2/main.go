package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 4
	//calc("0.txt")
	//calc("1.txt")
	//calc("3.txt")
	calc("4.txt")
	// 4
	//calc("5.txt")
	//calc("6.txt")
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
	H9 := len(mz) * 3
	W9 := len(mz[0]) * 3
	mz9 := make([]int, H9*W9)
	const (
		S_NIL  = 0
		S_PATH = 1
		S_OUT  = 2
		S_X    = 3
	)
	mark := false
	get := func(pt Pt) uint8 {
		return mz[pt.y][pt.x]
	}
	markMz9 := func(pt Pt, v int, pts ...Pt) {
		/*if true {
			return
		}*/
		for _, p := range pts {
			mz9[(pt.y*(W9/3)+pt.x)*9+p.y*3+p.x] = v
		}
	}
	getMz9Mark := func(pt Pt) int {
		return mz9[((pt.y/3)*(W9/3)+(pt.x/3))*9+(pt.y%3)*3+(pt.x%3)]
	}
	setMz9Mark := func(pt Pt) {
		mz9[((pt.y/3)*(W9/3)+(pt.x/3))*9+(pt.y%3)*3+(pt.x%3)] = S_OUT
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
			if mark {
				markMz9(pt, S_PATH, Pt{1, 0}, Pt{1, 1}, Pt{0, 1})
			}
			if dirPt.x == 1 {
				return moveLoop(pt, 0, -1)
			}
			return moveLoop(pt, -1, 0)
		}
		if ch == 'F' {
			if mark {
				markMz9(pt, S_PATH, Pt{2, 1}, Pt{1, 1}, Pt{1, 2})
			}
			if dirPt.x == -1 {
				return moveLoop(pt, 0, 1)
			}
			return moveLoop(pt, 1, 0)
		}
		if ch == '7' {
			if mark {
				markMz9(pt, S_PATH, Pt{0, 1}, Pt{1, 1}, Pt{1, 2})
			}
			if dirPt.x == 1 {
				return moveLoop(pt, 0, 1)
			}
			return moveLoop(pt, -1, 0)
		}
		if ch == 'L' {
			if mark {
				markMz9(pt, S_PATH, Pt{1, 0}, Pt{1, 1}, Pt{2, 1})
			}

			// ch == 'L'
			if dirPt.x == -1 {
				return moveLoop(pt, 0, -1)
			}
			return moveLoop(pt, 1, 0)
		}
		if ch == '|' {
			if mark {
				markMz9(pt, S_PATH, Pt{1, 0}, Pt{1, 1}, Pt{1, 2})
			}

			return moveLoop(pt, 0, dirPt.y)
		}
		if ch == '-' {
			if mark {
				markMz9(pt, S_PATH, Pt{0, 1}, Pt{1, 1}, Pt{2, 1})
			}

			return moveLoop(pt, dirPt.x, 0)
		}
		panic(fmt.Sprintf("Unknown char '%v'", string(ch)))
	}
	find := func() int {
		i := 1
		//mark := false
		for i < 3 {
			pt := anPt
			var dirPt Pt
			var stDir Pt
			if i == 0 {
				//pt, dirPt = moveLoop(pt, -1, 0)
				stDir = Pt{-1, 0}
			} else if i == 1 {
				stDir = Pt{1, 0}
				//pt, dirPt = moveLoop(pt, 1, 0)
			} else if i == 2 {
				stDir = Pt{0, 1}

			}
			pt, dirPt = moveLoop(pt, stDir.x, stDir.y)
			moves := 1
			for pt.x != -1 {
				s := get(pt)
				if s == 'S' {
					markMz9(pt, S_X, Pt{1, 1})
					markMz9(pt, S_X, Pt{1 + stDir.x, 1 + stDir.y})
					markMz9(pt, S_X, Pt{1 - dirPt.x, 1 - dirPt.y})
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
				i++
				continue
			}
			if !mark {
				//markMz9(Pt{1, 1}, S_PATH, Pt{1, 1})
				mark = true
				continue
			}
			stPt := Pt{0, 0}
			toCheck := make([]Pt, 0)
			toCheck = append(toCheck, stPt)
			for len(toCheck) > 0 {
				c := toCheck[0]
				toCheck = toCheck[1:]
				if getMz9Mark(c) == S_NIL {
					if true {
						setMz9Mark(c)
					}
					if c.y > 0 {
						toCheck = append(toCheck, Pt{c.x, c.y - 1})
					}
					if c.y < H9-1 {
						toCheck = append(toCheck, Pt{c.x, c.y + 1})
					}
					if c.x > 0 {
						toCheck = append(toCheck, Pt{c.x - 1, c.y})
					}
					if c.x < W9-1 {
						toCheck = append(toCheck, Pt{c.x + 1, c.y})
					}
				}
			}
			cnt := 0
			fmt.Println(H9, W9)
			if true {
				for y := 0; y < H9; y++ {
					for x := 0; x < W9; x++ {
						s := getMz9Mark(Pt{x, y})
						switch s {
						case S_NIL:
							fmt.Print("-")
						case S_PATH:
							fmt.Print("+")
						case S_OUT:
							fmt.Print(" ")
						case S_X:
							fmt.Print("x")
						default:
							fmt.Print("?")
						}
					}
					fmt.Println()
				}

			}
			for y := 0; y < H9/3; y++ {
				for x := 0; x < W9/3; x++ {
					k := 0
					for ; k < 9; k++ {
						//fmt.Println(x, y, k)
						if mz9[(y*(W9/3)+x)*9+k] != S_NIL {
							break
						}
					}
					if k == 9 {
						cnt++
					}
				}
			}
			fmt.Println("Count", cnt)
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
