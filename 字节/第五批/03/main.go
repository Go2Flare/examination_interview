package main

/*
手机游戏有战力很多角色，角色可设定战力值，为了平衡每个角色的能力，提升玩家的体验，策划往往会对角色的战力值通过一定规则进行限制
某款手游有n个角色，这些角色从0 到 n-1进行编号排成一列，角色的战力值按照一下规则进行限定：
	1.第一个角色的初始战力为0
	2.每个角色的战力值是一个非负的整数
	3.相邻角色的战力值不超过1（0，+1，-1）
基于以上规则还要再加一定的限制，，设定最大战力值，这些限制会以若干数对的形式给出，每一个数对定义如下
limit[i] = [index, maxPower] (index != 0 && i<n-1) 表示编号为index的角色的最大战力值，不超过maxPower(非负）
由于第一个角色有了初始的战力值，所以不会再对编号0的角色进行战力值的限定
	基于以上规则，计算出单个角色能达到的最大战力值

	输入描述：
第一行为两个整数n,m(2<=n<=10^6, 1<=m<=10^5) ,其中n为游戏种角色的总个数，m为限制了最大战力的角色数
后面的m行，每一行都有两个整数，index,maxPower(index!=0 && index<n-1), index是角色编号，index!=0即
不会对编号0的角色进行战力限定，maxPower是该角色被限定的最大战斗值

	输出描述：
输出1个整数，表示单个角色能到达的最大战斗值

实例1：
输入：
3 2
1 3
2 2

输出：
2

说明：
输入：
第一行表示游戏中有3个角色，对其中2个角色限制了最大的战斗值
接下来的2行是具体的限制，对编号为1的角色限制了最大战斗力3， 对编号为2的角色限制了为2
输出：
[0,1,2]是满足规则约束的一个战力值设定方案，相邻角色的战力值之差不超过1， 最大战力值是最后一个角色，最大战力值为2

实例2：
输入：
5 3
1 1
2 3
4 3

输出：
3

说明：
输入：
第一行表示游戏中有5个角色，对其中3个角色限制了最大的战斗值
接下来的2行是具体的限制，对编号为1的角色限制了最大战斗值1， 对编号为2的角色限制了最大战斗值为3， 对编号为4的角色限制了最大战斗值为3
输出：
[0,1,2,3,3]是满足规则约束的一个战力值设定方案，相邻角色的战力值之差不超过1， 最大战力值是最后一个角色，最大战力值为2


*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Reader

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
	nums := utilsReadLine(" ", 1).([]int)
	n, m := nums[0], nums[1]
	p := make([]int, n)
	for i := 0; i < m; i++ {
		nums = utilsReadLine(" ", 1).([]int)
		p[nums[0]] = nums[1]
	}

	moutain(p)
}

func moutain(p []int) {
	//p1的i是战力值，p1[i]是对应的坐标数组，是递增的
	p1 := make([][]int, 100000+1)
	for i := 0; i < len(p); i++ {
		if p[i] > 0 {
			p1[p[i]] = append(p1[p[i]], i)
		}
	}

	//最后的战力值数组
	powerNums := make([]int, len(p))
	start := 1

	//i是战力值, 对应数组是战力值的坐标
	for i := 1; i < len(p1); i++ {
		if p1[i] == nil{
			continue
		}
		for j := 0; j < len(p1[i]); j++ {

			maxPower := i
			if maxPower != 0 {
				end := p1[i][j]
				//start前的都固定增加好战力

				if end >= start { //?                     //powerNums[start]所在的战力高度,要实时更新
					for start < len(powerNums) && powerNums[start] < maxPower && start <= end { //起始到终点之间，起始点一定要小于最右边限制的maxPower
						//add刷一层的增加
						for add := start; add < len(powerNums); add++ {
							powerNums[add]++
						}
						start++
					}
					start = end + 1
				}
			}
		}
	}
	fmt.Println(powerNums)
	curMax := powerNums[0]
	for _,v:=range powerNums[1:]{
		if v>curMax{
			curMax = v
		}
	}
	fmt.Println("当前最大战力",curMax)
}
