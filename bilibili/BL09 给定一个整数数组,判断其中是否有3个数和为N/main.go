package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ThreeSum(nums []int, target int) {
	sort.Ints(nums)
	stand := 0
	//     只要第一层去重逻辑就可以了
	for stand=0; stand<len(nums) && nums[stand] <= target; stand++{
		if stand == 0 || nums[stand] != nums[stand-1]{
			targetNext, i, j:= target-nums[stand], stand+1, len(nums)-1
			for i<j{
				//                 这里只要有一种情况符合即可
				if nums[i] + nums[j] == targetNext{
					fmt.Println("True")
					os.Exit(0)
				}else if nums[i] + nums[j] < targetNext {
					i++
				}else if nums[i] + nums[j] > targetNext{
					j--
				}
			}
		}
	}
	fmt.Println("False")

}

var r *bufio.Reader
var nums []int

var target int
func splitStdString(b byte , s string) []string {
	r = bufio.NewReader(os.Stdin)
	str,_ := r.ReadString(b)
	str = strings.Trim(str, string(b))
	return strings.Split(str, s)
}

//我的方法
func myPower() {
	numsStr := splitStdString(',', " ")
	nums = make([]int, len(numsStr))
	for i:=0; i<len(numsStr); i++{
		nums[i],_ = strconv.Atoi(numsStr[i])
	}

	num,_ := r.ReadBytes('\n')
	numStr := strings.Trim(string(num), "\n")
	target,_ = strconv.Atoi(numStr)

	ThreeSum(nums, target)
}


//排行榜老哥
func higherPower() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	list := strings.Split(input.Text(), ",")
	n, _ := strconv.Atoi(list[1])
	nums := make([]int, 0)
	for _, str := range strings.Split(list[0], " ") {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	numsLen := len(nums)
	sort.Ints(nums)
	for i := 1; i < numsLen; i++ {
		first, last := 0, numsLen-1
		for first < i && last > i {
			sum := nums[first] + nums[i] + nums[last]
			//fmt.Println(first, "-", nums[first], i, "-", nums[i], last, "-", nums[last], sum)
			if sum == n {
				fmt.Println("True")
				os.Exit(0)
			} else if sum < n {
				first += 1
			} else {
				last -= 1
			}
		}
	}
	fmt.Println("False")
}

func main() {
	myPower()
	higherPower()
}