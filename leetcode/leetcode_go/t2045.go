package main

import (
	"fmt"
)

/*

 */

func main() {
	// 1 -> (1!) 2 -> (2) 1  -> (3) 2
	check(2, [][]int{{1, 2}}, 1, 100, 3)
	// 1 -> (1) 2 -> (2) 3
	// 1 -> (1) 2 -> (2) 3 (=>4)-> (5) 2 -> (6) 3
	check(3, [][]int{{1, 2}, {2, 3}}, 1, 2, 6)
	// 1 -> (1) 2 -> 3
	// 1 -> (1) 2 -> (2) 3
	check(3, [][]int{{1, 2}, {1, 3}}, 1, 100, 3)
}

func check(n int, edges [][]int, time int, change int, expected int) {
	r := secondMinimum(n, edges, time, change)
	good := true
	if r != expected {
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

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		if adj[e[0]] == nil {
			adj[e[0]] = make([]int, 0)
		}
		if adj[e[1]] == nil {
			adj[e[1]] = make([]int, 0)
		}
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	//fmt.Println(n)
	q := [][]int{{1, 1}}
	d1 := make([]int, n+1)
	d2 := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		d1[i] = -1
		d2[i] = -1
	}
	d1[1] = 0
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		t := d1[v[0]]
		if v[1] == 2 {
			t = d2[v[0]]
		}
		if (t/change)%2 > 0 {
			//t += (change - t%change)
			t = change * (t/change + 1)
		}
		t += time
		if adj[v[0]] == nil {
			continue
		}
		// change = 3, 0..2 -> t += time
		// 3..5 -> t += (change - t % change) + time
		// 0..2->0, 3..5 -> 1, 6..8->0, 9..11->1 ...
		// (v / 3) % 2 == 1
		//fmt.Println(t, v, adj[v[0]])

		for _, e := range adj[v[0]] {
			if d1[e] == -1 {
				d1[e] = t
				q = append(q, []int{e, 1})
			} else if d2[e] == -1 && d1[e] != t {
				if e == n {
					return t
				}
				d2[e] = t
				q = append(q, []int{e, 2})
			}
		}
	}
	return -1
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// dfs algorithm
	//fmt.Println(n)
	q := [][]int{{-1, 3}, {1, 1}}
	d1 := make([]int, n+1)
	d2 := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		d1[i] = -1
		d2[i] = -1
	}
	d1[1] = 0
	adj := make([][]int, n+1)
	for _, e := range edges {
		if adj[e[0]] == nil {
			adj[e[0]] = make([]int, 0)
		}
		if adj[e[1]] == nil {
			adj[e[1]] = make([]int, 0)
		}
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	t := 0
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		if v[1] == 3 {
			if len(q) == 0 {
				break
			}
			if (t/change)%2 == 1 {
				t += (change - t%change)
			}
			t += time
			q = append(q, []int{-1, 3})
			continue
		}
		if adj[v[0]] == nil {
			continue
		}
		// change = 3, 0..2 -> t += time
		// 3..5 -> t += (change - t % change) + time
		// 0..2->0, 3..5 -> 1, 6..8->0, 9..11->1 ...
		// (v / 3) % 2 == 1
		//fmt.Println(t, v, adj[v[0]])

		for _, e := range adj[v[0]] {
			if d1[e] == -1 && v[1] == 1 {
				d1[e] = t
				if e == n {
					q = append(q, []int{e, 2})
				} else {
					q = append(q, []int{e, 1})
				}
			} else if d2[e] == -1 && d1[e] != t {
				d2[e] = t
				if e == n {
					return t
				}
				q = append(q, []int{e, 2})
			}
		}
	}
	return -1
}
