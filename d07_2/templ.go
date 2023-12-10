package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	calc("0.txt")
	calc("1.txt")
	calc("2.txt")
}

type Data struct {
	code string
	typ  int
	v    int
}

func calc(fn string) {
	f, e := os.Open(fn)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	a := make([]Data, 0)
	for scanner.Scan() {
		s := scanner.Text()
		s1 := strings.Split(s, " ")
		//fmt.Println(s1)
		s1v, _ := strconv.Atoi(s1[1])
		dd := Data{code: s1[0], v: s1v}
		dd.typ = calcType(s1[0])
		a = append(a, dd)

		//calcType(s1[0])
		//fmt.Println(dd)
		//break
	}
	or := []uint8{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	orm := make(map[uint8]int)
	for i, c := range or {
		orm[c] = i
	}
	slices.SortFunc(a, func(a, b Data) int {
		d := -cmp.Compare(a.typ, b.typ)
		if d != 0 {
			return d
		}
		if a.code == b.code {
			fmt.Println(a.code)
		}
		for i := 0; i < 5; i++ {
			var r int

			r = -cmp.Compare(orm[a.code[i]], orm[b.code[i]])

			if r != 0 {
				return r
			}
		}
		return 0
	})
	s := int64(0)
	// not 250673552 (high)
	//     250612989 (high)
	//     250612989

	for i, v := range a {
		s += int64((i + 1) * v.v)
		//fmt.Println(i, v)
	}
	fmt.Println(s)
}

func calcType(s string) int {
	mp := make(map[int32]int, 0)
	for _, b := range s {
		//fmt.Println(b)
		mp[b] += 1
	}
	mpa := make([]int, 0)
	for _, v := range mp {
		mpa = append(mpa, v)
	}
	slices.Sort(mpa)
	//fmt.Println(mpa)
	//fmt.Println(len(mpa))

	if len(mpa) == 1 {
		// 5
		return 0
	}
	jc, _ := mp['J']
	if len(mpa) == 2 {
		if mpa[1] == 4 {
			// (4 1) ?
			if jc == 1 {
				return 0 // (5)
			}
			if jc == 4 {
				return 0
			}
			return 1 // (4 1)
		}
		// 3 + 2 ?
		if jc == 2 || jc == 3 {
			return 0 // (5)
		}
		// 3 + 2
		return 2
	}
	if len(mpa) == 3 {
		if mpa[2] == 3 {
			// 3 1 1
			if jc == 3 {
				return 1 // (4 1)
			}
			if jc == 1 {
				return 1 // (4 1)
			}
			return 3 // (3 1 1)
		}
		if mpa[2] == 2 {
			// 2 2 1
			if jc == 2 {
				return 1 // (4 1)
			}
			if jc == 1 {
				return 2 // (3 2)
			}
			return 4
		}
		panic(mpa)
	}
	if len(mpa) == 4 {
		// 2 1 1 1
		if jc == 2 {
			return 3 // (3 1 1)
		}
		if jc == 1 {
			return 3 // (3 1 1)
		}
		return 6
	}
	if jc == 1 {
		return 6 //2
	}
	return 7
}
