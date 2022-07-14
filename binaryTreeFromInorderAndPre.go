package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	tests_preorder := [][]int{
		{3, 9, 20, 15, 7},
		{1, 2, 3, 4},
		{0, 1, 2, 3, 4, 5},
	}

	tests_inorder := [][]int{
		{9, 3, 15, 20, 7},
		{1, 3, 2, 4},
		{0, 1, 2, 4, 3, 5},
	}

	for i := 0; i < len(tests_preorder); i++ {
		print_preorder(buildTree(tests_preorder[i], tests_inorder[i]))
		fmt.Println()
	}
}

/* "leetcode - 105 ; O(n) tc iterative solution without map"*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	var root *TreeNode
	stack := []*TreeNode{}
	var i, p int
	for p, i = len(preorder)-1, len(inorder)-1; i > -1 || p > -1; {
		if i > -1 && preorder[p] == inorder[i] {
			root = &TreeNode{Val: preorder[p], Right: root}
			i--
			p--
		} else if i > -1 && (len(stack) == 0 || preorder[p] != stack[len(stack)-1].Val) {
			stack = append(stack, &TreeNode{Val: inorder[i], Right: root})
			root = nil
			i--
		} else { // preorder[p] == stack[len(stack)-1].Val
			stack[len(stack)-1].Left = root
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p--
		}
	}
	return root
}

func print_preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	print_preorder(root.Left)
	print_preorder(root.Right)
}
