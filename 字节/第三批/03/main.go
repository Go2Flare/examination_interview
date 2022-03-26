package main

/*
分配房间
编程题结果 未编译
占用内存 0 KB
代码耗时 0 MS
代码规范 N/A
尝试编译运行 次
切出编程界面 0 次
作答时间 N/A 分钟
题目描述

有n个房间，现在i号房间里的人需要被重新分配，分配的规则是这样的：先让i号房间里的人全都出来，接下来按照 i+1, i+2, i+3, ... 的顺序依此往这些房间里放一个人，n号房间的的下一个房间是1号房间，直到所有的人都被重新分配。

现在告诉你分配完后每个房间的人数以及最后一个人被分配的房间号x，你需要求出分配前每个房间的人数。数据保证一定有解，若有多解输出任意一个解。

输入

第一行两个整数n, x (2<=n<=10^5, 1<=x<=n)，代表房间房间数量以及最后一个人被分配的房间号；
第二行n个整数 a_i(0<=a_i<=10^9) ，代表每个房间分配后的人数。

输出

输出n个整数，代表每个房间分配前的人数。

题目限制

时间限制：C/C++语言 1000 MS；其他语言 3000 MS
内存限制：C/C++语言 65536KB；其他语言 589824KB

样例输入

3 1 6 5 1
样例输出

4 4 4

*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var br *bufio.Reader

func utilsReadLine(s string, t int) (itemSlice interface{}){
	//s:切割的分隔符, t:类型{0:字符串切片, 1:数组切片, 2只读一个数字}
	bytes,_,_ := br.ReadLine()//读出的是[]byte
	strSlice := strings.Split(string(bytes), s)
	if t == 0{
		itemSlice = strSlice
	}else if t == 1{
		nums:= make([]int,len(strSlice))
		for i,v := range strSlice{
			nums[i], _ = strconv.Atoi(v)
		}
		itemSlice = nums
	}else if t == 2{
		num, _ := strconv.Atoi(strSlice[0])
		itemSlice = num
	}
	return
}

func main() {
	br = bufio.NewReader(os.Stdin)
	nums := utilsReadLine(" ", 1).([]int)
	n,index := nums[0],nums[1]
	nums = make([]int,n)
	nums = utilsReadLine(" ", 1).([]int)

	findOrigin(nums, index-1)
}

func findOrigin(nums []int, index int) {
	beSplitPeople:=0
	//人数都被分配的房间
	i := index
	for {
		if nums[i] == 0{//找到被分配的房间
			break
		}
		nums[i]--
		beSplitPeople++
		i = (i-1+len(nums))%len(nums)
	}
	nums[i] = beSplitPeople
	//fmt.Println(nums)
	fmt.Print(nums[0])
	for _,v:=range nums[1:]{
		fmt.Print(" ",v)
	}
	fmt.Println()
}
