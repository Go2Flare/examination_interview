package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//小A参加了一个n人的活动，每个人都有一个唯一编号i(i>=0 & i<n)，其中m对相互认识，在活动中两个人可以通过互相都认识的一个人介绍认识。现在问活动结束后，小A最多会认识多少人？
//
//输入描述：
//第一行聚会的人数：n（n>=3 & n<10000）；
//第二行小A的编号: ai（ai >= 0 & ai < n)；
//第三互相认识的数目: m（m>=1 & m
//< n(n-1)/2）；
//第4到m+3行为互相认识的对，以','分割的编号。
//输出描述：
//输出小A最多会新认识的多少人？
//示例1
//输入：
//7
//5
//6
//1,0
//3,1
//4,1
//5,3
//6,1
//6,5
//复制
//输出：
//3

var inputReader *bufio.Reader
var input string
var err error

var r *bufio.Scanner

//并查集
var fa []int

//祖父数组
func faInit(n int) {
	for i := 1; i <= n; i++ {
		fa[i] = i
	}
}

//   查
func find(x int) int {
	if x == fa[x] {
		//找到祖父
		return x
	} else {
		//不是祖父节点, 向下查找,
		fa[x] = find(fa[x])
		return fa[x]
		//不要写成这种形式, 我们查找到的节点
		//return find(fa[x])
	}
}

//   并
func union(x int, y int) {
	fx := find(x)
	fy := find(y)
	fa[fx] = fy
}



func findFriend2() {
	inputReader = bufio.NewReader(os.Stdin)
	//     长度，当前结点，关系数量
	var n, ai, m int
	condition := [3]string{}
	for i := 0; i < 3; i++ {
		str, _ := inputReader.ReadString('\n')
		condition[i] = strings.Trim(str, "\n")
	}

	n, _ = strconv.Atoi(condition[0])
	ai, _ = strconv.Atoi(condition[1])
	m, _ = strconv.Atoi(condition[2])

	fa = make([]int, n+1)
	//     初始化
	faInit(n)

	//     结果
	var res, already int

	var x, y int
	for i := 0; i < m; i++ {
		input, _ = inputReader.ReadString('\n')
		str := strings.Trim(input, "\n")
		//         fmt.Println(str)
		sSplited := strings.Split(str, ",")
		x, _ = strconv.Atoi(sSplited[0])
		y, _ = strconv.Atoi(sSplited[1])

		//      有关当前结点的关系，先提前算，最后减掉
		if x == ai || y == ai {
			already++
		}
		//       union
		if x != y {
			union(x, y)
		}
		fmt.Println(fa)
	}
	for i := 0; i < n; i++ {
		//         当前ai的祖父和i的祖父相交
		if find(ai) == find(i) {
			res++
		}
	}
	fmt.Println(res - already - 1)
}

func main() {
	//并查集
	//findFriend1()
	//邻接表
	//findFriend2()
	//review
	findFriend3()
}


func initFa(N int) {
	for i := 1; i <= N; i++ {
		fa[i] = i
	}
}
func findFa(x int) int {
	if x == fa[x] {
		return x
	} else {
		//每个经过节点都要接收父亲,建立一一指向
		fa[x] = findFa(fa[x])
		return fa[x]
	}
}

func unionFa(x1, x2 int) {
	x1Fa := findFa(x1)
	x2Fa := findFa(x2)
	fa[x1Fa] = x2Fa
}

func findFriend3() {
	inputReader = bufio.NewReader(os.Stdin)
	//	初始化数组fa
	var n, ai, m int
	conditions := [3]string{}
	for i := 0; i < 3; i++ {
		str, _ := inputReader.ReadString('\n')
		conditions[i] = strings.Trim(str, "\n")
	}
	n, _ = strconv.Atoi(conditions[0])
	ai, _ = strconv.Atoi(conditions[1])
	m, _ = strconv.Atoi(conditions[2])

	fa = make([]int, n+1)
	// 初始化fa
	initFa(n)
	fmt.Println(fa)

	var res, already int
	for i := 0; i < m; i++ {
		strs, _ := inputReader.ReadString('\n')
		strs = strings.Trim(strs, "\n")
		strsSplit := strings.Split(strs, ",")
		a1, _ := strconv.Atoi(strsSplit[0])
		a2, _ := strconv.Atoi(strsSplit[1])
		//fmt.Println(a1, a2)
		if a1 == ai || a2 == ai {
			already++
		}
		//执行并操作,将关系
		//	并操作 将每行的相关两个节点建立关系
		//	 两个节点中有一个是要查的节点, 先记录
		if a1 != a2 {
			unionFa(a1, a2)
		}
		//fmt.Println(fa)
	}

	for i:=0; i<n; i++{
		if findFa(ai) == findFa(i) {
			res++
		}
	}
	fmt.Println(res- already -1)
}

func findFriend1() {
	//长度，当前结点，关系数量
	var n, ai, m int
	r = bufio.NewScanner(os.Stdin)
	conditions := [3]string{}
	for i := 0; i < 3; i++ {
		r.Scan()
		conditions[i] = r.Text()
	}

	n, _ = strconv.Atoi(conditions[0])
	ai, _ = strconv.Atoi(conditions[1])
	m, _ = strconv.Atoi(conditions[2])
	//     fmt.Println(n,ai,m)
	G := make([][]int, n)

	var j, k int
	//     生成邻接表
	for i := 0; i < m; i++ {
		r.Scan()
		numSlice := strings.Split(r.Text(), ",")
		j, _ = strconv.Atoi(numSlice[0])
		k, _ = strconv.Atoi(numSlice[1])
		//         fmt.Println(rStr,numSlice,j,k)
		G[j] = append(G[j], k)
		G[k] = append(G[k], j)
	}

	//     fmt.Println(G)
	var res int

	var dfs func(ai int, v []bool)
	dfs = func(ai int, v []bool) {
		for i := 0; i < len(G[ai]); i++ {
			if !v[G[ai][i]] {
				v[G[ai][i]] = true
				res++
				dfs(G[ai][i], v)
			}
		}
		return
	}

	//     所有元素的备忘录
	visited := make([]bool, n)
	//     邻接表记忆化递归查找
	dfs(ai, visited)
	//     当前结点的邻居和自身排除
	already := len(G[ai]) + 1
	fmt.Println(res - already)
}
