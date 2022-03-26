package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//1196多个数组的共有最小数
//1,2
//2,4
//3,5
//1,3

var r *bufio.Scanner

func main() {
	r = bufio.NewScanner(os.Stdin)
	var all [][]int
	for r.Scan() {
		str := r.Text()
		strSlice := strings.Split(str, ",")
		nums := make([]int, len(strSlice))
		for i := 0; i < len(nums); i++ {
			nums[i], _ = strconv.Atoi(strSlice[i])
		}
		all = append(all, nums)
	}

	fmt.Println(findE(all))
}

func findE(all [][]int) (res int) {
	if len(all) == 0{
		return 0
	}
	n := len(all)
	var min =  math.MinInt64
	for {
		for i := 0; i < n; i++ {
			//选出第一个元素最大
			if len(all[i]) != 0 && all[i][0] > min {
				min = all[i][0]
			}
		}

		/*排除最大的元素之前的元素*/
		for i := 0; i < n; i++ {
			for len(all[i]) != 0 && all[i][0] < min {
				all[i] = all[i][1:]
			}
		}

		//有一个数组为空就返回-1
		for i:=0; i<len(all); i++{
			if len(all[i]) == 0{
				return -1
			}
		}

		flag := true
		for i := 0; i < n-1; i++ {
			if len(all[i]) != 0 &&len(all[i+1]) != 0 && all[i][0] != all[i+1][0] {
				flag = false
				break
			}
		}

		if flag {
			var e int
			for i:=0; i<len(all); i++{
				if len(all[i])!=0{
						e = all[i][0]
					break
				}
			}
			return e
		}
	}
}
