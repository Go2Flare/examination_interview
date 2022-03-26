package main

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
	calculate(nums)
	fmt.Println(count)
}

func a1(nums []int) {
	temp1, temp2 := nums[0], nums[2]

	nums[0], nums[2] = nums[6], nums[12]
	nums[6], nums[12] = nums[16], nums[18]
	nums[16], nums[18] = nums[20], nums[22]
	nums[20], nums[22] = temp1, temp2

	//左面旋转
	temp1 = nums[4]
	nums[4] = nums[5]
	nums[5] = nums[11]
	nums[11] = nums[10]
	nums[10] = temp1
}

func a2 (nums []int) {
	temp1, temp2 := nums[1], nums[3]

	nums[1], nums[3] = nums[7], nums[13]
	nums[7], nums[13] = nums[17], nums[19]
	nums[17], nums[19] = nums[21], nums[23]
	nums[21], nums[23] = temp1, temp2
	//右面旋转
	temp1 = nums[8]
	nums[8] = nums[14]
	nums[14] = nums[15]
	nums[15] = nums[9]
	nums[9] = temp1
}

func b1(nums []int) {

	temp1, temp2 := nums[2], nums[3]

	nums[2], nums[3] = nums[8], nums[14]
	nums[8], nums[14] = nums[17], nums[16]
	nums[17], nums[16] = nums[11], nums[5]
	nums[11], nums[5] = temp1, temp2
	//前面旋转
	temp1 = nums[6]
	nums[6] = nums[7]
	nums[7] = nums[13]
	nums[13] = nums[12]
	nums[12] = temp1
}

func b2 (nums []int) {
	temp1, temp2 := nums[0], nums[1]

	nums[0], nums[1] = nums[9], nums[15]
	nums[9], nums[15] = nums[19], nums[18]
	nums[19], nums[18] = nums[10], nums[4]
	nums[10], nums[4] = temp1, temp2
	//后面旋转
	temp1 = nums[23]
	nums[23] = nums[21]
	nums[21] = nums[20]
	nums[20] = nums[22]
	nums[22] = temp1
}

func c1(nums []int) {
	temp1, temp2 := nums[4], nums[5]

	nums[4], nums[5] = nums[6], nums[7]
	nums[6], nums[7] = nums[8], nums[9]
	nums[8], nums[9] = nums[23], nums[22]
	nums[23], nums[22] = temp1, temp2
	//上面旋转
	temp1 = nums[0]
	nums[0] = nums[1]
	nums[1] = nums[3]
	nums[3] = nums[2]
	nums[2] = temp1

}
func c2 (nums []int) {
	temp1, temp2 := nums[10], nums[11]

	nums[10], nums[11] = nums[12], nums[13]
	nums[12], nums[13] = nums[14], nums[15]
	nums[14], nums[15] = nums[21], nums[20]
	nums[21], nums[20] = temp1, temp2
	//后面旋转
	temp1 = nums[16]
	nums[16] = nums[18]
	nums[18] = nums[19]
	nums[19] = nums[17]
	nums[17] = temp1
}

type funcChoice []func(nums []int)

var res int
var count int
func calculate(nums []int) {
	//6种翻转方法
	fc := funcChoice{a1, a2, b1, b2, c1, c2}

	var dfs func(nums []int, index int)
	dfs = func(nums []int, index int) {
		if index == 3{
			culMax(nums)
			return
		}
		for i := 0; i < len(fc); i++ {
			cpNums := make([]int, len(nums))
			copy(cpNums, nums)
			//fmt.Println("操作前",nums)
			fc[i](nums)
			//fmt.Println(i, fc[i], index, countAll(nums))
			dfs(nums, index+1)
			copy(nums, cpNums)
			//fmt.Println("回溯后",nums)
		}
	}
	dfs(nums, 0)
	fmt.Println(res)
}

func countAll(nums []int) int{
	return nums[4]*nums[5]*nums[10]*nums[11] +  nums[0]*nums[1]*nums[2]*nums[3]+
		nums[6]*nums[7]*nums[12]*nums[13] + nums[8]*nums[9]*nums[14]*nums[15]+
		nums[16]*nums[17]*nums[18]*nums[19] + nums[20]*nums[21]*nums[22]*nums[23]
}

func culMax(nums []int){
	cur := countAll(nums)
	fmt.Println("curmax",cur, res, nums)
	if cur > res {
		res = cur
	}
}