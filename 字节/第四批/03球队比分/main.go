package main

/*
编程题|20分3/3
球队比分
时间限制： 3000MS
内存限制： 589824KB
题目描述：
有三只球队，每只球队编号分别为球队1，球队2，球队3，这三只球队一共需要进行 n 场比赛。现在已经踢完了k场比赛，每场比赛不能打平，踢赢一场比赛得一分，输了不得分不减分。已知球队1和球队2的比分相差d1分，球队2和球队3的比分相差d2分，每场比赛可以任意选择两只队伍进行。求如果打完最后的 (n-k) 场比赛，有没有可能三只球队的分数打平。



输入描述
第一行包含一个数字 t (1 <= t <= 10) 接下来的t行每行包括四个数字 n, k, d1, d2(1 <= n <= 10^12; 0 <= k <= n, 0 <= d1, d2 <= k)

输出描述
每行的比分数据，最终三只球队若能够打平，则输出“yes”，否则输出“no”


样例输入
2
3 3 0 0
3 3 3 3
样例输出
yes
no
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
	n := utilsReadLine(" ", 2).(int)
	mat := make([][]int, n)
	for i:=0; i<n; i++{
		mat[i] = utilsReadLine(" ", 1).([]int)
		winwin(mat[i])
	}
	//fmt.Println(n, mat)

}

func winwin(nums []int) {
	if nums[0]-nums[1]>= nums[2]+nums[3] {
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
}