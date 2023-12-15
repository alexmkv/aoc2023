package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	do("0.txt")
	do("1.txt")
}

type It struct {
	label string
	pw    int
}

type ItP struct {
	it     It
	hs     int
	action rune
}

func do(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	//a := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), ",")
		bx := make([][]It, 256)
		for _, p := range parts {
			pi := calcP(p)
			//fmt.Println(pi)
			if pi.action == '=' {
				b := &bx[pi.hs]
				found := false
				for i := 0; i < len(*b); i++ {
					if (*b)[i].label == pi.it.label {
						found = true
						(*b)[i].pw = pi.it.pw
					}
				}
				if !found {
					bx[pi.hs] = append(bx[pi.hs], pi.it)
				}
			} else if pi.action == '-' {
				b := &bx[pi.hs]
				toD := 0
				for i := 0; i < len(*b); i++ {
					if (*b)[i].label == pi.it.label {
						toD++
					} else if toD > 0 {
						(*b)[i-toD] = (*b)[i]
					}
				}
				if toD > 0 {
					bx[pi.hs] = (*b)[:len(*b)-toD]
				}
			}
		}
		ss := int64(0)
		for bi, b := range bx {
			for si, s := range b {
				ss += int64((bi + 1) * (si + 1) * s.pw)
			}
		}
		fmt.Println(ss)
	}
}

func calcP(p string) ItP {
	rs := ItP{}
	stage := 0
	num := ""
	for _, r := range p {
		if stage == 1 {
			num += string(r)
			continue
		}
		if r == '-' || r == '=' {
			rs.action = r
			stage = 1
			continue
		}
		rs.it.label += string(r)
		rs.hs += int(r)
		rs.hs *= 17
		rs.hs %= 256
	}
	rs.it.pw, _ = strconv.Atoi(num)
	return rs
}
