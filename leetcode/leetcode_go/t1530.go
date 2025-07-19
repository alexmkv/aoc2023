package t1530

import "fmt"

/*
You are given the root of a binary tree and an integer distance. A pair of two different leaf nodes of a binary tree is said to be good if the length of the shortest path between them is less than or equal to distance.

Return the number of good leaf node pairs in the tree.

Input: root = [1,2,3,null,4], distance = 3
Output: 1
Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.

Input: root = [1,2,3,4,5,6,7], distance = 3
Output: 2
Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.
Example 3:

Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
Output: 1
Explanation: The only good pair is [2,5].
ld

Constraints:

The number of nodldes in the tree is in the range [1, 210].
1 <= Node.val <= 100

1 <= distance <= 10


*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {

	check(0, []any{}, 5)
	check(0, []any{1}, 5)
	check(1, []any{1, 2, 3}, 5)
	check(3, []any{1, 2, 3, 4, 5, 6, nil}, 50)
	check(6, []any{1, 21, 22, 31, 32, 33, 34}, 50)
	check(6, []any{1, 21, 22, 31, 32, 33, 34, 41}, 50)
	check(4, []any{1, 21, 22, 31, 32, 33, 34, 41}, 4)

	check(1, []any{1, 2, 3, nil, 4}, 3)
	check(2, []any{1, 2, 3, 4, 5, 6, 7}, 3)
	check(1, []any{7, 1, 4, 6, nil, 5, 3, nil, nil, nil, nil, nil, 2}, 3)
}

func countPairs(root *TreeNode, distance int) int {
	r, _ := countPairsI(root, distance)
	return r
}

func countPairsI(root *TreeNode, distance int) (int, []int) {
	if root == nil {
		return 0, nil
	}
	if root.Left == nil && root.Right == nil {
		return 0, []int{1}
	}
	lr, ld := countPairsI(root.Left, distance)
	rr, rd := countPairsI(root.Right, distance)
	if len(ld) < len(rd) {
		ld, rd = rd, ld
	}
	cd := 0
	for i := 0; i < len(ld); i++ {
		lc := ld[len(ld)-i-1]
		for j := 0; j < min(len(rd), distance-i-1); j++ {
			rv := rd[len(rd)-j-1]
			cd += rv * lc
		}
	}
	dl := len(ld) - len(rd)
	for i, v := range rd {
		ld[i+dl] += v
	}
	ld = append(ld, 0)
	if len(ld) >= distance {
		ld = ld[1:]
	}
	return lr + rr + cd, ld
}

func check(expected int, tree []any, distance int) {
	r := countPairs(makeTree(tree, 0), distance)
	if r == expected {
		fmt.Print("OK ")
	} else {
		fmt.Print("FAIL ")
	}
	fmt.Println(expected, tree, distance, " RESULT: ", r)
}

func makeTree(tree []any, i int) *TreeNode {
	if len(tree) <= i {
		return nil
	}
	v := tree[i]
	if v == nil {
		return nil
	}
	return &TreeNode{Val: v.(int), Left: makeTree(tree, (i+1)*2-1), Right: makeTree(tree, (i+1)*2)}
}
