package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, e := os.Open("d04_2")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		res := doRegexp(`Card\W+[0-9]+: (.*) \| (.*)$`, s)
		winning := make(map[int]bool)
		fmt.Println(s)
		fmt.Println(res)
		for _, n := range doRegexpAll(`[0-9]*`, res[1]) {
			v, _ := strconv.Atoi(n)
			if v == 0 {
				continue
			}
			winning[v] = true
		}
		i := 0
		for _, n := range doRegexpAll(`[0-9]*`, res[2]) {
			v, _ := strconv.Atoi(n)
			if winning[v] {
				//fmt.Println(v)
				if i == 0 {
					i = 1
				} else {
					i *= 2
				}
			}
		}
		//fmt.Println(s)
		//fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

func doRegexp(reg string, s string) []string {
	re := regexp.MustCompile(reg)
	return re.FindStringSubmatch(s)
}

func doRegexpAll(reg string, s string) []string {
	re := regexp.MustCompile(reg)
	return re.FindAllString(s, -1)
}
