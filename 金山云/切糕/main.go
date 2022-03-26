package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
X星切糕
时间限制： 3000MS
内存限制： 589824KB
题目描述：
据说在X星上有一种切糕，这种切糕呈细长形状，并且按长度来进行售卖
。 更有意思的是，不同长度的切糕其价格不一样。
定价规则如下：
1单位长度：价值为3枚X星币。
2单位长度：价值为7枚X星币。
3单位长度：价值为11枚X星币。
4单位长度：价值为15枚X星币。
5单位长度：价值为20枚X星币。
现有一块N个单位长度的切糕，需要将其切成若干小段，每一小段的长度均
不超过5个单位长度，请问可以得到的最大价值是多少枚X星币？


输入描述
单组输入。 输入一个正整数N表示切糕的总长度，N<=10^6。

输出描述
输出切成小段之后可以得到的最大价值（X星币的枚数）。

样例输入
12
样例输出
47
*/

//完全背包

var weight = []int{1,2,3,4,5}
var value = []int{3,7,11,15,20}
var r *bufio.Reader
func getValue(n int) {
	if n == 0 {
		fmt.Println(0)
		return
	}
	//dp := make([][]int, 5)
	//for i:=0; i<5; i++{
	//	dp[i]=make([]int, n+1)
	//}
	dp := make([]int, n+1)
	//初始化
	//for i:=0 ;i<5; i++{
	//
	//}
	//遍历背包
	for j:=0; j<n+1; j++{
		//物品
		for i:=0; i<len(weight); i++{
			if j>=weight[i]{
				dp[j] = max(dp[j-weight[i]]+value[i], dp[j])
			}
		}
	}
	fmt.Println(dp[n])
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func main() {
	r = bufio.NewReader(os.Stdin)
	str, _ := r.ReadString('\n')
	str = str[:len(str)-1]
	n,_ := strconv.Atoi(str)
	getValue(n)
}