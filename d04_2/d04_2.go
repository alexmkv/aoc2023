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
	cards := make([]int, 0)
	id := 0
	for scanner.Scan() {
		if id == len(cards) {
			cards = append(cards, 1)
		}
		id++
		s := scanner.Text()
		res := doRegexp(`Card\W+[0-9]+: (.*) \| (.*)$`, s)
		winning := make(map[int]bool)
		//fmt.Println(s)
		//fmt.Println(res)
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
				i++
			}
		}
		//fmt.Println(id, i, cards)
		for j := 0; j < i; j++ {
			idn := id + j
			if idn >= len(cards) {
				cards = append(cards, 1)
			}
			cards[idn] += cards[id-1]
		}
		//fmt.Println(id, i, cards)
		//fmt.Println(s)d04_1
		//fmt.Println(i)
	}
	sum = 0
	for _, v := range cards {
		//fmt.Println(v)
		sum += v
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
