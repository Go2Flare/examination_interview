package main

/*
物价上涨，良心杂货铺涨价
1.每次价格都是正整数
2.每次涨价的正好为x% (x=1,2,3, ... , 100)
3.价格区间[a,b]
求最多涨价的次数，和该次数对应的所有可能的价格组合的数目
初始价格不必为a,末尾价格不必为b
输入描述：
两个正整数a,b
30%数据a<=b<=10^2
100%数据a<=b<=10^8
输出描述：
两个正整数，分别表示最多涨价的次数，和该次数对应的所有可能的价格组合的数目

输入：
3 10
输出：
3 2
说明：
4 -> 5 -> 8 -> 10
4 -> 5 -> 6 -> 9
输入：
输出：
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
	s,e := nums[0],nums[1]
	//原先价格的可能的涨价数组
	nums = make([]int, e-s+1)
	start := s
	for i:=range nums{
		nums[i] = start
		start++
	}
	dp := make([]int, e-s+1)

	for i:= 1; i<len(dp); i++{
		for j:=0; j<i; j++{
			if isIntegerMul(nums[j], nums[i]){
				dp[i] = dp[j]+1
			}
		}
	}
	fmt.Println(dp)
	curMax := dp[0]
	for _,v:=range dp[1:]{
		if v>curMax{
			curMax = v
		}
	}
	countMax := 0
	for i:= range dp{
		if dp[i] == curMax{
			countMax++
		}
	}
	fmt.Printf("%v %v\n",curMax, countMax)
}

//判断y是否是x的整数百分比倍
func isIntegerMul(x,y int) bool {
	X := float64(x)
	Y := float64(y)
	d := Y-X//多的部分
	if (d/X * 100) - float64(int(d/X * 100)) < 1e-9{
		return true
	}
	return false
}
