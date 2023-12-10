package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, e := os.Open("2.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var r1 []int64
	var r2 []int64
	for scanner.Scan() {
		s := scanner.Text()
		if r1 == nil {
			r1 = getInts(s)
		} else {
			r2 = getInts(s)
		}
		//fmt.Println(s)
	}
	res := int64(1)
	for i := 0; i < len(r1); i++ {
		t := r1[i]
		r := r2[i]
		s := solve(float64(t), float64(r))
		//fmt.Println(s)
		res *= s
	}
	fmt.Println(res)
	//fmt.Println(r2)
}

func solve(t float64, r float64) int64 {
	v1f := (t - math.Sqrt(t*t-4*r)) / 2
	v1 := int64(math.Ceil(v1f))
	if v1f == math.Ceil(v1f) {
		//fmt.Println(v1f, math.Ceil(v1f))
		v1++
	}
	v2f := (t + math.Sqrt(t*t-4*r)) / 2
	v2 := int64(math.Floor(v2f))
	if v2f == math.Floor(v2f) {
		//fmt.Println(v2f, math.Floor(v2f))
		v2--
	}
	//fmt.Println(v1, (t-math.Sqrt(t*t-4*r))/2, v2, (t+math.Sqrt(t*t-4*r))/2)
	return v2 - v1 + 1
}

func doRegexpAll(reg string, s string) []string {
	re := regexp.MustCompile(reg)
	return re.FindAllString(s, -1)
}

func getInts(s string) []int64 {
	r := make([]int64, 0)
	seedsSS := doRegexpAll(`([0-9]+)`, s)
	for _, s := range seedsSS {
		x, _ := strconv.ParseInt(s, 10, 64)
		r = append(r, x)
	}
	return r
}
