package main

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max(values... int) int {
	max := values[0]
	for _, val := range values {
		if max < val {
			max = val
		}
	}
	return max
}

func findMax(root *TreeNode, answer *int) int {
	if root == nil {
		return math.MinInt16
	}

	left := findMax(root.Left, answer)
	right := findMax(root.Right, answer)
	if left == math.MinInt16 && right == math.MinInt16 {
		return root.Val
	}
	cur := left
	if right > cur {
		cur = right
	}
	*answer = max(*answer, cur, root.Val, cur + root.Val, left + root.Val + right)
	return max(root.Val, cur + root.Val)
}

func maxPathSum(root *TreeNode) int {
	var answer int = root.Val

	findMax(root, &answer)
	return answer
}

func main() {

}
