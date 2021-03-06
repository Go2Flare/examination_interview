package main

//22娘和33娘接到了小电视君的扭蛋任务：
//一共有两台扭蛋机，编号分别为扭蛋机2号和扭蛋机3号，22娘使用扭蛋机2号，33娘使用扭蛋机3号。
//扭蛋机都不需要投币，但有一项特殊能力：
//扭蛋机2号：如果塞x（x范围为>=0正整数）个扭蛋进去，然后就可以扭到2x+1个
//扭蛋机3号：如果塞x（x范围为>=0正整数）个扭蛋进去，然后就可以扭到2x+2个
//22娘和33娘手中没有扭蛋，需要你帮她们设计一个方案，两人“轮流扭”（谁先开始不限，扭到的蛋可以交给对方使用），用“最少”的次数，使她们能够最后恰好扭到N个交给小电视君。
//输入描述：
//输入一个正整数，表示小电视君需要的N个扭蛋。
//输出描述：
//输出一个字符串，每个字符表示扭蛋机，字符只能包含"2"和"3"。
//示例1
//输入：
//10
//复制
//输出：
//233
//复制
//备注：
//1<=N<=1e9

import (
	"fmt"
	//     "bufio"
	//     "os"
	//     "strconv"
)

func getEggStr(N int) string {
	var res string
	//     greedy
	for N!=0{
		//偶数
		if N%2 == 0{
			N = (N-2)>>1
			res = "3"+res
		}else {//奇数
			N = (N-1)>>1
			res = "2"+res
		}
	}
	return res
}

func main() {
	var wantEgg int
	//     fmt.Scanf("%v", &wantEgg)
	fmt.Scanln(&wantEgg)

	fmt.Println(getEggStr(wantEgg))
}


// func 02mapInterface() {
// 	s := []string{}
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		s1 := input.Text()
// 		s = append(s, s1)
// 	}
// 	for _,v := range s{
// //         逻辑
//         N,_ := strconv.Atoi(v)
//         fmt.Println(getEggStr(N))
// 	}
// }