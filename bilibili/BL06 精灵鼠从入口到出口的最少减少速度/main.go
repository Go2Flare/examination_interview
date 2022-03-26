package main


//猛兽侠中精灵鼠在利剑飞船的追逐下逃到一个n*n的建筑群中，精灵鼠从（0,0）的位置进入建筑群，建筑群的出口位置为（n-1,n-1），建筑群的每个位置都有阻碍，每个位置上都会相当于给了精灵鼠一个固定值减速，因为精灵鼠正在逃命所以不能回头只能向前或者向下逃跑，现在问精灵鼠最少在减速多少的情况下逃出迷宫？
//
//输入描述：
//第一行迷宫的大小: n >=2 & n <= 10000；
//第2到n+1行，每行输入为以','分割的该位置的减速,减速f >=1 & f < 10。
//输出描述：
//BL06 精灵鼠从入口到出口的最少减少速度？
//示例1
//输入：
//3
//5,5,7
//6,7,8
//2,2,4
//复制
//输出：
//19

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r *bufio.Reader

func getShortestPath(m [][]int){
	//     初始化
	row := len(m)
	col := len(m[0])
	dp:= make([][]int, row)
	for i,_:=range dp{
		dp[i] = make([]int, col)
	}
	copy(dp, m)
	//     dp初始化
	for i := range dp[0]{
		if i>0{
			dp[0][i] += m[0][i-1]
		}
	}
	for i := range dp{
		if i>0{
			dp[i][0] += m[i-1][0]
		}
	}
	for i:=1; i<row; i++{
		for j:=1; j<col; j++{
			dp[i][j] = min(dp[i-1][j], dp[i][j-1])+m[i][j]
		}
	}
	fmt.Println(dp[row-1][col-1])
}
func min(a,b int ) int{
	if a<b{
		return a
	}
	return b
}

func splitStdString(b byte, s string) []string{
	str,_ := r.ReadString(b)
	str = strings.Trim(str, string(b))
	strS := strings.Split(str, s)
	return strS
}

func main() {
	r = bufio.NewReader(os.Stdin)
	nStr := splitStdString('\n'," ")
	n,_ := strconv.Atoi(nStr[0])
	//     矩阵
	m := make([][]int, n)
	for i:=0; i<n; i++{
		row := splitStdString('\n', ",")
		m[i] = make([]int, n)
		for j:=0; j<n; j++{
			e,_ := strconv.Atoi(row[j])
			m[i][j] = e
		}
	}
	getShortestPath(m)
	//     fmt.Println(m)
}