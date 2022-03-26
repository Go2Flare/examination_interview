package main

import (
	"fmt"
	//     "bufio"
	//     "os"
	//     "strings"
	//     "strconv"
)

// var r *bufio.Scanner
// var ipStr []string
// var ipPart int
// func 02mapInterface(){
//     r = bufio.NewScanner(os.Stdin)
//     for r.Scan(){
//         ipStr = strings.Split(r.Text(), ".")
//         wangGuan := make([]int, len(ipStr))
//         for i:=0; i<len(ipStr); i++{
//             ipPart, _ := strconv.Atoi(ipStr[i])
// //             fmt.Println(ipPart)
//             wangGuan[i] = ipPart
//         }
//         if wangGuan[0] == 10 || (wangGuan[0] ==192 && wangGuan[1] ==168) || ((wangGuan[1] >=16 && wangGuan[1] <32)&&wangGuan[0] ==172){
//             fmt.Println(1)
//         }else{
//             fmt.Println(0)
//         }
// //         fmt.Println(wangGuan)
//     }
// }

var a,b,c,d int
func main() {
	fmt.Scanf("%v.%v.%v.%v", &a, &b, &c, &d)
	if a==10 || a==192&&b==168 || (b>=16 && b<32)&&a == 172{
		fmt.Println(1)
	}else{
		fmt.Println(0)
	}
}