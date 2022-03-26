package main

//描述
//av394281 中，充满威严的蕾米莉亚大小姐因为触犯某条禁忌，被隙间妖怪八云紫（紫m……èi）按住头在键盘上滚动。
//同样在弹幕里乱刷梗被紫姐姐做成罪袋的你被指派找到大小姐脸滚键盘打出的一行字中的第 `k` 个仅出现一次的字。
//(为简化问题，大小姐没有滚出 ascii 字符集以外的字)
//
//
//输入描述：
//每个输入都有若干行，每行的第一个数字为`k`，表示求第`k`个仅出现一次的字。然后间隔一个半角空格，之后直到行尾的所有字符表示大小姐滚出的字符串`S`。
//输出描述：
//输出的每一行对应输入的每一行的答案，如果无解，输出字符串`Myon~`
//
//(请不要输出多余的空行）
//
//为了方便评测，如果答案存在且为c，请输出[c]
//示例1
//输入：
//2 misakamikotodaisuki
//3 !bakabaka~ bakabaka~ 1~2~9!
//3 3.1415926535897932384626433832795028841971693993751o582097494459211451488946419191919l91919hmmhmmahhhhhhhhhh
//7 www.bilibili.com/av170001
//1 111
//复制
//输出：
//[d]
//[9]
//[l]
//[7]
//Myon~
//复制
//备注：
//字符串S仅包含可见ascii码，长度不超过100000

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)
var r *bufio.Reader
var r1 *bufio.Scanner

func ReadString1() {
	r = bufio.NewReader(os.Stdin)
	for {
		strs, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		strs = strs[:len(strs)-1]

		ss := strings.SplitN(strs, " ", 2)
		n, _ := strconv.Atoi(ss[0])

		strs = ss[1]
		m := make([]int, 256)
		for i, v := range strs {
			p := m[v]
			//三种状态0，-1，1
			if p != 0 {
				if p > 0 {
					m[v] = -1
				}
			} else {
				m[v] = i + 1
			}
		}
		sort.Ints(m)

		for _, v := range m {
			if v > 0 {
				n--
				if n == 0 {
					fmt.Printf("[%c]\n", strs[v-1])
					break
				}
			}
		}
		if n != 0 {
			fmt.Println("Myon~")
		}
	}
}

func ReadLine2() {
	r = bufio.NewReader(os.Stdin)

	for {
		str, _, err := r.ReadLine()
		//不能用Scanner，有可能是Scanner扫描不了太长的行
		if err == io.EOF {
			return
		}

		ss := strings.SplitN(string(str), " ", 2)
		n, err := strconv.Atoi(string(ss[0]))

		strs := ss[1]
		m := make([]int, 256)
		for i, v := range strs {
			p := m[v]
			//三种状态0，-1，1
			if p != 0 {
				if p > 0 {
					m[v] = -1
				}
			} else {
				m[v] = i + 1
			}
		}
		sort.Ints(m)

		for _, v := range m {
			if v > 0 {
				n--
				if n == 0 {
					fmt.Printf("[%c]\n", strs[v-1])
					break
				}
			}
		}
		if n != 0 {
			fmt.Println("Myon~")
		}
	}
}

func ReadScanner3() {
	r1 = bufio.NewScanner(os.Stdin)
	//最后一个用例不能直接用Scanner
	//Scanner扫描不了太长的行,除非我们能预设token大小的缓冲池
	r1.Buffer(make([]byte, 1), 1.3*64500) //切片是设置分隔符的缓冲大小,试了一会的最小token长度
	for r1.Scan() {
		str := r1.Text()
		ss := strings.SplitN(str, " ", 2)
		n, _ := strconv.Atoi(ss[0])

		strs := ss[1]
		m := make([]int, 256)
		for i, v := range strs {
			p := m[v]
			//三种状态0，-1，1
			if p != 0 {
				if p > 0 {
					m[v] = -1
				}
			} else {
				m[v] = i + 1
			}
		}
		sort.Ints(m)

		for _, v := range m {
			if v > 0 {
				n--
				if n == 0 {
					fmt.Printf("[%c]\n", strs[v-1])
					break
				}
			}
		}
		if n != 0 {
			fmt.Println("Myon~")
		}
	}
}

func main() {
	//ReadString1()
	//ReadLine2()
	ReadScanner3()
}