package main

/*
序列解构
时间限制： 5000MS
内存限制： 589824KB
题目描述：
       已知一个长度为n的序列A。我们求出了它的前缀序列，但是却不小心把它的原序列丢失了，请你从它的前缀序列中推出它的原序列，并输出出来。
       这里的前缀序列B定义为，B[i]为所有下标小于等于i中的元素中，所有奇数位置的和减去所有偶数位置的和。A序列和B序列的下标均从1开始。

输入描述
      输入第一行仅包含一个正整数n，表示序列的长度。(1<=n<=10000)
      输入第二行包含n个整数，空格隔开，表示一个前缀序列，所有整数的绝对值都小于等于10^9。

输出描述
输出仅包含一行，包含n个整数，表示原序列，中间用空格隔开。

样例输入
4
1 -1 2 -2
样例输出
1 2 3 4*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r *bufio.Reader

func recovery(nums []int){

	originNums := make([]int, len(nums))
	for i:=len(nums)-1; i>=0; i--{
		if (i+1)%2 == 1{
			//i==0
			if i==0{
				originNums[i] = nums[i]
				continue
			}
			originNums[i] = nums[i]-nums[i-1]
		}else if i>=1 && (i+1)%2 == 0{
			originNums[i] = nums[i-1]-nums[i]
		}
	}
	//打印
	for i:=0; i<len(originNums); i++{
		if i==0 {
			fmt.Print(originNums[i])
		}else{
			fmt.Print(" ",originNums[i])
		}
	}
	fmt.Println()
}

func main(){
	r = bufio.NewReader(os.Stdin)
	readStr,_ := r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	n, _ := strconv.Atoi(readStr)
	readStr,_ = r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	nums := make([]int, n)
	readStrs := strings.Split(readStr, " ")
	for i:=0; i<n; i++{
		nums[i],_ = strconv.Atoi(readStrs[i])
	}
	recovery(nums)

}