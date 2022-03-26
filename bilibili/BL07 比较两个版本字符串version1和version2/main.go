package main

//描述
//如果version1 > version2 返回1，如果 version1 < version2 返回-1，不然返回0.
//
//输入的version字符串非空，只包含数字和字符.。.字符不代表通常意义上的小数点，只是用来区分数字序列。例如字符串2.5并不代表二点五，只是代表版本是第一级版本号是2，第二级版本号是5.
//
//输入描述：
//两个字符串，用空格分割。
//每个字符串为一个version字符串，非空，只包含数字和字符.
//输出描述：
//只能输出1, -1，或0
//示例1
//输入：
//0.1 1.1
//复制
//输出：
//-1
//复制
//备注：
//version1和version2的长度不超过1000，由小数点'.'分隔的每个数字不超过256。


import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var r *bufio.Scanner

func main() {
	r = bufio.NewScanner(os.Stdin)
label1:
	for r.Scan(){
		versionStr := r.Text()
		versionSlice := strings.Split(versionStr, " ")
		version1 := strings.Split(versionSlice[0],".")
		version2 := strings.Split(versionSlice[1],".")
		v1Index, v2Index := 0, 0
		for v1Index<len(version1) && v2Index<len(version2) {
			version1digit,_ := strconv.Atoi(version1[v1Index])
			version2digit,_ := strconv.Atoi(version2[v2Index])
			if version1digit < version2digit{
				fmt.Println(-1)
				break label1
			}else if version1digit > version2digit{
				fmt.Println(1)
				break label1
			}
			v1Index++
			v2Index++
		}
		if len(version1[v1Index:])>0{
			fmt.Println(1)
			break
		}else if len(version2[v2Index:])>0{
			fmt.Println(-1)
			break
		}

		fmt.Println(0)
		//         fmt.Println(version1, version2)
	}

}