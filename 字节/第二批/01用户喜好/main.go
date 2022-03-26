package main

/*
用户喜好
时间限制： 3000MS
内存限制： 589824KB
题目描述：
为了不断优化推荐效果，今日头条每天要存储和处理海量数据。假设有这样一种场景：我们对用户按照它们的注册时间先后来标号，对于一类文章，每个用户都有不同的喜好值，我们会想知道某一段时间内注册的用户（标号相连的一批用户）中，有多少用户对这类文章喜好值为k。因为一些特殊的原因，不会出现一个查询的用户区间完全覆盖另一个查询的用户区间(不存在L1<=L2<=R2<=R1)。


输入描述
第1行为n代表用户的个数

第2行为n个整数，第i个代表用户标号为i的用户对某类文章的喜好度

第3行为一个正整数q代表查询的组数

第4行到第（3+q）行，每行包含3个整数l,r,k代表一组查询，即标号为l<=i<=r的用户中对这类文章喜好值为k的用户的个数。

数据范围n <= 300000,q<=300000 k是整型

输出描述
一共q行，每行一个整数代表喜好值为k的用户的个数
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Reader

//可以读取长度很长的行
func utilsReadLine(s string, t int) (itemSlice interface{}) {
	//s:切割的分隔符, t:类型{0:字符串切片, 1:数组切片, 2只读一个数字}
	bytes, _, _ := br.ReadLine() //读出的是[]byte
	strSlice := strings.Split(string(bytes), s)
	if t == 0 {
		itemSlice = strSlice
	} else if t == 1 {
		nums := make([]int, len(strSlice))
		for i, v := range strSlice {
			nums[i], _ = strconv.Atoi(v)
		}
		itemSlice = nums
	} else if t == 2 {
		num, _ := strconv.Atoi(strSlice[0])
		itemSlice = num
	}
	return
}

func main() {
	br = bufio.NewReader(os.Stdin)
	utilsReadLine(" ", 2)
	like := utilsReadLine(" ", 1).([]int)
	q := utilsReadLine(" ", 2).(int)
	var nums [][]int
	for i:=0; i<q; i++{
		nums = append(nums, utilsReadLine(" ", 1).([]int))
	}

	res := []int{}
	for i:= 0; i<len(nums); i++{
		l,r,k := nums[i][0]-1, nums[i][1]-1, nums[i][2]
		cnt := 0
		for j:=l; j<=r; j++{
			if like[j] == k{
				cnt++
			}
		}
		res = append(res, cnt)
	}
	for _,v := range res{
		fmt.Println(v)
	}
	//fmt.Println(n, like, q, nums)
}
