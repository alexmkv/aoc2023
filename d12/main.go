package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calc("0.txt", 1)
	calc("1.txt", 1)
	calc("0.txt", 5)
	calc("1.txt", 5)
}

type Key struct {
	p1i    int
	p2i    int
	head   int
	active bool
}

var mp map[Key]int64

func calc(fname string, repeat int) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	sc := bufio.NewScanner(f)
	sum := int64(0)
	line := 0
	for sc.Scan() {
		t := sc.Text()
		sp := strings.Split(t, " ")
		p1r := sp[0]
		p1 := ""
		for i := 0; i < repeat; i++ {
			if i != 0 {
				p1 += "?"
			}
			p1 += p1r
		}
		p2r := strings.Split(sp[1], ",")
		p2 := make([]int, 0)
		for i := 0; i < repeat; i++ {
			for _, p := range p2r {
				i, _ := strconv.Atoi(p)
				p2 = append(p2, i)
			}
		}
		mp = make(map[Key]int64)
		v := options(p1, p2, 0, 0, -1, false)
		sum += v
		line++
	}
	fmt.Println(sum)
}

func options(p1 string, p2 []int, p1i int, p2i int, head int, active bool) int64 {
	k := Key{p1i, p2i, head, active}
	v, ok := mp[k]
	if ok {
		return v
	}
	r := options1(p1, p2, p1i, p2i, head, active)
	mp[k] = r
	return r
}

func options1(p1 string, p2 []int, p1i int, p2i int, head int, active bool) int64 {
	if len(p1) == p1i && len(p2) == p2i {
		return 1
	}
	if len(p1) == p1i && len(p2) > p2i {
		if len(p2)-1 == p2i && head == 0 {
			return 1
		}
		return 0
	}
	if len(p2) == p2i {
		if p1[p1i] == '#' {
			return 0
		}
		return options(p1, p2, p1i+1, p2i, -1, false)
	}

	num := head
	if num == -1 {
		num = p2[p2i]
	}
	fact := func() int64 {
		return options(p1, p2, p1i+1, p2i, num-1, true)
	}
	factf := func() int64 {
		if num > 0 {
			return fact()
		}
		if num == 0 {
			return options(p1, p2, p1i+1, p2i+1, -1, false)
		}
		panic(`unexpected factf`)
	}

	if active {
		if p1[p1i] == '.' && num > 0 {
			return 0
		}
		if p1[p1i] == '#' && num == 0 {
			return 0
		}
		return factf()
	}
	if p1[p1i] == '.' {
		return options(p1, p2, p1i+1, p2i, -1, false)
	}
	if p1[p1i] == '#' {
		return factf()
	}
	if p1[p1i] == '?' {
		r1 := options(p1, p2, p1i+1, p2i, -1, false)
		r2 := factf()
		return r1 + r2
	}
	panic(`something unexpected`)
}
