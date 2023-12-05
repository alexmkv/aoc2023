package main

import (
	"bufio"
	"fmt"
	tree "github.com/ugurcsen/gods-generic/trees/redblacktree"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//processFile("0.txt")
	processFile("1.txt")
}
func processFile(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	convs := make([]*Comp, 7)
	for i := 0; i < 7; i++ {
		convs[i] = initComp()
	}
	//m.Put(5, "x")
	//m.Put(10, "y")
	//fmt.Println(m.Floor(11))
	scanner := bufio.NewScanner(f)
	stage := 100
	seeds := make([]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if stage == 100 {
			stage = 99
			seedsS := doRegexp(`seeds: (.*)`, s)[1]
			seeds = getInts(seedsS)
		} else if stage == 99 {
			if s == "" {
				continue
			}
			switch s {
			case "seed-to-soil map:":
				stage = 0
			case "soil-to-fertilizer map:":
				stage = 1
			case "fertilizer-to-water map:":
				stage = 2
			case "water-to-light map:":
				stage = 3
			case "light-to-temperature map:":
				stage = 4
			case "temperature-to-humidity map:":
				stage = 5
			case "humidity-to-location map:":
				stage = 6
			default:
				panic(s)
			}
		} else {
			if s == "" {
				stage = 99
				continue
			}
			cc := convs[stage]
			vals := getInts(s)
			r := Range{vals[0], vals[2]}
			cc.add(vals[1], r)
		}
	}
	fmt.Println(seeds)
	min := -1
	for _, s := range seeds {
		v := s
		for i := 0; i < 7; i++ {
			v = convs[i].conv(v)
			//fmt.Println(v)
		}
		//break
		//fmt.Println(v)
		if min == -1 || min > v {
			min = v
		}
	}
	fmt.Println(min)
}

type Range struct {
	start int
	len   int
}

type Comp struct {
	rg *tree.Tree[int, Range]
}

func initComp() *Comp {
	c := Comp{}
	c.rg = tree.NewWithNumberComparator[Range]()
	return &c
}

func (c *Comp) add(k int, r Range) {
	c.rg.Put(k, r)
}

func (c *Comp) conv(v int) int {
	fl, ok := c.rg.Floor(v)
	if !ok {
		return v
	}
	//fmt.Println(",,", fl.Key)
	diff := v - fl.Key
	if diff >= fl.Value.len {
		return v
	}
	return fl.Value.start + diff
}

func doRegexp(reg string, s string) []string {
	re := regexp.MustCompile(reg)
	return re.FindStringSubmatch(s)
}

func doRegexpAll(reg string, s string) []string {
	re := regexp.MustCompile(reg)
	return re.FindAllString(s, -1)
}

func getInts(s string) []int {
	r := make([]int, 0)
	seedsSS := doRegexpAll(`([0-9]+)`, s)
	for _, s := range seedsSS {
		x, _ := strconv.Atoi(s)
		r = append(r, x)
	}
	return r
}
