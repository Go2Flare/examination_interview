package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
时间限制： 3000MS
内存限制： 589824KB
题目描述：
有一个仅包含’a’和’b’两种字符的字符串s，长度为n，每次操作可以把一个字符做一次转换（把一个’a’设置为’b’，或者把一个’b’置成’a’)；但是操作的次数有上限m，问在有限的操作数范围内，能够得到最大连续的相同字符的子串的长度是多少。



输入描述
第一行两个整数 n , m (1<=m<=n<=50000)，第二行为长度为n且只包含’a’和’b’的字符串s。

输出描述
输出在操作次数不超过 m 的情况下，能够得到的 最大连续 全’a’子串或全’b’子串的长度。


样例输入
8 1
aabaabaa
样例输出
5*/

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

func main(){
	br = bufio.NewReader(os.Stdin)
	nums := utilsReadLine(" ", 1).([]int)
	_,m := nums[0], nums[1]
	strs := utilsReadLine(" ", 0).([]string)
	str := strs[0]

	maxStr2(str,m)
}

func maxStr2(str string, m int){
	cA, cB := 0, 0
	for i:= 0; i<len(str); i++{
		if str[i] == 'a'{
			cA++
		}else{
			cB++
		}
	}
	e := 'b'
	if cA<cB{e='a'}

	res := 1

	l,r := 0,0
	//滑窗向右探测最大边界
	for l<len(str){
		curK:=m
		for r<len(str) && (curK>=0 || str[r] == byte(e)){
			if str[r] == byte(e){
				curK--
			}
			if curK<0{break}
			r++
		}
		if r-l > res {
			res = r-l
		}
		for l<len(str) && str[l] != byte(e){
			l++
		}
		l = l+1
		r = l
	}
	fmt.Println(res)
}


//
//func maxStr (str string, k int) {
//	cA, cB := 0, 0
//	for i:= 0; i<len(str); i++{
//		if str[i] == 'a'{
//			cA++
//		}else{
//			cB++
//		}
//	}
//	e := 'b'
//	if cA<cB{e='a'}
//
//	dp := make([][]int, len(str))
//	for i:=range dp{
//		dp[i] = make([]int, len(str))
//		dp[i][i] = 1
//	}
//	fmt.Println(k)
//	for i:=0; i<len(dp); i++{
//		curK := k
//		for j :=i+1; j<len(dp); j++{
//			if str[j] != byte(e){
//
//			}else{
//				if curK>0{
//					dp[i][j] = dp[i][j-1]+1
//					curK--
//				}else{
//					dp[i][j] = dp[i][j-1]
//				}
//			}
//		}
//	}
//	fmt.Println(dp)
//}