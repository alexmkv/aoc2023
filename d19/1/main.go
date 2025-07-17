package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	calc("0.txt")
	calc("1.txt")
}

type Workflows struct {
	a map[string]Workflow
}

type Workflow struct {
	r []Rule
}

const (
	R_ALW  = 0
	R_LESS = 1
	R_GR   = 2
)

type Rule struct {
	param string
	rule  string
	value int64
	out   string
}

func checkRule(r *Rule, e *El) bool {
	if r.rule == "" {
		return true
	}
	v, ex := e.vals[r.param]
	if !ex {
		panic(r)
	}
	if r.rule == ">" {
		return r.value < v
	}
	if r.rule == "<" {
		return r.value > v
	}
	panic(r)
}

func findOutput(w *Workflow, e *El) string {
	for _, r := range w.r {
		if checkRule(&r, e) {
			return r.out
		}
	}
	fmt.Println(w)
	panic(w)
}

type El struct {
	vals map[string]int64
}

func calc(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	scn := bufio.NewScanner(f)
	ph := 0
	wf := make(map[string]Workflow)
	sum := int64(0)
	for scn.Scan() {
		//fmt.Println(scn.Text())
		t := scn.Text()
		//res := regexp.MustCompile(`(.) ([0-9]+) \(#(.*)\)`).FindStringSubmatch(scn.Text())
		if ph == 0 {
			if t == "" {
				ph = 1
				continue
			}
			w := Workflow{}
			res := regexp.MustCompile(`([a-z]+)\{(.*)\}`).FindStringSubmatch(t)
			//fmt.Println(t)
			name := res[1]
			res1 := strings.Split(res[2], ",")
			for _, s := range res1 {
				res2 := regexp.MustCompile(`([a-z]+)([><])([0-9]+)\:([a-zA-Z]+)`).FindStringSubmatch(s)
				r := Rule{}
				if len(res2) == 0 {
					r.out = s
					//fmt.Println(s)
				} else {
					r.param = res2[1]
					r.rule = res2[2]
					r.value, _ = strconv.ParseInt(res2[3], 10, 64)
					r.out = res2[4]
				}
				w.r = append(w.r, r)
			}
			wf[name] = w
		} else {
			res := regexp.MustCompile(`([a-z]+)=([0-9]+)`).FindAllStringSubmatch(t, -1)
			//fmt.Println(res)
			el := El{make(map[string]int64)}
			for _, r0 := range res {
				el.vals[r0[1]], _ = strconv.ParseInt(r0[2], 10, 64)
			}
			r := "in"
			for r != "A" && r != "R" {
				ww, _ := wf[r]
				//fmt.Println(r)
				r = findOutput(&ww, &el)
			}
			f := func(e *El, k string) int64 {
				return e.vals[k]
			}
			if r == "A" {
				sum += el.vals["x"] + f(&el, "m") + f(&el, "a") + f(&el, "s")
			}
			//break
		}
	}
	fmt.Println(sum)
}
