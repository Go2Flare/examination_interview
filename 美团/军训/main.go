package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*

 正在录屏
美团 2022届秋招 技术综合-运维方向在线考试
编程题|20.0分4/4
军训
时间限制： 3000MS
内存限制： 1048576KB
题目描述：
小团去参加军训了。军训快要结束了，长官想要把大家一排n个人分成m组，然后让每组分别去参加阅兵仪式。阅兵仪式上会进行打分，
其中有一个奇怪的扣分点是每组的极差，即每组最大值减去最小值。长官想要让这分成的m组总扣分量最小，即这m组分别的极差之和最小
。长官正在思索如何安排中，就让小团来帮帮他吧。
给出n个数a[1],a[2],…,a[n]，分别表示一排n个人的身高，以及1个数m，表示要将他们分成m组。每组要求都是连续的一个区间，
比如，可以将某一组安排为第2,3,4,5个人，但不能安排为第2,3,4,6个人，因为第4个人和第6个人之间有还有一个人，这样的区间是不连续的。

输入描述
第一行两个数n, m，空格隔开，表示一共有n个人以及要分成m组。

第二行n个数，a[1],a[2],…,a[n]，空格隔开，依次从左到右表示这排人的身高。

1<=m<=n<=500，1<=a[i]<=1,000,000,000

输出描述
输出一个值，表示分成的m组人，各组极差之和。


样例输入
3 2
5 1 2
样例输出
1

5 2
5 4 3 1 2

提示
将第1个人单独分成一组，第2、3个人单独分成一组。

此时第一组极差为5 – 5 = 0

第二组极差为 2 – 1 = 1
极差之和为1*/

var r *bufio.Reader

func getMin(nums []int, m int) int {
	if m == 0 || len(nums) == 1 { //
		return 0
	}
	var minMinux1, minMinux2 int

	minus := math.MaxInt64
	minMM := make(map[int]int)
	for i:=1; i<len(nums); i++{
		nums1 := make([]int, len(nums[:i]))
		nums2 := make([]int, len(nums[i:]))
		copy(nums1,nums[:i])
		copy(nums2,nums[i:])

		for i:=0; i<len(nums1); i++{
			sort.Ints(nums1)
			minMinux1 = nums1[len(nums1)-1] - nums1[0]
		}
		for i:=0; i<len(nums2); i++{
			sort.Ints(nums2)
			minMinux2 = nums2[len(nums2)-1] - nums2[0]
		}
		//极差map
		//fmt.Println(nums1, nums2)
		//fmt.Println("mm", minMinux1, minMinux2)
		minMM[minMinux1+minMinux2] = i
	}


	//fmt.Println(minMM)
	for i,_ := range minMM{
		if i<minus{
			minus = i
		}
	}
	//分界元素
	pivot := minMM[minus]
///*	for _,v := range minMM{
//		if v<minus{
//			minus = v
//		}
//	}*/
//	fmt.Println("pm",pivot, minus)
//	fmt.Println(nums[:pivot], nums[pivot:])
	return minus+getMin(nums[:pivot], m-1)+getMin(nums[pivot+1:], m-1)
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func main() {
	r = bufio.NewReader(os.Stdin)
	readStr, _ := r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	readStrs := strings.Split(readStr, " ")
	n , _:= strconv.Atoi(readStrs[0])
	m , _ :=strconv.Atoi(readStrs[1])

	readStr, _ = r.ReadString('\n')

	readStr = readStr[:len(readStr)-1]
	readStrs = strings.Split(readStr, " ")
	nums := make([]int, n)
	for i:=0; i<n; i++{
		nums[i],_ = strconv.Atoi(readStrs[i])
	}
	fmt.Println(getMin(nums, m))
}