package main

/*

现在你在玩一款游戏，叫做节奏小师。它有三种判定
P : Perfect完美，加200分。
G : Great很棒，加100分。
M : Miss错过，不加分也不扣分，但累计三次Miss就会输掉游戏。
 另外有一种奖励是连击奖励。一旦连续三个Perfect之后，后续连击的Perfect分数将变成250分，但一旦打出了Great或者Miss则连击数将重新开始计算。
你的任务是根据游戏记录计算分数。特别地，失败记为零分。
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const(
	Perfect = 200
	SuperPerfect = 250
	Great = 100
)

var r *bufio.Reader

func getScore(str string)(res int){
	miss := 0//miss3次返回0
	perfect := 0
	for i:=0; i<len(str); i++{
		if miss == 3 {
			//输掉游戏
			res = 0
			break
		}else if str[i]=='P'{
			if perfect >= 3 {
				res += SuperPerfect
			}else{
				res += Perfect
			}
			perfect++
		}else if str[i]=='M'{
			//if perfect >= 3{
				perfect =0
			//}
			miss++
		}else if str[i]=='G'{
			//if perfect >= 3{
				perfect =0
			//}
			//是否+else
			res += Great
		}
	}
	return
}

func main() {
	r = bufio.NewReader(os.Stdin)
	for {
		readStr,err := r.ReadString('\n')
		if err == io.EOF{
			break
		}
		readStr = readStr[:len(readStr)-1]

		fmt.Println(getScore(readStr))
	}

}
/*
func 02mapInterface() {
    str := ""
    fmt.Scanf("%s", str)
    fmt.Println(getScore(str))
}
*/

//package 02mapInterface
//import "fmt"

func getReward(s string) int {
	countM := 0
	combo := 0
	ans := 0
	for i := range s{
		if s[i] == 'M'{
			combo = 0
			countM++
			if countM >= 3 {
				return 0
			}
		} else if s[i] == 'P'{
			combo++
			if combo > 3 {
				ans += 250
			}else {
				ans += 200
			}
		} else {
			combo = 0
			ans += 100
		}
	}
	return ans
}

func main(){
	var s string
	fmt.Scanln(&s)
	fmt.Println(getReward(s))
}