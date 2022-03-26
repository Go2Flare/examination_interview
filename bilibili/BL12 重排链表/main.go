package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type listNode struct{
	Val int
	Next *listNode
}

var r *bufio.Reader
var nums []int

func constuctor(nums []int) *listNode{
	dummy := &listNode{}
	cur := dummy
	for i:=0; i<len(nums); i++{
		cur.Next= &listNode{
			Val : nums[i],
			Next : nil,
		}
		cur = cur.Next
	}
	return dummy.Next
}



func findMiddleNode(head *listNode) *listNode{
	fast, slow := head, head
	for fast.Next != nil &&fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *listNode) *listNode{
	if head == nil || head.Next == nil{
		return head
	}
	cur := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}


func mergeList(l1 *listNode, l2 *listNode){
	var l1Next, l2Next *listNode
	//     遍历到两个链表都为空
	for l1 != nil && l2 != nil{
		l1Next, l2Next = l1.Next, l2.Next
		l1.Next = l2
		l1 = l1Next
		l2.Next = l1
		l2 = l2Next
	}
}

func reSortLinkList(head *listNode){

	//     双指针
	mid := findMiddleNode(head)

	l1 := head
	dummy := &listNode{Next:l1}//哨兵
	l2 := mid.Next
	mid.Next = nil

	l2 = reverseList(l2)
	mergeList(l1, l2)
	for l1 = dummy.Next ;l1!= nil; l1= l1.Next{
		if l1.Next == nil{fmt.Print(l1.Val)}else{fmt.Print(l1.Val,",")}
	}
}

// func 02mapInterface() {

//     r = bufio.NewReader(os.Stdin)
//         for {


//     readStr,err := r.ReadString('\n')
//             if err == io.EOF{
//                 return
//             }
//     readStr = readStr[:len(readStr)-1]
//     readStrSplit := strings.Split(readStr, ",")
//     nums = make([]int, len(readStrSplit))
//     if len(readStrSplit) == 0{fmt.Println()}
//     for i:=0; i<len(readStrSplit); i++{
//         nums[i],_ = strconv.Atoi(readStrSplit[i])
//     }
//     head := constuctor(nums)
//     reSortLinkList(head)
//                 }

// }
func main() {
	nums = []int{}
	r = bufio.NewReader(os.Stdin)
	for {
		readOneStr, err := r.ReadString(',')
		readOneStr = readOneStr[:len(readOneStr)-1]
		num, _ := strconv.Atoi(readOneStr)
		nums = append(nums, num)
		if err == io.EOF{
			break
		}
	}

	head := constuctor(nums)
	reSortLinkList(head)

}
