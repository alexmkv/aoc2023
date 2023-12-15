package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	do("0.txt")
	do("1.txt")
}

func do(fname string) {
	f, e := os.Open(fname)
	if e != nil {
		panic(e)
	}
	//a := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		//a = append(a, sc.Text())
		parts := strings.Split(sc.Text(), ",")
		var sum int64
		for _, p := range parts {
			sum += calcP(p)
		}
		fmt.Println(sum)
	}
}

func calcP(p string) int64 {
	var sum int64
	for _, r := range p {
		sum += int64(r)
		sum *= 17
		sum %= 256
	}
	return sum
}
