package main

import (
	"container/heap"
	"fmt"
)

/*

 */

func main() {
	//check(2, [][]int{{0, 1}}, []float64{0.5}, 0, 1, 0.5)
	check(5, [][]int{{2, 3}, {1, 2}, {3, 4}, {1, 3}, {1, 4}, {0, 1}, {2, 4}, {0, 4}, {0, 2}},
		[]float64{0.06, 0.26, 0.49, 0.25, 0.2, 0.64, 0.23, 0.21, 0.77}, 0, 3, 0.16000)

}

func check(n int, edges [][]int, succProb []float64, start_node int, end_node int, expected float64) {
	r := maxProbability(n, edges, succProb, start_node, end_node)
	good := true
	if abs(r-expected) > 0.0000001 {
		good = false
	} else {
		good = true
	}

	if good {
		fmt.Print("OK ")
	} else {
		fmt.Print("FAIL ")
	}
	fmt.Println(expected, "RESULT: ", r)
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	nodeProb := make([]float64, n)
	nodeProb[start_node] = 1
	hps := make([]PrA, n)
	for i, v := range edges {
		if hps[v[0]] == nil {
			hps[v[0]] = make([]Pr, 0)
		}
		if hps[v[1]] == nil {
			hps[v[1]] = make([]Pr, 0)
		}
		hps[v[1]] = append(hps[v[1]], Pr{v[0], succProb[i]})
		hps[v[0]] = append(hps[v[0]], Pr{v[1], succProb[i]})
	}
	for _, v := range hps {
		heap.Init(&v)
	}
	nds := PrA{{start_node, 1}}
	heap.Init(&nds)
	for len(nds) > 0 {
		pn := heap.Pop(&nds).(Pr)
		if pn.pr < nodeProb[pn.to] {
			continue
		}
		edgs := hps[pn.to]
		for _, e := range edgs {
			np := e.pr * pn.pr
			if np > nodeProb[e.to] {
				nodeProb[e.to] = np
				heap.Push(&nds, Pr{e.to, np})
			}
		}
	}

	return nodeProb[end_node]
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

type Pr struct {
	to int
	pr float64
}

type PrA []Pr

func (p PrA) Len() int {
	return len(p)
}

func (p PrA) Less(i, j int) bool {
	if p[i].pr != p[j].pr {
		return p[i].pr > p[j].pr
	}
	return p[i].to < p[j].to
}

func (p PrA) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PrA) Push(x any) {
	*p = append(*p, x.(Pr))
}

func (p *PrA) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
