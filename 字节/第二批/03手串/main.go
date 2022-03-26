package main
/*
手串
时间限制： 3000MS
内存限制： 589824KB
题目描述：
作为一个手串艺人，有金主向你订购了一条包含n个杂色串珠的手串——每个串珠要么无色，要么涂了若干种颜色。为了使手串的色彩看起来不那么单调，金主要求，手串上的任意一种颜色（不包含无色），在任意连续的m个串珠里至多出现一次（注意这里手串是一个环形）。手串上的颜色一共有c种。现在按顺时针序告诉你n个串珠的手串上，每个串珠用所包含的颜色分别有哪些。请你判断该手串上有多少种颜色不符合要求。即询问有多少种颜色在任意连续m个串珠中出现了至少两次。



输入描述
第一行输入n，m，c三个数，用空格隔开。(1 <= n <= 10000, 1 <= m <= 1000, 1 <= c <= 50) 接下来n行每行的第一个数num_i(0 <= num_i <= c)表示第i颗珠子有多少种颜色。接下来依次读入num_i个数字，每个数字x表示第i颗柱子上包含第x种颜色(1 <= x <= c)

输出描述
一个非负整数，表示该手链上有多少种颜色不符需求。


样例输入
5 2 3
3 1 2 3
0
2 2 3
1 2
1 3
样例输出
2
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
	n,m,c := nums[0], nums[1], nums[2]
	mat := make([][]int, n)
	for i:=0; i<n; i++{
		nums = utilsReadLine(" ", 1).([]int)
		mat[i] = make([]int, nums[0])
		copy(mat[i], nums[1:])
	}

	findColor(mat,m, c)
}

func findColor(mat [][]int, m, c int){
	color := make([]int, c+1)


	for k:=0; k<len(mat);k++{
		m1 := make(map[int]int)
		for i:=0; i<m; i++{
			curI := (k+i)%len(mat)
			for j:=0; j<len(mat[curI]); j++{
				m1[mat[curI][j]]++
			}
		}
		for i,v :=range m1{
			if v>=2{
				//不符合的color出现
				color[i]++
			}
		}

	}
	countUnfit:=0
	for i:=1; i<len(color); i++{
		if color[i] >0{
			countUnfit++
		}
	}
	fmt.Println(countUnfit)
}