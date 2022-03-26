package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
假象一个绝对光滑的，高度很高的盒子，长度为，宽度为1。在其中，有许多的光滑铁块，铁块的每一个角一定位于整数坐标上。

由于宽度为1，我们可以用二维字符图来表示每一个铁块的位置。如下图是一个长度为4，其中有7个铁块的二维字符图（'x'代表铁块，'o'代表没有铁块）：
oooo
xooo
xoxo
xxxx
每一列的铁块数分别为3,1,2,1
由于重力的缘故，所有的铁块要么下面是盒子底面，要么下面是另一个铁块。现在，在盒子的右边增加一个强磁铁。所有右边没有其他铁块或边界的铁块会向右移动，直到撞上一个铁块和边界停下。
在上一张二维图上，加入磁铁后的字符图会变为：
oooo
ooox
ooxx
xxxx
每一列的铁块数分别为。可以证明，这样操作后所有铁块要么下面是盒子底面，要么下面是另一个铁块。
现在给你初始每一列有多少个铁块，请你计算，加入磁铁后每一列有多少铁块。
*/


var r *bufio.Reader
func main() {
	r = bufio.NewReader(os.Stdin)
	readStr,_ := r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	n,_ := strconv.Atoi(readStr)

	readStr,_ = r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	readStrs := strings.Split(readStr," ")
	nums := make([]int, n)
	for i:=0;i<n; i++{
		nums[i], _ = strconv.Atoi(readStrs[i])
	}
	//fmt.Println(n, nums)
	sort.Ints(nums)
	for i, v:= range nums{
		if i == 0{
			fmt.Print(v)
		}else{
			fmt.Print(" ",v)
		}

	}
	fmt.Println()
}
