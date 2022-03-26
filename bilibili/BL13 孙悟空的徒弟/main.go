package main

//描述
//打败魔人布欧以后，孙悟空收了n个徒弟，每个徒弟战斗力各不相同。他教导所有的徒弟和体术，合体后战斗力为原战斗力相乘。任何两个徒弟都可以合体，所以一共有n*(n-1)/2种合体徒弟。有一天，他想考验一下孙悟天战斗力如何，希望在所有n*(n-1)/2种合体徒弟中选择战斗力第k高的，与孙悟天对战。可是孙悟空徒弟太多了，他已然懵逼，于是找到了你，请你帮他找到对的人。
//
//输入描述：
//第一行两个int。徒弟数量：n <= 1*10^6；战斗力排名:k <= n*(n-1)/2
//第二行空格分隔n个int，表示每个徒弟的战斗力。
//输出描述：
//战斗力排名k的合体徒弟战斗力。
//示例1
//输入：
//5 2
//1 3 4 5 9
//复制
//输出：
//36

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var r *bufio.Reader
var nums []int

func ReadStdString(b byte, s string) []int{
	str,_ := r.ReadString(b)
	str = str[:len(str)-1]
	strs := strings.Split(str, s)
	nums = make([]int, len(strs))
	for i:=0; i<len(strs); i++{
		nums[i], _ = strconv.Atoi(strs[i])
	}
	return nums
}

func main() {
	r = bufio.NewReader(os.Stdin)
	readIn := ReadStdString('\n', " ")
	n, k := readIn[0], readIn[1]
	//     fmt.Println(n,k)
	nums = ReadStdString('\n', " ")
	//     fmt.Println(nums)
	//     二分查找
	sort.Ints(nums)
	l,r := 1, nums[n-1]*nums[n-2]
	for l<r {
		m := (l+r+1)>>1
		count :=0
		i, j := 0, n-1
		for i<j&&count<k{
			//             找到接近中间线的乘积的排行
			for i<j&&nums[i]*nums[j] <m{
				i++
			}
			//             count加上当前这个循环大于m的排名
			count += j-i
			//             j还要往下循环，因为还存在可能会有大于mid的元素
			j--
		}
		if count >=k{//排行在k之后的的，说明还不够，左边赋值中间线
			l = m//最后找到的m就是第k个
		}else if count < k{
			r = m-1
		}
	}
	fmt.Println(l)
}