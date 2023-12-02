package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ColorCount struct {
	blue  int
	red   int
	green int
}

type Game struct {
	num    int
	rounds []ColorCount
}

func parseGame(s string) Game {
	g := Game{}

	p0 := strings.Split(s, ":")
	g.num, _ = strconv.Atoi(p0[0][5:])
	p1 := strings.Split(p0[1], ";")
	g.rounds = make([]ColorCount, len(p1))
	//fmt.Println(p1)
	for i, v := range p1 {
		g.rounds[i] = parseRound(v)
	}
	return g
}

func parseRound(s string) ColorCount {
	c := ColorCount{}
	p := strings.Split(s, ",")
	re := regexp.MustCompile("([0-9]*) (red|green|blue)")
	for _, v := range p {
		res := re.FindStringSubmatch(v)
		//fmt.Println(res[1])
		cnt, _ := strconv.Atoi(res[1])
		switch res[2] {
		case "red":
			c.red = cnt
		case "blue":
			c.blue = cnt
		case "green":
			c.green = cnt
		default:
			panic(res[2])
		}
	}
	//fmt.Println(p)
	return c
}

func possibleGame(r ColorCount, c ColorCount) bool {
	if r.red > c.red {
		return false
	}
	if r.green > c.green {
		return false
	}
	if r.blue > c.blue {
		return false
	}
	return true
}

func main() {
	cond := ColorCount{
		blue: 14, red: 12, green: 13}
	fmt.Println(cond)
	file, _ := os.Open("d02_1/02_11")
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		g := parseGame(s)
		//fmt.Println(parseGame(s))
		possible := true
		for _, r := range g.rounds {
			if !possibleGame(r, cond) {
				possible = false
				break
			}
		}
		if possible {
			sum += g.num
		}
	}
	fmt.Println(sum)
}
