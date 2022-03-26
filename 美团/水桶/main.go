package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*水桶
时间限制： 3000MS
内存限制： 1048576KB
题目描述：
小美有一个水桶，它由n个木板构成。我们知道水桶效应，一个水桶能装多少水由最短的木板决定。现在小美有一个神奇的魔法，可以指定一个木板让其长度增加1，小美最多用m次这样的魔法。她想知道如果她合理使用这个魔法，这个水桶能最多装多高的水。

形式化地，给出n个数a[1],a[2],…a[n],表示n个木板的高度，每次操作可以选定一个i ,1<=i<=n,使得a[i]变成a[i]+1，这样的操作最多进行m次。问最优情况下，min{a[1],a[2],…a[n]}最大是多少。



输入描述
第一行两个空格隔开的整数n和m，表示木板数量和魔法最多使用次数。

第二行n个数a[1],a[2],…a[n],表示木板高度，空格隔开。

1<=n<=100,000，1<=m<=1,000,000,000，

1<=a[i]<= 1,000,000,000

输出描述
一行一个数，代表木桶最多能装的水的高度。


样例输入
6 9
6 7 3 4 9 6
样例输出
7

提示
在第1块木板上使用1次魔法，长度从6变为7。
在第3块木板上使用4次魔法，长度从3变为7。
在第4块木板上使用3次魔法，长度从4变为7。
在第6块木板上使用1次魔法，长度从6变为7。
共使用9次魔方，此时最短木桶长度为7，即木桶最多能装的水的高度为7。
*/
var r *bufio.Reader


func main() {
	r = bufio.NewReader(os.Stdin)
	readStr, _ := r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	readStrs := strings.Split(readStr, " ")
	n , _:= strconv.Atoi(readStrs[0])
	m , _ :=strconv.Atoi(readStrs[1])

	readStr, _ = r.ReadString('\n')
	readStr = readStr[:len(readStr)-1]
	readStrs = strings.Split(readStr, " ")
	nums := make([]int, n)
	for i:=0; i<n; i++{
		nums[i],_ = strconv.Atoi(readStrs[i])
	}
	for m!=0{
		sort.Ints(nums)
		nums[0]++
		m--
	}

	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums[0])
}