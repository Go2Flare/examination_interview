package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*
房间传送门
时间限制： 3000MS
内存限制： 589824KB
题目描述：
存在n+1个房间，每个房间依次为房间1 2 3...i，每个房间都存在一个传送门，i房间的传送门可以把人传送到房间pi(1<=pi<=i),现在路人甲从房间1开始出发(当前房间1即第一次访问)，每次移动他有两种移动策略：     A. 如果访问过当前房间 i 偶数次，那么下一次移动到房间i+1；     B. 如果访问过当前房间 i 奇数次，那么移动到房间pi； 现在路人甲想知道移动到房间n+1一共需要多少次移动；



输入描述
第一行包括一个数字n(30%数据1<=n<=100，100%数据 1<=n<=1000)，表示房间的数量，接下来一行存在n个数字 pi(1<=pi<=i), pi表示从房间i可以传送到房间pi。

输出描述
输出一行数字，表示最终移动的次数，最终结果需要对1000000007 (10e9 + 7) 取模。


样例输入
2
1 2
样例输出
4
*/

var br *bufio.Reader
var bs *bufio.Scanner

//scan读取的行不能超过64*1024,否则就得启用buffer来扩大读取的token
func utilsScanner(s string, t int) (itemSlice interface{}){
	//s:切割的分隔符, t:类型{0:字符串切片, 1:数组切片, 2只读一个数字}
	//bs.Buffer(make([]byte, 1), 2*bufio.MaxScanTokenSize)//扩容2倍
	str := bs.Text()
	strSlice := strings.Split(str, s)
	if t == 0{
		itemSlice = strSlice
	}else if t == 1{
		nums := make([]int, len(strSlice))
		for i,v := range strSlice{
			nums[i],_ = strconv.Atoi(v)
		}
		itemSlice = nums
	}else if t==2{
		num,_ := strconv.Atoi(strSlice[0])
		itemSlice = num
	}
	return
}
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
//
//func numsSum(){
//	//br = bufio.NewReader(os.Stdin)
//	bs = bufio.NewScanner(os.Stdin)
//	bs.Buffer(make([]byte, 1), 2*bufio.MaxScanTokenSize)//扩容2倍
//	for bs.Scan(){
//		n := utilsScanner(" ", 2).(int)
//		nums := make([]int, n)
//		nums = utilsScanner(" ", 1).([]int)
//		dpMovement(nums)
//	}
//	//fmt.Println(nums)
//
//	//dpMovement(nums)
//}

func main(){
	//br = bufio.NewReader(os.Stdin)
	//_ = utilsReadLine(" ", 2).(int)
	//nums := utilsReadLine(" ", 1).([]int)
	n := 0
	s := ""
	fmt.Scanf("%v", &n)
	fmt.Scanf("%s", &s)
	ss := strings.Split(s," ")
	nums := make([]int, len(ss))
	for i:=range ss{
		nums[i],_ = strconv.Atoi(ss[i])
	}
	//fmt.Println(nums)
	//movement(nums)
	dpMovement(nums)
}

func dpMovement(nums []int){
	nums = append([]int{0}, nums...)
	dp := make([]int, len(nums))
//	初始化，只有一个房间至少要2次才能出去
	dp[1] = 2
	for i:=2; i<len(dp); i++{
		//偶数的两次 -
		dp[i] = 2*dp[i-1] - dp[nums[i-1]] +2
	}
	//fmt.Println(dp)
	fmt.Println(dp[len(dp)-1]%(10e9+7))
}

//模拟
func movement(nums []int) {
	count := make([]int, len(nums))
	step := 0
	i := 0//移动指针
	var mod int = 10e9+7
	for i<len(nums){
		count[i]++
		if count[i]%2 == 1{
			i= nums[i]-1
		}else if count[i] %2 == 0{
			i++
		}
		step = step % mod + 1
	}
	fmt.Println(step)
}