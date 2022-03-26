package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


/*
机甲移动
时间限制： 3000MS
内存限制： 589824KB
题目描述：
     一台机甲在一个二维空间里移动，它的移动受到一串事先写好的指令的控制，每条指令会让它往特定方向移动一格。该空间为矩形，四周都是墙壁，无法通过，空间中每个格子位置的阻力不同，我们用一个数值来表示它的阻力，还有一些格子存在障碍无法通过。机器人每走一格，需要消耗能量，能量数为走之前和走之后位置的阻力数值中较大的那个数。当它的方向改变时，消耗能量 x，当它走向墙壁或者障碍物时，会停在原地并消耗能量 y（如果需要，仍会先改变方向）。一开始它位于整个空间的左上角，第一步移动不需要花能量改变方向，请计算它最后总共消耗了多少能量。

输入描述
    第一行四个整数 n, m, x, y，（0 < n, m <= 100，1 <= x, y <= 100000）
    后面 n 行，每行 m 个数，形成一个方阵，表示各个格子位置的阻力数值，如果为 -1，表示该位置无法通过。各个位置的数值范围为 [-1, 100000]。
    最后一行一个字符串，由 hjkl 四种字母组成，表示指令序列。h 表示向左移动一格，j 表示向下移动一格，k 表示向上移动一格，l 表示向右移动一格。字符串长度不超过 100000。
输出描述
一个整数，表示消耗的总能量。

样例输入
2 2 2 2
1 2
-1 3
kjljk
样例输出
20
*/

var r *bufio.Reader
type pos struct{
	x int
	y int
}

var allEnergy int

func getValue(mat [][]int, step, mstep int, str string) {
	n,m := len(mat), len(mat[0])
	p := &pos{
		x:0,
		y:0,
	}
	for i:=0; i<len(str); i++{
		if str[i] == 'h'{//左
			p.x -= 1
			if stop(p,n,m,mat) {
				p.x += 1
				allEnergy += mstep
			}else{
				allEnergy += step
			}
		}else if str[i] == 'j'{//下
			p.y += 1
			if stop(p,n,m,mat) {
				p.y -= 1
				allEnergy += mstep
			}else{
				allEnergy += step
			}

		}else if str[i] == 'l'{//右
			p.x += 1
			if stop(p,n,m,mat) {
				p.x -= 1
				allEnergy += mstep
			}else{
				allEnergy += step
			}
		}else if str[i] == 'k'{//上
			p.y -= 1
			if stop(p,n,m,mat) {
				p.y += 1
				allEnergy += mstep
			}else{
				allEnergy += step
			}
		}
	}
	fmt.Println(allEnergy)
}

func stop(p *pos, n,m int, mat [][]int) bool{
	if (p.x<0||p.x>=n) || (p.y<0||p.y>=m) || mat[p.x][p.y] == -1{
		return true
	}
	return false
}


func ReadStdString(b byte, s string) []int{

	readStr,_ := r.ReadString(b)
	readStr = readStr[:len(readStr)-1]
	readStrs := strings.Split(readStr, s)

	readNums := make([]int, len(readStrs))
	for i:=0; i<len(readStrs); i++{
		readNums[i],_ = strconv.Atoi(readStrs[i])
	}
	return readNums
}

func main(){

	allEnergy = 0//初始化
	r = bufio.NewReader(os.Stdin)
	readNums := ReadStdString('\n', " ")
	n, m, x, y := readNums[0],readNums[1],readNums[2],readNums[3]

	mat := make([][]int, n)
	for i:=0; i<n; i++{
		mat[i] = make([]int, m)
	}
	for i:=0; i<n; i++{
		//一行
		readNums := ReadStdString('\n', " ")
		for j:=0; j<m; j++{
			mat[i][j] = readNums[j]
		}
	}

	readStr,_ := r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]


	getValue(mat, x, y, readStr)

}