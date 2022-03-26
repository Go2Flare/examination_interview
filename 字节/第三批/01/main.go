package main

/*
推箱子
时间限制： 3000MS
内存限制： 589824KB
题目描述：
有一个推箱子的游戏, 一开始的情况如下图:



上图中, '.' 表示可到达的位置, '#' 表示不可到达的位置，其中 S 表示你起始的位置, 0表示初始箱子的位置, E表示预期箱子的位置，你可以走到箱子的上下左右任意一侧, 将箱子向另一侧推动。如下图将箱子向右推动一格;

..S0.. -> ...S0.

注意不能将箱子推动到'#'上, 也不能将箱子推出边界;

现在, 给你游戏的初始样子, 你需要输出最少几步能够完成游戏, 如果不能完成, 则输出-1。



输入描述
第一行为2个数字,n, m, 表示游戏盘面大小有n 行m 列(5< n, m < 50); 后面为n行字符串,每行字符串有m字符, 表示游戏盘面;

输出描述
一个数字,表示最少几步能完成游戏,如果不能,输出-1;

例
输入：
3 6
.S#..E
.#.0..
......

输出：
11

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
	m, n := nums[0], nums[1]
	mat := make([][]byte, m)
	start := [2]int{}
	box := [2]int{}
	//end := [2]int{}
	for i:=0; i<len(mat); i++{
		mat[i] = make([]byte, n)
		strs := utilsReadLine(" ", 0).([]string)
		for j:=0; j<len(mat[i]);j++{
			mat[i][j] = strs[0][j]
			if strs[0][j] == 'S'{start = [2]int{i,j}}
			if strs[0][j] == '0'{box = [2]int{i,j}}
			//if strs[0][j] == 'E'{end = [2]int{i,j}}
		}
	}

	//判断是否合法坐标
	reasonablePos := func(x, y int) bool{
		return x>=0&&x<len(mat) && y>=0 && y<len(mat[0]) && mat[x][y] != '#'
	}

	//访问的复合坐标
	type pos struct{
		ax,ay,bx,by int
	}
	//标记访问过的复合坐标
	vis := make(map[pos]bool)

	//复合坐标
	type node struct{
		ax,ay,bx,by,step int
	}

	var (
		dx = [4]int{0,-1,0,1}
		dy = [4]int{-1,0,1,0}
	)
	var bfs func(startX, startY, boxX, boxY int)int
	bfs = func(startX, startY, boxX, boxY int)int{
		e0 := node{startX, startY, boxX, boxY, 0}
		queue := []node{e0}
		vis[pos{startX, startY, boxX, boxY}] = true

		for len(queue) >0{
			e := queue[0]
			queue = queue[1:]
			ax, ay, bx, by, step := e.ax, e.ay, e.bx, e.by, e.step
			for i:=0; i<4; i++{
				max,may, mbx,mby := ax+dx[i], ay+dy[i], bx+dx[i], by+dy[i]
				if !reasonablePos(max, may) {
					continue
				}
				//移动到箱子前
				_,ok1 := vis[pos{max,may,bx,by}]
				_,ok2 := vis[pos{max,may,mbx,mby}]
				if (max!=bx || may!=by)&&!ok1{
					vis[pos{max,may,bx,by}] = true
					queue = append(queue, node{max, may, bx, by, step+1})
				//人和箱子一起移动
				}else if (max == bx&&may==by&&reasonablePos(mbx,mby))&&!ok2{
					vis[pos{max,may,mbx,mby}]= true
					//fmt.Println(mat[mbx][mby],mbx,mby,step)
					if mat[mbx][mby] == 'E'{
						return step+1
					}
					queue = append(queue, node{max,may,mbx,mby,step+1})
				}

			}
		}
		return -1
	}

	fmt.Println(bfs(start[0], start[1], box[0], box[1]))
	//fmt.Println(vis)
}
