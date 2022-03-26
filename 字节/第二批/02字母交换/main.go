package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var br *bufio.Reader


//分割读出的行,返回int切片或字符串切片
func utilsReadString(b byte, s string, t int) (itemSlice interface{}){
	//b:读取的结尾, s:切割的分隔符, t:类型{0:字符串切片, 1:数组切片, 2只读一个数字}
	str,_ := br.ReadString(b)
	str = str[:len(str)-1]
	strSlice := strings.Split(str, s)
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
	str := utilsReadString('\n', " ", 0).([]string)
	s := str[0]
	n,_ := strconv.Atoi(str[1])
	fmt.Println(DP(s,n))
}

func DP(s string, n int) int{
	res :=1
	for c:='a'; c<='z'; c++{
		pos := []int{}
		for i:=0; i<len(s); i++{
			if s[i] == uint8(c){
				pos = append(pos, i)
			}
		}

		if len(pos) < 2{
			continue
		}

		dp := make([][]int, len(pos))
		for i:=range dp{
			dp[i]= make([]int, len(pos))
		}
		//	存取当前重复字母的下标数组
		//	解题关键是，相同元素的数组，要使其排列在一起的操作次数怎么计算
		//	dp[l][r]是区间内两个数要靠近的最少移动次数，与dp[l+1][r-1]相关，就是左右都减一位的区间
		//	将l数与r数移动到连续的中间需要 pos[r]-pos[l]-1次，排除r-l-1个中间元素（少移动的次数）
		//	dp[l][r] = dp[l+1][r-1] + pos[r] - pos[l] -1 -(r-l-1)

		//先将相邻区域初始化
		for i:=1; i<len(pos); i++{
			dp[i-1][i] = pos[i] - pos[i-1] - 1
		}
		for rig:=2; rig<len(pos); rig++{//区间长度
			for l:=0; l<len(pos)-rig; l++{//区间起始坐标
				r := l+rig
				//当前区间的操作次数
				dp[l][r] = dp[l+1][r-1] + (pos[r]-pos[l]- 1) - (r-l-1)
			}
		}
		for i:=0; i<len(dp); i++{
			for j:=i+1; j<len(dp[0]); j++{
				if dp[i][j] <= n{
					res = max(res, j-i+1)
				}
			}
		}
	}
	return res
}

func max(a,b int) int{
	if a>b{return a}
	return b
}
