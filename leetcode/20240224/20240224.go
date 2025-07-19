package main

import (
	"sort"
)

type Info struct {
	tm   int
	with int
}

const NO_FIND = 1 << 30

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make(map[int][]Info)
	for _, m := range meetings {
		if graph[m[0]] == nil {
			graph[m[0]] = make([]Info, 0)
		}
		graph[m[0]] = append(graph[m[0]], Info{m[2], m[1]})
		if graph[m[1]] == nil {
			graph[m[1]] = make([]Info, 0)
		}
		graph[m[1]] = append(graph[m[1]], Info{m[2], m[0]})
	}
	earliest := make([]int, n)
	for i := 0; i < len(earliest); i++ {
		earliest[i] = NO_FIND
	}
	queue := make([]Info, 0)
	setEarliest := func(with int, tm int) {
		earliest[with] = tm
		queue = append(queue, Info{tm, with})
	}
	setEarliest(0, 0)
	setEarliest(firstPerson, 0)
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]
		if el.tm > earliest[el.with] {
			continue
		}
		for _, el2 := range graph[el.with] {
			if el2.tm >= el.tm && earliest[el2.with] > el2.tm {
				setEarliest(el2.with, el2.tm)
			}
		}
	}
	res := make([]int, 0, n)
	for i, v := range earliest {
		if v != NO_FIND {
			res = append(res, i)
		}
	}
	return res
}

func findAllPeople_my(n int, meetings [][]int, firstPerson int) []int {
	res := make([]bool, n)
	meetRes := make([]bool, len(meetings))
	res[0] = true
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	res[firstPerson] = true
	currentTime := 0
	recheckPos := -1
	checkMeeting := func(i int) int {
		if meetRes[i] {
			return 2
		}
		if res[meetings[i][0]] && res[meetings[i][1]] {
			meetRes[i] = true
			return 2
		}
		if res[meetings[i][0]] && !res[meetings[i][1]] {
			res[meetings[i][1]] = true
			meetRes[i] = true
			return 1
		}
		if res[meetings[i][1]] && !res[meetings[i][0]] {
			res[meetings[i][0]] = true
			meetRes[i] = true
			return 1
		}
		return 0
	}
	needRecheck := false
	for i := 0; i <= len(meetings); i++ {
		time := -1
		if i != len(meetings) {
			time = meetings[i][2]
		}
		if currentTime != time {
			for needRecheck && recheckPos != -1 {
				j := recheckPos
				recheckPos = -1
				needRecheck = false
				for ; j < i; j++ {
					checkRes := checkMeeting(j)
					if checkRes == 0 && recheckPos == -1 {
						recheckPos = j
					}
					if checkRes == 1 {
						needRecheck = true
					}
				}
			}
			currentTime = time
			needRecheck = false
			recheckPos = -1
		}
		if time == -1 {
			break
		}
		checkRes := checkMeeting(i)
		if checkRes == 1 {
			needRecheck = true
		} else if checkRes == 0 && recheckPos == -1 {
			recheckPos = i
		}
	}
	ress := make([]int, len(res))
	idx := 0
	for pId, _ := range res {
		ress[idx] = pId
		idx++
	}
	sort.Ints(ress)
	return ress
}
