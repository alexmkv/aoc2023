package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func testUni(t *testing.T, input string, output string) {
	re := regexp.MustCompile("n = (\\d+), meetings = (.*), firstPerson = (\\d+)")
	res := re.FindStringSubmatch(input)
	resShould := strToArray2(output)
	//fmt.Println(res[1:])
	resMy := findAllPeople(strToI(res[1]), strToArray1(res[2]), strToI(res[3]))
	fFail := func(expl string) {
		log.Fatalf("%v: My: %v, should: %v (%v)", t.Name(), resMy, resShould, expl)
	}
	if len(resMy) != len(resShould) {
		fFail("len don't match")
	}
	for i, v := range resShould {
		if resMy[i] != v {
			fFail(fmt.Sprintf("%v item don't match", i))
		}
	}
}

func strToArray1(s string) [][]int {
	res := make([][]int, 0)
	sTrimmedBraces := s[1 : len(s)-1]
	re := regexp.MustCompilePOSIX("\\[([0-9]+)\\,([0-9]+)\\,([0-9]+)\\]")
	reRes := re.FindAllStringSubmatch(sTrimmedBraces, -1)
	for _, m := range reRes {
		res = append(res, []int{strToI(m[1]), strToI(m[2]), strToI(m[3])})
	}
	return res
}

func strToArray2(s string) []int {
	res := make([]int, 0)
	sTrimmedBraces := s[1 : len(s)-1]
	splittedS := strings.Split(sTrimmedBraces, ",")
	for _, m := range splittedS {
		res = append(res, strToI(m))
	}
	return res
}

func strToI(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Test1(t *testing.T) {
	testUni(t, "n = 6, meetings = [[1,2,5],[2,3,8],[1,5,10]], firstPerson = 1", "[0,1,2,3,5]")
}

func Test2(t *testing.T) {
	testUni(t, "n = 4, meetings = [[3,1,3],[1,2,2],[0,3,3]], firstPerson = 3", "[0,1,3]")
}

func Test3(t *testing.T) {
	testUni(t, "n = 5, meetings = [[3,4,2],[1,2,1],[2,3,1]], firstPerson = 1", "[0,1,2,3,4]")
}

func Test_4(t *testing.T) {
	testUni(t, "n = 5, meetings = [[1,4,3],[0,4,3]], firstPerson = 3", "[0,1,3,4]")
}
