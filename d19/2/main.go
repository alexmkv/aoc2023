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
	//calc("1.txt")
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

type Path struct {
	rs []Rule
}

func prependPath(r Rule, p0 *Path) Path {
	p := Path{}
	p.rs = []Rule{r}
	p.rs = append(p.rs, p0.rs...)
	return p
}

func findPath(m map[string]Workflow, el string) []Path {
	if el == "R" {
		return nil
	}
	if el == "A" {
		return []Path{{rs: []Rule{{rule: ""}}}}
	}
	w, _ := m[el]
	rr := []Path{}
	for _, r := range w.r {
		res := findPath(m, r.out)
		if res != nil {
			for _, r0 := range res {
				rr = append(rr, prependPath(r, &r0))
			}
		}
	}
	if len(rr) == 0 {
		return nil
	}
	return rr
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

type Range struct {
	l, r int64
}

type Ranges = map[string]Range

func defRange() Range {
	return Range{1, 4000}
}

func defRanges() *Ranges {
	r := Ranges{}
	ks := []string{"x", "m", "a", "s"}
	for _, k := range ks {
		r[k] = defRange()
	}
	return &r
}

func applyRule(r Rule, rng *Ranges) {
	if r.rule == "" {
		return
	}
	rngV := (*rng)[r.param]
	if r.rule == "<" {
		if rngV.r > r.value-1 {
			rngV.r = r.value - 1
		}
	} else if r.rule == ">" {
		if rngV.l < r.value+1 {
			rngV.l = r.value + 1
		}
	} else {
		panic(r)
	}
}

func applyPath(p *Path, rng *Ranges) {
	for _, p0 := range p.rs {
		applyRule(p0, rng)
	}
}

func rngValue(rng *Ranges) int64 {
	i := int64(1)
	for _, v := range *rng {
		if v.l > v.r {
			return 0
		}
		i *= (v.l - v.r + 1)
	}
	return i
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
		}
	}
	//fmt.Println(wf)
	rr := findPath(wf, "in")
	rng := []*Ranges{}
	for i, r0 := range rr {
		rng = append(rng, defRanges())
		applyPath(&r0, rng[0])
		sum += rngValue(rng[0])
		fmt.Println(rngValue(rng[0]))
		fmt.Println(r0)
	}
	//fmt.Println(rr)

	fmt.Println(sum)
}
