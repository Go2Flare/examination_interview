package main

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
	strs := utilsReadLine(" ", 0).([]string)
	nums := utilsReadLine(" ", 1).([]int)
	n := utilsReadLine(" ", 2).(int)
	fmt.Println(strs, nums, n)
}