package main

import "fmt"

/*
二叉树不相连的点的最大值
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func eating(root *TreeNode) (res int) {
	traversal(root)

	return root.Val
}

//
func traversal(root *TreeNode) {
	if root == nil {
		return
	}
	//想通过自底向上，用后序遍历逻辑啊，不用花里胡哨的接收这个左右
	traversal(root.Left)
	traversal(root.Right)

	if root.Left != nil && root.Right != nil {
		root.Val = max(root.Left.Val+root.Right.Val, root.Val)
	} else if root.Left != nil && root.Right == nil {
		root.Val = max(root.Left.Val, root.Val)
	} else if root.Left == nil && root.Right != nil {
		root.Val = max(root.Right.Val, root.Val)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//返回建立的新节点
func newNode(val int) *TreeNode{
	 return &TreeNode{
		 Val:val,
	 }
}

func buildTree(arr []int) *TreeNode {
	root := newNode(arr[0])
	queue:=[]*TreeNode{root}
	k:=1
	label1:
	for len(queue)>0{
		l := len(queue)
		for i:=0; i<l; i++{
			e := queue[0]
			queue = queue[1:]
			//左
			var left *TreeNode
			if arr[k] != -1{left = newNode(arr[k])}
			e.Left = left
			queue = append(queue, left)
			k++
			if k == len(arr){break label1}
//右
			var right *TreeNode
			if arr[k] != -1{right = newNode(arr[k])}
			e.Right = right
			queue = append(queue, right)
			k++
			if k == len(arr){break label1}
		}
	}
	return root
}

func main() {
	//直接通过创建节点来建树
	root := buildDummily()
	//层序遍历构建节点
	arr := []int{4, 5, 10, -1, 8, 3, -1}
	tree := buildTree(arr)
	fmt.Println(tree, tree.Left, tree.Right, tree.Left.Right, tree.Right.Left)
	fmt.Println(root, root.Left, root.Right, root.Left.Right, root.Right.Left)
	fmt.Println(eating(tree))
}

func buildDummily() *TreeNode{
	root := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	return root
}