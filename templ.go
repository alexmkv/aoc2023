package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("d03_1/00")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}
