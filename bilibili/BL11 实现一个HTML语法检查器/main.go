package main

//描述
//BL11 实现一个HTML语法检查器。HTML语法规则简化如下：标签必须闭合，可以由开始和结束两个标签闭合，如<div></div>，也可以自闭合，
//如<div />
//标签可以嵌套如<div><a></a></div>或者 <div><a/></div>，但是标签不能交叉：<div><a></div></a>是不允许的标签里可以有属性
//如<div id="a<1"></div>
//属性的规则是name="任意非引号字符"，多属性声明之间必须有空格，属性声明不符合规则时，整段HTML都算语法错误
//输入文本只会出现字母a-z和<>"=
//请用任意语言实现一个HTML语法检查器函数，有语法错误返回1，没有语法错误返回0
//输入描述：
//一行，一个HTML字符串
//输出描述：
//有语法错误返回1，没有语法错误返回0
//示例1
//输入：
//<div><a></a></div>
//输出：
//0
//示例2
//输入：
//<div><a></div></a>
//输出：
//1
//备注：
//字符串长度不超过100


import (
	"bufio"
	"fmt"
	"os"
)

var r *bufio.Reader
var st []string
var elem []byte
func isHTML(line []byte) {
	for i:=0; i<len(line); i++{
		if line[i] == '<'{
			elem = elem[:0]
			i++
			var check bool
			if line[i] == '/'{
				check = true
				i++
			}
			for i<len(line)&&isChar(line[i]) {
				elem =append(elem, line[i])
				i++
			}
			if string(elem) ==""{
				fmt.Println(1)
				return
			}
			//             如果有/，出栈检查是否元素匹配
			if check {
				if st[len(st)-1]==string(elem) {
					st = st[:len(st)-1]
				}else {
					fmt.Println(1)
				}
			}else{
				st = append(st, string(elem))
			}
			i--
		}else if line[i] == '>' {
			if line[i-1] == '/'{
				st = st[:len(st)-1]
			}
		}else if isChar(line[i]){
			for i<len(line) && isChar(line[i]){
				i++
			}
			if line[i] != '='{
				for line[i] != '>'{
					i++
				}
				i--
				continue
			}
			i++
			if line[i] == '"'{
				i++
				for i<len(line) && line[i] != '"'{
					i++
				}
			}
		}
	}
	//     如果栈不为空，没有完全匹配
	if len(st) > 0{
		fmt.Println(1)
		return
	}
	fmt.Println(0)

}

func isChar(b byte) bool {
	return  'a' <=b && b<= 'z'
}

func main() {
	r = bufio.NewReader(os.Stdin)
	line, _, _ := r.ReadLine()
	//     fmt.Println(string(line))
	if len(line)>0{

		isHTML(line)
		//         fmt.Println(line)

	}else{
		fmt.Println(0)
	}
}