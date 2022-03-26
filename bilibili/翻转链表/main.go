package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Val int
	Next *Node
}

func mergeNode(l1 *Node, l2 *Node) {
	for l1 != nil && l2 != nil{
		l1Next, l2Next := l1.Next, l2.Next
		l1.Next = l2
		l1 = l1Next
		l2.Next = l1
		l2 = l2Next
	}
}

func findMid(head *Node) *Node{
	s, f:= head, head
	for f.Next != nil&& f.Next.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	return s
}

func reverse(head *Node) *Node{
	cur := head
	var pre *Node
	for cur!=nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func contructor(arr []int) *Node{
	head := &Node{}
	dummy := head
	for i:=0; i<len(arr); i++{
		NewNode := &Node{
			Val: arr[i],
			Next: nil,
		}
		head.Next = NewNode
		head = head.Next
	}
	return dummy.Next
}
var r *bufio.Scanner

func main() {
	r = bufio.NewScanner(os.Stdin)
	for r.Scan(){
		str := r.Text()
		strSlice := strings.Split(str,",")
		nums := make([]int, len(strSlice))
		for i:=0; i<len(nums); i++{
			nums[i],_ = strconv.Atoi(strSlice[i])
		}

		head := contructor(nums)
		l1 := head
		mid := findMid(head)
		l2 := mid.Next
		mid.Next = nil
		l2 = reverse(l2)
		mergeNode(l1, l2)
		fmt.Print(l1.Val)
		l1 = l1.Next
		for cur:=l1; cur!=nil; cur=cur.Next{
			fmt.Print(",",cur.Val)
		}
		fmt.Println()
	}
}
