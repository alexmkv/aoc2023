package main

import (
	"fmt"
)

/*

 */

func main() {
	// n = 4, edges = [[0,1,3],[1,2,1],[1,3,4],[2,3,1]], distanceThreshold = 4
	check(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4, 3)
	//n = 5, edges = [[0,1,2],[0,4,8],[1,2,3],[1,4,2],[2,3,1],[3,4,1]], distanceThreshold = 2
	check(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2, 0)
	check(2, [][]int{{0, 1, 200}}, 20, 1)
	check(2, [][]int{{0, 1, 2}}, 20, 1)
	check(3, [][]int{{0, 1, 2}}, 20, 2)

	edges := make([][]int, 0)
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			edges = append(edges, []int{i, j, 1})
		}
	}
	check(100, edges, 10000, 0)

}

func check(n int, edges [][]int, distanceThreshold int, expected int) {
	r := findTheCity(n, edges, distanceThreshold)
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

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	//fmt.Println(n, distanceThreshold)
	MAX := distanceThreshold + 1
	m := make([][]int, n)
	for i, _ := range m {
		m[i] = make([]int, n)
		mi := m[i]
		for j, _ := range mi {
			if i == j {
				mi[j] = 0
			} else {
				mi[j] = MAX
			}
		}
	}
	for _, e := range edges {
		//fmt.Println(e)
		m[e[0]][e[1]] = e[2]
		m[e[1]][e[0]] = e[2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				nv := m[i][k] + m[k][j]
				if m[i][j] > nv {
					m[i][j] = nv
					m[j][i] = nv
				}
			}
		}
	}
	//fmt.Println(m)
	mi := -1
	mv := 0

	for i, r := range m {
		cnt := 0
		for j, c := range r {
			if i != j && c <= distanceThreshold {
				cnt++
			}
		}
		//fmt.Println(i, cnt)
		if mi == -1 || cnt <= mv {
			mv = cnt
			mi = i
		}
	}
	return mi
}

func findTheCityO(n int, edges [][]int, distanceThreshold int) int {
	ed := make(map[int]map[int]int)
	add := func(v1, v2, w int) {
		if ed[v1] == nil {
			ed[v1] = make(map[int]int)
		}
		ed[v1][v2] = w
	}
	for _, v := range edges {
		add(v[0], v[1], v[2])
		add(v[1], v[0], v[2])
	}
	//fmt.Println(ed)
	d := make(map[int]map[int]int)
	mi := -1
	m := 0
	for i := 0; i < n; i++ {
		fill(i, i, ed, d, distanceThreshold)

		edn := d[i]
		//fmt.Println(i, edn)
		cn := 0
		if edn != nil {
			cn = len(edn)
		}
		if mi == -1 || cn <= m {
			m = cn
			mi = i
		}
	}
	return mi
}

func fill(i, i0 int, ed map[int]map[int]int, d map[int]map[int]int, distanceThreshold int) {
	edn := ed[i0]
	if edn == nil {
		return
	}
	for i1, v := range edn {
		if i1 == i {
			continue
		}
		ndt := distanceThreshold - v
		if ndt < 0 {
			continue
		}
		if d[i] == nil {
			d[i] = make(map[int]int)
		}
		/*if d[i1] == nil {
			d[i11] = make(map[int]int)
		}*/

		dt, ok := d[i][i1]
		if !ok || ndt > dt {
			d[i][i1] = ndt
			//d[i1][i] = ndt
			if ndt > 0 {
				fill(i, i1, ed, d, ndt)
			}
		}
	}
}
