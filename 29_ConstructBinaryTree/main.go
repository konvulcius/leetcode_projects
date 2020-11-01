package main

//Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	tmpStart := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			tmpStart = i
			break
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:tmpStart+1], inorder[:tmpStart]),
		Right: buildTree(preorder[tmpStart+1:], inorder[tmpStart+1:]),
	}
}

func main() {
	buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	return
}
