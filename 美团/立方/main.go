package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*立方
时间限制： 3000MS
内存限制： 1048576KB
题目描述：
小团正在周游世界。他来到了立方之国，想要通过城门必须通过守卫的考验。
守卫给出了n个位置，每个位置上有一个数字，以及m个立方之碑，他可以将立方之碑放在这n个位置上，一
放上去那个位置的数就会变成它原本的立方，比如，2变成8，-3变成-27。但因为位置比较小，每个位置最多放一个。
守卫想知道小团能否通过放置最多m个（也可以不放）立方之碑，让n个位置对应的数之和成为指定的数val。如果方案存在，请输出YES，否则输出NO。

输入描述
第一行一个整数T，表示接下来有T组数据

每组数据第一行三个整数n，m，val，分别表示位置的数量，立方之碑的数量以及指定的数。
每组数据第二行n个整数，a[1],a[2],…,a[n]，表示这n个位置初始的数是多少。
数字间两两空格隔开。
其中，60%的数据保证，1<=T<=5，1<=m<=n<=10，-100<=a[i]<=100，-10,000,000 <=val<=10,000,000
100%的数据保证，1<=T<=10，1<=m<=n<=20，-1,000<=a[i]<=1,000，-10,000,000,000 <=val<=10,000,000,000

输出描述
输出T行，每行输出一个字符串，YES或NO，表示这个方案存在或不存在。

样例输入
2
2 1 10
1 2
4 3 37
1 1 2 3

样例输出
NO
YES

提示
样例中T为2，一共有两组数据。
第一组，一共有2个位置，1个立方之碑，无论立方之碑放在第一个位置还是第二个位置还是不放立方之碑，都无法满足要求，输出NO。
第二组，一共有4个位置，3个立方之碑，如果在第三个和第四个位置放置了立方之碑，这4个位置上的数就会变成1 1 8 27，它们的和是37，满足要求，输出YES。*/

var r *bufio.Reader

func StdReadString(b byte, s string) []int {
	readStr, _ := r.ReadString(b)
	readStr = readStr[:len(readStr)-1]
	readStrs := strings.Split(readStr, s)
	nums:=make([]int, len(readStrs))
	for i:= 0; i<len(readStrs); i++{
		nums[i], _ =strconv.Atoi(readStrs[i])
	}
	return nums
}

func sum(nums []int)(res int) {

	for i:=0; i<len(nums); i++{
		res += nums[i]
	}
	return
}

func main() {
	r = bufio.NewReader(os.Stdin)
	t := StdReadString('\n', " ")
	for i:=0; i<t[0]; i++{
		str := StdReadString('\n', " ")
		n, m, val := str[0], str[1], str[2]
		nums := StdReadString('\n', " ")
		label1:
		for i:=0; i<m; i++ {
			sort.Ints(nums)

			for j:=n-1; j>=0; j--{
				numsCp1 := make([]int, len(nums[:j]))
				numsCp2 := make([]int, len(nums[j+1:]))
				copy(numsCp1, nums[:j])
				copy(numsCp2, nums[j+1:])

				//fmt.Println("j", j, nums, sum(append(numsCp1, numsCp2...)))
				target := val - sum(append(numsCp1, numsCp2...))
				powNumsI := int(math.Pow(float64(nums[j]),3))
				//fmt.Println(j, nums, powNumsI, target,powNumsI< target )
				if  powNumsI == target {
					nums[j] = powNumsI
					fmt.Println("YES")
					break label1
				}else if powNumsI< target {
					nums[j] = powNumsI
					break
				}else if j == 0{
					fmt.Println("NO")
					break label1
				}

			}
		}
		if sum(nums) != val {
			fmt.Println("NO")
		}
	}
}
