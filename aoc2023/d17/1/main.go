package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println(Calc("0.txt"))
	fmt.Println(Calc("1.txt"))
}

func Calc(fname string) int {
	f, e := os.Open(fname)
	if e != nil {
		return 1
	}
	scn := bufio.NewScanner(f)
	fld := Field{}
	for scn.Scan() {
		fld.addLine(scn.Text())
	}
	fld.initField()
	return fld.calcPath()
}

type Field struct {
	heat [][]int
	f    [][]FldInfo
	W    int
	H    int
}

func (f *Field) initField() {
	f.W = len(f.f[0])
	f.H = len(f.f)
}

func (fld *Field) addLine(s string) {
	fld.heat = append(fld.heat, []int{})
	fld.f = append(fld.f, []FldInfo{})
	for _, r := range s {
		h, _ := strconv.Atoi(string(r))
		fld.heat[len(fld.f)-1] = append(fld.heat[len(fld.f)-1], h)
		fld.f[len(fld.f)-1] = append(fld.f[len(fld.f)-1], FldInfo{loss: make(map[int]map[int]int)})
	}
}

func (fld *Field) heatLoss(x, y int) int {
	return fld.heat[y][x]
}

const (
	DIR_N = 0
	DIR_E = 1
	DIR_S = 2
	DIR_W = 3
)

func (fld *Field) calcPath() int {
	toCheck := []ToCheck{ToCheck{0, 0, 1, 3, -fld.heatLoss(0, 0)}}
	for len(toCheck) > 0 {
		chk := toCheck[0]
		toCheck = toCheck[1:]
		if chk.x < 0 || chk.y < 0 || chk.x >= fld.W || chk.y >= fld.H {
			continue
		}
		new_loss := chk.loss + fld.heatLoss(chk.x, chk.y)
		for dd := -1; dd <= 1; dd++ {
			left := 3
			if dd == 0 {
				left = chk.left
			}
			chk_d := (chk.dir + dd + 4) % 4
			fldInf := &fld.f[chk.y][chk.x]
			if fldInf.checkAndAdd(chk_d, left, new_loss) {
				if left == 0 {
					continue
				}
				switch chk_d {
				case DIR_N:
					toCheck = append(toCheck, ToCheck{chk.x, chk.y - 1, chk_d, left - 1, new_loss})
				case DIR_S:
					toCheck = append(toCheck, ToCheck{chk.x, chk.y + 1, chk_d, left - 1, new_loss})
				case DIR_E:
					toCheck = append(toCheck, ToCheck{chk.x + 1, chk.y, chk_d, left - 1, new_loss})
				case DIR_W:
					toCheck = append(toCheck, ToCheck{chk.x - 1, chk.y, chk_d, left - 1, new_loss})
				}
			}
		}
	}
	return fld.f[fld.H-1][fld.W-1].minWeight()
}

type ToCheck struct {
	x, y int
	dir  int
	left int
	loss int
}

type FldInfo struct {
	//data FldInfoData
	loss map[int]map[int]int
	min  int
}

type FldInfoData struct {
	loss map[int]map[int]int
}

func (fi *FldInfo) checkAndAdd(dir int, moves int, loss int) bool {
	toAdd := true
	res := false
	_, ok := fi.loss[dir]
	if !ok {
		fi.loss[dir] = make(map[int]int)
	}
	//toDel := []int{}
	for fid_moves, fid_loss := range fi.loss[dir] {
		if fid_moves == moves {
			if fid_loss > loss {
				fi.loss[dir][moves] = loss
				res = true
			}
			toAdd = false
		} else if fid_moves < moves {
			if fid_loss >= loss {
				delete(fi.loss[dir], fid_moves)
			}
		} else { // fid.moves > moves
			if fid_loss <= loss {
				toAdd = false
			}
		}
	}
	if toAdd {
		fi.loss[dir][moves] = loss
		res = true
	}
	//fi.min = fi.minWeight()
	return res
}

func (fldInfo *FldInfo) minWeight() int {
	minLoss := 1 << 29
	for _, fi := range fldInfo.loss {
		for _, ls := range fi {
			if minLoss > ls {
				minLoss = ls
			}
		}
	}
	return minLoss
}
