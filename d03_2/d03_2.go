package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ObjType int

const (
	NUM ObjType = iota
	SYMB
	NONE
)

type Obj struct {
	t      ObjType
	val    int
	marked bool
}

func main() {
	f, e := os.Open("01")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	objs := make([]*Obj, 0)
	field := make(map[int]*Obj, 0)
	W := 0
	y := 0

	for scanner.Scan() {
		s := scanner.Text()
		W = len(s)
		cur_numb := ""
		var num_obj *Obj = nil
		for i, v := range s {
			if v >= '0' && v <= '9' {
				if cur_numb == "" {
					num_obj = &Obj{t: NUM}
					objs = append(objs, num_obj)
				}
				cur_numb = cur_numb + string(v)
				field[i+y*W] = num_obj
			} else {
				if cur_numb != "" {
					num_obj.val, _ = strconv.Atoi(cur_numb)
					cur_numb = ""
				}
				if v == '.' {
					o := Obj{t: NONE}
					objs = append(objs, &o)
					field[i+y*W] = &o
				} else {
					o := Obj{t: SYMB}
					objs = append(objs, &o)
					field[i+y*W] = &o
				}
			}
		}
		if cur_numb != "" {
			num_obj.val, _ = strconv.Atoi(cur_numb)
			cur_numb = ""
		}
		//fmt.Println(s)
		y += 1
	}
	H := y
	safeMark := func(x int, y int, cnt *int, mult *int64, uniq map[*Obj]bool) {
		if x >= 0 && x < W && y >= 0 && y < H {
			o := field[x+y*W]
			if o.t == NUM && !uniq[o] {
				uniq[o] = true
				*cnt++
				*mult *= int64(field[x+y*W].val)
			}
		}

	}
	var sum int64
	for y = 0; y < H; y++ {
		for x := 0; x < W; x++ {
			o1 := field[x+y*W]
			if o1.t == SYMB {
				var cnt int
				var mult int64 = 1
				uniq := make(map[*Obj]bool)
				safeMark(x-1, y-1, &cnt, &mult, uniq)
				safeMark(x, y-1, &cnt, &mult, uniq)
				safeMark(x+1, y-1, &cnt, &mult, uniq)
				safeMark(x-1, y, &cnt, &mult, uniq)
				safeMark(x+1, y, &cnt, &mult, uniq)
				safeMark(x-1, y+1, &cnt, &mult, uniq)
				safeMark(x, y+1, &cnt, &mult, uniq)
				safeMark(x+1, y+1, &cnt, &mult, uniq)
				if cnt == 2 {
					fmt.Println(x, y)
					fmt.Println(cnt)
					sum += mult
				}
			}
		}
	}
	fmt.Println(sum)
}
