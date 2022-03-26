package main

//描述
//给定一个合法的表达式字符串，其中只包含非负整数、加法、减法以及乘法符号（不会有括号），例如7+3*4*5+2+4-3-1，请写程序计算该表达式的结果并输出；
//
//输入描述：
//输入有多行，每行是一个表达式，输入以END作为结束
//输出描述：
//每行表达式的计算结果
//示例1
//输入：
//7+3*4*5+2+4-3-1
//2-3*1
//END
//输出：
//69
//-1
//备注：
//每个表达式的长度不超过10000，保证所有中间结果和最后结果在[-2e9,2e9]范围内

import (
"bufio"
"fmt"
"os"
"strconv"
)

var r *bufio.Scanner

//     符号栈
var sign []byte

// 数字栈
var res []int

// 比较优先级
func priorCompare(c byte, top byte) bool {
//     字符串末尾也要踢掉符号栈，因为肯定有一个符号留到最后，是最后一步
if c == '\n' {
return false
}
if c == '*' {
if top == '*' {
return false
} else {
return true
}
//         加减必须出栈一个
} else {
return false
}
}

// 计算结果
func getValue(a, b int, sign byte) int {
if sign == '*' {
return b * a
} else if sign == '+' {
return b + a
} else {
return b - a
}
}

func getRevPoland(s string) {
//     字符组成数字
digits := []byte{}
for i := 0; i < len(s); i++ {
c := s[i]
if '0' <= c && c <= '9' {
digits = append(digits, c)
} else {
num, _ := strconv.Atoi(string(digits))
res = append(res, num)
digits = []byte{}

fmt.Println(res, sign, len(sign) != 0 && !priorCompare(c, sign[len(sign)-1]))
//             遇到符号，数字栈取出两个数，用当前符号运算
//             条件是：符号栈不空，当前符号的可以踢掉符号栈的栈顶符号
for len(sign) != 0 && !priorCompare(c, sign[len(sign)-1]) {
//                 出栈两个数
nums1 := res[len(res)-1]
res = res[:len(res)-1]
nums2 := res[len(res)-1]
res = res[:len(res)-1]
fmt.Printf("当前运算数字和符号%v%c%v\n", nums1, byte(sign[len(sign)-1]), nums2)
// 两个数和符号栈顶的符号的运算结果加入数字栈
res = append(res, getValue(nums1, nums2, sign[len(sign)-1]))
sign = sign[:len(sign)-1]
}
sign = append(sign, c)
}
}
}

func main() {
r := bufio.NewScanner(os.Stdin)
for r.Scan() {
sign = []byte{}
res = []int{}
if r.Text() == "END" {
break
}
getRevPoland(r.Text() + "\n")
//最后数字栈剩下的为结果
fmt.Println("结果为：", res[0])
}
}
